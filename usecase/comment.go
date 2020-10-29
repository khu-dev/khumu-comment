package usecase

import (
	"encoding/json"
	"github.com/khu-dev/khumu-comment/repository"
	"github.com/labstack/echo/v4"
	"log"
	"reflect"
)
import "github.com/khu-dev/khumu-comment/model"


type CommentUseCase struct{
	Repository repository.CommentRepository
}

type Comments interface{}

type ParentComment struct{
	*model.Comment
	Children []*ChildComment `gorm:"-" json:"children"`
	Author *model.SimpleKhumuUser `json:"author"`
}

type ChildComment struct{
	*model.Comment
	Author *model.SimpleKhumuUser `json:"author"`
	// DeleteFields string `json:"ParentID,omitempty"`
	// DeleteFields에 zero value를 넘김으로써 json으로 직렬화될 때 뺼 수는 있지만,
	// 바람직하지 않아보임.
}

func NewChildComment(comment *model.Comment) *ChildComment{
	var tmp []byte
	tmp, err := json.Marshal(comment)
	if err!=nil{log.Println(err)}

	var child *ChildComment
	err = json.Unmarshal(tmp, &child)
	if err!=nil{log.Println(tmp)}

	child.Author = &model.SimpleKhumuUser{&comment.AuthorUsername, "ACTIVE"}

	t := reflect.TypeOf(*child)
	for i := 0; i < t.NumField(); i++{
		field := t.Field(i)
		embedTag := field.Tag.Get("embed")
		if embedTag == "zeroValue"{
			reflect.Indirect(reflect.ValueOf(child)).Field(i).Set(reflect.Zero(field.Type))
		}
	}

	return child
}

func NewParentComment(comment *model.Comment) *ParentComment{
	var tmp []byte
	tmp, err := json.Marshal(comment)
	if err!=nil{log.Println(err)}

	var parent *ParentComment
	err = json.Unmarshal(tmp, &parent)
	if err!=nil{log.Println(tmp)}
	parent.Author = &model.SimpleKhumuUser{&comment.AuthorUsername, "ACTIVE"}
	parent.Children = []*ChildComment{}

	t := reflect.TypeOf(*parent)
	for i := 0; i < t.NumField(); i++{
		field := t.Field(i)
		embedTag := field.Tag.Get("embed")
		if embedTag == "zeroValue"{
			reflect.Indirect(reflect.ValueOf(parent)).Field(i).Set(reflect.Zero(field.Type))
		}
	}

	return parent
}


func (uc *CommentUseCase) List(c echo.Context) Comments {
	log.Println("CommentUseCase List")
	comments := uc.Repository.List(&repository.CommentQueryOption{})
	//for _, c := range comments{
	//	model.PrintModel(c)
	//}
	parents := uc.listParentWithChildren(comments)
	return parents
}

func (uc *CommentUseCase) listParentWithChildren(allComments []*model.Comment) []*model.Comment{
	var parents []*model.Comment

	for _, comment := range allComments{
		if comment.ParentID == 0{
			parents = append(parents, comment)
		}

	}

	return parents
}

func (uc *CommentUseCase) Get(c echo.Context) *model.Comment {
	log.Println("CommentUseCase Get")
	comment := uc.Repository.Get(c.Param("id"))
	return comment
}
