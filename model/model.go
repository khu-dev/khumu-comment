package model

import (
	"encoding/json"
	"fmt"
	"time"
)

type ModelPrettyPrinter interface{
	prettyPrint()
}

func PrintModel(p ModelPrettyPrinter){
	p.prettyPrint()
}

type ModelStringer interface{
	toString() string
}

func String(s ModelStringer) string{
	return s.toString()
}

type KhumuUser struct {
	//gorm.Model
	Username string `gorm:"primaryKey"`
	Email string
	IsActive bool
	Type string
	Nickanme string
	StudentNumber string
	Memo string
	CreatedAt time.Time
}

func (m *KhumuUser) prettyPrint() {
	s, _ := json.MarshalIndent(m, "", "    ")
    fmt.Print(string(s))
}

func (*KhumuUser) TableName() string {
    return "user_khumuuser"
}

type Article struct {
	ArticleID uint `gorm:"column:id"`
	BoardID uint
	Board Board `gorm:"foreignKey:BoardID"`
	Title string
	AuthorUsername string
	Author KhumuUser `gorm:"foreignKey:AuthorUsername"`
	Content string
	CreatedAt time.Time
}

func (*Article) TableName() string {
    return "article_article"
}

type Board struct {
	BoardID uint `gorm:"column:id"`
	ShortName string
	LongName string
	Name string
	Description string
	AdminUsername string
	Admin KhumuUser `gorm:"foreignKey:AdminUsername"`
}

func (*Board) TableName() string{
	return "board_board"
}

type Comment struct {
	ID uint `gorm:"column:id"`
	AuthorID string `gorm:"column:author_id"`
	ArticleID uint `gorm:"column:article_id"`
	//Article Article `gorm:"foreignKey:ArticleID"`
	Content string
	Type string `gorm:"-"`
	ParentID uint `gorm:"column:parent_id"`
	//Parent *Comment `gorm:"foreignKey:ParentID"`
	CreatedAt time.Time
}

func (*Comment) TableName() string{
	return "comment_comment"
}

func (m *Comment) prettyPrint() {
	s, _ := json.MarshalIndent(m, "", "    ")
    fmt.Print(string(s))
}

func (m *Comment) toString() string{
	s, _ := json.MarshalIndent(m, "", "    ")
    return string(s)
}