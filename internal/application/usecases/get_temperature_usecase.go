package usecases

import (
	"errors"

	"github.com/Frank-Macedo/lab-forecast/internal/domain/entities"
	"github.com/Frank-Macedo/lab-forecast/internal/domain/interfaces"
	valueobject "github.com/Frank-Macedo/lab-forecast/internal/domain/valueObject"
)

type GetTemperatureUseCase struct {
	temperatureInterface interfaces.WeatherInterface
	cepInterface         interfaces.CepInterface
}

func NewGetTemperatureUseCase(temperatureInterface interfaces.WeatherInterface, cepInterface interfaces.CepInterface) *GetTemperatureUseCase {
	return &GetTemperatureUseCase{
		temperatureInterface: temperatureInterface,
		cepInterface:         cepInterface,
	}
}

func (uc *GetTemperatureUseCase) Execute(cep valueobject.Cep) (*entities.Weather, error) {
	address, err := uc.cepInterface.GetAddress(cep)
	if err != nil {
		return nil, errors.New("Invalid zipcode")
	}
	if address.Erro != "" {
		return nil, errors.New("can not find zipcode")
	}

	weather, err := uc.temperatureInterface.GetTemperature(*address)
	if err != nil {
		return nil, err
	}
	return weather, nil
}
