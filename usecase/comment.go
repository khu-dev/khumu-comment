package usecase

import (
	"github.com/khu-dev/khumu-comment/data"
	"github.com/khu-dev/khumu-comment/data/mapper"
	"github.com/khu-dev/khumu-comment/ent"
	"github.com/khu-dev/khumu-comment/errorz"
	"github.com/khu-dev/khumu-comment/infra/khumu"
	"github.com/khu-dev/khumu-comment/infra/message"
	"github.com/khu-dev/khumu-comment/repository"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	DeletedCommentContent    string = "삭제된 댓글입니다."
	AnonymousCommentUsername string = "익명"
	AnonymousCommentNickname string = "익명"
	DeletedCommentUsername   string = "삭제된 댓글의 작성자"
	DeletedCommentNickname   string = "삭제된 댓글의 작성자"
	DeletedUserUsername      string = "탈퇴한 유저"
	DeletedUserNickname      string = "탈퇴한 유저"
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
	Repo            repository.CommentRepository
	likeRepo        repository.LikeCommentRepository
	snsClient       message.MessagePublisher
	khumuAPIAdapter khumu.KhumuAPIAdapter
}

type LikeCommentUseCase struct {
	Repo        repository.LikeCommentRepository
	CommentRepo repository.CommentRepository
}

func NewCommentUseCase(
	repo repository.CommentRepository,
	likeRepo repository.LikeCommentRepository,
	snsClient message.MessagePublisher,
	khumuAPIAdapter khumu.KhumuAPIAdapter) CommentUseCaseInterface {
	return &CommentUseCase{
		Repo:            repo,
		likeRepo:        likeRepo,
		snsClient:       snsClient,
		khumuAPIAdapter: khumuAPIAdapter,
	}
}

func (uc *CommentUseCase) Create(username string, commentInput *data.CommentInput) (*data.CommentOutput, error) {
	logrus.Infof("Start Create Comment(%+v)", commentInput)
	//articleId := 1
	if commentInput.Author == "" {
		return nil, errors.WithStack(errorz.ErrNoCommentAuthor)
	}

	if commentInput.Article == nil && commentInput.StudyArticle == nil {
		return nil, errors.WithStack(errorz.ErrNoArticleIDInput)
	}

	if *commentInput.Article <= 0 {
		return nil, errors.WithStack(errorz.ErrWrongArticle)
	}

	newComment, err := uc.Repo.Create(commentInput)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// TODO: 이렇게 new comment에 대한 이벤트들은 chan fan in fan out 처럼 구현하는 게 더 Go스럽게 좋을 듯
	// 현재 같은 절차지향 방식보다는.
	// 절차지향방식은 의존성 주입도 많이 받아야함.
	go uc.snsClient.Publish(uc.modelToOutput(commentInput.Author, newComment, nil))
	// cache invalidate

	// SNS에 Publish한 output을 hide하면 hide 된 채 Publish 될 수 있다는 이슈가 있어서
	// 이렇게 두 번 output을 따로 생성한다.
	output := uc.modelToOutput(commentInput.Author, newComment, nil)
	uc.hideFieldOfCommentOutput(username, output)
	return output, nil
}

func (uc *CommentUseCase) List(username string, opt *CommentQueryOption) ([]*data.CommentOutput, error) {
	logrus.WithField("username", username).Infof("Start List CommentQueryOption(%#v)", opt)
	var (
		parents []*ent.Comment
		err     error
	)

	if opt.AuthorUsername != "" {
		parents, err = uc.Repo.FindAllParentCommentsByAuthorID(opt.AuthorUsername)
	}
	if opt.ArticleID != 0 {
		parents, err = uc.Repo.FindAllParentCommentsByArticleID(opt.ArticleID)
	}
	//if opt.StudyArticleID != 0 {
	//	parents, err = uc.Repo.FindAllParentCommentsByStudyArticleID(opt.StudyArticleID)
	//}

	if err != nil {
		e := errors.Wrap(err, "댓글 조회에 실패했습니다.")
		logrus.Errorf("%+v", e)
		return nil, e
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
	com, err := uc.Repo.Get(id)

	if err != nil {
		e := errors.Wrap(err, "댓글 조회에 실패했습니다.")
		logrus.Errorf("%+v", e)
		return nil, e
	}

	output := uc.modelToOutput(username, com, nil)
	uc.hideFieldOfCommentOutput(username, output)
	return output, nil
}

func (uc *CommentUseCase) Update(username string, id int, opt map[string]interface{}) (*data.CommentOutput, error) {
	logrus.WithField("username", username).WithField("id", id).Infof("Start Get CommentQueryOption(%#v)", opt)
	_, err := uc.Repo.Get(id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.Wrap(errorz.ErrResourceNotFound, "존재하지 않는 댓글입니다.")
		}

		return nil, errors.WithStack(err)
	}

	com, err := uc.Repo.Update(id, opt)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	output := uc.modelToOutput(username, com, nil)
	uc.hideFieldOfCommentOutput(username, output)
	return output, nil
}

// 실제로 Delete 하지는 않고 State를 "deleted"로 변경
func (uc *CommentUseCase) Delete(username string, id int) error {
	logrus.Infof("Start Get Comment(id:%#v)", id)
	commentExisting, err := uc.Repo.Get(id)
	// 해당 아이디의 엔티티 존재 X
	if err != nil {
		if ent.IsNotFound(err) {
			return errors.Wrap(errorz.ErrResourceNotFound, "존재하지 않는 댓글입니다.")
		}

		return errors.WithStack(err)
	}

	if commentExisting.Edges.Author.ID != username {
		return errors.WithStack(errorz.ErrUnauthorized)
	}

	// 대댓글이 없는 댓글 => 삭제 가능
	if len(commentExisting.Edges.Children) == 0 {
		err = uc.Repo.Delete(id)
		if err != nil {
			return errors.WithStack(err)
		}
	} else {
		updateInput := map[string]interface{}{
			"state": "deleted",
		}
		_, err = uc.Repo.Update(id, updateInput)
		if err != nil {
			return errors.WithStack(err)
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

	if comment.Edges.Parent != nil {
		output.Parent = &comment.Edges.Parent.ID
	}

	if username == comment.Edges.Author.ID {
		output.IsAuthor = true
	}

	output.CreatedAt = mapper.NewCreatedAtExpression(comment.CreatedAt)
	likes, err := uc.likeRepo.FindAllByCommentID(comment.ID)
	if err != nil {
		logrus.Error(err)
	} else {
		output.LikeCommentCount = len(likes)
		output.Liked = data.LikeCommentEntities(likes).GetLiked(username)
	}

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
	} else if output.Author.Status == "deleted" {
		output.Author.Username = DeletedUserUsername
		output.Author.Nickname = DeletedUserNickname
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

func NewLikeCommentUseCase(
	repo repository.LikeCommentRepository,
	commentRepo repository.CommentRepository) LikeCommentUseCaseInterface {
	return &LikeCommentUseCase{
		Repo:        repo,
		CommentRepo: commentRepo,
	}
}

func (uc *LikeCommentUseCase) Toggle(input *data.LikeCommentInput) (bool, error) {
	var err error
	commentExisting, err := uc.CommentRepo.Get(input.Comment)
	if err != nil {
		if ent.IsNotFound(err) {
			return false, errors.Wrap(errorz.ErrResourceNotFound, "존재하지 않는 댓글에 대한 좋아요 작업입니다.")
		}

		return false, errors.WithStack(err)
	}
	if commentExisting.Edges.Author.ID == input.User {
		return false, errorz.ErrSelfLikeComment
	}

	hisLikes, err := uc.Repo.FindAllByUserIDAndCommentID(input.User, input.Comment)
	if err != nil {
		return false, errors.WithStack(err)
	}

	// 길이가 1보다 크거나 같으면 삭제. 1인 경우는 정상적으로 하나만 있을 때,
	// 1보다 큰 경우는 비정상적으로 여러개 존재할 때
	if len(hisLikes) >= 1 {
		logrus.Infof("Comment(%d)에 대한 좋아요를 삭제합니다.", input.Comment)
		for _, like := range hisLikes {
			err := uc.Repo.Delete(like.ID)
			if err != nil {
				return false, errors.WithStack(err)
			}
		}
		// 정상적으로 삭제한 경우
		return false, nil
	} else {
		// 생성
		logrus.Infof("Comment(%d)에 대한 좋아요를 생성.", input.Comment)
		_, err := uc.Repo.Create(input)
		if err != nil {
			logrus.Error(err)
			return false, errors.WithStack(err)
		}
		// 정상적으로 생성한 경우
		return true, nil
	}
}
