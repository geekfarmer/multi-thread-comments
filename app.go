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
	database           *mongo.Database               = data.Connect()
	userRepository     repository.UserRepository     = repository.NewUserMongoRepository(database)
	appRepository      repository.AppRepository      = repository.NewAppMongoRepository(database)
	userService        service.UserService           = service.NewUserService(userRepository)
	appService         service.AppService            = service.NewAppService(appRepository)
	userController     controller.UserController     = controller.NewUserController(userService)
	appController      controller.AppController      = controller.NewAppController(appService, userService)
	currencyController controller.CurrencyController = controller.NewCurrencyController()
	httpRouter         router.Router                 = router.NewMuxRouter()
	initializer        router.Init                   = router.Initialize(userController, appController, currencyController, httpRouter)
)

func main() {
	initializer.Run()
}
