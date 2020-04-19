// main
package main

import (
	"log"
	"os"

	router "bcompanion/http"
	userservice "bcompanion/user"
	"bcompanion/user/controller"
	userrepository "bcompanion/user/repository"

	place "bcompanion/place"
)

var (
	userRepository userrepository.UserRepository = userrepository.NewMongoRepository()
	userService    userservice.UserService       = userservice.NewUserService(userRepository)
	userController controller.UserController     = controller.NewUserController(userService)

	placeRepository place.PlaceRepository = place.NewMongoRepository()
	placeService    place.PlaceService    = place.NewPlaceService(placeRepository)
	placeController place.PlaceController = place.NewPlaceController(placeService)

	httpRouter router.Router = router.NewMuxRouter()
)

func main() {
	port := os.Getenv("PORT")
	log.Printf("port: " + port)

	httpRouter.POST("/users/authorize", userController.SignUser)
	httpRouter.GET("/users/getUser", userController.FindUser)

	httpRouter.POST("/city/add", placeController.AddCity)
	httpRouter.GET("/cities", placeController.GetCities)

	httpRouter.SERVE(port)
}
