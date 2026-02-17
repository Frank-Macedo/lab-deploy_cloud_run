package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/Frank-Macedo/lab-forecast/internal/domain/entities"
)

const WeatherURL = "http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no"

type WeatherClient struct {
}

func NewWeatherClient() *WeatherClient {
	return &WeatherClient{}
}

func (wc *WeatherClient) GetTemperature(address entities.Address) (*entities.Weather, error) {

	key := os.Getenv("API_KEY")

	url := fmt.Sprintf(WeatherURL, key, url.QueryEscape(address.Localidade+","+address.Uf))

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get weather data: status code %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var weather entities.Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return nil, err
	}

	return &weather, nil
}
