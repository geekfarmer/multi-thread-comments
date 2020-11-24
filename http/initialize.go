package router

import "github.com/geekfarmer/multi-thread-comments/controller"

type initialize struct {
	userController     controller.UserController
	appController      controller.AppController
	currencyController controller.CurrencyController
	httpRouter         Router
}

type Init interface {
	Run()
}

func Initialize(userController controller.UserController, appController controller.AppController, currencyController controller.CurrencyController, httpRouter Router) Init {
	return &initialize{userController, appController, currencyController, httpRouter}
}

func (i *initialize) Run() {
	//User controller
	i.httpRouter.POST("/user", i.userController.CreateUser)
	i.httpRouter.GET("/users", i.userController.FindAll)
	i.httpRouter.PUT("/user", i.userController.UpdateUser)
	i.httpRouter.GET("/user/{id}", i.userController.FindAll)

	//Currency routes
	i.httpRouter.GET("/currency", i.currencyController.GetCurrency)

	//App routes
	i.httpRouter.GET("/apps", i.appController.FindAllApps)
	i.httpRouter.POST("/app", i.appController.CreateApp)
	i.httpRouter.POST("/request-app", i.appController.RequestNewApp)
	i.httpRouter.GET("/request-app", i.appController.GetAllRequestedApps)
	i.httpRouter.POST("/apps", i.appController.CreateApps)
	i.httpRouter.PUT("/user/sync", i.appController.SyncInstalledApps)
	i.httpRouter.PUT("/user/app", i.appController.AddInstalledApp)

	i.httpRouter.SERVE("2000")
}
