package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sriharivishnu/shopify-challenge/external"
	models "github.com/sriharivishnu/shopify-challenge/models/api"
)

type WeatherLayer interface {
	FetchWeather(city string, lon, lat float32) (models.WeatherResponse, error)
}

type WeatherService struct {
	Cache      external.Cache
	HttpClient *http.Client
}

// Also caches the requests for a given city. In production, common cities could be pre-fetched
// to avoid latency in the future.
func (service *WeatherService) FetchWeather(city string, lon, lat float32) (models.WeatherResponse, error) {
	if weather := service.Cache.Get(city); weather != nil {
		return weather.(models.WeatherResponse), nil
	}
	appId := "dfc4083c5474f48352a4dfc1b72c771f"

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%.6f&lon=%.6f&appid=%s", lat, lon, appId)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return models.WeatherResponse{}, err
	}

	res, err := service.HttpClient.Do(req)
	if err != nil {
		return models.WeatherResponse{}, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return models.WeatherResponse{}, readErr
	}

	fmt.Println(string(body))

	weather := models.WeatherResponse{}
	jsonErr := json.Unmarshal(body, &weather)
	if jsonErr != nil {
		return models.WeatherResponse{}, jsonErr
	}

	service.Cache.Set(city, weather, 3600)
	return weather, jsonErr
}
