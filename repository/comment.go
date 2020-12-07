package repository

import (
	"errors"
	"github.com/khu-dev/khumu-comment/model"
	"gorm.io/gorm"
)

type CommentRepositoryInterface interface {
	Create(comment *model.Comment) (*model.Comment, error)
	List(opt *CommentQueryOption) []*model.Comment
	Get(id int) (*model.Comment, error)
	Update(id int, opt map[string]interface{}) (*model.Comment, error)
	Delete(id int) (*model.Comment, error)
}

type LikeCommentRepositoryInterface interface{
	Create(like *model.LikeComment) (*model.LikeComment, error)
	List(opt *LikeCommentQueryOption) ([]*model.LikeComment)
	Delete(id int) error
}

type CommentRepositoryGorm struct {
	DB *gorm.DB
}

type CommentQueryOption struct {
	ArticleID int
	AuthorUsername  string
	CommentID int
}

type LikeCommentRepositoryGorm struct{
	DB *gorm.DB
}
type LikeCommentQueryOption struct{
	CommentID int
	Username string
}

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
	// 리턴할 땐 다시 그 정보 복사
	if comment.Author.Username == "" && comment.AuthorUsername != ""{
		comment.Author.Username = comment.AuthorUsername
	} else if comment.Author.Username != "" && comment.AuthorUsername == ""{
		comment.AuthorUsername = comment.Author.Username
	} else if comment.Author.Username == "" && comment.AuthorUsername == ""{
		return nil, errors.New("Please input author username")
	}
	// 기본 값
	comment.Author.State = "active"

	tmpStoreUser := comment.Author
	comment.Author = nil
	err := r.DB.Create(comment).Error
	if err != nil{return nil, err}

	comment.Author = tmpStoreUser
	return comment, err
}


// List 할 때에는 Author{}가 아닌 AuthorUsername을 이용해 List 한다.
func (r *CommentRepositoryGorm) List(opt *CommentQueryOption) []*model.Comment {
	conditions := make(map[string]interface{})
	// option 인자 정리
	if opt.ArticleID != 0 {
		conditions["article_id"] = opt.ArticleID
	}
	if opt.AuthorUsername != "" {
		conditions["author_id"] = opt.AuthorUsername
	}

	var comments []*model.Comment
	preloaded := r.DB.Preload("Author"). // Author가 사용하는 foreignKey를 이용해 Preload
		Preload("Children.Author"). // nested preload
		Preload("Children.Children")
	if len(conditions) == 0 {
		preloaded.Find(&comments)
	} else {
		preloaded.Find(&comments, conditions)
	}

	for _, c := range comments{
		// copy 작업을 해주지 않으면 같은 author는 같은 주소값을 참조하게됨.
		tmpAuthor := *(c.Author)
		c.Author = &tmpAuthor
	}
	return comments
}

func (r *CommentRepositoryGorm) Get(id int) (*model.Comment, error) {
	var tmp *model.Comment = &model.Comment{}
	err := r.DB.Preload("Author").First(tmp, id).Error
	if err != nil{
		return nil, err
	}
	return tmp, nil
}

func (r *CommentRepositoryGorm) Update(id int, opt map[string]interface{}) (*model.Comment, error) {
	var tmp *model.Comment = &model.Comment{}
	err := r.DB.Preload("Author").First(tmp, id).Error
	if err != nil{
		return nil, err
	}

	// update 된 내용은 tmp에 저장됨.
	err = r.DB.Model(tmp).Updates(opt).Error
	if err != nil{
		return nil, err
	}

	return tmp, nil
}

func (r *CommentRepositoryGorm) Delete(id int) (*model.Comment, error) {
	var tmp *model.Comment = &model.Comment{}
	err := r.DB.First(tmp, id).Error
	if err != nil{
		return nil, gorm.ErrRecordNotFound
	}

	err = r.DB.Delete(tmp, id).Error

	if tmp.ID == 0{
		return nil, gorm.ErrRecordNotFound
	}

	return tmp, nil
}

func (r *LikeCommentRepositoryGorm) Create(like *model.LikeComment) (*model.LikeComment, error) {
	like.Comment = nil
	like.User = nil
	err := r.DB.Save(like).Error
	return like, err
}

func (r *LikeCommentRepositoryGorm) List(opt *LikeCommentQueryOption) []*model.LikeComment {
	var conditions map[string]interface{} = map[string]interface{}{}
	var likes []*model.LikeComment
	if opt.CommentID != 0{
		conditions["comment_id"] = opt.CommentID
	}
	if opt.Username != ""{
		conditions["user_id"] = opt.Username
	}

	if len(conditions) == 0{
		r.DB.Find(&likes)
	} else{
		r.DB.Find(&likes, conditions)
	}

	return likes
}

func (r *LikeCommentRepositoryGorm) Delete(id int) error {
	err := r.DB.Delete(&model.LikeComment{}, id).Error
	return err
}