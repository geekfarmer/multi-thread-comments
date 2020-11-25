package repository

import (
	"encoding/json"

	"github.com/geekfarmer/multi-thread-comments/entity"
)

type UserRepository interface {
	Create(user *entity.User) (*entity.User, error)
	FindAll() ([]entity.User, error)
	FindByID(id string) (*entity.User, error)
	Delete(id string) error
	UpdateUser(*entity.User) (*entity.User, error)
}

type CommentRepository interface {
	CreateComment(comment *entity.Comment, userID string) (*entity.Comment, error)
	FindAllComments() ([]entity.Comment, error)
	FindCommentByID(id string) (*entity.Comment, error)
	FindCommentsByPostID(id string) ([]entity.Comment, error)
	DeleteComment(id string) error
	UpdateComment(*entity.Comment) (*entity.Comment, error)
	UpdateChildComments(id string, comment *entity.Comment) (*entity.Comment, error)
}

func StructToMap(x interface{}) (newMap map[string]interface{}, err error) {
	data, err := json.Marshal(x)

	if err != nil {
		return
	}

	err = json.Unmarshal(data, &newMap) // Convert to a map
	return
}
