package main

import (
	"fmt"
)

func main() {
	// 1 - Створити структуру Game, в якій буде 5 полей. Одне з яких TotalPlayers з типом int. Інші будь які.
	type Game struct {
		TotalPlayers int
		Difficulty string
		Online bool
		WorldRank int
		Price int
	}

	// 2 - Створити структуру TableGame  - де буде вбудована структура Game та ще одне будь яке поле.
	type TableGame struct {
		game Game
		Name string
	}

	// 3 - Створити обʼєкт та заповнити всі поля.
	var gloomhaven = TableGame {
		Name: "Gloomhaven",
		game: Game{
			TotalPlayers: 4,
			Difficulty: "Very hard",
			WorldRank: 18,
			Online: true,
			Price: 200,
		},
	}

	// 4 - Створити слайс з типом TableGame додати туди обʼєкт з 3 (третього) пункту.
	gamesSlice := make([]TableGame, 1, 3)
	gamesSlice[0] = gloomhaven

	// 5 - Створити ще 2 обʼєкти і також додати їх до слайсу.
	root := TableGame{}

	root.Name = "Root"
	root.game.Difficulty = "Hard"
	root.game.Online = true
	root.game.Price = 100
	root.game.TotalPlayers = 6
	root.game.WorldRank = 1

	var duneImperium TableGame = TableGame{Name: "Dune Imperium"}
	
	duneImperium.game = Game {
		Difficulty: "Easy",
		Online: false,
		WorldRank: 72,
		Price: 150,
		TotalPlayers: 4,
	}

	gamesSlice = append(gamesSlice, root, duneImperium)

	// 6 - Переконатись що при виконанні 5-го пункта в слайсі не буде виникати аллокація памʼяті
	//  (якщо буде, то змінити код так щоб аллокація не виникала)
	fmt.Printf("Len - %d Cap - %d\n", len(gamesSlice), cap(gamesSlice)) // Len - 3 Cap - 3

	// 7 - Всі елементи зі слайсу записати в мапу map[int]TableGame - де ключом буде індекс елементу зі слайса. 
	// За допогою for{}
	var gamesMap = make(map[int]TableGame)

	for k, v := range gamesSlice {
		gamesMap[k] = v
	}

	// 8 - вивести на екран мапу (кожне значення з нового рядка)
	for _, v := range gamesMap {
		fmt.Println(v)
	}

	// 9 - За допомогою циклу знайти сумму всіх TotalPlayers мапі (в усіх елементах) 
	var totalPlayers int

	for _, v := range gamesMap {
		totalPlayers += v.game.TotalPlayers
	}

	// 10 - Сумму вивести на екран
	fmt.Print(totalPlayers)
}