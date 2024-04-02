package main

import (
	"fmt"
	"hillel_go_cource/lesson_5/generators"
	"hillel_go_cource/lesson_5/generators/interfaces"
)

func main() {
	generators := generators.NewGenerators()
	groupedGenerators := []interfaces.Generator{
		generators.Bool,
		generators.Float,
		generators.Integer,
		generators.Str,
	}

	for _, v := range groupedGenerators {
		interfaces.TestGenerator(v, 1)
		fmt.Println()
	}

	
	
}
