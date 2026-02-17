package usecases

import (
	"errors"
	"testing"

	"github.com/Frank-Macedo/lab-forecast/internal/domain/entities"
	valueobject "github.com/Frank-Macedo/lab-forecast/internal/domain/valueObject"
)

// MockCepInterface implements interfaces.CepInterface
type MockCepInterface struct {
	GetAddressFunc func(cep valueobject.Cep) (*entities.Address, error)
}

func (m *MockCepInterface) GetAddress(cep valueobject.Cep) (*entities.Address, error) {
	return m.GetAddressFunc(cep)
}

// MockWeatherInterface implements interfaces.WeatherInterface
type MockWeatherInterface struct {
	GetTemperatureFunc func(address entities.Address) (*entities.Weather, error)
}

func (m *MockWeatherInterface) GetTemperature(address entities.Address) (*entities.Weather, error) {
	return m.GetTemperatureFunc(address)
}

func TestExecute_Success(t *testing.T) {
	mockCep := &MockCepInterface{
		GetAddressFunc: func(cep valueobject.Cep) (*entities.Address, error) {
			return &entities.Address{Erro: ""}, nil
		},
	}

	location := entities.Location{}
	current := entities.Current{TempC: 25.0, TempF: 25.0*1.8 + 32}

	mockWeather := &MockWeatherInterface{
		GetTemperatureFunc: func(address entities.Address) (*entities.Weather, error) {
			return &entities.Weather{Location: location, Current: current}, nil
		},
	}
	cep, _ := valueobject.NewCep("12345678")
	uc := NewGetTemperatureUseCase(mockWeather, mockCep)
	weather, err := uc.Execute(cep)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if weather == nil || weather.Current.TempC != 25.0 {
		t.Fatalf("expected weather with temperature 25.0, got %+v", weather)
	}
}

func TestExecute_InvalidZipcodeError(t *testing.T) {
	mockCep := &MockCepInterface{
		GetAddressFunc: func(cep valueobject.Cep) (*entities.Address, error) {
			return nil, errors.New("some error")
		},
	}
	mockWeather := &MockWeatherInterface{}
	uc := NewGetTemperatureUseCase(mockWeather, mockCep)

	cep, _ := valueobject.NewCep("invalid")
	weather, err := uc.Execute(cep)

	if err == nil || err.Error() != "Invalid zipcode" {
		t.Fatalf("expected 'Invalid zipcode' error, got %v", err)
	}
	if weather != nil {
		t.Fatalf("expected nil weather, got %+v", weather)
	}
}

func TestExecute_CannotFindZipcodeError(t *testing.T) {
	mockCep := &MockCepInterface{
		GetAddressFunc: func(cep valueobject.Cep) (*entities.Address, error) {
			return &entities.Address{Erro: "true"}, nil
		},
	}
	mockWeather := &MockWeatherInterface{}
	uc := NewGetTemperatureUseCase(mockWeather, mockCep)

	cep, _ := valueobject.NewCep("00000000")
	weather, err := uc.Execute(cep)
	if err == nil || err.Error() != "can not find zipcode" {
		t.Fatalf("expected 'can not find zipcode' error, got %v", err)
	}
	if weather != nil {
		t.Fatalf("expected nil weather, got %+v", weather)
	}
}

func TestExecute_WeatherInterfaceError(t *testing.T) {
	mockCep := &MockCepInterface{
		GetAddressFunc: func(cep valueobject.Cep) (*entities.Address, error) {
			return &entities.Address{Erro: ""}, nil
		},
	}
	mockWeather := &MockWeatherInterface{
		GetTemperatureFunc: func(address entities.Address) (*entities.Weather, error) {
			return nil, errors.New("weather service error")
		},
	}
	uc := NewGetTemperatureUseCase(mockWeather, mockCep)

	cep, _ := valueobject.NewCep("12345678")
	weather, err := uc.Execute(cep)
	if err == nil || err.Error() != "weather service error" {
		t.Fatalf("expected 'weather service error', got %v", err)
	}
	if weather != nil {
		t.Fatalf("expected nil weather, got %+v", weather)
	}
}
