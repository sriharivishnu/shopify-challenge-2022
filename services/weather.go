package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sriharivishnu/shopify-challenge/config"
	"github.com/sriharivishnu/shopify-challenge/external"
	models "github.com/sriharivishnu/shopify-challenge/models/api"
)

type WeatherLayer interface {
	FetchWeather(city string) (models.WeatherResponse, error)
}

type WeatherService struct {
	Cache      external.Cache
	HttpClient *http.Client
}

type CityMap map[string]struct {
	lon float32
	lat float32
}

// can implement this function to make another API call (with caching)
// to support any cities
func (service *WeatherService) GetLonLat(city string) (float32, float32) {
	mapping := CityMap{
		"Toronto":   {lon: -79.38, lat: 43.65},
		"Montreal":  {lon: -73.57, lat: 45.50},
		"Vancouver": {lon: -123.12, lat: 49.28},
		"Calgary":   {lon: -114.05, lat: 51.05},
		"Waterloo":  {lon: -80.47, lat: 43.47},
	}

	return mapping[city].lon, mapping[city].lat
}

// Also caches the requests for a given city. In production, common cities could be pre-fetched
// to avoid latency in the future.
func (service *WeatherService) FetchWeather(city string) (models.WeatherResponse, error) {
	if weather := service.Cache.Get(city); weather != nil {
		return weather.(models.WeatherResponse), nil
	}
	appId := config.Config.WEATHER_API_KEY

	lon, lat := service.GetLonLat(city)

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

	weather := models.WeatherResponse{}
	jsonErr := json.Unmarshal(body, &weather)
	if jsonErr != nil {
		return models.WeatherResponse{}, jsonErr
	}

	service.Cache.Set(city, weather, 3600)
	return weather, jsonErr
}
