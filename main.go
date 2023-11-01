package main

import (
	"fmt"
	"math/rand"
)

func main() {
	playerCount := 3
	diceCount := 4
	GenerateGame(playerCount, diceCount)
}

func GenerateGame(playerCount, diceCount int) {
	endGame := false
	playersData := make(map[string]map[string]interface{})
	round := 1
	var filteredDice []int

	for i := 1; i <= playerCount; i++ {
		playerName := fmt.Sprintf("Player %d", i)
		playersData[playerName] = map[string]interface{}{
			"score": 0,
			"dice":  make([]int, diceCount),
		}
	}

	for !endGame {
		for key := range playersData {
			diceCount := len(playersData[key]["dice"].([]int))
			playersData[key]["dice"] = rollDice(diceCount)
		}

		fmt.Printf("ROUND %d\n", round)
		fmt.Printf("Kocok Dadu %d\n", round)
		fmt.Println("========================\n")
		for key, playerData := range playersData {
			fmt.Printf("%s(%d): %v\n", key, playerData["score"], playerData["dice"])
		}
		fmt.Println("======================== \n")
		playersMovementCount := make(map[string]int)

		for key, playerData := range playersData {
			currentPlayerKey := key[len(key)-1 : len(key)]
			nextPlayerKey := "1"
			playerCount := len(playersData)
			if currentPlayerKey != fmt.Sprintf("%d", playerCount) {
				nextPlayerKey = fmt.Sprintf("%d", (int(currentPlayerKey[0]-'0') + 1))
			}

			diceArray := playerData["dice"].([]int)
			/* Cek Dadu apakah ada 1 & 6 */
			for _, diceNumber := range diceArray {
				if diceNumber == 1 {
					playersMovementCount[fmt.Sprintf("Player %s", nextPlayerKey)]++

				} else if diceNumber == 6 {
					playerData["score"] = playerData["score"].(int) + 1
				}
				if diceNumber != 1 && diceNumber != 6 {
					filteredDice = append(filteredDice, diceNumber)
				}
			}
			playerData["dice"] = filteredDice
			filteredDice = nil
		}

		// Proses Pemindahan Dadu
		for key, movementCount := range playersMovementCount {
			for i := 0; i < movementCount; i++ {
				diceArray := playersData[key]["dice"].([]int)
				diceArray = append(diceArray, 1)
				playersData[key]["dice"] = diceArray
			}
		}
		// Cek Apakah Game Sudah Berakhir
		playerHaveDiceCount := 0
		for _, playerData := range playersData {
			diceArray := playerData["dice"].([]int)
			if len(diceArray) > 0 {
				playerHaveDiceCount++
			} else {
				playerData["play"] = false
			}
		}
		if playerHaveDiceCount < 2 {
			endGame = true
		}

		fmt.Printf("Evaluasi %d\n", round)
		fmt.Println("========================")
		for key, playerData := range playersData {
			fmt.Printf("%s(%d): %v\n", key, playerData["score"], playerData["dice"])
		}
		fmt.Println("========================")

		round++
	}
	highestScore := -1
	playersChampions := []string{}

	for playerName, playerData := range playersData {
		score := playerData["score"].(int)
		if score > highestScore {
			highestScore = score
			playersChampions = []string{playerName}
		} else if score == highestScore {
			playersChampions = append(playersChampions, playerName)
		}
	}

	fmt.Println("========================")
	fmt.Printf("Score Tertinggi Adalah %d \n", highestScore)
	fmt.Printf("Dan Pemenangnya Adalah %s\n", playersChampions)
	fmt.Println("========================")
}

func rollDice(total int) []int {
	var dice []int
	for i := 0; i < total; i++ {
		dice = append(dice, rand.Intn(6)+1)
	}
	return dice
}
