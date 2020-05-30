// main
package main

import (
	"log"
	"os"

	router "bcompanion/http"
	userservice "bcompanion/user"
	"bcompanion/user/controller"
	userrepository "bcompanion/user/repository"

	"bcompanion/group"
	"bcompanion/hike"
	"bcompanion/place"
	place1 "bcompanion/place/place"
)

var (
	userRepository userrepository.UserRepository = userrepository.NewMongoRepository()
	userService    userservice.UserService       = userservice.NewUserService(userRepository)
	userController controller.UserController     = controller.NewUserController(userService)

	placeRepository place.PlaceRepository  = place.NewMongoRepository()
	placeService    place.PlaceService     = place.NewPlaceService(placeRepository)
	cityController  place.CityController   = place.NewCityController(placeService)
	placeController place1.PlaceController = place1.NewPlaceController(placeService)

	groupRepository group.GroupRepository = group.NewMongoRepository()
	groupService    group.GroupService    = group.NewGroupService(groupRepository)
	groupController group.GroupController = group.NewGroupController(groupService)

	hikeRepository hike.HikeRepository = hike.NewMongoRepository()
	hikeService    hike.HikeService    = hike.NewHikeService(hikeRepository)
	hikeController hike.HikeController = hike.NewHikeController(hikeService)

	httpRouter router.Router = router.NewMuxRouter()
)

func main() {
	port := os.Getenv("PORT")
	log.Printf("port: " + port)

	httpRouter.POST("/users/authorize", port, userController.SignUser)
	httpRouter.GET("/users/getUser", port, userController.FindUser)
	httpRouter.GET("/users/getToken", port, userController.FindToken)

	httpRouter.POST("/city/add", port, cityController.AddCity)
	httpRouter.GET("/cities", port, cityController.GetCities)
	httpRouter.GET("/cities/citiesName", port, cityController.GetCitiesName)
	httpRouter.GET("/cities/cityCoordinates", port, cityController.GetCityCoordinates)

	httpRouter.POST("/place/add", port, placeController.AddPlace)
	httpRouter.GET("/places/byCity", port, placeController.GetPlaces)
	httpRouter.GET("/places/placesName", port, placeController.GetPlacesName)

	httpRouter.POST("/placeDescription/add", port, placeController.AddPlaceDescription)
	httpRouter.GET("/placeDescription", port, placeController.GetPlaceDescription)

	httpRouter.GET("/placeRoute", port, placeController.GetPlaceRoute)
	httpRouter.GET("/placesRoutesByCity", port, placeController.GetPlacesRoutesByCity)

	httpRouter.POST("/groups/createGroup", port, groupController.AddGroup)
	httpRouter.GET("/groups/getByUser", port, groupController.GetUserGroups)
	httpRouter.GET("/groups/getByActivity", port, groupController.GetAllGroups)
	httpRouter.GET("/groups/getGroup", port, groupController.GetGroup)

	httpRouter.POST("/groups/joinGroup", port, groupController.JoinGroup)
	httpRouter.POST("/groups/leaveGroup", port, groupController.LeaveGroup)

	httpRouter.POST("/hikes/createHike", port, hikeController.AddHike)
	httpRouter.GET("/hikes/getHike", port, hikeController.GetHike)

	httpRouter.GET("/hikes/getHikes", port, hikeController.GetHikes)
	httpRouter.GET("/hikes/getUpcomingAll", port, hikeController.GetUpcomingHikes)
	httpRouter.GET("/hikes/getUpcomingByUser", port, hikeController.GetUpcomingHikesByUser)
	httpRouter.GET("/hikes/GetUpcomingByPlace", port, hikeController.GetUpcomingHikesByPlace)
	httpRouter.GET("/hikes/getPastByUser", port, hikeController.GetPastHikesByUser)

	httpRouter.POST("/hikes/joinHike", port, hikeController.JoinHike)
	httpRouter.POST("/hikes/leaveHike", port, hikeController.LeaveHike)

}
