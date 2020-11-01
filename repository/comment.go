package repository

import (
	"github.com/khu-dev/khumu-comment/model"
	"gorm.io/gorm"
)

type CommentRepositoryInterface interface {
	Create(comment *model.Comment) error
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

func (r *CommentRepositoryGorm) Create(comment *model.Comment) error {
	err := r.DB.Create(comment).Error

	return err
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

/*
func (r *CommentRepositoryGorm) Get(id string) *model.Comment{
	log.Println("CommentRepositoryGorm Get")
	//var comment *model.Comment
	//idInt, _ := strconv.Atoi(id)
	var tmp *model.Comment = &model.Comment{}
	r.DB.First(tmp)
	return tmp
}
이슈

var tmp *model.Comment = &model.Comment{}
하면 되는데
var tmp *model.Comment로 하면 안 됨.
이유는 nil pointer 이기때문.

*/
