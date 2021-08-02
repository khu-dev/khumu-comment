package usecase

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/khu-dev/khumu-comment/data"
	"github.com/khu-dev/khumu-comment/data/mapper"
	"github.com/khu-dev/khumu-comment/ent"
	"github.com/khu-dev/khumu-comment/ent/article"
	"github.com/khu-dev/khumu-comment/ent/comment"
	"github.com/khu-dev/khumu-comment/ent/khumuuser"
	"github.com/khu-dev/khumu-comment/ent/likecomment"
	"github.com/khu-dev/khumu-comment/ent/studyarticle"
	"github.com/khu-dev/khumu-comment/errorz"
	"github.com/khu-dev/khumu-comment/external"
	"github.com/sirupsen/logrus"
	"reflect"
)

var (
	DeletedCommentContent    string = "삭제된 댓글입니다."
	AnonymousCommentUsername string = "익명"
	AnonymousCommentNickname string = "익명"
	DeletedCommentUsername   string = "삭제된 댓글의 작성자"
	DeletedCommentNickname   string = "삭제된 댓글의 작성자"
)

type CommentQueryOption struct {
	AuthorUsername string
	ArticleID      int
	StudyArticleID int
	CommentId      int
	PostKind       *string
}

type CommentUseCaseInterface interface {
	Create(username string, commentInput *data.CommentInput) (*data.CommentOutput, error)
	List(username string, opt *CommentQueryOption) ([]*data.CommentOutput, error)
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

func (uc *CommentUseCase) Create(username string, commentInput *data.CommentInput) (*data.CommentOutput, error) {
	logrus.Infof("Start Create Comment(%#v)", commentInput)
	//articleId := 1
	if commentInput.Author == "" {
		logrus.Error("댓글 생성에 대한 author가 존재하지 않습니다.")
		return nil, errorz.ErrUnauthorized
	}

	if commentInput.Article == nil && commentInput.StudyArticle == nil {
		logrus.Error("커뮤니티 게시글 ID나 스터디 게시글 ID가 입력되지 않았습니다.")
		return nil, errorz.ErrNoArticleIDInput
	}

	newComment, err := uc.Repo.Comment.Create().
		SetNillableArticleID(commentInput.Article).
		SetNillableStudyArticleID(commentInput.StudyArticle).
		SetNillableParentID(commentInput.Parent).
		SetAuthorID(commentInput.Author).
		SetContent(commentInput.Content).
		SetState("exists").
		SetNillableKind(commentInput.Kind).
		Save(context.Background())

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	newComment, err = uc.Repo.Comment.Query().
		// 댓글 작성자
		WithAuthor().
		WithArticle(func(query *ent.ArticleQuery) {
			query.WithAuthor()
		}).
		WithStudyArticle(func(query *ent.StudyArticleQuery) {
			query.WithAuthor()
		}).
		// 대댓글. 근데 어차피 새 댓글에는 대댓글이 없긴 하다.
		WithChildren(
			func(query *ent.CommentQuery) {
				query.
					// 대댓글의 작성자
					WithAuthor().
					// 대댓글의 게시글
					WithArticle(func(query *ent.ArticleQuery) {
						query.WithAuthor()
					}).
					// 대댓글의 게시글
					WithStudyArticle(func(query *ent.StudyArticleQuery) {
						query.WithAuthor()
					})
			},
		).
		Where(comment.ID(newComment.ID)).
		Only(context.TODO())
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	err = uc.SnsClient.PublishMessage(uc.modelToOutput(commentInput.Author, newComment, nil))
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	// SNS에 Publish한 output을 hide하면 hide 된 채 Publish 될 수 있다는 이슈가 있어서
	// 이렇게 두 번 output을 따로 생성한다.
	output := uc.modelToOutput(commentInput.Author, newComment, nil)
	uc.hideFieldOfCommentOutput(username, output)
	return output, nil
}

func (uc *CommentUseCase) List(username string, opt *CommentQueryOption) ([]*data.CommentOutput, error) {
	logrus.WithField("username", username).Infof("Start List CommentQueryOption(%#v)", opt)
	query := uc.Repo.Comment.Query()
	if opt.AuthorUsername != "" {
		query.Where(comment.HasAuthorWith(khumuuser.ID(opt.AuthorUsername)))
	}
	if opt.ArticleID != 0 {
		query.Where(comment.HasArticleWith(article.ID(opt.ArticleID)))
	}
	if opt.StudyArticleID != 0 {
		query.Where(comment.HasStudyArticleWith(studyarticle.ID(opt.StudyArticleID)))
	}

	parents, err := appendQueryForComment(query).
		Where(comment.Not(comment.HasParent())).
		All(context.TODO())

	if err != nil {
		logrus.Errorf("comments 쿼리 도중 오류 발생. QueryOption(%+v)", opt)
		return nil, err
	}

	outputs := make([]*data.CommentOutput, 0)
	for _, parent := range parents {
		output := uc.modelToOutput(username, parent, nil)
		uc.hideFieldOfCommentOutput(username, output)
		outputs = append(outputs, output)
	}

	return outputs, nil
}

// 지금의 Get은 Children은 가져오지 못함
func (uc *CommentUseCase) Get(username string, id int) (*data.CommentOutput, error) {
	logrus.WithField("username", username).Infof("Start Get Comment(id:%#v)", id)
	ctx := context.Background()

	comment, err := appendQueryForComment(uc.Repo.Comment.Query()).
		Where(comment.ID(id)).
		Only(ctx)

	if err != nil {
		logrus.Errorf("comment 쿼리 도중 오류 발생.")
		return nil, err
	}

	output := uc.modelToOutput(username, comment, nil)
	uc.hideFieldOfCommentOutput(username, output)
	return output, nil
}

func (uc *CommentUseCase) Update(username string, id int, opt map[string]interface{}) (*data.CommentOutput, error) {
	logrus.WithField("username", username).WithField("id", id).Infof("Start Get CommentQueryOption(%#v)", opt)
	ctx := context.Background()
	commentExisting, err := uc.Repo.Comment.Get(ctx, id)
	if err != nil {
		logrus.Errorf("Comment(%d)를 찾는 도중 에러가 발생했습니다.", id)
		return nil, err
	}
	// pk로 검색하게 하고 내용만 바꿔놓으면 저장 안 됨.
	// JPA랑 다름.
	updater := uc.Repo.Comment.UpdateOne(commentExisting)

	contentValue, contentExists := opt["content"]
	if contentExists {
		updater.SetContent(contentValue.(string))
	}
	kindValue, kindExists := opt["kind"]
	if kindExists {
		updater.SetKind(kindValue.(string))
	}
	_, err = updater.Save(ctx)
	if err != nil {
		return nil, err
	}

	commentUpdated, err := appendQueryForComment(uc.Repo.Comment.Query()).
		Where(comment.ID(id)).
		Only(ctx)

	output := uc.modelToOutput(username, commentUpdated, nil)
	uc.hideFieldOfCommentOutput(username, output)
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
		WithChildren(func(query *ent.CommentQuery) {
			query.Select("id")
		}).
		Select("id").
		Where(comment.ID(id)).First(ctx)
	// 해당 아이디의 엔티티 존재 X
	if err != nil {
		if reflect.TypeOf(err).ConvertibleTo(reflect.TypeOf(&ent.NotFoundError{})) {
			logrus.Error("Here!")
			return errorz.ErrResourceNotFound
		}
		return err
	}

	if commentExisting.Edges.Author.ID != username {
		return errorz.ErrUnauthorized
	}

	// 대댓글이 없는 댓글 => 삭제 가능
	if len(commentExisting.Edges.Children) == 0 {
		logrus.Info("부모 댓글이 없어 댓글 자체를 삭제하는 작업을 시작합니다.")
		tx, err := uc.Repo.BeginTx(ctx, new(sql.TxOptions))
		defer func() {
			if err = tx.Commit(); err != nil {
				logrus.Error(err)
			}
		}()
		if err != nil {
			logrus.Error(err)
			return err
		}

		n, err := tx.LikeComment.Delete().
			Where(likecomment.HasAboutWith(comment.ID(commentExisting.ID))).
			Exec(ctx)
		if err != nil {
			logrus.Error(err)
			return err
		}
		logrus.Info(commentExisting, "에 대한 좋아요를 ", n, " 개 삭제했습니다.")

		err = tx.Comment.DeleteOneID(id).Exec(ctx)
		if err != nil {
			logrus.Error(err)
			return err
		}
		logrus.Infof("Comment(id=%d)를 삭제했습니다.", id)

		return nil
	} else {
		_, err = uc.Repo.Comment.Update().
			Where(comment.ID(id)).
			SetState("deleted").
			SetContent(DeletedCommentContent).
			Save(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

// CommentOutput.Children까지 재귀적으로 CommentOutput으로 만든다.
// mapper의 단순 mapping 작업 뿐만 아니라 서비스 로직이 깃든다.
// username: 요청자
// comment: 원본 모델 댓글
// output: 결과물 output 댓글. create if nil
func (uc *CommentUseCase) modelToOutput(username string, comment *ent.Comment, outputRef *data.CommentOutput) *data.CommentOutput {
	output := outputRef
	if output == nil {
		output = &data.CommentOutput{}
	}

	mapper.CommentModelToOutput(comment, output)
	if comment.Edges.Article != nil {
		if username == comment.Edges.Article.Edges.Author.ID {
			output.IsAuthorOfArticle = true
		}
	} else if comment.Edges.StudyArticle != nil {
		if username == comment.Edges.StudyArticle.Edges.Author.ID {
			output.IsAuthorOfArticle = true
		}
	}
	if comment.Edges.Parent != nil {
		output.Parent = &comment.Edges.Parent.ID
	}

	if username == comment.Edges.Author.ID {
		output.IsAuthor = true
	}

	output.CreatedAt = mapper.NewCreatedAtExpression(comment.CreatedAt)
	output.LikeCommentCount = uc.getLikeCommentCount(comment.ID)
	output.Liked = uc.getLiked(comment.ID)

	if comment.Edges.Children != nil {
		for _, child := range comment.Edges.Children {
			output.Children = append(output.Children, uc.modelToOutput(username, child, nil))
		}
	}

	return output
}

// 재귀적으로 output.Children의 field도 숨긴다.
func (uc *CommentUseCase) hideFieldOfCommentOutput(username string, output *data.CommentOutput) {
	// hide author
	if output.State == "deleted" {
		output.Author.Username = DeletedCommentUsername
		output.Author.Nickname = DeletedCommentNickname
		output.Content = DeletedCommentContent
	} else if output.Kind == "anonymous" {
		output.Author.Username = AnonymousCommentUsername
		output.Author.Nickname = AnonymousCommentNickname
	}

	if output.Children != nil {
		for _, child := range output.Children {
			uc.hideFieldOfCommentOutput(username, child)
		}
	}
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
		return false, errorz.ErrSelfLikeComment
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
