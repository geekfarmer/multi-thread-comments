package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/geekfarmer/multi-thread-comments/entity"
	"github.com/geekfarmer/multi-thread-comments/errors"
	"github.com/geekfarmer/multi-thread-comments/service"
)

var appService service.AppService
var uService service.UserService

type appController struct {
}
type AppController interface {
	FindAllApps(response http.ResponseWriter, request *http.Request)
	GetAllRequestedApps(response http.ResponseWriter, request *http.Request)
	CreateApp(response http.ResponseWriter, request *http.Request)
	CreateApps(response http.ResponseWriter, request *http.Request)
	SyncInstalledApps(response http.ResponseWriter, request *http.Request)
	AddInstalledApp(response http.ResponseWriter, request *http.Request)
	RequestNewApp(response http.ResponseWriter, request *http.Request)
}

//NewAppController creates an object for the same.
func NewAppController(service service.AppService, u service.UserService) AppController {
	appService = service
	uService = u
	return &appController{}
}

//Find all apps
func (*appController) FindAllApps(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var app []entity.App
	_ = json.NewDecoder(request.Body).Decode(&app)
	app, err := appService.FindAllApps()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.GenericError{Message: "Error getting the apps"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(app)
}

//Find all apps
func (*appController) GetAllRequestedApps(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var app []entity.App
	_ = json.NewDecoder(request.Body).Decode(&app)
	app, err := appService.GetAllRequestedApps()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.GenericError{Message: "Error getting the apps"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(app)
}

//Create individual app
func (*appController) CreateApp(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var app *entity.App
	_ = json.NewDecoder(request.Body).Decode(&app)
	app.ID = primitive.NewObjectID()
	app.CreatedAt = time.Now()
	app.UpdatedAt = time.Now()
	app, err := appService.CreateApp(app)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.GenericError{Message: "Error getting the apps"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(app)
}

//Create multiple apps by passing multiple apps
func (*appController) CreateApps(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var apps []entity.App
	fmt.Println(apps)
	_ = json.NewDecoder(request.Body).Decode(&apps)
	app, err := appService.CreateApps(apps)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.GenericError{Message: "Error getting the apps"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(app)
}

func (*appController) SyncInstalledApps(response http.ResponseWriter, request *http.Request) {
	user, err := getUser(request)
	if err != nil {
		response.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(response).Encode(errors.GenericError{Message: "User not found"})
		return
	}
	var apps []entity.InstalledApps
	_ = json.NewDecoder(request.Body).Decode(&apps)
	user.Apps = apps
	u, err := appService.UpdateAppsForUser(user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.GenericError{Message: "Error updating the apps"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(u)

}

func (*appController) AddInstalledApp(response http.ResponseWriter, request *http.Request) {
	user, err := getUser(request)

	if err != nil {
		response.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(response).Encode(errors.GenericError{Message: "User not found"})
		return
	}

	var alreadyPresentApps = user.Apps
	var apps entity.InstalledApps

	_ = json.NewDecoder(request.Body).Decode(&apps)
	if len(apps.ID) == 0 {
		alreadyPresentApps = append(alreadyPresentApps, apps)
		user.Apps = alreadyPresentApps
	} else {
		response.WriteHeader(http.StatusForbidden)
		json.NewEncoder(response).Encode(errors.GenericError{Message: "Please enter an App"})
		return
	}

	u, err := appService.UpdateAppsForUser(user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.GenericError{Message: "Error updating the app"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(u)
}

func getUser(request *http.Request) (*entity.User, error) {
	authToken := request.Header.Get("Authorization")
	splitAuthToken := strings.Split(authToken, " ")
	token := splitAuthToken[1]
	user, err := uService.FindByToken(token)
	return user, err
}

func (*appController) RequestNewApp(response http.ResponseWriter, request *http.Request) {
	user, err := getUser(request)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(errors.GenericError{Message: "User not found"})
		return
	}
	var app entity.App
	_ = json.NewDecoder(request.Body).Decode(&app)

	if len(app.Name) == 0 {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(errors.GenericError{Message: "Name cannot be empty"})
		return
	}
	if len(app.AppPackageName) == 0 {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(errors.GenericError{Message: "package name cannot be empty"})
		return
	}
	if len(app.Description) == 0 {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(errors.GenericError{Message: "description cannot be empty"})
		return
	}

	app.ID = primitive.NewObjectID()
	app.CreatedAt = time.Now()
	app.UpdatedAt = time.Now()
	app.AddedBy = user.Phone
	app.ActiveStatus = "pending"
	appService.CreateApp(&app)
}
