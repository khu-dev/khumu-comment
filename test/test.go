// test를 진행할 때 필요한 기본 데이터를 메모리에 적재한다.
// repository 계층에선 필요 없는 편이다.
// 다른 계층에선 repository의 실제 데이터 대신 이곳의 데이터를 이용할 수 있다.

package test

import (
	"context"
	"github.com/khu-dev/khumu-comment/ent"
	"github.com/sirupsen/logrus"
)

var (
	UserJinsu    *ent.KhumuUser
	UserSomebody *ent.KhumuUser
	UserPuppy    *ent.KhumuUser
	Users        ent.KhumuUsers

	Articles ent.Articles

	Comment1JinsuAnonymous               *ent.Comment
	Comment2JinsuNamed                    *ent.Comment
	Comment3SomebodyAnonymous             *ent.Comment
	Comment4PuppyAnonymous                *ent.Comment
	Comment5JinsuAnonymousFromComment1    *ent.Comment
	Comment6JinsuNamedFromComment1        *ent.Comment
	Comment7SomebodyAnonymousFromComment1 *ent.Comment
	Comment8PuppyAnonymousFromComment1    *ent.Comment

	Comments ent.Comments
	// parent Comment의 ID는 모두 1
)

func init() {
}

// test 진행 시에 각 step에서 사용할 초기 데이터를 만든다.
func SetUpUsers(client *ent.Client) {
	ctx := context.Background()
	var err error
	UserJinsu, err = client.KhumuUser.Create().
		SetID("jinsu").
		SetPassword("123123").
		SetNickname("진수짱짱맨").
		SetState("active").
		Save(ctx)
	if err != nil {
		logrus.Panic(err)
	}

	UserSomebody, err = client.KhumuUser.Create().
		SetID("somebody").
		SetPassword("123123").
		SetNickname("썸바디").
		SetState("active").
		Save(ctx)
	if err != nil {
		logrus.Panic(err)
	}

	UserPuppy, err = client.KhumuUser.Create().
		SetID("puppy").
		SetPassword("123123").
		SetNickname("댕댕이").
		SetState("active").
		Save(ctx)
	if err != nil {
		logrus.Panic(err)
	}

	Users, err = client.KhumuUser.Query().All(ctx)
	if err != nil {
		logrus.Panic(err)
	}
}

func SetUpArticles(client *ent.Client) {
	ctx := context.TODO()
	_, err := client.Article.Create().
		SetID(1).
		SetTitle("1번 게시물입니다.").
		SetImages(&[]string{}).
		SetAuthor(UserJinsu).
		Save(ctx)
	if err != nil {
		logrus.Panic(err)
	}
	_, err = client.Article.Create().
		SetID(2).
		SetTitle("2번 게시물입니다.").
		SetImages(&[]string{}).
		SetAuthor(UserSomebody).
		Save(ctx)
	if err != nil {
		logrus.Panic(err)
	}
	_, err = client.Article.Create().
		SetID(3).
		SetTitle("3번 게시물입니다.").
		SetImages(&[]string{}).
		SetAuthor(UserPuppy).
		Save(ctx)
	if err != nil {
		logrus.Panic(err)
	}

	Articles, err = client.Article.Query().All(ctx)
	if err != nil {
		logrus.Panic(err)
	}
}

// 좋아요 기능 및 댓글에 대한 수정, 삭제 작업에 사용할 Fixture comment들
func SetUpComments(client *ent.Client) {
	var err error
	ctx := context.TODO()
	Comment1JinsuAnonymous, err = client.Comment.Create().
		SetArticleID(1).
		SetAuthorID("jinsu").
		SetKind("anonymous").
		SetContent("테스트로 작성한 jinsu의 익명 코멘트").
		Save(ctx)
	if err != nil {
		logrus.Panic(err)
	}

	Comment2JinsuNamed, err = client.Comment.Create().
		SetArticleID(1).
		SetAuthorID("jinsu").
		SetKind("named").
		SetContent("테스트로 작성한 jinsu의 기명 코멘트").
		Save(ctx)
	if err != nil {
		logrus.Panic(err)
	}

	Comment3SomebodyAnonymous, err = client.Comment.Create().
		SetArticleID(1).
		SetAuthorID("somebody").
		SetKind("anonymous").
		SetContent("테스트로 작성한 somebody의 익명 코멘트").
		Save(ctx)
	if err != nil {
		logrus.Panic(err)
	}

	Comment4PuppyAnonymous, err = client.Comment.Create().
		SetArticleID(1).
		SetAuthorID("puppy").
		SetKind("anonymous").
		SetContent("테스트로 작성한 puppy의 익명 코멘트").
		Save(ctx)
	if err != nil {
		logrus.Panic(err)
	}

	Comment5JinsuAnonymousFromComment1, err = client.Comment.Create().
		SetArticleID(1).
		SetAuthorID("jinsu").
		SetKind("anonymous").
		SetContent("테스트로 작성한 jinsu의 익명 대댓글").
		SetParent(Comment1JinsuAnonymous).
		Save(ctx)
	if err != nil {
		logrus.Panic(err)
	}

	Comment6JinsuNamedFromComment1, err = client.Comment.Create().
		SetArticleID(1).
		SetAuthorID("jinsu").
		SetKind("named").
		SetContent("테스트로 작성한 jinsu의 기명 대댓글").
		SetParent(Comment1JinsuAnonymous).
		Save(ctx)
	if err != nil {
		logrus.Panic(err)
	}

	Comment7SomebodyAnonymousFromComment1, err = client.Comment.Create().
		SetArticleID(1).
		SetAuthorID("somebody").
		SetKind("anonymous").
		SetContent("테스트로 작성한 somebody의 익명 코멘트").
		SetParent(Comment1JinsuAnonymous).
		Save(ctx)
	if err != nil {
		logrus.Panic(err)
	}

	Comment8PuppyAnonymousFromComment1, err = client.Comment.Create().
		SetArticleID(1).
		SetAuthorID("puppy").
		SetKind("anonymous").
		SetContent("테스트로 작성한 puppy의 익명 코멘트").
		SetParent(Comment1JinsuAnonymous).
		Save(ctx)
	if err != nil {
		logrus.Panic(err)
	}

	Comments, err = client.Comment.Query().All(ctx)
	if err != nil {
		logrus.Panic(err)
	}
}
