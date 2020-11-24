package service

import (
	"errors"

	"github.com/geekfarmer/multi-thread-comments/entity"
	"github.com/geekfarmer/multi-thread-comments/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type service struct{}

var repo repository.UserRepository

func NewUserService(repository repository.UserRepository) UserService {
	repo = repository
	return &service{}
}

func (*service) Validate(user *entity.User) error {

	if user == nil {
		err := errors.New("The User is empty")
		return err
	}

	return nil
}

func (*service) Create(user *entity.User) (*entity.User, error) {
	user.ID = primitive.NewObjectID()
	return repo.Create(user)
}

func (*service) FindAll() ([]entity.User, error) {
	return repo.FindAll()
}

func (*service) Delete(user *entity.User) error {
	return repo.Delete(user)
}

func (*service) FindByToken(token string) (*entity.User, error) {
	user, _ := repo.FindByToken(token)
	return user, nil
}
func (*service) UpdateUser(user *entity.User) (*entity.User, error) {
	u, _ := repo.UpdateUser(user)
	return u, nil
}
