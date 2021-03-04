// test를 진행할 때 필요한 기본 데이터를 메모리에 적재한다.
// repository 계층에선 필요 없는 편이다.
// 다른 계층에선 repository의 실제 데이터 대신 이곳의 데이터를 이용할 수 있다.

package test

import (
	"github.com/khu-dev/khumu-comment/model"
)

var (
	UserJinsu    *model.KhumuUserSimple
	UserSomebody *model.KhumuUserSimple
	UserPuppy    *model.KhumuUserSimple
	Users        []*model.KhumuUserSimple

	BoardFree                          *model.Board
	BoardDepartmentComputerEngineering *model.Board
	Boards                             []*model.Board

	Articles []*model.Article

	Comment1JinsuAnnonymous               *model.Comment
	Comment2JinsuNamed                    *model.Comment
	Comment3SomebodyAnonymous             *model.Comment
	Comment4PuppyAnonymous                *model.Comment
	Comment5JinsuAnonymousFromComment1    *model.Comment
	Comment6JinsuNamedFromComment1        *model.Comment
	Comment7SomebodyAnonymousFromComment1 *model.Comment
	Comment8PuppyAnonymousFromComment1    *model.Comment

	Comments []*model.Comment
	// parent Comment의 ID는 모두 1
)

func init() {
}

// test 진행 시에 각 step에서 사용할 초기 데이터를 만든다.
func SetUp() {

	Users = make([]*model.KhumuUserSimple, 0)
	Boards = make([]*model.Board, 0)
	Articles = make([]*model.Article, 0)
	Comments = make([]*model.Comment, 0)

	UserJinsu = &model.KhumuUserSimple{
		Username: "jinsu",
		Nickname: "진수짱짱맨",
		State:    "active",
	}
	Users = append(Users, UserJinsu)

	UserSomebody = &model.KhumuUserSimple{
		Username: "somebody",
		Nickname: "썸바디",
		State:    "active",
	}
	Users = append(Users, UserSomebody)

	UserPuppy = &model.KhumuUserSimple{
		Username: "puppy",
		Nickname: "댕댕이",
		State:    "active",
	}
	Users = append(Users, UserPuppy)

	BoardFree = &model.Board{Name: "free", DisplayName: "자유게시판"}
	Boards = append(Boards, BoardFree)
	BoardDepartmentComputerEngineering = &model.Board{Name: "department_computer_engineering", DisplayName: "컴퓨터공학과"}
	Boards = append(Boards, BoardDepartmentComputerEngineering)

	Articles = make([]*model.Article, 0)
	Articles = append(Articles,
		&model.Article{
			ArticleID:      1,
			BoardName:      "free",
			Title:          "1번 게시물입니다.",
			AuthorUsername: "jinsu",
			Content:        "이것은 1번 게시물!",
		})
	Articles = append(Articles,
		&model.Article{
			ArticleID:      2,
			BoardName:      "free",
			Title:          "2번 게시물입니다.",
			AuthorUsername: "somebody",
			Content:        "이것은 2번 게시물!",
		})
	Articles = append(Articles,
		&model.Article{
			ArticleID:      3,
			BoardName:      "free",
			Title:          "3번 게시물입니다.",
			AuthorUsername: "puppy",
			Content:        "이것은 3번 게시물!",
		})

	Comments = make([]*model.Comment, 0)

	Comment1JinsuAnnonymous = &model.Comment{
		ID:             1,
		Kind:           "anonymous",
		AuthorUsername: "jinsu",
		Author: &model.KhumuUserSimple{
			Username: "jinsu",
			Nickname: "진수짱짱맨",
			State:    "active",
		},
		ArticleID: 1,
		Content:   "테스트로 작성한 jinsu의 익명 코멘트",
		ParentID:  nil,
	}
	Comments = append(Comments, Comment1JinsuAnnonymous)

	Comment2JinsuNamed = &model.Comment{
		ID:             2,
		Kind:           "named",
		AuthorUsername: "jinsu",
		Author: &model.KhumuUserSimple{
			Username: "jinsu",
			Nickname: "진수짱짱맨",
			State:    "active",
		},
		ArticleID: 1,
		Content:   "테스트로 작성한 jinsu의 기명 코멘트",
		ParentID:  nil,
	}
	Comments = append(Comments, Comment2JinsuNamed)

	Comment3SomebodyAnonymous = &model.Comment{
		ID:             3,
		Kind:           "anonymous",
		AuthorUsername: "somebody",
		Author: &model.KhumuUserSimple{
			Username: "somebody",
			Nickname: "썸바디",
			State:    "active",
		},
		ArticleID: 1,
		Content:   "테스트로 작성한 somebody의 익명 코멘트",
		ParentID:  nil,
	}
	Comments = append(Comments, Comment3SomebodyAnonymous)

	Comment4PuppyAnonymous = &model.Comment{
		ID:             4,
		Kind:           "anonymous",
		AuthorUsername: "puppy",
		Author: &model.KhumuUserSimple{
			Username: "puppy",
			Nickname: "댕댕이",
			State:    "active",
		},
		ArticleID: 1,
		Content:   "테스트로 작성한 puppy의 익명 코멘트",
		ParentID:  nil,
	}
	Comments = append(Comments, Comment4PuppyAnonymous)

	parentID := 1
	Comment5JinsuAnonymousFromComment1 = &model.Comment{
		ID:             5,
		Kind:           "anonymous",
		AuthorUsername: "jinsu",
		Author: &model.KhumuUserSimple{
			Username: "jinsu",
			Nickname: "진수짱짱맨",
			State:    "active",
		},
		ArticleID: 1,
		Content:   "테스트로 작성한 jinsu의 익명 대댓글",
		ParentID:  &parentID,
	}
	Comments = append(Comments, Comment5JinsuAnonymousFromComment1)

	Comment6JinsuNamedFromComment1 = &model.Comment{
		ID:             6,
		Kind:           "anonymous",
		AuthorUsername: "jinsu",
		Author: &model.KhumuUserSimple{
			Username: "jinsu",
			Nickname: "진수짱짱맨",
			State:    "active",
		},
		ArticleID: 1,
		Content:   "테스트로 작성한 jinsu의 기명 대댓글",
		ParentID:  &parentID,
	}
	Comments = append(Comments, Comment6JinsuNamedFromComment1)

	Comment7SomebodyAnonymousFromComment1 = &model.Comment{
		ID:             7,
		Kind:           "anonymous",
		AuthorUsername: "somebody",
		Author: &model.KhumuUserSimple{
			Username: "somebody",
			Nickname: "썸바디",
			State:    "active",
		},
		ArticleID: 1,
		Content:   "테스트로 작성한 somebody의 익명 코멘트",
		ParentID:  &parentID,
	}
	Comments = append(Comments, Comment7SomebodyAnonymousFromComment1)

	Comment8PuppyAnonymousFromComment1 = &model.Comment{
		ID:             8,
		Kind:           "anonymous",
		AuthorUsername: "puppy",
		Author: &model.KhumuUserSimple{
			Username: "puppy",
			Nickname: "댕댕이",
			State:    "active",
		},
		ArticleID: 1,
		Content:   "테스트로 작성한 puppy의 익명 코멘트",
		ParentID:  &parentID,
	}
	Comments = append(Comments, Comment8PuppyAnonymousFromComment1)
}
