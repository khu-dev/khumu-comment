package repository

import (
	"github.com/khu-dev/khumu-comment/model"
	"github.com/khu-dev/khumu-comment/test"
    "github.com/sirupsen/logrus"
    "go.uber.org/dig"
    "gorm.io/gorm"
    "testing"
)

var (
	db                    *gorm.DB
	commentRepository     CommentRepositoryInterface
	likeCommentRepository LikeCommentRepositoryInterface
	eventMessageRepository EventMessageRepository
)

func TestMain(m *testing.M) {
	cont := BuildContainer()
	err := cont.Invoke(func(database *gorm.DB, cr CommentRepositoryInterface, lcr LikeCommentRepositoryInterface, evtr EventMessageRepository) {
		db = database
		commentRepository = cr
		likeCommentRepository = lcr
		eventMessageRepository = evtr
	})
	if err != nil {
		logrus.Fatal(err)
	}

	m.Run()
}

// B는 Before each의 acronym
func B(tb testing.TB) {
	test.SetUp(db)
	// test 진행 시에 각 step에서 진행한 내용을 초기화한다.
	//func MigrateAll(db *gorm.DB){
	//	err := db.AutoMigrate(&model.Board{}, &model.KhumuUser{}, &model.Article{}, &model.Comment{}, &model.LikeComment{})
	//	if err != nil {
	//		logrus.Fatal(err)
	//	}
	//}()
}

}

// A는 After each의 acronym
func A(tb testing.TB) {
	//test.CleanUp(db)
	//func CleanUp(db *gorm.DB) {
	//	err := db.Migrator().DropTable(&model.LikeComment{}, &model.Comment{}, &model.Article{}, &model.KhumuUserSimple{}, &model.Board{})
	//	if err != nil {
	//		logrus.Fatal(err)
	//	}
	//}
}

// 처음 Test 환경을 만들 때에 필요한 SetUp 작업들
// 주로 IoC 컨테이너
func BuildContainer() *dig.Container {
	cont := dig.New()
	err := cont.Provide(NewTestGorm)
	if err != nil {
		logrus.Fatal(err)
	}
	err = cont.Provide(NewCommentRepositoryGorm)
	if err != nil {
		logrus.Fatal(err)
	}
	err = cont.Provide(NewLikeCommentRepositoryGorm)
	if err != nil {
		logrus.Fatal(err)
	}
	err = cont.Provide(NewRedisEventMessageRepository)
	if err != nil {
		logrus.Fatal(err)
	}

	return cont
}
