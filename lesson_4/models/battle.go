package models

import (
	"fmt"
	"hillel_go_cource/lesson_4/stash"
	"hillel_go_cource/lesson_4/utils"
	"math/rand"
)


type Battle struct {
	hero Hero
	hydra Hydra
}

func NewBattle(hero *Hero, hydra *Hydra) *Battle {
	return &Battle{
		hero: *hero,
		hydra: *hydra,
	}
}

func (b *Battle) StartBattle() {
	var round int

	for {
		round++
		damage := b.hero.Attack()


		b.hydra.LostHeads(damage)

		fmt.Printf("Round: %d. Hydra lost %d heads, %d left\n", round, damage ,b.hydra.NumOfHeads)

		if damage > 0 {
			recoverHeads := utils.RecoverHeads(rand.Intn(101) , damage)
			b.hydra.NumOfHeads += recoverHeads
			fmt.Printf("Hydra recover %d heads\n", recoverHeads)
		}

		if damage == 0 {
			b.hydra.Shout = utils.ReleaseShout(stash.ListOfHydraShouts)
			fmt.Printf("Hydra: %s\n", b.hydra.Shout)
		}

		if damage >= 4 {
			b.hero.Shout = utils.ReleaseShout(stash.ListOfHeroShouts)
			fmt.Printf("%s: %s\n", b.hero.Name,  b.hero.Shout)
		}
		
		if b.hydra.NumOfHeads <= 0 {
			fmt.Printf("The winner is %s !!!", b.hero.Name)
			break
		}
			
		if b.hydra.NumOfHeads >= 200 {
			panic(fmt.Sprintf("Hydra has too many heads, %s retreat", b.hero.Name))
		}
	}
}

