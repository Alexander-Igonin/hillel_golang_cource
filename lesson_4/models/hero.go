package models

import (
	"math/rand"
)

type Hero struct {
	Name string
	Shout string
}

func NewHero(name string) *Hero {
	return &Hero{
		Name: name,
	}
}

func (h Hero) Attack() int {
	return rand.Intn(6)
}
