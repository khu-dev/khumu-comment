package usecase

import (
	"context"
	"github.com/khu-dev/khumu-comment/data"
	"github.com/khu-dev/khumu-comment/data/mapper"
	"github.com/khu-dev/khumu-comment/ent"
	"github.com/khu-dev/khumu-comment/ent/comment"
	"github.com/khu-dev/khumu-comment/ent/likecomment"
	"github.com/khu-dev/khumu-comment/errorz"
	"github.com/khu-dev/khumu-comment/external"
	"github.com/khu-dev/khumu-comment/external/khumu"
	"github.com/khu-dev/khumu-comment/repository"
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
	Repo            repository.CommentRepository
	entclient       *ent.Client
	SnsClient       external.SnsClient
	khumuAPIAdapter khumu.KhumuAPIAdapter
	redisAdapter    external.RedisAdapter
}

type LikeCommentUseCase struct {
	Repo         repository.LikeCommentRepository
	CommentRepo  repository.CommentRepository
	redisAdapter external.RedisAdapter
}

func NewCommentUseCase(
	repo repository.CommentRepository,
	entclient *ent.Client,
	snsClient external.SnsClient,
	khumuAPIAdapter khumu.KhumuAPIAdapter,
	redisAdapter external.RedisAdapter) CommentUseCaseInterface {
	return &CommentUseCase{
		Repo:            repo,
		entclient:       entclient,
		SnsClient:       snsClient,
		khumuAPIAdapter: khumuAPIAdapter,
		redisAdapter:    redisAdapter,
	}
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
	isWrittenByArticleAuthor := <-uc.khumuAPIAdapter.IsAuthor(*commentInput.Article, commentInput.Author)
	newComment, err := uc.Repo.Create(commentInput, isWrittenByArticleAuthor)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	newComment.IsWrittenByArticleAuthor = isWrittenByArticleAuthor

	// TODO: 이렇게 new comment에 대한 이벤트들은 chan fan in fan out 처럼 구현하는 게 더 Go스럽게 좋을 듯
	// 현재 같은 절차지향 방식보다는.
	// 절차지향방식은 의존성 주입도 많이 받아야함.
	go uc.SnsClient.PublishMessage(uc.modelToOutput(commentInput.Author, newComment, nil))
	// cache invalidate
	go uc.redisAdapter.RefreshCommentsByArticle(*commentInput.Article)

	// SNS에 Publish한 output을 hide하면 hide 된 채 Publish 될 수 있다는 이슈가 있어서
	// 이렇게 두 번 output을 따로 생성한다.
	output := uc.modelToOutput(commentInput.Author, newComment, nil)
	uc.hideFieldOfCommentOutput(username, output)
	return output, nil
}

func (uc *CommentUseCase) List(username string, opt *CommentQueryOption) ([]*data.CommentOutput, error) {
	//start := time.Now()
	//logrus.Warn("1  ", time.Now().Sub(start))
	//start = time.Now()
	logrus.WithField("username", username).Infof("Start List CommentQueryOption(%#v)", opt)
	var (
		parents []*ent.Comment
		err     error
	)

	//logrus.Warn("2  ", time.Now().Sub(start))
	//start = time.Now()

	if opt.AuthorUsername != "" {
		parents, err = uc.Repo.FindAllParentsByAuthorID(opt.AuthorUsername)
	}
	if opt.ArticleID != 0 {
		//parents, err = uc.Repo.FindAllParentsByArticleID(opt.ArticleID)
		parents = uc.redisAdapter.GetCommentsByArticle(opt.ArticleID)
	}
	if opt.StudyArticleID != 0 {
		parents, err = uc.Repo.FindAllParentsByStudyArticleID(opt.StudyArticleID)
	}

	//logrus.Warn("3  ", time.Now().Sub(start))
	//start = time.Now()

	if err != nil {
		logrus.Errorf("comments 쿼리 도중 오류 발생. QueryOption(%+v)", opt)
		return nil, err
	}

	//logrus.Warn("4  ", time.Now().Sub(start))
	//start = time.Now()

	outputs := make([]*data.CommentOutput, 0)
	for _, parent := range parents {
		output := uc.modelToOutput(username, parent, nil)
		//logrus.Warn("modelToOuptut  ", time.Now().Sub(start))
		//start = time.Now()
		uc.hideFieldOfCommentOutput(username, output)
		//logrus.Warn("hidField  ", time.Now().Sub(start))
		//start = time.Now()
		outputs = append(outputs, output)
	}
	//logrus.Warn("5  ", time.Now().Sub(start))
	//start = time.Now()

	return outputs, nil
}

// 지금의 Get은 Children은 가져오지 못함
func (uc *CommentUseCase) Get(username string, id int) (*data.CommentOutput, error) {
	logrus.WithField("username", username).Infof("Start Get Comment(id:%#v)", id)
	com, err := uc.Repo.Get(id)

	if err != nil {
		logrus.Errorf("comment 쿼리 도중 오류 발생.")
		return nil, err
	}

	output := uc.modelToOutput(username, com, nil)
	uc.hideFieldOfCommentOutput(username, output)
	return output, nil
}

func (uc *CommentUseCase) Update(username string, id int, opt map[string]interface{}) (*data.CommentOutput, error) {
	logrus.WithField("username", username).WithField("id", id).Infof("Start Get CommentQueryOption(%#v)", opt)
	_, err := uc.Repo.Get(id)
	if err != nil {
		logrus.Errorf("Comment(%d)를 찾는 도중 에러가 발생했습니다.", id)
		return nil, err
	}

	com, err := uc.Repo.Update(id, opt)
	if err != nil {
		return nil, err
	}
	go uc.redisAdapter.RefreshCommentsByArticle(com.Edges.Article.ID)
	output := uc.modelToOutput(username, com, nil)
	uc.hideFieldOfCommentOutput(username, output)
	return output, err
}

// 실제로 Delete 하지는 않고 State를 "deleted"로 변경
func (uc *CommentUseCase) Delete(username string, id int) error {
	logrus.Infof("Start Get Comment(id:%#v)", id)
	commentExisting, err := uc.Repo.Get(id)
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
		err = uc.Repo.Delete(id)
		if err != nil {
			return err
		}
		go uc.redisAdapter.RefreshCommentsByArticle(commentExisting.Edges.Article.ID)
	} else {
		updateInput := map[string]interface{}{
			"state": "deleted",
		}
		_, err = uc.Repo.Update(id, updateInput)
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

	if comment.Edges.Parent != nil {
		output.Parent = &comment.Edges.Parent.ID
	}

	if username == comment.Edges.Author.ID {
		output.IsAuthor = true
	}

	//start := time.Now()
	output.CreatedAt = mapper.NewCreatedAtExpression(comment.CreatedAt)
	//logrus.Warn("A ", time.Now().Sub(start))
	//start = time.Now()
	likes := uc.redisAdapter.GetLikeCommentsByComment(comment.ID)
	output.LikeCommentCount = len(likes)
	//logrus.Warn("B ", time.Now().Sub(start))
	//start = time.Now()
	output.Liked = likes.GetLiked(username)
	//logrus.Warn("C ", time.Now().Sub(start))
	//start = time.Now()

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
	likes, err := uc.entclient.LikeComment.Query().
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
	likes, err := uc.entclient.LikeComment.Query().
		Where(likecomment.HasAboutWith(comment.ID(commentID))).
		All(ctx)
	if err != nil {
		logrus.Error(err, "그냥 liked를 false로 처리")
		return false
	}
	return len(likes) != 0
}

func NewLikeCommentUseCase(
	repo repository.LikeCommentRepository,
	commentRepo repository.CommentRepository,
	redisAdapter external.RedisAdapter) LikeCommentUseCaseInterface {
	return &LikeCommentUseCase{
		Repo:         repo,
		CommentRepo:  commentRepo,
		redisAdapter: redisAdapter,
	}
}

func (uc *LikeCommentUseCase) Toggle(input *data.LikeCommentInput) (bool, error) {
	var err error
	commentExisting, err := uc.CommentRepo.Get(input.Comment)
	if err != nil {
		logrus.Error(err)
		return false, err
	}
	if commentExisting.Edges.Author.ID == input.User {
		return false, errorz.ErrSelfLikeComment
	}

	hisLikes, err := uc.Repo.FindAllByUserIDAndCommentID(input.User, input.Comment)
	if err != nil {
		logrus.Error(err)
		return false, err
	}

	// 길이가 1보다 크거나 같으면 삭제. 1인 경우는 정상적으로 하나만 있을 때,
	// 1보다 큰 경우는 비정상적으로 여러개 존재할 때
	if len(hisLikes) >= 1 {
		logrus.Infof("Comment(%d)에 대한 좋아요를 삭제합니다.", input.Comment)
		for _, like := range hisLikes {
			err := uc.Repo.Delete(like.ID)
			if err != nil {
				logrus.Error(err)
				return false, err
			}
		}
		uc.redisAdapter.RefreshLikeCommentsByComment(input.Comment)
		// 정상적으로 삭제한 경우
		return false, nil
	} else {
		// 생성
		logrus.Infof("Comment(%d)에 대한 좋아요를 생성.", input.Comment)
		_, err := uc.Repo.Create(input)
		if err != nil {
			logrus.Error(err)
			return false, err
		}
		uc.redisAdapter.RefreshLikeCommentsByComment(input.Comment)
		// 정상적으로 생성한 경우
		return true, nil
	}
}
