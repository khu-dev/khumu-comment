package usecase

import (
	"context"
	"errors"
	"github.com/khu-dev/khumu-comment/data"
	"github.com/khu-dev/khumu-comment/data/mapper"
	"github.com/khu-dev/khumu-comment/ent"
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
)

type CommentUseCaseInterface interface {
	Create(commentInput *data.CommentInput) (*data.CommentOutput, error)
	List(username string, opt *repository.CommentQueryOption) ([]*model.Comment, error)
	Get(username string, id int) (*model.Comment, error)
	Update(username string, id int, opt map[string]interface{}) (*model.Comment, error)
	Delete(id int) (*model.Comment, error)
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
	EntRepository *ent.Client
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
	return &CommentUseCase{Repository: repository, LikeCommentRepository: likeRepository, SnsClient: snsClient, EntRepository: repo}
}

func (uc *CommentUseCase) Create(commentInput *data.CommentInput) (*data.CommentOutput, error) {
	logrus.Infof("Start Create Comment(%#v)", commentInput)
	//articleId := 1
	newComment, err := uc.EntRepository.Comment.Create().
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
	newComment.Edges.Article = newComment.QueryArticle().Select("id").OnlyX(context.Background())
	newComment.Edges.Author = newComment.QueryAuthor().Select("username").OnlyX(context.Background())

	return mapper.CommentModelToOutput(newComment, nil), nil
}

func (uc *CommentUseCase) List(username string, opt *repository.CommentQueryOption) ([]*model.Comment, error) {
	logrus.WithField("username", username).Infof("Start List CommentQueryOption(%#v)", opt)
	comments := uc.Repository.List(opt)
	parents := uc.listParentWithChildren(comments)

	for _, p := range parents {
		uc.handleComment(p, username, 0)
	}

	return parents, nil
}

// 지금의 Get은 Children은 가져오지 못함
func (uc *CommentUseCase) Get(username string, id int) (*model.Comment, error) {
	logrus.WithField("username", username).Infof("Start Get Comment(id:%#v)", id)
	comment, err := uc.Repository.Get(id)
	if err != nil {
		return nil, err
	}

	uc.handleComment(comment, username, 0) // ""는 어떠한 author username과도 다르기때문에 숨겨진다.
	return comment, nil
}

func (uc *CommentUseCase) Update(username string, id int, opt map[string]interface{}) (*model.Comment, error) {
	logrus.WithField("username", username).WithField("id", id).Infof("Start Get CommentQueryOption(%#v)", opt)
	updated, err := uc.Repository.Update(id, opt)
	uc.handleComment(updated, username, 0)
	return updated, err
}

// 실제로 Delete 하지는 않고 State를 "deleted"로 변경
func (uc *CommentUseCase) Delete(id int) (*model.Comment, error) {
	logrus.Infof("Start Get Comment(id:%#v)", id)
	comment, err := uc.Repository.Update(id, map[string]interface{}{
		"state":   "deleted",
		"content": DeletedCommentContent,
	})
	if err != nil {
		return nil, err
	}

	uc.handleComment(comment, comment.AuthorUsername, 0)
	return comment, nil
}

func (uc *CommentUseCase) listParentWithChildren(allComments []*model.Comment) []*model.Comment {
	var parents []*model.Comment

	for _, comment := range allComments {
		if comment.ParentID.Valid == false {
			parents = append(parents, comment)
		}
	}

	return parents
}

// 대부분의 comment usecase에서 사용되는 로직을 담당한다. 재귀적으로 자식 코멘트들에게도 적용된다.
func (uc *CommentUseCase) handleComment(c *model.Comment, username string, currentDepth int) {
	//const maxDepth = 1
	//if c.AuthorUsername == username {
	//	c.IsAuthor = true
	//}
	//if c.Kind == "anonymous" || c.State == "deleted" {
	//	uc.hideAuthor(c)
	//}
	//if c.State == "deleted" {
	//	c.Content = DeletedCommentContent
	//}
	//
	//likeCount := uc.getLikeCommentCount(c.ID)
	//c.LikeCommentCount = likeCount
	//likes := uc.LikeCommentRepository.List(&repository.LikeCommentQueryOption{CommentID: c.ID, Username: username})
	//if len(likes) >= 1 {
	//	c.Liked = true
	//}
	//uc.setCreatedAtExpression(c)
	//if currentDepth < maxDepth {
	//	for _, child := range c.Children {
	//		uc.handleComment(child, username, currentDepth+1)
	//	}
	//}
}

// mapper의 단순 mapping 작업 뿐만 아니라 서비스 로직이 깃든다.
func (uc *CommentUseCase) Model2Output(comment *ent.Comment) *data.CommentOutput {
	ctx := context.Background()
	comment.Edges.Article = comment.QueryArticle().Select("id").OnlyX(ctx)
	comment.Edges.Author = comment.QueryAuthor().Select("username").OnlyX(ctx)

	output := mapper.CommentModelToOutput(comment, nil)

	// hide author
	if comment.State == "deleted" {
		output.Author.Username = DeletedCommentUsername
		output.Author.Nickname = DeletedCommentNickname
	} else if comment.Kind == "anonymous" {
		output.Author.Username = AnonymousCommentUsername
		output.Author.Nickname = AnonymousCommentNickname
	}

	output.CreatedAt = mapper.NewCreatedAtExpression(comment.CreatedAt)

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
