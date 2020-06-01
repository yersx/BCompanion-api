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
	"bcompanion/weather"
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

	weatherRepository weather.WeatherRepository = weather.NewMongoRepository()
	weatherService    weather.WeatherService    = weather.NewWeatherService(weatherRepository)
	weatherController weather.WeatherController = weather.NewWeatherController(weatherService)

	hikeRepository hike.HikeRepository = hike.NewMongoRepository()
	hikeService    hike.HikeService    = hike.NewHikeService(hikeRepository)
	hikeController hike.HikeController = hike.NewHikeController(hikeService)

	httpRouter router.Router = router.NewMuxRouter()
)

func main() {
	port := os.Getenv("PORT")
	log.Printf("port: " + port)

	httpRouter.POST("/users/authorize", userController.SignUser)
	httpRouter.GET("/users/getUser", userController.FindUser)
	httpRouter.GET("/users/getToken", userController.FindToken)
	httpRouter.GET("/users/getUserProfile", userController.FindUserProfile)
	httpRouter.POST("/users/updateUserPhoto", userController.UpdateUserPhoto)

	httpRouter.POST("/city/add", cityController.AddCity)
	httpRouter.GET("/cities", cityController.GetCities)
	httpRouter.GET("/cities/citiesName", cityController.GetCitiesName)
	httpRouter.GET("/cities/cityCoordinates", cityController.GetCityCoordinates)

	httpRouter.POST("/place/add", placeController.AddPlace)
	httpRouter.GET("/places/byCity", placeController.GetPlaces)
	httpRouter.GET("/places/placesName", placeController.GetPlacesName)

	httpRouter.POST("/placeDescription/add", placeController.AddPlaceDescription)
	httpRouter.GET("/placeDescription", placeController.GetPlaceDescription)

	httpRouter.GET("/placeRoute", placeController.GetPlaceRoute)
	httpRouter.GET("/placesRoutesByCity", placeController.GetPlacesRoutesByCity)

	httpRouter.POST("/groups/createGroup", groupController.AddGroup)
	httpRouter.GET("/groups/getByUser", groupController.GetUserGroups)
	httpRouter.GET("/groups/getByActivity", groupController.GetAllGroups)
	httpRouter.GET("/groups/getGroup", groupController.GetGroup)

	httpRouter.POST("/groups/joinGroup", groupController.JoinGroup)
	httpRouter.POST("/groups/leaveGroup", groupController.LeaveGroup)

	httpRouter.POST("/hikes/createHike", hikeController.AddHike)
	httpRouter.GET("/hikes/getHike", hikeController.GetHike)

	httpRouter.GET("/hikes/getHikes", hikeController.GetHikes)
	httpRouter.GET("/hikes/getUpcomingAll", hikeController.GetUpcomingHikes)
	httpRouter.GET("/hikes/getUpcomingByUser", hikeController.GetUpcomingHikesByUser)
	httpRouter.GET("/hikes/GetUpcomingByPlace", hikeController.GetUpcomingHikesByPlace)
	httpRouter.GET("/hikes/getPastByUser", hikeController.GetPastHikesByUser)

	httpRouter.POST("/hikes/joinHike", hikeController.JoinHike)
	httpRouter.POST("/hikes/leaveHike", hikeController.LeaveHike)

	httpRouter.GET("/weather/week", weatherController.GetWeekWeather)
	httpRouter.GET("/weather/day", weatherController.GetDayWeather)

	httpRouter.SERVE(port)
}
