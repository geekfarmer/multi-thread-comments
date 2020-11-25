package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/geekfarmer/multi-thread-comments/entity"
	"github.com/geekfarmer/multi-thread-comments/errors"
	"github.com/geekfarmer/multi-thread-comments/service"
	"github.com/gorilla/mux"
)

type controller struct {
}

var userService service.UserService

type UserController interface {
	CreateUser(response http.ResponseWriter, request *http.Request)
	FindAll(response http.ResponseWriter, request *http.Request)
	GetUser(response http.ResponseWriter, request *http.Request)
	DeleteUser(response http.ResponseWriter, request *http.Request)
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

	user, err := userService.Create(user)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(errors.GenericError{Message: `Error creating user.`})
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

func (*controller) GetUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	id := vars["id"]
	user, err := userService.FindByID(id)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.GenericError{Message: "Error getting the user"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(user)
}

func (*controller) DeleteUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	id := vars["id"]
	err := userService.Delete(id)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.GenericError{Message: "Error deleting user"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(nil)
}

func (*controller) UpdateUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	id := vars["id"]
	user, err := userService.FindByID(id)

	if err != nil {
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
