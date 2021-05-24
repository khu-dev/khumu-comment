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
	"github.com/khu-dev/khumu-comment/ent/likecomment"
	"github.com/khu-dev/khumu-comment/external"
	"github.com/khu-dev/khumu-comment/repository"
	"github.com/sirupsen/logrus"
)

var (
	DeletedCommentContent    string = "삭제된 댓글입니다."
	AnonymousCommentUsername string = "익명"
	AnonymousCommentNickname string = "익명"
	DeletedCommentUsername   string = "삭제된 댓글의 작성자"
	DeletedCommentNickname   string = "삭제된 댓글의 작성자"
	ErrUnAuthorized                 = errors.New("권한이 존재하지 않습니다")
	ErrSelfLikeComment              = errors.New("본인의 댓글은 좋아요할 수 없습니다")
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
	Toggle(like *data.LikeCommentInput) (bool, error)
}

type CommentUseCase struct {
	Repo      *ent.Client
	SnsClient external.SnsClient
}

type LikeCommentUseCase struct {
	Repo *ent.Client
}

func NewCommentUseCase(
	repo *ent.Client,
	snsClient external.SnsClient) CommentUseCaseInterface {
	return &CommentUseCase{Repo: repo, SnsClient: snsClient}
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
	if opt.ArticleId != 0 {
		query.Where(comment.HasArticleWith(article.ID(opt.ArticleId)))
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
	if err != nil {
		logrus.Errorf("Comment(%d)를 찾는 도중 에러가 발생했습니다.", id)
		return nil, err
	}
	// pk로 검색하게 하고 내용만 바꿔놓으면 저장 안 됨.
	// JPA랑 다름.
	updater := uc.Repo.Comment.UpdateOne(comment)

	contentValue, contentExists := opt["content"]
	if contentExists {
		updater.SetContent(contentValue.(string))
	}
	kindValue, kindExists := opt["kind"]
	if kindExists {
		updater.SetKind(kindValue.(string))
	}
	commentUpdated, err := updater.Save(ctx)

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
		return err
	}

	if commentExisting.Edges.Author.ID != username {
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
	if username == comment.Edges.Author.ID {
		output.IsAuthor = true
	}

	// hide author
	if comment.State == "deleted" {
		output.Author.Username = DeletedCommentUsername
		output.Author.Nickname = DeletedCommentNickname
		output.Content = DeletedCommentContent
	} else if comment.Kind == "anonymous" {
		output.Author.Username = AnonymousCommentUsername
		output.Author.Nickname = AnonymousCommentNickname
	}

	output.CreatedAt = mapper.NewCreatedAtExpression(comment.CreatedAt)
	output.LikeCommentCount = uc.getLikeCommentCount(comment.ID)
	if comment.Edges.Children != nil {
		for _, c := range comment.Edges.Children {
			output.Children = append(output.Children, uc.ModelToOutput(username, c, nil))
		}
	}
	output.Liked = uc.getLiked(comment.ID)

	return output
}

func (uc *CommentUseCase) getLikeCommentCount(commentID int) int {
	ctx := context.Background()
	likes, err := uc.Repo.LikeComment.Query().
		Where(likecomment.HasAboutWith(comment.ID(commentID))).
		All(ctx)
	if err != nil {
		logrus.Error(err, "그냥 like comment count를 0으로 처리")
		return 0
	}
	return len(likes)
}

func (uc *CommentUseCase) getLiked(commentID int) bool {
	ctx := context.Background()
	likes, err := uc.Repo.LikeComment.Query().
		Where(likecomment.HasAboutWith(comment.ID(commentID))).
		All(ctx)
	if err != nil {
		logrus.Error(err, "그냥 liked를 false로 처리")
		return false
	}
	return len(likes) != 0
}

func NewLikeCommentUseCase(
	repo *ent.Client) LikeCommentUseCaseInterface {
	return &LikeCommentUseCase{Repo: repo}
}

func (uc *LikeCommentUseCase) Toggle(input *data.LikeCommentInput) (bool, error) {
	var err error
	ctx := context.Background()
	commentExisting, err := uc.Repo.Comment.Query().
		WithAuthor().
		Where(comment.ID(input.Comment)).
		Only(ctx)
	if err != nil {
		logrus.Error(err)
		return false, err
	}
	if commentExisting.Edges.Author.ID == input.User {
		return false, ErrSelfLikeComment
	}

	likeIDs, err := uc.Repo.LikeComment.Query().
		Where(likecomment.HasAboutWith(comment.ID(input.Comment))).
		IDs(ctx)
	if err != nil {
		logrus.Error(err)
		return false, err
	}

	// 길이가 1보다 크거나 같으면 삭제. 1인 경우는 정상적으로 하나만 있을 때,
	// 1보다 큰 경우는 비정상적으로 여러개 존재할 때
	if len(likeIDs) >= 1 {
		logrus.Infof("Comment(%d)에 대한 좋아요를 삭제합니다.", input.Comment)
		_, err := uc.Repo.LikeComment.Delete().Where(likecomment.IDIn(likeIDs...)).Exec(ctx)
		if err != nil {
			logrus.Error(err)
			return false, err
		}

		return false, nil
	} else {
		// 생성
		logrus.Infof("Comment(%d)에 대한 좋아요를 생성.", input.Comment)
		_, err := uc.Repo.LikeComment.Create().SetAboutID(input.Comment).SetLikedByID(input.User).Save(ctx)
		if err != nil {
			logrus.Error(err)
			return false, err
		}

		return true, nil
	}
}
