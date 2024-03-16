package utils

import (
	"errors"
	"math/rand"
)


type RandomNameGenerator struct {
	Name string
	Options [][]string
}

func NewRandomNameGenerator() *RandomNameGenerator {
	return &RandomNameGenerator{}
}


func (r *RandomNameGenerator) GenerateName() (string, error) {
	if len(r.Options) == 0 {
		return "", errors.New("slice with names can't be empty")
	}
	var name string

	for _, sl := range r.Options {
		name += sl[rand.Intn(len(sl))] + " "
	}

	return name, nil
}
