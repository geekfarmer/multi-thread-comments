package repository

import (
	"encoding/json"

	"github.com/geekfarmer/multi-thread-comments/entity"
)

type UserRepository interface {
	Create(user *entity.User) (*entity.User, error)
	FindAll() ([]entity.User, error)
	FindByToken(token string) (*entity.User, error)
	Delete(*entity.User) error
	UpdateUser(*entity.User) (*entity.User, error)
}

type AppRepository interface {
	CreateApp(app *entity.App) (*entity.App, error)
	FindAllApps() ([]entity.App, error)
	GetAllRequestedApps() ([]entity.App, error)
	CreateApps(app []entity.App) ([]entity.App, error)
	UpdateAppsForUser(user *entity.User) (*entity.User, error)

	// FindByUserID(id string) (*entity.App, error)
	// Delete(*entity.App) error
}

func StructToMap(x interface{}) (newMap map[string]interface{}, err error) {
	data, err := json.Marshal(x)

	if err != nil {
		return
	}

	err = json.Unmarshal(data, &newMap) // Convert to a map
	return
}
