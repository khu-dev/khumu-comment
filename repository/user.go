package repository

import (
	"fmt"
	"github.com/khu-dev/khumu-comment/model"
	"gorm.io/gorm"
)

type UserRepository interface{
	Get(username string) *model.KhumuUser
	GetUserForAuth(username string) *model.KhumuUserAuth
}

type UserRepositoryGorm struct{
	DB *gorm.DB
}

func (r *UserRepositoryGorm) Get(username string) *model.KhumuUser{
	var user []*model.KhumuUser
	r.DB.Find(&user, "username",username)
	fmt.Print("Found user ", user)
	return user[0]
}

func (r *UserRepositoryGorm) GetUserForAuth(username string) *model.KhumuUserAuth{
	var users []*model.KhumuUserAuth
	r.DB.Find(&users, "username",username)
	fmt.Print("Found user ", users)
	if len(users) == 0{
		return nil
	}
	return users[0]
}