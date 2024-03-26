package float

import (
	"errors"
	"math/rand"
)

func (f Float) Generate() interface{} {
	return f.Float
}

func (f Float) GenerateSlice() []interface{} {
	var Float = f.Float
	var result []interface{}

	result = append(result, Float)

	return result
}

func (f Float) GenerateWithParam(param int) (interface{}, error) {
	if param <= 0 {
		return nil, errors.New("param must be > 0")
	}

	return float32(rand.Intn(param)) + rand.Float32(), nil
}
