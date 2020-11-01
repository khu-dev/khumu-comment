package model

import (
	"encoding/json"
	"fmt"
	"time"
)

type ModelPrettyPrinter interface {
	prettyPrint()
}

func PrintModel(p ModelPrettyPrinter) {
	p.prettyPrint()
}

type ModelStringer interface {
	toString() string
}

func String(s ModelStringer) string {
	return s.toString()
}

type KhumuUser struct {
	//gorm.Model
	Username      string `gorm:"primaryKey"`
	Email         string
	IsActive      bool
	Type          string
	Nickanme      string
	StudentNumber string
	Memo          string
	CreatedAt     time.Time
}

func (*KhumuUser) TableName() string {
	return "user_khumuuser"
}

func (m *KhumuUser) prettyPrint() {
	s, _ := json.MarshalIndent(m, "", "    ")
	fmt.Print(string(s))
}

type KhumuUserAuth struct {
	Username string `gorm:"primaryKey"`
	Password string `gorm:"password"`
}

func (*KhumuUserAuth) TableName() string {
	return "user_khumuuser"
}

type KhumuUserSimple struct {
	Username string `gorm:"primaryKey" json:"username"`
	Type     string `gorm:"-" json:"type"`
}

func (*KhumuUserSimple) TableName() string {
	return "user_khumuuser"
}

type Article struct {
	ArticleID      uint `gorm:"column:id"`
	BoardID        uint
	Board          Board `gorm:"foreignKey:BoardID"`
	Title          string
	AuthorUsername string    `gorm:"column:author_id"`
	Author         KhumuUser `gorm:"foreignKey:AuthorUsername"`
	Content        string
	CreatedAt      time.Time
}

func (*Article) TableName() string {
	return "article_article"
}

type Board struct {
	BoardID       uint `gorm:"column:id"`
	ShortName     string
	LongName      string
	Name          string
	Description   string
	AdminUsername string
	Admin         KhumuUser `gorm:"foreignKey:AdminUsername"`
}

func (*Board) TableName() string {
	return "board_board"
}

type Comment struct {
	ID             uint             `gorm:"column:id" json:"id"`
	Kind           string           `gorm:"column:kind" json:"kind"`
	Author         *KhumuUserSimple `gorm:"foreignKey:AuthorUsername;references:Username" json:"author"`
	AuthorUsername string           `gorm:"column:author_id" json:"-"`
	ArticleID      uint             `gorm:"column:article_id" json:"article"`
	//Article Article `gorm:"foreignKey:ArticleID"`
	Content   string     `json:"content"`
	ParentID  *uint      `gorm:"column:parent_id;default:null" json:"parent"`
	Children  []*Comment `gorm:"foreignKey:ParentID;references:ID" json:"children"` //Has-Many relationship => Preload 필요
	CreatedAt time.Time  `json:"created_at"`
}

func (*Comment) TableName() string {
	return "comment_comment"
}

func (m *Comment) prettyPrint() {
	s, _ := json.MarshalIndent(m, "", "    ")
	fmt.Print(string(s))
}

func (m *Comment) toString() string {
	s, _ := json.MarshalIndent(m, "", "    ")
	return string(s)
}

type LikeComment struct{
	ID int `gorm:"primaryKey"`
	CommentID int `gorm:"column:comment_id" json:"comment"`
	Username string `gorm:"column:user_id" json:"username"`
}

func (*LikeComment) TableName() string {
	return "comment_likecomment"
}
