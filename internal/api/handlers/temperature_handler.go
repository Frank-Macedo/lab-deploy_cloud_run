package handlers

import (
	"fmt"
	"net/http"

	"github.com/Frank-Macedo/lab-forecast/internal/application/usecases"
	valueobject "github.com/Frank-Macedo/lab-forecast/internal/domain/valueObject"
	viacep "github.com/Frank-Macedo/lab-forecast/internal/infrastructure/via_cep"
	"github.com/Frank-Macedo/lab-forecast/internal/infrastructure/weather"
	"github.com/gorilla/mux"
)

func GetTemperature(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	cep, err := valueobject.NewCep(vars["cep"])
	if err != nil {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	cepClient := viacep.NewViaCepClient()
	weatherClient := weather.NewWeatherClient()

	TempUseCase := usecases.NewGetTemperatureUseCase(weatherClient, cepClient)
	weather, err := TempUseCase.Execute(cep)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v°C, %v°F, %vK", weather.Current.TempC, weather.Current.TempF, weather.Current.TempC+273.15)))
}
