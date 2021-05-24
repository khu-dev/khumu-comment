package usecase

import (
	"context"
	"errors"
	"github.com/khu-dev/khumu-comment/data"
	"github.com/khu-dev/khumu-comment/data/mapper"
	"github.com/khu-dev/khumu-comment/ent"
	"github.com/khu-dev/khumu-comment/ent/article"
	"github.com/khu-dev/khumu-comment/ent/comment"
	"github.com/khu-dev/khumu-comment/ent/khumuuser"
	"github.com/khu-dev/khumu-comment/external"
	"github.com/khu-dev/khumu-comment/repository"
	"github.com/sirupsen/logrus"
)
import "github.com/khu-dev/khumu-comment/model"

var (
	DeletedCommentContent    string = "삭제된 댓글입니다."
	AnonymousCommentUsername string = "익명"
	AnonymousCommentNickname string = "익명"
	DeletedCommentUsername   string = "삭제된 댓글의 작성자"
	DeletedCommentNickname   string = "삭제된 댓글의 작성자"
	ErrUnAuthorized = errors.New("권한이 존재하지 않습니다")
)

type CommentUseCaseInterface interface {
	Create(commentInput *data.CommentInput) (*data.CommentOutput, error)
	List(username string, opt *repository.CommentQueryOption) ([]*data.CommentOutput, error)
	Get(username string, id int) (*data.CommentOutput, error)
	Update(username string, id int, opt map[string]interface{}) (*data.CommentOutput, error)
	Delete(username string, id int) error
}

type LikeCommentUseCaseInterface interface {
	// return value 중 bool이 true면 생성, false면 삭제
	Toggle(like *model.LikeComment) (bool, error)
}

type CommentUseCase struct {
	Repository            repository.CommentRepositoryInterface
	LikeCommentRepository repository.LikeCommentRepositoryInterface
	//EventMessageRepository repository.EventMessageRepository
	SnsClient     external.SnsClient
	Repo *ent.Client
}

type LikeCommentUseCase struct {
	Repository             repository.LikeCommentRepositoryInterface
	CommentRepository      repository.CommentRepositoryInterface
	EventMessageRepository repository.EventMessageRepository
}

type SomeoneLikesHisCommentError string

func NewCommentUseCase(repository repository.CommentRepositoryInterface,
	likeRepository repository.LikeCommentRepositoryInterface,
	snsClient external.SnsClient,
	repo *ent.Client) CommentUseCaseInterface {
	return &CommentUseCase{Repository: repository, LikeCommentRepository: likeRepository, SnsClient: snsClient, Repo: repo}
}

func (uc *CommentUseCase) Create(commentInput *data.CommentInput) (*data.CommentOutput, error) {
	logrus.Infof("Start Create Comment(%#v)", commentInput)
	//articleId := 1
	if commentInput.Author == "" {
		logrus.Error("댓글 생성에 대한 author가 존재하지 않습니다.")
		return nil, ErrUnAuthorized
	}
	newComment, err := uc.Repo.Comment.Create().
		SetNillableArticleID(commentInput.Article).
		//SetArticleID(&articleId).
		SetAuthorID(commentInput.Author).
		//SetAuthorID("bo314").
		SetContent(commentInput.Content).
		//SetContent("hello").
		SetState("exists").
		Save(context.Background())
	if err != nil {
		logrus.Error(newComment, err)
		return nil, err
	}

	return uc.ModelToOutput(commentInput.Author, newComment, nil), nil
}

func (uc *CommentUseCase) List(username string, opt *repository.CommentQueryOption) ([]*data.CommentOutput, error) {
	logrus.WithField("username", username).Infof("Start List CommentQueryOption(%#v)", opt)
	ctx := context.Background()
	query := uc.Repo.Comment.Query()
	if opt.AuthorUsername != "" {
		query.Where(comment.HasAuthorWith(khumuuser.ID(opt.AuthorUsername)))
	}
	if opt.ArticleID != 0 {
		query.Where(comment.HasArticleWith(article.ID(opt.ArticleID)))
	}

	parents, err := query.
		WithChildren().
		Where(comment.Not(comment.HasParent())).
		All(ctx)
	if err != nil {
		logrus.Errorf("comments 쿼리 도중 오류 발생. QueryOption(%+v)", opt)
		return nil, err
	}

	outputs := make([]*data.CommentOutput, 0)
	for _, parent := range parents {
		outputs = append(outputs, uc.ModelToOutput(username, parent, nil))
	}

	return outputs, nil
}

// 지금의 Get은 Children은 가져오지 못함
func (uc *CommentUseCase) Get(username string, id int) (*data.CommentOutput, error) {
	logrus.WithField("username", username).Infof("Start Get Comment(id:%#v)", id)
	ctx := context.Background()
	comment, err := uc.Repo.Comment.Query().
		WithChildren().
		Where(comment.ID(id)).
		Only(ctx)
	if err != nil {
		logrus.Errorf("comment 쿼리 도중 오류 발생.")
		return nil, err
	}

	output := uc.ModelToOutput(username, comment, nil)

	return output, nil
}

func (uc *CommentUseCase) Update(username string, id int, opt map[string]interface{}) (*data.CommentOutput, error) {
	logrus.WithField("username", username).WithField("id", id).Infof("Start Get CommentQueryOption(%#v)", opt)
	ctx := context.Background()
	comment, err := uc.Repo.Comment.Get(ctx, id)
	if err != nil{
		logrus.Errorf("Comment(%d)를 찾는 도중 에러가 발생했습니다.", id)
		return nil, err
	}
	// pk로 검색하게 하고 내용만 바꿔놓으면 저장 안 됨.
	// JPA랑 다름.
	updater := uc.Repo.Comment.UpdateOne(comment)

	contentValue, contentKeyExists := opt["content"]
	if contentKeyExists{
		updater.SetContent(contentValue.(string))
	}
	commentUpdated, err:= updater.Save(ctx)

	if err != nil {
		return nil, err
	}

	output := uc.ModelToOutput(username, commentUpdated, nil)
	return output, err
}

// 실제로 Delete 하지는 않고 State를 "deleted"로 변경
func (uc *CommentUseCase) Delete(username string, id int) error {
	ctx := context.Background()
	logrus.Infof("Start Get Comment(id:%#v)", id)
	commentExisting, err := uc.Repo.Comment.Query().
		WithAuthor(func(query *ent.KhumuUserQuery) {
			query.Select("username")
		}).
		Select("id").
		Where(comment.ID(id)).First(ctx)
	// 해당 아이디의 엔티티 존재 X
	if err != nil {
		return  err
	}

	if commentExisting.Edges.Author.ID != username{
		return ErrUnAuthorized
	}

	_, err = uc.Repo.Comment.Update().
		Where(comment.ID(id)).
		SetState("deleted").
		SetContent(DeletedCommentContent).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (uc *CommentUseCase) listParentWithChildren(allComments []*data.CommentOutput) []*data.CommentOutput {
	var parents []*data.CommentOutput

	for _, comment := range allComments {
		if comment.Parent == nil {
			parents = append(parents, comment)
		}
	}
	for _, comment := range allComments {
		if comment.Parent != nil {

		}
	}

	return parents
}

// mapper의 단순 mapping 작업 뿐만 아니라 서비스 로직이 깃든다.
// username: 요청자
// comment: 원본 모델 댓글
// output: 결과물 output 댓글. create if nil
func (uc *CommentUseCase) ModelToOutput(username string, comment *ent.Comment, output *data.CommentOutput) *data.CommentOutput {
	ctx := context.Background()
	if output == nil {
		output = &data.CommentOutput{}
	}

	comment.Edges.Article = comment.QueryArticle().Select("id").OnlyX(ctx)
	comment.Edges.Author = comment.QueryAuthor().Select("username").OnlyX(ctx)

	mapper.CommentModelToOutput(comment, output)

	// hide author
	if comment.State == "deleted" {
		output.Author.Username = DeletedCommentUsername
		output.Author.Nickname = DeletedCommentNickname
	} else if comment.Kind == "anonymous" {
		output.Author.Username = AnonymousCommentUsername
		output.Author.Nickname = AnonymousCommentNickname
	}

	output.CreatedAt = mapper.NewCreatedAtExpression(comment.CreatedAt)
	if comment.Edges.Children != nil 	{
		for _, c := range comment.Edges.Children {
			output.Children = append(output.Children, uc.ModelToOutput(username, c, nil))
		}
	}

	return output
}



func (uc *CommentUseCase) getLikeCommentCount(commentID int) int {
	likes := uc.LikeCommentRepository.List(&repository.LikeCommentQueryOption{CommentID: commentID})
	return len(likes)
}

func NewLikeCommentUseCase(
	likeRepo repository.LikeCommentRepositoryInterface,
	commentRepo repository.CommentRepositoryInterface) LikeCommentUseCaseInterface {
	return &LikeCommentUseCase{Repository: likeRepo, CommentRepository: commentRepo}
}

func NewLikeCommentUseCaseImpl(
	likeRepo repository.LikeCommentRepositoryInterface,
	commentRepo repository.CommentRepositoryInterface) *LikeCommentUseCase {
	return &LikeCommentUseCase{Repository: likeRepo, CommentRepository: commentRepo}
}

func (uc *LikeCommentUseCase) Toggle(like *model.LikeComment) (bool, error) {
	var err error
	logger := logrus.WithField("CommentID", like.CommentID)
	logger.Debug("Toggle LikeComment")
	likes := uc.Repository.List(&repository.LikeCommentQueryOption{CommentID: like.CommentID, Username: like.Username})
	// 길이가 1보다 크거나 같으면 삭제. 1인 경우는 정상적으로 하나만 있을 때,
	// 1보다 큰 경우는 비정상적으로 여러개 존재할 때
	if len(likes) >= 1 {
		for _, like := range likes {
			err = uc.Repository.Delete(like.ID)
			if err != nil {
				logger.Panic(false, err)
			}
		}
		return false, err
	} else {
		// 생성
		comment, err := uc.CommentRepository.Get(like.CommentID)
		if err != nil {
			return false, err
		}
		if comment.AuthorUsername == like.Username {
			return false, errors.New("Error: " + like.Username + " requested to like his comment.")
		}
		_, err = uc.Repository.Create(like)
		if err != nil {
			return false, err
		}

		return true, err
	}
}
