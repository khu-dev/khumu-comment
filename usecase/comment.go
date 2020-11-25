package usecase

import (
	"errors"
	"github.com/khu-dev/khumu-comment/repository"
	"log"
)
import "github.com/khu-dev/khumu-comment/model"

type CommentUseCaseInterface interface {
	Create(comment *model.Comment) (*model.Comment, error)
	List(username string, opt *repository.CommentQueryOption) ([]*model.Comment, error)
	Get(id int) (*model.Comment, error)
	Update(id int, opt map[string]interface{}) (*model.Comment, error)
	Delete(id int) (*model.Comment, error)
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
	comments := uc.Repository.List(opt)
	parents := uc.listParentWithChildren(comments)

	for _, p := range parents {
		for _, c := range p.Children {
			uc.handleComment(c, username)
		}
		uc.handleComment(p, username)
	}

	return parents, nil
}

func (uc *CommentUseCase) Get(id int) (*model.Comment, error) {
	log.Println("CommentUseCase_Get")
	comment, err := uc.Repository.Get(id)
	if err != nil { return nil, err}

	uc.hideAuthor(comment, "") // ""는 어떠한 author username과도 다르기때문에 숨겨진다.
	return comment, nil
}

func (uc *CommentUseCase) Update(id int, opt map[string]interface{}) (*model.Comment, error) {
	updated, err := uc.Repository.Update(id, opt)
	uc.hideAuthor(updated, "")
	return updated, err
}

// 실제로 Delete 하지는 않고 State를 "deleted"로 변경
func (uc *CommentUseCase) Delete(id int) (*model.Comment, error) {
	log.Println("CommentUseCase_Delete")
	comment, err := uc.Repository.Update(id, map[string]interface{}{
		"state": "deleted",
	})
	if err != nil { return nil, err}

	uc.hideAuthor(comment, "") // ""는 어떠한 author username과도 다르기때문에 숨겨진다.
	return comment, nil
}


func (uc *CommentUseCase) listParentWithChildren(allComments []*model.Comment) []*model.Comment {
	var parents []*model.Comment

	for _, comment := range allComments {
		if comment.ParentID == 0 {
			parents = append(parents, comment)
		}
	}

	return parents
}

// 대부분의 comment usecase에서 사용되는 로직을 담당한다.
func (uc *CommentUseCase) handleComment(c *model.Comment, username string){
	uc.hideAuthor(c, username)
	likeCount := uc.getLikeCommentCount(c.ID)
	c.LikeCommentCount = likeCount
	likes := uc.LikeCommentRepository.List(&repository.LikeCommentQueryOption{CommentID: c.ID, Username: username})
	if len(likes) == 1{
		c.Liked = true
	} else if len(likes) > 1{
		log.Print("[ERROR] ", c, "에 대한", username, "의 like가 1개 이상")
	}
}

// username이 author의 username과 일치하면 hide
// 그냥 무조건 hide 하고싶다면 username을 ""으로 전달
func (uc *CommentUseCase) hideAuthor(c *model.Comment, username string) {
	if c.State == "deleted" {
		c.AuthorUsername = "삭제된 댓글의 작성자"
		c.Author.Username = "삭제된 댓글의 작성자"
	} else if c.Kind == "anonymous" && c.AuthorUsername != username {
		c.AuthorUsername = "익명"
		c.Author.Username = "익명"
	}
}

func (uc *CommentUseCase) getLikeCommentCount(commentID int) int{
	likes := uc.LikeCommentRepository.List(&repository.LikeCommentQueryOption{CommentID: commentID})
	return len(likes)
}

func NewLikeCommentUseCase(
	likeRepo repository.LikeCommentRepositoryInterface,
	commentRepo repository.CommentRepositoryInterface) LikeCommentUseCaseInterface{
	return &LikeCommentUseCase{Repository: likeRepo, CommentRepository: commentRepo}
}

func (uc *LikeCommentUseCase) Toggle(like *model.LikeComment) (bool, error){
	var err error
	likes := uc.Repository.List(&repository.LikeCommentQueryOption{CommentID: like.CommentID, Username: like.Username})
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
		comment, err := uc.CommentRepository.Get(like.CommentID)
		if err != nil { return false, err}
		if comment.AuthorUsername == like.Username{
			return false, errors.New("Error: " + like.Username + " requested to like his comment.")
		}
		_, err = uc.Repository.Create(like)
		if err != nil{return false, err}

		return true, err
	}
}
