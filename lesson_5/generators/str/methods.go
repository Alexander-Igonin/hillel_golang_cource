package str

import "errors"

func (s Str) Generate() interface{} {
	return s.String
}

func (s Str) GenerateSlice() []interface{} {
	var Str Str
	var result []interface{}

	result = append(result,  Str)

	return result
}

func (s Str) GenerateWithParam(param int) (interface{}, error) {
	if param <= 0 {
		return nil, errors.New("param must be > 0")
	}

	result := ""

	for i := 0; i < param; i++ {
		result += " "
	}

	return result, nil
}
