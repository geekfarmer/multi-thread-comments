package service

import "github.com/geekfarmer/multi-thread-comments/entity"

type AppService interface {
	CreateApp(user *entity.App) (*entity.App, error)
	FindAllApps() ([]entity.App, error)
	GetAllRequestedApps() ([]entity.App, error)
	CreateApps(app []entity.App) ([]entity.App, error)
	UpdateAppsForUser(user *entity.User) (*entity.User, error)
	// FindByUserID(id string) (*entity.App, error)
	// Delete(*entity.App) error
}

type UserService interface {
	Create(user *entity.User) (*entity.User, error)
	FindAll() ([]entity.User, error)
	FindByToken(id string) (*entity.User, error)
	Delete(*entity.User) error
	UpdateUser(*entity.User) (*entity.User, error)
}
