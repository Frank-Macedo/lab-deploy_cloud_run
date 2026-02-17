package interfaces

import (
	"github.com/Frank-Macedo/lab-forecast/internal/domain/entities"
)

type WeatherInterface interface {
	GetTemperature(address entities.Address) (*entities.Weather, error)
}
