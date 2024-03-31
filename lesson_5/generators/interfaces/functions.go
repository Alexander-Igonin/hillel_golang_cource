package interfaces

import (
	"github.com/sirupsen/logrus"
)

func TestGenerator (generator Generator, param int) {
	generate := generator.Generate()
	generateSlice := generator.GenerateSlice()
	generateWithParam, err := generator.GenerateWithParam(param)
	if err != nil {
		logrus.Error(err)
	}

	results := make([]interface{}, 0, 3)
	results = append(results, generate, generateSlice, generateWithParam)

	for _, v := range results {
		switch result := v.(type) {
		case int:
			logrus.Info("type int was generated")
		case string:
			logrus.Info("type string was generated")
		case []interface{}:
			logrus.Info("type slice was generated")
		case float32:
			logrus.Info("type float was generated")
		case bool:
			logrus.Info("type bool was generated")
		default:
			logrus.Errorf("Type %s is unknown", result)
		}
	}
}