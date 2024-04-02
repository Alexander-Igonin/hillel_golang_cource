package integer

import (
	"errors"
	"math/rand"
)


func (i Integer) Generate() interface{} {
	return i.Integer
}

func (i Integer) GenerateSlice() []interface{} {
	var Integer = i.Integer
	var result []interface{}

	result = append(result, Integer)

	return result
}

func (f Integer) GenerateWithParam(param int) (interface{}, error) {
	if param <= 0 {
		return nil, errors.New("param must be > 0")
	}

	return rand.Intn(param), nil
}

