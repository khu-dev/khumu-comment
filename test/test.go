// test를 진행할 때 필요한 기본 데이터를 제공한다.
//
// 기본적인 흐름
// 3 유저 jinsu, somebody, puppy 존재
// 1번 article에 3 user가 각각 comment 단다.
// 이 흐름들은 주로 test 코드의 Create 부분에서 실행된다.
// like-comment의 경우 Toggle인 경우도 있다.

package test

import "github.com/khu-dev/khumu-comment/model"

var CommentsData map[string]*model.Comment
var UsersData map[string]*model.KhumuUserSimple

func init(){
	CommentsData = make(map[string]*model.Comment)
	UsersData = make(map[string]*model.KhumuUserSimple)

	var id int = 1
	CommentsData["JinsuAnonymousComment"] = &model.Comment{
		Kind:           "anonymous",
		//AuthorUsername: "jinsu",
		Author: &model.KhumuUserSimple{Username: "jinsu"},
		ArticleID:      1,
		Content:        "테스트로 작성한 jinsu의 익명 코멘트",
		ParentID:       0,
	}
	id++
	CommentsData["JinsuNamedComment"] = &model.Comment{
		Kind:           "named",
		Author: &model.KhumuUserSimple{Username: "jinsu"},
		ArticleID:      1,
		Content:        "테스트로 작성한 jinsu의 기명 코멘트",
		ParentID:       0,
	}
	id++
	CommentsData["SomebodyAnonymousComment"] = &model.Comment{
		Kind:           "anonymous",
		Author: &model.KhumuUserSimple{Username: "somebody"},
		ArticleID:      1,
		Content:        "테스트로 작성한 somebody의 익명 코멘트",
		ParentID:       0,
	}
	UsersData["Jinsu"] = &model.KhumuUserSimple{Username: "jinsu", Nickname: "진수짱짱맨"}
	UsersData["Somebody"] = &model.KhumuUserSimple{Username: "somebody", Nickname: "썸바디"}
	UsersData["Puppy"] = &model.KhumuUserSimple{Username: "puppy", Nickname: "댕댕이"}
}
