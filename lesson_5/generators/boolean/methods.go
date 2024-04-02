package boolean

import "errors"


func (b Boolean) Generate() interface{} {
	return b.Boolean
}

func (b Boolean) GenerateSlice() []interface{} {
	var Bool = b.Boolean
	var result []interface{}

	result = append(result, Bool)

	return result
}

func (b Boolean) GenerateWithParam(param int) (interface{}, error) {
	switch param {
	case 0:
		return false, nil
	case 1:
		return true, nil
	}

	return nil, errors.New("param must be 0 or 1")
}
