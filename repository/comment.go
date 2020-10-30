package repository

import (
	"github.com/khu-dev/khumu-comment/model"
	"gorm.io/gorm"
	"log"
)

type CommentRepositorySQLite3 struct {
	DB *gorm.DB
}

func (r *CommentRepositorySQLite3) List(opt *CommentQueryOption) []*model.Comment{
	log.Println("CommentRepositorySQLite3 List")
	var comments []*model.Comment
	r.DB.Preload("Author").
		Preload("Children.Author"). // nested preload
		Preload("Children.Children").
		Find(&comments)
	return comments
}

func (r *CommentRepositorySQLite3) Get(id string) *model.Comment{
	log.Println("CommentRepositorySQLite3 Get")
	//var comment *model.Comment
	//idInt, _ := strconv.Atoi(id)
	var tmp *model.Comment = &model.Comment{}
	r.DB.First(tmp)
	return tmp
}

/*
func (r *CommentRepositorySQLite3) Get(id string) *model.Comment{
	log.Println("CommentRepositorySQLite3 Get")
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