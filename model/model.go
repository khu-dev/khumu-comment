package model

import (
	"encoding/json"
	"gopkg.in/guregu/null.v4"
	"time"
)

type KhumuUser struct {
	//gorm.Model
	Username      string `gorm:"primaryKey"`
	Email         string
	IsActive      bool
	State         string `gorm:"column:State"`
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
	Nickname string `gorm:"unique" json:"nickname"`
	State    string `gorm:"column:state" json:"state"`
}

func (*KhumuUserSimple) TableName() string {
	return "user_khumuuser"
}

type Board struct {
	Name        string `gorm:"primaryKey"`
	DisplayName string `gorm:"column:display_name"`
}

func (*Board) TableName() string {
	return "board_board"
}

type Article struct {
	ArticleID      int `gorm:"column:id"`
	BoardName      string
	Board          Board `gorm:"foreignKey:BoardName"`
	Title          string
	AuthorUsername string    `gorm:"column:author_id"`
	Author         KhumuUser `gorm:"foreignKey:AuthorUsername"`
	Content        string
	CreatedAt      time.Time
}

func (*Article) TableName() string {
	return "article_article"
}

type StudyArticle struct {
	Id             int    `gorm:"column:id"`
	AuthorUsername string `gorm:"column:author_id"`
	CreatedAt      time.Time
}

func (*StudyArticle) TableName() string {
	return "article_studyarticle"
}

type Comment struct {
	ID int `gorm:"column:id" json:"id"`
	// Kind: (anonymous, named)
	Kind string `gorm:"column:kind; default:anonymous" json:"kind"`
	// State: (exists, deleted)
	State          string           `gorm:"column:state; default:exists" json:"state"`
	Author         *KhumuUserSimple `gorm:"foreignKey:AuthorUsername; references:Username; constraint:OnDELETE:CASCADE" json:"author"`
	AuthorUsername string           `gorm:"column:author_id" json:"-"`
	ArticleID      null.Int         `gorm:"column:article_id" json:"article"`
	StudyArticleID null.Int         `gorm:"column:study_article_id" json:"study_article"`
	Content        string           `json:"content"`
	ParentID       null.Int         `gorm:"column:parent_id;default:null" json:"parent"`
	Parent         *Comment         `gorm:"foreignKey: ParentID;constraint: OnDelete: CASCADE" json:",omitempty"`
	Children       []*Comment       `gorm:"foreignKey:ParentID;references:ID" json:"children"` //Has-Many relatio
	CreatedAt      time.Time        `gorm:"autoCreateTime" json:"-"`                           // nship => Preload 필요
	// 여기서 부턴 gorm과 상관 없는 field
	IsAuthor            bool   `gorm:"-" json:"is_author"`
	LikeCommentCount    int    `gorm:"-" json:"like_comment_count"`
	Liked               bool   `gorm:"-" json:"liked"`
	CreatedAtExpression string `gorm:"-" json:"created_at"`
}

func (*Comment) TableName() string {
	return "comment_comment"
}

type LikeComment struct {
	ID        int              `gorm:"primaryKey"`
	CommentID int              `gorm:"column:comment_id" json:"comment"`
	Comment   *Comment         `gorm:"foreignKey: CommentID; references:ID; constraint:OnDelete:CASCADE;" json:"-"`
	Username  string           `gorm:"column:user_id" json:"username"`
	User      *KhumuUserSimple `gorm:"foreignKey: Username; references:Username; constraint:OnDelete:CASCADE" json:"-"`
}

func (*LikeComment) TableName() string {
	return "comment_likecomment"
}

// eventmessage가 event를 publish 할 때 사용
type EventMessage struct {
	ResourceKind string   `json:"resource_kind"`
	EventKind    string   `json:"event_kind"`
	Resource     *Comment `json:"resource"`
}

// redis는 encoding.BinaryMarshaler를 구현한 type만을 Marshal 가능하다.
func (e *EventMessage) MarshalBinary() (data []byte, err error) {
	return json.Marshal(e)
}
