package main

import (
	"hillel_go_cource/lesson_4/helpers"
	"hillel_go_cource/lesson_4/models"
	"hillel_go_cource/lesson_4/stash"
	"hillel_go_cource/lesson_4/utils"

	"github.com/sirupsen/logrus"
)

func main() {
	possibleNames, err := helpers.NewPossibleNames(
		stash.NamePart1,
		stash.NamePart2,
		stash.NamePart3,
	)
	if err != nil {
		logrus.Fatal(err)
	}
	
	names := possibleNames.GroupUpNames()

	nameGenerator := utils.NewRandomNameGenerator()
	nameGenerator.Options = names
	nameGenerator.Name, err = nameGenerator.GenerateName()
	if err != nil {
		logrus.Fatal(err)
	}
	

	hero := models.NewHero(nameGenerator.Name)

	hydraHeads, err := utils.InitHeads(50, 150)
	if err != nil {
		logrus.Fatal(err)
	}
	hydra:= models.NewHydra(hydraHeads)


	battle := models.NewBattle(hero, hydra)
	battle.StartBattle()
}
