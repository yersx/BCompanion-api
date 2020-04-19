// main
package main

import (
	"os"

	router "bcompanion/http"
	userservice "bcompanion/user"
	"bcompanion/user/controller"
	userrepository "bcompanion/user/repository"
)

var (
	userRepository userrepository.UserRepository = userrepository.NewMongoRepository()
	userService    userservice.UserService       = userservice.NewUserService(userRepository)
	userController controller.UserController     = controller.NewUserController(userService)
	httpRouter     router.Router                 = router.NewMuxRouter()
)

func main() {
	port := os.Getenv("PORT")

	httpRouter.POST("/users/authorize", userController.SignUser)
	httpRouter.GET("/users/{phone}", userController.FindUser)

	httpRouter.SERVE(port)
}
