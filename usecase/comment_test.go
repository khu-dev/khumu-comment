package usecase

//type CommentRepositoryMock struct{}
//
//var (
//	commentsMock []*model.Comment
//	mockRepository
//)
//
//// QyeryOption기능은 제외하고 mock
//func(r *CommentRepositoryMock) List(opt *repository.CommentQueryOption) []*model.Comment{
//	return mockComments
//}
//
//func(r *CommentRepositoryMock) Get(id int) *model.Comment{
//	for _, comment := range mockComments{
//		if int(comment.ArticleID) == id{
//			return comment
//		}
//	}
//	return nil
//}
//
//func TestInit(t *testing.T){
//	for i := 0; i<5; i++{
//		c := &model.Comment{ID: uint(i)}
//		mockComments = append(mockComments, c)
//	}
//	assert.NotEmpty(t, mockComments)
//}
