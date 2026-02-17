package viacep

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Frank-Macedo/lab-forecast/internal/domain/entities"
	valueobject "github.com/Frank-Macedo/lab-forecast/internal/domain/valueObject"
)

const ViaCepURL = "https://viacep.com.br/ws/%s/json/"

type ViaCepClient struct {
}

func NewViaCepClient() *ViaCepClient {
	return &ViaCepClient{}
}

func (c *ViaCepClient) GetAddress(cep valueobject.Cep) (*entities.Address, error) {

	url := fmt.Sprintf(ViaCepURL, cep)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var address entities.Address
	err = json.Unmarshal(body, &address)
	if err != nil {
		return nil, err
	}
	return &address, nil
}
