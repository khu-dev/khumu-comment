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

	Comment1JinsuAnnonymous               *ent.Comment
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
func SetUp(client *ent.Client) {
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

	_, err = client.Article.Create().
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

	//Comments = make([]*ent.Comment, 0)
	//
	//Comment1JinsuAnnonymous = &ent.Comment{
	//	ID:   1,
	//	Kind: "anonymous",
	//	Content: "테스트로 작성한 jinsu의 익명 코멘트",
	//	Edges: ent.CommentEdges{
	//		Author: UserJinsu,
	//		Article: Articles[0],
	//		Parent: nil,
	//		Children: []*ent.Comment{},
	//	},
	//}
	//Comments = append(Comments, Comment1JinsuAnnonymous)
	//
	//Comment2JinsuNamed = &ent.Comment{
	//	ID:             2,
	//	Kind:           "named",
	//	Content: "테스트로 작성한 jinsu의 기명 코멘트",
	//	Edges: ent.CommentEdges{
	//		Author: UserJinsu,
	//		Article: Articles[0],
	//		Parent: nil,
	//		Children: []*ent.Comment{},
	//	},
	//}
	//Comments = append(Comments, Comment2JinsuNamed)
	//
	//Comment3SomebodyAnonymous = &ent.Comment{
	//	ID:             3,
	//	Kind:           "anonymous",
	//	Content: "테스트로 작성한 somebody의 익명 코멘트",
	//	Edges: ent.CommentEdges{
	//		Author: UserSomebody,
	//		Article: Articles[0],
	//		Parent: nil,
	//		Children: []*ent.Comment{},
	//	},
	//}
	//Comments = append(Comments, Comment3SomebodyAnonymous)
	//
	//Comment4PuppyAnonymous = &ent.Comment{
	//	ID:             4,
	//	Kind:           "anonymous",
	//	Content:   "테스트로 작성한 puppy의 익명 코멘트",
	//	Edges: ent.CommentEdges{
	//		Author: UserPuppy,
	//		Article: Articles[0],
	//		Parent: nil,
	//		Children: []*ent.Comment{},
	//	},
	//}
	//Comments = append(Comments, Comment4PuppyAnonymous)
	//
	//Comment5JinsuAnonymousFromComment1 = &ent.Comment{
	//	ID:             5,
	//	Kind:           "anonymous",
	//	Content:   "테스트로 작성한 jinsu의 익명 대댓글",
	//	Edges: ent.CommentEdges{
	//		Author: UserJinsu,
	//		Article: Articles[0],
	//		Parent: Comments[0],
	//		Children: []*ent.Comment{},
	//	},
	//}
	//Comments = append(Comments, Comment5JinsuAnonymousFromComment1)
	//
	//Comment6JinsuNamedFromComment1 = &ent.Comment{
	//	ID:             6,
	//	Kind:           "named",
	//	Content:   "테스트로 작성한 jinsu의 기명 대댓글",
	//	Edges: ent.CommentEdges{
	//		Author: UserJinsu,
	//		Article: Articles[0],
	//		Parent: Comments[0],
	//		Children: []*ent.Comment{},
	//	},
	//}
	//Comments = append(Comments, Comment6JinsuNamedFromComment1)
	//
	//Comment7SomebodyAnonymousFromComment1 = &ent.Comment{
	//	ID:             7,
	//	Kind:           "anonymous",
	//	Content:   "테스트로 작성한 somebody의 익명 코멘트",
	//	Edges: ent.CommentEdges{
	//		Author: UserSomebody,
	//		Article: Articles[0],
	//		Parent: Comments[0],
	//		Children: []*ent.Comment{},
	//	},
	//}
	//Comments = append(Comments, Comment7SomebodyAnonymousFromComment1)
	//
	//Comment8PuppyAnonymousFromComment1 = &ent.Comment{
	//	ID:             8,
	//	Kind:           "anonymous",
	//	Content:   "테스트로 작성한 puppy의 익명 코멘트",
	//	Edges: ent.CommentEdges{
	//		Author: UserPuppy,
	//		Article: Articles[0],
	//		Parent: Comments[0],
	//		Children: []*ent.Comment{},
	//	},
	//}
	//Comments = append(Comments, Comment8PuppyAnonymousFromComment1)
}
