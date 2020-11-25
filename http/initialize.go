package router

import "github.com/geekfarmer/multi-thread-comments/controller"

type initialize struct {
	userController    controller.UserController
	commentController controller.CommentController
	httpRouter        Router
}

type Init interface {
	Run()
}

func Initialize(userController controller.UserController, commentController controller.CommentController, httpRouter Router) Init {
	return &initialize{userController, commentController, httpRouter}
}

func (i *initialize) Run() {
	//User controller
	i.httpRouter.POST("/user", i.userController.CreateUser)
	i.httpRouter.GET("/users", i.userController.FindAll)
	i.httpRouter.PUT("/user/{id}", i.userController.UpdateUser)
	i.httpRouter.GET("/user/{id}", i.userController.GetUser)
	i.httpRouter.DELETE("/user/{id}", i.userController.DeleteUser)

	//Comment controller
	i.httpRouter.POST("/comment/{userID}", i.commentController.CreateComment)
	i.httpRouter.PUT("/comment/{id}", i.commentController.UpdateComment)
	i.httpRouter.GET("/comment/{postID}", i.commentController.GetCommentByPost)
	i.httpRouter.DELETE("/comment/{id}", i.commentController.DeleteComment)

	i.httpRouter.SERVE("2000")
}
