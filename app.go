package main

import (
	"github.com/geekfarmer/multi-thread-comments/controller"
	"github.com/geekfarmer/multi-thread-comments/data"
	router "github.com/geekfarmer/multi-thread-comments/http"
	"github.com/geekfarmer/multi-thread-comments/repository"
	"github.com/geekfarmer/multi-thread-comments/service"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	database          *mongo.Database              = data.Connect()
	userRepository    repository.UserRepository    = repository.NewUserMongoRepository(database)
	userService       service.UserService          = service.NewUserService(userRepository)
	userController    controller.UserController    = controller.NewUserController(userService)
	commentRepository repository.CommentRepository = repository.NewCommentMongoRepository(database)
	commentService    service.CommentService       = service.NewCommentService(commentRepository)
	commentControllet controller.CommentController = controller.NewCommentController(commentService)
	httpRouter        router.Router                = router.NewMuxRouter()
	initializer       router.Init                  = router.Initialize(userController, commentControllet, httpRouter)
)

func main() {
	initializer.Run()
}
