// test를 진행할 때 필요한 기본 데이터를 제공한다.
//
// 기본적인 흐름
// 3 유저 jinsu, somebody, puppy 존재
// 1번 article에 3 user가 각각 comment 단다.
// 이 흐름들은 주로 test 코드의 Create 부분에서 실행된다.
// like-comment의 경우 Toggle인 경우도 있다.

package test

import (
	"github.com/khu-dev/khumu-comment/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	UsersData map[string]*model.KhumuUserSimple
	BoardData []*model.Board
	ArticleData []*model.Article
	CommentsData map[string]*model.Comment
	// parent Comment의 ID는 모두 1
	ReplyCommentsData map[string]*model.Comment
)

func init(){
	UsersData = map[string]*model.KhumuUserSimple{
		"Jinsu": &model.KhumuUserSimple{
			Username: "jinsu",
			Nickname: "진수짱짱맨",
		},
		"Somebody": &model.KhumuUserSimple{
			Username: "somebody",
			Nickname: "썸바디",
		},
		"Puppy": &model.KhumuUserSimple{
			Username: "puppy",
			Nickname: "댕댕이",
		},
	}
	
	BoardData = []*model.Board{
		&model.Board{Name: "free", DisplayName: "자유게시판"},
		&model.Board{Name: "department_computer_engineering", DisplayName: "컴퓨터공학과"},
	}

	ArticleData = []*model.Article{
		&model.Article{
			ArticleID: 1,
			BoardName: "free",
			Title: "1번 게시물입니다.",
			AuthorUsername: "jinsu",
			Content: "이것은 1번 게시물!",
		},
		&model.Article{
			ArticleID: 2,
			BoardName: "free",
			Title: "2번 게시물입니다.",
			AuthorUsername: "somebody",
			Content: "이것은 2번 게시물!",
		},
		&model.Article{
			ArticleID: 3,
			BoardName: "free",
			Title: "3번 게시물입니다.",
			AuthorUsername: "puppy",
			Content: "이것은 3번 게시물!",
		},
	}

	CommentsData = map[string]*model.Comment{
		"JinsuAnonymousComment": &model.Comment{
			Kind: "anonymous",
			//AuthorUsername: "jinsu",
			AuthorUsername: "jinsu",
			ArticleID: 1,
			Content:   "테스트로 작성한 jinsu의 익명 코멘트",
			ParentID:  nil,
		},
		"JinsuNamedComment": &model.Comment{
			Kind:           "named",
			AuthorUsername: "jinsu",
			ArticleID:      1,
			Content:        "테스트로 작성한 jinsu의 기명 코멘트",
			ParentID:       nil,
		},
		"SomebodyAnonymousComment": &model.Comment{
			Kind:           "anonymous",
			AuthorUsername: "somebody",
			ArticleID:      1,
			Content:        "테스트로 작성한 somebody의 익명 코멘트",
			ParentID:       nil,
		},
		"PuppyAnonymousComment": &model.Comment{
			Kind:           "anonymous",
			AuthorUsername: "puppy",
			ArticleID:      1,
			Content:        "테스트로 작성한 puppy의 익명 코멘트",
			ParentID:       nil,
		},
	}

	parentID := 1
	ReplyCommentsData = map[string]*model.Comment{
		"JinsuAnonymousReplyComment": &model.Comment{
			Kind:           "anonymous",
			AuthorUsername: "jinsu",
			ArticleID:      1,
			Content:        "테스트로 작성한 jinsu의 익명 대댓글",
			ParentID:       &parentID,
		},
		"JinsuNamedReplyComment": &model.Comment{
			Kind:           "anonymous",
			AuthorUsername: "jinsu",
			ArticleID:      1,
			Content:        "테스트로 작성한 jinsu의 기명 대댓글",
			ParentID:       &parentID,
		},
		"SomebodyAnonymousReplyComment": &model.Comment{
			Kind:           "anonymous",
			AuthorUsername: "somebody",
			ArticleID:      1,
			Content:        "테스트로 작성한 somebody의 익명 코멘트",
			ParentID:       &parentID,
		},
		"PuppyAnonymousReplyComment": &model.Comment{
			Kind:           "anonymous",
			AuthorUsername: "puppy",
			ArticleID:      1,
			Content:        "테스트로 작성한 puppy의 익명 코멘트",
			ParentID:       &parentID,
		},
	}
}

// test 진행 시에 각 step에서 사용할 초기 데이터를 만든다.
func SetUp(db *gorm.DB){
	MigrateAll(db)
	for _, userData := range UsersData {
		err := db.Create(userData).Error
		if err != nil{
			logrus.Fatal(err)
		}
	}
	for _, boardData := range BoardData{
		err := db.Create(boardData).Error
		if err != nil{
			logrus.Fatal(err)
		}
	}
	for _, articleData := range ArticleData{
		err := db.Omit("Board", "Author").Create(articleData).Error
		if err != nil{
			logrus.Fatal(err)
		}
	}
	for _, commentData := range CommentsData{
		err := db.Omit("Author", "Parent", "Children").Create(commentData).Error
		if err != nil{
			logrus.Fatal(err)
		}
	}
	// reply comment는 comment에 의존성이있다.
	for _, rcData := range ReplyCommentsData{
		err := db.Omit("Author", "Parent", "Children").Create(rcData).Error
		if err != nil{
			logrus.Fatal(err)
		}
	}
}

// test 진행 시에 각 step에서 진행한 내용을 초기화한다.
func CleanUp(db *gorm.DB){
	err := db.Migrator().DropTable(&model.LikeComment{}, &model.Comment{}, &model.Article{}, &model.KhumuUserSimple{}, &model.Board{})
	if err != nil{
		logrus.Fatal(err)
	}
}

func MigrateAll(db *gorm.DB){
	err := db.AutoMigrate(&model.Board{}, &model.KhumuUser{}, &model.Article{}, &model.Comment{}, &model.LikeComment{})
	if err != nil{
		logrus.Fatal(err)
	}
}