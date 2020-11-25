package service

import (
	"errors"
	"time"

	"github.com/geekfarmer/multi-thread-comments/entity"
	"github.com/geekfarmer/multi-thread-comments/repository"
	sid "github.com/lithammer/shortuuid"
)

type commentService struct{}

var commentRepo repository.CommentRepository

func NewCommentService(repository repository.CommentRepository) CommentService {
	commentRepo = repository
	return &service{}
}

func (*service) ValidateComment(comment *entity.Comment) error {

	if comment == nil {
		err := errors.New("The comment is empty")
		return err
	}

	return nil
}

func (*service) CreateComment(comment *entity.Comment, userID string) (*entity.Comment, error) {
	comment.ID = sid.New()
	comment.CreatedAt = time.Now().UTC()
	comment.UpdatedAt = time.Now().UTC()
	return commentRepo.CreateComment(comment, userID)
}

func (*service) FindAllComments() ([]entity.Comment, error) {
	return commentRepo.FindAllComments()
}

func (*service) DeleteComment(id string) error {
	return commentRepo.DeleteComment(id)
}

func (*service) UpdateComment(comment *entity.Comment) (*entity.Comment, error) {
	u, _ := commentRepo.UpdateComment(comment)
	return u, nil
}

func (*service) UpdateChildComments(id string, comment *entity.Comment) (*entity.Comment, error) {
	u, _ := commentRepo.UpdateChildComments(id, comment)
	return u, nil
}

func (*service) FindCommentByID(id string) (*entity.Comment, error) {
	comment, _ := commentRepo.FindCommentByID(id)
	return comment, nil
}

func (*service) FindCommentsByPostID(id string) ([]entity.Comment, error) {
	return commentRepo.FindCommentsByPostID(id)
}
