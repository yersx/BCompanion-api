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
	place1 "bcompanion/place/place"
)

var (
	userRepository userrepository.UserRepository = userrepository.NewMongoRepository()
	userService    userservice.UserService       = userservice.NewUserService(userRepository)
	userController controller.UserController     = controller.NewUserController(userService)

	placeRepository place.PlaceRepository = place.NewMongoRepository()
	placeService    place.PlaceService    = place.NewPlaceService(placeRepository)
	cityController  place.CityController  = place.NewCityController(placeService)

	placeController place1.PlaceController = place1.NewPlaceController(placeService)

	httpRouter router.Router = router.NewMuxRouter()
)

func main() {
	port := os.Getenv("PORT")
	log.Printf("port: " + port)

	httpRouter.POST("/users/authorize", userController.SignUser)
	httpRouter.GET("/users/getUser", userController.FindUser)

	httpRouter.POST("/city/add", cityController.AddCity)
	httpRouter.GET("/cities", cityController.GetCities)

	httpRouter.POST("/place/add", placeController.AddPlace)
	httpRouter.GET("/places/byCity", placeController.GetPlaces)

	httpRouter.SERVE(port)
}
