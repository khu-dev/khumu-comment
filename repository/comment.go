package repository

import (
	"github.com/khu-dev/khumu-comment/model"
	"gorm.io/gorm"
)

type CommentRepositoryInterface interface {
	Create(comment *model.Comment) (*model.Comment, error)
	List(opt *CommentQueryOption) []*model.Comment
	Get(id int) *model.Comment
}

type LikeCommentRepositoryInterface interface{
	Create(like *model.LikeComment) (*model.LikeComment, error)
	//List(opt *LikeCommentQueryOption) []*model.LikeComment
}

type CommentRepositoryGorm struct {
	DB *gorm.DB
}

type CommentQueryOption struct {
	ArticleID uint
	AuthorID  string
}

type LikeCommentRepositoryGorm struct{
	DB *gorm.DB
}
type LikeCommentQueryOption struct{}

func NewCommentRepositoryGorm(db *gorm.DB) CommentRepositoryInterface{
	return &CommentRepositoryGorm{DB: db}
}

func NewLikeCommentRepositoryGorm(db *gorm.DB) LikeCommentRepositoryInterface{
	return &LikeCommentRepositoryGorm{DB: db}
}

// Comment를 Create 할 때에는 AuthorUsername field가 비어있고, Author field에
// Author에 대한 정보가 담겨있다. AuthorUsername은 json 통신할 때에도 사용하지 않기때문에
// 입력받을 수 없다. 하지만 리턴할 때에는 상위 계층이 잘 사용할 수 있게끔 해당 값도 입력해서 전달해줄 것이다.
// 따라서 입력받을 땐 AuthorUsername은 비어있고 Author에만 유효한 값이 들어있고, 리턴할 땐 둘 다 유효한 값이 들어있다.
func (r *CommentRepositoryGorm) Create(comment *model.Comment) (*model.Comment, error) {
	// Author field가 남아있으면 그걸 기준으로 Author 필드의 데이터도 업데이트시키려고하기때문에
	// 단순히 foreignKey field만 남긴다.
	// 리턴할 땐 다시 그 정보 복
	comment.AuthorUsername = comment.Author.Username
	tmpStoreUser := comment.Author
	comment.Author = nil
	err := r.DB.Save(comment).Error
	if err != nil{return nil, err}

	comment.Author = tmpStoreUser
	return comment, err
}

func (r *CommentRepositoryGorm) List(opt *CommentQueryOption) []*model.Comment {
	conditions := make(map[string]interface{})
	if opt.ArticleID != 0 {
		conditions["article_id"] = opt.ArticleID
	}
	if opt.AuthorID != "" {
		conditions["author_id"] = opt.AuthorID
	}
	var comments []*model.Comment
	preloaded := r.DB.Preload("Author").
		Preload("Children.Author"). // nested preload
		Preload("Children.Children")
	if len(conditions) == 0 {
		preloaded.Find(&comments)
	} else {
		preloaded.Find(&comments, conditions)
	}
	return comments
}

func (r *CommentRepositoryGorm) Get(id int) *model.Comment {
	var tmp *model.Comment = &model.Comment{}
	r.DB.First(tmp)
	return tmp
}

func (r *LikeCommentRepositoryGorm) Create(like *model.LikeComment) (*model.LikeComment, error) {
	err := r.DB.Save(like).Error
	return like, err
}