// Deprecated. 이제 수작업으로 Gorm을 통한 Repository 계층을 구현하지 않고
// ent가 자동으로 생성하는 코드로 Repository 계층을 이용.

package _repository
//
//import (
//	"github.com/khu-dev/khumu-comment/model"
//	"github.com/sirupsen/logrus"
//	"gorm.io/gorm"
//)
//
//type CommentRepositoryInterface interface {
//	Create(comment *model.Comment) (*model.Comment, error)
//	List(opt *CommentQueryOption) []*model.Comment
//	Get(id int) (*model.Comment, error)
//	Update(id int, opt map[string]interface{}) (*model.Comment, error)
//	Delete(id int) (*model.Comment, error)
//}
//
//type LikeCommentRepositoryInterface interface {
//	Create(like *model.LikeComment) (*model.LikeComment, error)
//	List(opt *LikeCommentQueryOption) []*model.LikeComment
//	Delete(id int) error
//}
//
//type CommentRepositoryGorm struct {
//	DB *gorm.DB
//}
//
//type CommentQueryOption struct {
//	ArticleID      int
//	AuthorUsername string
//	CommentID      int
//}
//
//type LikeCommentRepositoryGorm struct {
//	DB *gorm.DB
//}
//
//type LikeCommentQueryOption struct {
//	CommentID int
//	Username  string
//}
//
//func NewCommentRepositoryGorm(db *gorm.DB) CommentRepositoryInterface {
//	return &CommentRepositoryGorm{DB: db}
//}
//
//func NewLikeCommentRepositoryGorm(db *gorm.DB) LikeCommentRepositoryInterface {
//	return &LikeCommentRepositoryGorm{DB: db}
//}
//
//// 기본적으로는 AuthorUsername만을 이용해서 쿼리하되 Author 필드를 채워주긴 할 것임.
//func (r *CommentRepositoryGorm) Create(comment *model.Comment) (*model.Comment, error) {
//	// Omit을 통해 해당 field들이 upsert 되지 않고, AuthorUsername, ParentID등만을 이용하도록 함.
//	err := r.DB.Omit("Author", "Parent", "Children").Create(comment).Error
//	if err != nil {
//		logrus.Error(err)
//		return nil, err
//	}
//	// Omit했던 녀석들에 대한 Join
//	err = r.DB.Preload("Author").Preload("Parent").Preload("Children").Find(comment).Error
//	if err != nil {
//		logrus.Error(err)
//		return nil, err
//	}
//
//	return comment, err
//}
//
//// List 할 때에는 Author{}가 아닌 AuthorUsername을 이용해 List 한다.
//func (r *CommentRepositoryGorm) List(opt *CommentQueryOption) []*model.Comment {
//	if opt == nil {
//		opt = &CommentQueryOption{}
//	}
//	conditions := make(map[string]interface{})
//	// option 인자 정리
//	if opt.ArticleID != 0 {
//		conditions["article_id"] = opt.ArticleID
//	}
//	if opt.AuthorUsername != "" {
//		conditions["author_id"] = opt.AuthorUsername
//	}
//	if opt.CommentID != 0 {
//		conditions["id"] = opt.CommentID
//	}
//
//	var comments []*model.Comment
//	preloaded := r.DB.Preload("Author"). // Author가 사용하는 foreignKey를 이용해 Preload
//						Preload("Children").
//						Preload("Children.Author"). // nested preload
//						Preload("Children.Children")
//	if len(conditions) == 0 {
//		preloaded.Find(&comments)
//	} else {
//		preloaded.Find(&comments, conditions)
//	}
//
//	for _, c := range comments {
//		// copy 작업을 해주지 않으면 같은 author는 같은 주소값을 참조하게됨.
//		if c.Author == nil {
//			logrus.Warn("Author가 nil입니다. 테스트 환경에서 SQLite3를 쓰는 경우 외엔 오류입니다.")
//		} else {
//			tmpAuthor := *(c.Author)
//			c.Author = &tmpAuthor
//		}
//
//		if len(c.Children) == 0 {
//			c.Children = make([]*model.Comment, 0)
//		} // slice를 초기화해주지 않으면 empty의 경우 null이 되어버림.
//		for _, child := range c.Children {
//			// Author를 복제하지 않으면 같은 Author를 갖는 애들이 모두 같은 Author 자체를 참조하게된다.
//			if c.Author == nil {
//				logrus.Warn("Child 댓글의 Author가 nil입니다. 테스트 환경에서 SQLite3를 쓰는 경우 외엔 오류입니다.")
//			} else {
//				tmpAuthor := *(child.Author)
//				child.Author = &tmpAuthor
//			}
//
//			// 현재 비즈니스 로직상에선 max depth가 1이므로 child의 Children은 존재하지 않는다.
//			child.Children = make([]*model.Comment, 0)
//		}
//	}
//
//	return comments
//}
//
//func (r *CommentRepositoryGorm) Get(id int) (*model.Comment, error) {
//	var c *model.Comment = &model.Comment{}
//	err := r.DB.Preload("Author").
//		Preload("Children").
//		Preload("Children.Author").
//		First(c, id).Error
//
//	if err != nil {
//		return nil, err
//	}
//	for _, child := range c.Children {
//		// 현재 비즈니스 로직상에선 max depth가 1이므로 child의 Children은 존재하지 않는다.
//		child.Children = make([]*model.Comment, 0)
//	}
//	// List의 경우 Children이 empty의 경우 empty slice로 할당해줘야 null이 아닌 []가 되는데
//	// Get은 왜 안 해줘도 []이 되는지는 모르겠음.
//	// c.Children = make([]*model.Comment, 0)
//	return c, nil
//}
//
//func (r *CommentRepositoryGorm) Update(id int, opt map[string]interface{}) (*model.Comment, error) {
//	var tmp *model.Comment = &model.Comment{}
//	err := r.DB.Preload("Author").First(tmp, id).Error
//	if err != nil {
//		return nil, err
//	}
//
//	// update 된 내용은 tmp에 저장됨.
//	err = r.DB.Model(tmp).Omit("Author").Updates(opt).Error
//	if err != nil {
//		return nil, err
//	}
//
//	return tmp, nil
//}
//
//// repository 레벨에서 까지 실제로 Delete 할 일은 아직 없다.
//func (r *CommentRepositoryGorm) Delete(id int) (*model.Comment, error) {
//	var tmp *model.Comment = &model.Comment{}
//	err := r.DB.First(tmp, id).Error
//	if err != nil {
//		return nil, gorm.ErrRecordNotFound
//	}
//
//	err = r.DB.Delete(tmp, id).Error
//
//	if tmp.ID == 0 {
//		return nil, gorm.ErrRecordNotFound
//	}
//
//	return tmp, nil
//}
//
//func (r *LikeCommentRepositoryGorm) Create(like *model.LikeComment) (*model.LikeComment, error) {
//	like.Comment = nil
//	like.User = nil
//	//like.User = &model.KhumuUserSimple{Username: like.Username}
//	err := r.DB.Omit("User", "Comment").Save(like).Error
//	return like, err
//}
//
//func (r *LikeCommentRepositoryGorm) List(opt *LikeCommentQueryOption) []*model.LikeComment {
//	var conditions map[string]interface{} = map[string]interface{}{}
//	var likes []*model.LikeComment
//	if opt == nil {
//		opt = &LikeCommentQueryOption{}
//	}
//	if opt.CommentID != 0 {
//		conditions["comment_id"] = opt.CommentID
//	}
//	if opt.Username != "" {
//		conditions["user_id"] = opt.Username
//	}
//
//	if len(conditions) == 0 {
//		r.DB.Find(&likes)
//	} else {
//		r.DB.Find(&likes, conditions)
//	}
//
//	return likes
//}
//
//func (r *LikeCommentRepositoryGorm) Delete(id int) error {
//	err := r.DB.Delete(&model.LikeComment{}, id).Error
//	return err
//}
