// Deprecated. 이제 수작업으로 Gorm을 통한 Repository 계층을 구현하지 않고
// ent가 자동으로 생성하는 코드로 Repository 계층을 이용.


package _repository

//import (
//	"github.com/khu-dev/khumu-comment/model"
//	"gorm.io/gorm"
//)
//
//type UserRepositoryInterface interface {
//	Get(username string) *model.KhumuUser
//	GetUserForAuth(username string) *model.KhumuUserAuth
//}
//
//type UserRepositoryGorm struct {
//	DB *gorm.DB
//}
//
//func NewUserRepositoryGorm(db *gorm.DB) UserRepositoryInterface {
//	return &UserRepositoryGorm{db}
//}
//
//func (r *UserRepositoryGorm) Get(username string) *model.KhumuUser {
//	var user []*model.KhumuUser
//	r.DB.Find(&user, "username", username)
//	return user[0]
//}
//
//func (r *UserRepositoryGorm) GetUserForAuth(username string) *model.KhumuUserAuth {
//	var users []*model.KhumuUserAuth
//	r.DB.Find(&users, "username", username)
//	if len(users) == 0 {
//		return nil
//	}
//	return users[0]
//}
