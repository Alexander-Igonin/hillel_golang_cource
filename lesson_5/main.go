package main

import (
	"hillel_go_cource/lesson_5/generators"
	"hillel_go_cource/lesson_5/generators/interfaces"
)

func main() {
	generators := generators.NewGenerators()
	interfaces.TestGenerator(generators.Bool, 5)
	
}
