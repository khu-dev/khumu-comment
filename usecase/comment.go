package usecase

import (
	"github.com/khu-dev/khumu-comment/repository"
	"log"
)
import "github.com/khu-dev/khumu-comment/model"

type CommentUseCaseInterface interface {
	Create(comment *model.Comment) (*model.Comment, error)
	List(username string, opt *repository.CommentQueryOption) ([]*model.Comment, error)
	Get(id int, opt *repository.CommentQueryOption) (*model.Comment, error)
}

type LikeCommentUseCaseInterface interface{
	// return value 중 bool이 true면 생성, false면 삭제
	Toggle(like *model.LikeComment) (bool, error)
}

type CommentUseCase struct {
	Repository repository.CommentRepositoryInterface
	LikeCommentRepository repository.LikeCommentRepositoryInterface
}

type LikeCommentUseCase struct{
	Repository repository.LikeCommentRepositoryInterface
	CommentRepository repository.CommentRepositoryInterface
}

type SomeoneLikesHisCommentError string

func NewCommentUseCase(repository repository.CommentRepositoryInterface, likeRepository repository.LikeCommentRepositoryInterface) CommentUseCaseInterface {
	return &CommentUseCase{Repository: repository, LikeCommentRepository: likeRepository}
}

func (uc *CommentUseCase) Create(comment *model.Comment) (*model.Comment, error) {
	log.Println("CommentUseCase_Create")
	newComment, err := uc.Repository.Create(comment)
	if err!=nil{return newComment, err}
	return newComment, nil
}

func (uc *CommentUseCase) List(username string, opt *repository.CommentQueryOption) ([]*model.Comment, error) {

	log.Println("CommentUseCase List")
	comments := uc.Repository.List(&repository.CommentQueryOption{})
	parents := uc.listParentWithChildren(comments)

	for _, p := range parents {
		for _, c := range p.Children {
			uc.hideAuthor(c, username)
			likeCount, err := uc.getLikeCommentCount(c.ID)
			if err != nil{return nil, err}
			c.LikeCommentCount = likeCount
		}
		uc.hideAuthor(p, username)
		likeCount, err := uc.getLikeCommentCount(p.ID)
		if err != nil{return nil, err}
		p.LikeCommentCount = likeCount
	}

	return parents, nil
}

func (uc *CommentUseCase) Get(id int, opt *repository.CommentQueryOption) (*model.Comment, error) {
	log.Println("CommentUseCase_Get")
	comment := uc.Repository.Get(id)
	return comment, nil
}

func (uc *CommentUseCase) listParentWithChildren(allComments []*model.Comment) []*model.Comment {
	var parents []*model.Comment

	for _, comment := range allComments {
		if comment.ParentID == nil {
			parents = append(parents, comment)
		}
	}

	return parents
}

func (uc *CommentUseCase) hideAuthor(c *model.Comment, requestUsername string) {
	if c.Kind == "anonymous" && c.AuthorUsername != requestUsername {
		c.AuthorUsername = "익명"
		c.Author.Username = "익명"
	}
}

func (uc *CommentUseCase) getLikeCommentCount(commentID int) (int, error){
	likes, err := uc.LikeCommentRepository.List(&repository.LikeCommentQueryOption{CommentID: commentID})
	if err != nil{return 0, err}
	return len(likes), nil
}

func NewLikeCommentUseCase(
	likeRepo repository.LikeCommentRepositoryInterface,
	commentRepo repository.CommentRepositoryInterface) LikeCommentUseCaseInterface{
	return &LikeCommentUseCase{Repository: likeRepo, CommentRepository: commentRepo}
}

func (uc *LikeCommentUseCase) Toggle(like *model.LikeComment) (bool, error){
	likes, err := uc.Repository.List(&repository.LikeCommentQueryOption{CommentID: like.CommentID, Username: like.Username})
	if err != nil{
		log.Panic(false, err)
	}
	// 길이가 1보다 크거나 같으면 삭제. 1인 경우는 정상적으로 하나만 있을 때,
	// 1보다 큰 경우는 비정상적으로 여러개 존재할 때
	if len(likes) >= 1 {
		for _, like := range likes {
			err = uc.Repository.Delete(like.ID)
			if err != nil {
				log.Panic(false, err)
			}
		}
		return false, err
	}else{
		// 생성
		comment := uc.CommentRepository.Get(like.CommentID)
		if comment.AuthorUsername == like.Username{
			return false, SomeoneLikesHisCommentError(like.Username)
		}
		_, err := uc.Repository.Create(like)
		if err != nil{return false, err}

		return true, err
	}
}

func (e SomeoneLikesHisCommentError) Error() string{
	return "Error: " + string(e) + " requested to like his comment."
}