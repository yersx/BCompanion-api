package weather

import (
	"bcompanion/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type WeatherService interface {
	GetWeekWeather(place string) (*model.Weather, error)
}

type service struct{}

var (
	weatherRepo WeatherRepository
)
var client = &http.Client{}

const (
	units  = "metric"
	apiKey = "0b8e291e166948ee62b043b7d37a9fa7"
	lang   = "ru"
	owAPI  = "https://api.openweathermap.org/data/2.5/onecall"
)

func NewWeatherService(repository WeatherRepository) WeatherService {
	weatherRepo = repository
	return &service{}
}

func (*service) GetWeekWeather(place string) (*model.Weather, error) {
	w, err := weatherRepo.GetPlaceCoordinates(place)
	if err != nil {
		return nil, err
	}
	lat := FloatToString(*w.Lattitude)
	long := FloatToString(*w.Longitude)

	query := url.Values{}
	query.Set("lat", lat)
	query.Set("lon", long)
	query.Set("appid", apiKey)
	query.Set("lang", lang)

	url := fmt.Sprintf("%s?%s", owAPI, query.Encode())

	log.Println("getting URL %s", url)

	var request *http.Request
	if request, err = http.NewRequest("GET", url, nil); err != nil {
		log.Println("request error %v", err)
		return nil, err
	}

	var resp *http.Response
	if resp, err = client.Do(request); err != nil {
		log.Println("response error %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, nil
	}

	log.Println("respo Body: %v", resp.Body)
	var we model.Weather
	if err = json.NewDecoder(resp.Body).Decode(&we); err != nil {
		log.Println("decoder error %v", err)
		return nil, err
	}
	return &we, nil
}

func FloatToString(input_num float64) string {
	// convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}
