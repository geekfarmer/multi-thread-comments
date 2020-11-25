package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/geekfarmer/multi-thread-comments/entity"
	"github.com/geekfarmer/multi-thread-comments/errors"
	"github.com/geekfarmer/multi-thread-comments/service"
	"github.com/gorilla/mux"
)

type commentController struct {
}

var commentService service.CommentService

type CommentController interface {
	CreateComment(response http.ResponseWriter, request *http.Request)
	FindAllComments(response http.ResponseWriter, request *http.Request)
	UpdateComment(response http.ResponseWriter, request *http.Request)
	GetCommentByPost(response http.ResponseWriter, request *http.Request)
	DeleteComment(response http.ResponseWriter, request *http.Request)
}

func NewCommentController(service service.CommentService) CommentController {
	commentService = service
	return &controller{}
}

func (*controller) CreateComment(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var comment *entity.Comment
	var parentComment *entity.Comment
	vars := mux.Vars(request)
	userID := vars["userID"]
	_ = json.NewDecoder(request.Body).Decode(&comment)
	comment.ChildComments = make([]string, 0)
	comment, err := commentService.CreateComment(comment, userID)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(errors.GenericError{Message: `Error creating comment.`})
		return
	}
	if comment.ParentID != "" {
		parentComment, err = commentService.UpdateChildComments(comment.ParentID, comment)
		fmt.Println("Parent Comment", parentComment)
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(comment)
}

func (*controller) FindAllComments(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var comments []entity.Comment
	_ = json.NewDecoder(request.Body).Decode(&comments)
	comments, err := commentService.FindAllComments()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.GenericError{Message: "Error getting the comments"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(comments)
}

func (*controller) UpdateComment(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	id := vars["id"]
	comment, err := commentService.FindCommentByID(id)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.GenericError{Message: "User not found"})
		return
	}

	_ = json.NewDecoder(request.Body).Decode(&comment)
	comment.UpdatedAt = time.Now()
	comment, e := commentService.UpdateComment(comment)

	if e != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.GenericError{Message: "Something went wrong"})
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(comment)
}

func (*controller) GetCommentByPost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	postID := vars["postID"]
	var comments []entity.Comment
	_ = json.NewDecoder(request.Body).Decode(&comments)
	comments, err := commentService.FindCommentsByPostID(postID)
	fmt.Println(err)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.GenericError{Message: "Error getting the comments"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(comments)
}

func (*controller) DeleteComment(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	id := vars["id"]
	err := commentService.DeleteComment(id)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.GenericError{Message: "Error deleting comment"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(nil)
}
