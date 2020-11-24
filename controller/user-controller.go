package controller

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/geekfarmer/multi-thread-comments/entity"
	"github.com/geekfarmer/multi-thread-comments/errors"
	"github.com/geekfarmer/multi-thread-comments/service"
)

type controller struct {
}

var userService service.UserService

type UserController interface {
	CreateUser(response http.ResponseWriter, request *http.Request)
	FindAll(response http.ResponseWriter, request *http.Request)
	UpdateUser(response http.ResponseWriter, request *http.Request)
}

func NewUserController(service service.UserService) UserController {
	userService = service
	return &controller{}
}

func (*controller) CreateUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var user *entity.User
	_ = json.NewDecoder(request.Body).Decode(&user)

	if len(user.UserToken) == 0 {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(errors.GenericError{Message: "Token can't be null"})
		return
	}
	if len(user.Phone) == 0 {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(errors.GenericError{Message: "Number can't be null"})
		return
	}

	user, err := userService.Create(user)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(errors.GenericError{Message: "Error getting the users"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(user)
}

func (*controller) FindAll(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var user []entity.User
	_ = json.NewDecoder(request.Body).Decode(&user)
	user, err := userService.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.GenericError{Message: "Error getting the users"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(user)
}

func (*controller) UpdateUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	user, err := getUserD(response, request)

	if err == "err" {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.GenericError{Message: "User not found"})
		return
	}

	_ = json.NewDecoder(request.Body).Decode(&user)
	user.UpdatedAt = time.Now()
	user, e := userService.UpdateUser(user)

	if e != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.GenericError{Message: "Something went wrong"})
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(user)
}

func getUserD(response http.ResponseWriter, request *http.Request) (*entity.User, string) {
	authToken := request.Header.Get("Authorization")
	splitAuthToken := strings.Split(authToken, " ")
	token := splitAuthToken[1]
	user, _ := uService.FindByToken(token)
	return user, "-"
}
