package service

import (
	"errors"
	"time"

	"github.com/geekfarmer/multi-thread-comments/entity"
	"github.com/geekfarmer/multi-thread-comments/repository"
	sid "github.com/lithammer/shortuuid"
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
	user.ID = sid.New()
	user.CreatedAt = time.Now().UTC()
	user.UpdatedAt = time.Now().UTC()
	return repo.Create(user)
}

func (*service) FindAll() ([]entity.User, error) {
	return repo.FindAll()
}

func (*service) Delete(id string) error {
	return repo.Delete(id)
}

func (*service) FindByID(id string) (*entity.User, error) {
	user, _ := repo.FindByID(id)
	return user, nil
}

func (*service) UpdateUser(user *entity.User) (*entity.User, error) {
	u, _ := repo.UpdateUser(user)
	return u, nil
}
