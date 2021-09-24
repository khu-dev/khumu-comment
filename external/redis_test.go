package external

import (
	_ "github.com/khu-dev/khumu-comment/config"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestRedisAdapterImpl_GetAllByArticle(t *testing.T) {
	//type fields struct {
	//	client *redis.Client
	//}
	//type args struct {
	//	articleID int
	//}
	//tests := []struct {
	//	name   string
	//	fields fields
	//	args   args
	//	want   []*ent.Comment
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		a := &RedisAdapterImpl{
	//			client: tt.fields.client,
	//		}
	//		if got := a.GetCommentsByArticle(tt.args.articleID); !reflect.DeepEqual(got, tt.want) {
	//			t.Errorf("GetCommentsByArticle() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
}

//func TestRedisAdapterImpl_SetNewComment(t *testing.T) {
//	db := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
//	repo := repository.NewCommentRepository(db)
//	a := NewRedisAdapter(repo)
//	test.SetUpUsers(db)
//	test.SetUpArticles(db)
//	test.SetUpComments(db)
//	a.RefreshCommentsByArticle(test.Articles[0].ID)
//	results := a.GetCommentsByArticle(test.Articles[0].ID)
//	log.Info(results)
//}
