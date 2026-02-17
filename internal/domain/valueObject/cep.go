package valueobject

import "fmt"

type Cep struct {
	value string
}

func NewCep(cep string) (Cep, error) {

	if len(cep) != 8 || !isNumeric(cep) {
		return Cep{}, fmt.Errorf("Invalid CEP format")
	}

	return Cep{value: cep}, nil

}

func isNumeric(cep string) bool {
	for _, r := range cep {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
func (c Cep) String() string {
	return c.value
}

func (c Cep) Equals(other Cep) bool {
	return c.value == other.value
}
