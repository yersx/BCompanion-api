package weather

import (
	"bcompanion/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type WeatherService interface {
	GetWeekWeather(place string) ([]*model.WeatherDay, error)
	GetDayWeather(place string, day string) ([]*model.WeatherHourResponse, error)
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

var owIconNames = map[string]string{
	"01d": "https://img.icons8.com/wired/64/000000/sun.png",
	"01n": "https://img.icons8.com/wired/64/000000/sun.png",
	"02d": "https://img.icons8.com/wired/64/000000/partly-cloudy-day.png",
	"02n": "https://img.icons8.com/wired/64/000000/partly-cloudy-night.png",
	"03d": "https://img.icons8.com/wired/64/000000/cloud.png",
	"03n": "https://img.icons8.com/wired/64/000000/cloud.png",
	"04d": "https://img.icons8.com/wired/64/000000/clouds.png",
	"04n": "https://img.icons8.com/wired/64/000000/clouds.png",
	"09d": "https://img.icons8.com/ios/50/000000/light-rain.png",
	"09n": "https://img.icons8.com/ios/50/000000/light-rain.png",
	"10d": "https://img.icons8.com/ios/50/000000/moderate-rain.png",
	"10n": "https://img.icons8.com/ios/50/000000/moderate-rain.png",
	"11d": "https://img.icons8.com/wired/64/000000/storm.png",
	"11n": "https://img.icons8.com/wired/64/000000/storm.png",
	"13d": "https://img.icons8.com/wired/64/000000/snow.png",
	"13n": "https://img.icons8.com/wired/64/000000/snow.png",
	"50d": "https://img.icons8.com/wired/64/000000/foggy-night-1.png",
	"50n": "https://img.icons8.com/wired/64/000000/foggy-night-1.png",
}

var owWeekDays = map[string]string{
	"Monday":    "Пн",
	"Tuesday":   "Вт",
	"Wednesday": "Ср",
	"Thursday":  "Чт",
	"Friday":    "Пт",
	"Saturday":  "Сб",
	"Sunday":    "Вс",
}

func NewWeatherService(repository WeatherRepository) WeatherService {
	weatherRepo = repository
	return &service{}
}

func (*service) GetWeekWeather(place string) ([]*model.WeatherDay, error) {
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
	query.Set("units", units)

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

	out := make([]*model.WeatherDay, len(we.Daily))

	for i, b := range we.Daily {
		dateTime := time.Unix(b.Date, 0)
		out[i] = &model.WeatherDay{
			PlaceName:   place,
			Day:         owWeekDays[dateTime.Weekday().String()],
			Date:        dateTime.Format("02.01.2006"),
			Image:       owIconNames[b.Weather[0].Icon],
			DayDegree:   fmt.Sprintf("%.1f", b.Temp.Day),
			NightDegree: fmt.Sprintf("%.1f", b.Temp.Night),
			Humidity:    strconv.Itoa(b.Humidity),
			Clouds:      strconv.Itoa(b.Clouds),
			WindSpeed:   fmt.Sprintf("%.1f", b.WindSpeed),
		}
	}

	return out, nil
}

func FloatToString(input_num float64) string {
	// convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

// func FloatToStringP1(input_num float32) string {
// 	// convert a float number to a string
// 	return strconv.FormatFloat(input_num, 'f', 1, 32)
// }

func (*service) GetDayWeather(place string, date string) ([]*model.WeatherHourResponse, error) {
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
	query.Set("units", units)
	query.Set("exclude", "current,daily")

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
	var we model.WeatherHour
	if err = json.NewDecoder(resp.Body).Decode(&we); err != nil {
		log.Println("decoder error %v", err)
		return nil, err
	}

	out := make([]*model.WeatherHourResponse, 24)

	for i, b := range we.Hourly {
		dateTime := time.Unix(b.Date, 0)
		if dateTime.Format("02.01.2006") == date {
			hours, minutes, _ := dateTime.Clock()
			timeInString := fmt.Sprintf("%d:%02d", hours, minutes)
			out[i] = &model.WeatherHourResponse{
				PlaceName:   place,
				Hour:        timeInString,
				Date:        dateTime.Format("02.01.2006"),
				Image:       owIconNames[b.Weather[0].Icon],
				Description: b.Weather[0].Description,
				Degree:      fmt.Sprintf("%.1f", b.Temp),
			}
		}
	}

	return out, nil
}
