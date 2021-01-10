package usecase

import (
	"errors"
	"github.com/khu-dev/khumu-comment/config"
	"github.com/khu-dev/khumu-comment/repository"
	"github.com/sirupsen/logrus"
	"time"
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
	newComment, err := uc.Repository.Create(comment)
	if err!=nil{return newComment, err}
	return newComment, nil
}

func (uc *CommentUseCase) List(username string, opt *repository.CommentQueryOption) ([]*model.Comment, error) {
	comments := uc.Repository.List(opt)
	parents := uc.listParentWithChildren(comments)

	for _, p := range parents {
		uc.handleComment(p, username, 0)
	}

	return parents, nil
}

func (uc *CommentUseCase) Get(id int) (*model.Comment, error) {
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
	comment, err := uc.Repository.Update(id, map[string]interface{}{
		"state": "deleted",
		"content": "삭제된 댓글입니다.",
	})
	if err != nil { return nil, err}

	uc.hideAuthor(comment, "") // ""는 어떠한 author username과도 다르기때문에 숨겨진다.
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

// 대부분의 comment usecase에서 사용되는 로직을 담당한다. 재귀적으로 자식 코멘트들에게도 적용된다.
func (uc *CommentUseCase) handleComment(c *model.Comment, username string, currentDepth int){
	const maxDepth = 1
	uc.hideAuthor(c, username)
	likeCount := uc.getLikeCommentCount(c.ID)
	c.LikeCommentCount = likeCount
	likes := uc.LikeCommentRepository.List(&repository.LikeCommentQueryOption{CommentID: c.ID, Username: username})
	if len(likes) >= 1{
		c.Liked = true
	}
	uc.setCreatedAtExpression(c)
	if currentDepth < maxDepth{
		for _, child := range c.Children{
			uc.handleComment(child, username, currentDepth + 1)
		}
	}
}

// username이 author의 username과 일치하면 hide
// 그냥 무조건 hide 하고싶다면 username을 ""으로 전달
func (uc *CommentUseCase) hideAuthor(c *model.Comment, username string) {
	if c.State == "deleted" {
		c.AuthorUsername = "삭제된 댓글의 작성자"
		c.Author.Username = "삭제된 댓글의 작성자"
		c.Author.Nickname = "삭제된 댓글의 작성자"
	} else if c.Kind == "anonymous" && c.AuthorUsername != username {
		c.AuthorUsername = "익명"
		c.Author.Username = "익명"
		c.Author.Nickname = "익명"
	}
}

// Comment.CreatedAt을 바탕으로 Comment.CreatedAtExpression에 올바른 값을 입력시킨다.
func (uc *CommentUseCase) setCreatedAtExpression(c *model.Comment){
	// UTC 시간을 단순 한국시간으로 변경
	now := time.Now().In(config.Location) // now는 근데 기본적으로 UTC긴한듯.
	nowYear, nowMonth, nowDate := now.Date()
	//log.Println(c.CreatedAt.In(repository.Location).Format("2006/01/02 15:04")) // => 한국시간대로 잘 나옴.
	createdAt := c.CreatedAt.In(config.Location)
	createdYear, createdMonth, createdDate := createdAt.Date()
	if now.Sub(c.CreatedAt).Minutes() < 5{
		c.CreatedAtExpression = "지금"
	} else if nowYear == createdYear && nowMonth == createdMonth && nowDate == createdDate{
		c.CreatedAtExpression = createdAt.Format("15:04")
	} else if nowYear == createdYear{
		c.CreatedAtExpression = createdAt.Format("01/02 15:04")
	} else{
		c.CreatedAtExpression = createdAt.Format("2006/01/02 15:04")
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
	logger := logrus.WithField("CommentID", like.CommentID)
	logger.Debug("Toggle LikeComment")
	likes := uc.Repository.List(&repository.LikeCommentQueryOption{CommentID: like.CommentID, Username: like.Username})
	// 길이가 1보다 크거나 같으면 삭제. 1인 경우는 정상적으로 하나만 있을 때,
	// 1보다 큰 경우는 비정상적으로 여러개 존재할 때
	if len(likes) >= 1 {
		for _, like := range likes {
			err = uc.Repository.Delete(like.ID)
			if err != nil {
				logger.Panic(false, err)
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
