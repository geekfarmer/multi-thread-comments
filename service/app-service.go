package service

import (
	"github.com/geekfarmer/multi-thread-comments/entity"
	"github.com/geekfarmer/multi-thread-comments/repository"
)

type appService struct{}

var appRepo repository.AppRepository

func NewAppService(repository repository.AppRepository) AppService {
	appRepo = repository
	return &appService{}
}

func (*appService) FindAllApps() ([]entity.App, error) {
	return appRepo.FindAllApps()
}

func (*appService) GetAllRequestedApps() ([]entity.App, error) {
	return appRepo.GetAllRequestedApps()
}

func (*appService) CreateApp(apps *entity.App) (*entity.App, error) {
	return appRepo.CreateApp(apps)
}

func (r *appService) CreateApps(apps []entity.App) ([]entity.App, error) {
	return appRepo.CreateApps(apps)
}

func (r *appService) UpdateAppsForUser(user *entity.User) (*entity.User, error) {
	return appRepo.UpdateAppsForUser(user)
}
