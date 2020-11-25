package model

import (
	"time"
)

type ModelPrettyPrinter interface {
	prettyPrint()
}

type KhumuUser struct {
	//gorm.Model
	Username      string `gorm:"primaryKey"`
	Email         string
	IsActive      bool
	Kind          string
	Nickname      string `gorm:"unique"`
	StudentNumber string
	Memo          string
	CreatedAt     time.Time
}

func (*KhumuUser) TableName() string {
	return "user_khumuuser"
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
	Nickname string `gorm:"unique json:"nickname`
	State     string `gorm:"column:kind" json:"state"`
}

func (*KhumuUserSimple) TableName() string {
	return "user_khumuuser"
}

type Article struct {
	ArticleID      int `gorm:"column:id"`
	BoardID        int
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
	ShortName     string
	LongName      string
	Name          string `gorm:"primaryKey"`
	Description   string
	AdminUsername string
	Admin         KhumuUser `gorm:"foreignKey:AdminUsername"`
}

func (*Board) TableName() string {
	return "board_board"
}

type Comment struct {
	ID             int             `gorm:"column:id" json:"id"`
	// Kind: (anonymous, named)
	Kind           string           `gorm:"column:kind; default:anonymous" json:"kind"`
	// State: (exists, deleted)
	State           string           `gorm:"column:state; default:exists" json:"state"`
	Author         *KhumuUserSimple `gorm:"foreignKey:AuthorUsername; references:Username; constraint:OnDELETE:CASCADE" json:"author"`
	AuthorUsername string           `gorm:"column:author_id" json:"-"`
	ArticleID      int             `gorm:"column:article_id" json:"article"`
	Content   string     `json:"content"`
	ParentID  int      `gorm:"column:parent_id;default:null" json:"parent"`
	Parent *Comment `gorm:"foreignKey: ParentID;constraint: OnDelete: CASCADE" json:",omitempty"`
	Children  []*Comment `gorm:"foreignKey:ParentID;references:ID" json:"children"` //Has-Many relationship => Preload 필요
	LikeCommentCount int `gorm:"-" json:"like_comment_count"`
	Liked bool `gorm:"-" json:"liked"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
}

func (*Comment) TableName() string {
	return "comment_comment"
}

type LikeComment struct{
	ID int `gorm:"primaryKey"`
	CommentID int `gorm:"column:comment_id" json:"comment"`
	Comment *Comment `gorm:"foreignKey: CommentID; references:ID; constraint:OnDelete:CASCADE;" json:"-"`
	Username string `gorm:"column:username" json:"username"`
	User *KhumuUserSimple `gorm:"foreignKey: Username; references:Username; constraint:OnDelete:CASCADE" json:"-"`
}

func (*LikeComment) TableName() string {
	return "comment_likecomment"
}
