package generators

import (
	"hillel_go_cource/lesson_5/generators/boolean"
	"hillel_go_cource/lesson_5/generators/float"
	"hillel_go_cource/lesson_5/generators/integer"
	"hillel_go_cource/lesson_5/generators/str"
)

type Generators struct {
	Bool *boolean.Boolean
	Float *float.Float
	Integer *integer.Integer
	Str *str.Str
}

func NewGenerators() *Generators {
	return &Generators{
		Bool: boolean.NewBoolean(),
		Float: float.NewFloat(),
		Integer: integer.NewInteger(),
		Str: str.NewStr(),
	}
}
