package interfaces

import (
	"github.com/Frank-Macedo/lab-forecast/internal/domain/entities"
	valueobject "github.com/Frank-Macedo/lab-forecast/internal/domain/valueObject"
)

type CepInterface interface {
	GetAddress(cep valueobject.Cep) (*entities.Address, error)
}
