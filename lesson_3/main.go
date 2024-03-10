package main

import (
	"flag"
	"fmt"
)

func main() {
	turtleSpeed := flag.Int("turtle", 40, "Turttle run speed")
	tigerSpeed := flag.Int("tiger", 70, "Tiger run speed")
	fishSpeed := flag.Int("fish", 10, "Fish run speed")
	flag.Parse()

	turtle := Turtle {
		ninja: true,
		favoritePizza: "4 cheese",
		animal: Animal{
			runSpeed: *turtleSpeed,
			name: "Turtle",
			inRace: true,
			winnerVoice: "COWABUNGA!!!!1111",
			looserVoice: "is that a pizza?!",
		},
	}

	tiger := Tiger{
		clawLenght: 5,
		fed: false,
		animal: Animal{
			runSpeed: *tigerSpeed,
			name: "Tiger",
			inRace: true,
			winnerVoice: "I'm the king here!",
			looserVoice: "this is unfair!",
		},
	}

	fish := Fish{}
	fish.runSpeed = *fishSpeed
	fish.name = "Fish"
	fish.Appearance.color = "blue"
	fish.Appearance.lines = true
	fish.inRace = true
	fish.winnerVoice = "i am a tiger now!"
	fish.looserVoice = "i forget how to swim!"

	race := Race{
		Distance: 1000,
	}
	race.CreateTeam(&turtle, &tiger, &fish)
	race.Start()

	for k, v := range race.leaderboard {
		place := k + 1
		
		switch v {
		case "Tiger":
			race.ResultMessage(place, &race.tiger.animal)
		case "Turtle":
			race.ResultMessage(place, &race.turtle.animal)
		case "Fish":
			race.ResultMessage(place, &race.fish.Animal)
		}
	}
}

type Animal struct {
	runSpeed int
	name string
	winnerVoice string
	looserVoice string
	distance int
	inRace bool
	iterations int
}

type Turtle struct {
	animal Animal
	ninja bool
	favoritePizza string
}

type Tiger struct {
	animal Animal
	clawLenght int
	fed bool
}

type Fish struct {
	Animal
	Appearance struct {
		color string
		lines bool
	}
}

type Race struct {
	Distance int
	turtle Turtle
	tiger Tiger
	fish Fish
	leaderboard []string

}


func (*Animal) FinisherSay(name, voice string) {
	switch name {
	case "Tiger":
		fmt.Printf("%s screams: %s\n", name, voice)
	case "Turtle":
		fmt.Printf("%s screams: %s\n", name, voice)
	case "Fish":
		fmt.Printf("%s screams: %s\n", name, voice)
	}
}

func (r *Race) Start() {
	var iterations int

	for {
		iterations += 1

		r.turtle.animal.distance += r.turtle.animal.runSpeed
		r.tiger.animal.distance += r.tiger.animal.runSpeed
		r.fish.distance += r.fish.runSpeed

		switch {
		case r.turtle.animal.distance >= r.Distance && r.turtle.animal.inRace == true:
			r.Finish(&r.turtle.animal, iterations)
		case r.tiger.animal.distance >= r.Distance && r.tiger.animal.inRace == true:
			r.Finish(&r.tiger.animal, iterations)
		case r.fish.distance >= r.Distance && r.fish.inRace == true:
			r.Finish(&r.fish.Animal, iterations)
		}

		if len(r.leaderboard) == 3 {
			break
		}
	}
}

func (r *Race) CreateTeam(turtle *Turtle, tiger *Tiger, fish *Fish) {
	r.fish = *fish
	r.tiger = *tiger
	r.turtle = *turtle
}

func (r *Race) ResultMessage(place int, animal *Animal) {
	switch place {
	case 1:
		fmt.Printf("The winner is %s. Total iterations: %d\n", animal.name, animal.iterations)
		animal.FinisherSay(animal.name, animal.winnerVoice)
	case 2:
		 fmt.Printf("Second place is %s, Total iterations: %d\n", animal.name, animal.iterations)
	case 3:
		 fmt.Printf("Last place is %s, Total iterations: %d\n", animal.name, animal.iterations)
		 animal.FinisherSay(animal.name, animal.looserVoice)
	}
}

func (r *Race) Finish(finisher *Animal, iterations int) {
	r.leaderboard = append(r.leaderboard, finisher.name)
	finisher.inRace = false
	finisher.iterations = iterations
}
