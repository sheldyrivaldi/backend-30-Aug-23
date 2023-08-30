package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var players int
	var dices int

	for {
		fmt.Print("Masukkan jumlah player: ")
		_, playerErr := fmt.Scanf("%d", &players)
		if playerErr != nil {
			fmt.Println("Error:", playerErr)
		} else {
			break
		}
	}

	for {
		fmt.Print("Masukkan jumlah dadu: ")
		_, dicesErr := fmt.Scanf("%d", &dices)
		if dicesErr != nil {
			fmt.Println("Error:", dicesErr)
		} else {
			break
		}
	}

	fmt.Println()

	diceApp(players, dices)
}

// Fungsi untuk menjalankan aplikasi "Dice Game"
func diceApp(players, dice int) {
	playerRolls := make([][]int, players)
	score := make([]int, players)
	round := 1

	for idx := range playerRolls {
		playerRolls[idx] = rollDice(dice)
	}
	fmt.Println("Lemparan Ke -", round, playerRolls)

	for {
		filteredArray, poin := filteredRollDice(playerRolls, score)
		score = poin
		fmt.Println("Setelah evaluasi", filteredArray)
		fmt.Println("Score", poin)
		fmt.Println()

		notEmpty := 0
		for _, v := range filteredArray {
			if len(v) > 0 {
				notEmpty++
			}
		}
		if notEmpty <= 1 {
			break
		}

		for idx := range playerRolls {
			playerRolls[idx] = rollDice(len(filteredArray[idx]))
		}
		round++
		fmt.Println("Lemparan Ke -", round, playerRolls)
	}

	theWinner := findTheWinnerIndex(score)

	fmt.Println("Pemenangnya adalah Player", theWinner)

}

// Function untuk mencari index pemenang
func findTheWinnerIndex(arr []int) int {
	index := 1
	max := arr[0]
	if len(arr) == 0 {
		return 0
	} else {
		for idx, num := range arr {
			if num > max {
				max = num
				index = idx + 1
			}
		}
	}
	return index
}

// Function untuk mengacak slice dadu berdasarkan jumlah dadu
func rollDice(dice int) []int {
	allRolls := make([]int, dice)

	for idx := range allRolls {
		allRolls[idx] = rand.Intn(6) + 1
	}
	return allRolls
}

// Function untuk memfilter hasil dadu player dan mengembalikan 2 nilai yaitu slice dadu player dan slice poin
func filteredRollDice(arr [][]int, score []int) ([][]int, []int) {
	playerRolls := arr
	poin := score
	except := 0

	for idx := range playerRolls {
		temp := []int{}
		if except != 0 {

			for i := 1; 1 <= except; i++ {
				temp = append(temp, 1)
				except -= 1
			}

			for idx2 := range playerRolls[idx] {
				if playerRolls[idx][idx2] == 1 {
					except += 1
					continue
				}
				if playerRolls[idx][idx2] == 6 {
					poin[idx] += 1
				} else {
					temp = append(temp, playerRolls[idx][idx2])
				}
			}
		} else {
			for idx2 := range playerRolls[idx] {
				if playerRolls[idx][idx2] == 1 {
					except += 1
					continue
				}
				if playerRolls[idx][idx2] == 6 {
					poin[idx] += 1
				} else {
					temp = append(temp, playerRolls[idx][idx2])
				}
			}
		}

		playerRolls[idx] = temp

	}

	temp := []int{}
	if except != 0 {
		for idx := range playerRolls[0] {
			if playerRolls[0][idx] == 1 {
				except += 1
				continue
			}
			if playerRolls[0][idx] == 6 {
				poin[0] += 1
			} else {
				temp = append(temp, playerRolls[0][idx])
			}
		}

		for i := 1; 1 <= except; i++ {
			temp = append(temp, 1)
			except -= 1
		}

		playerRolls[0] = temp
	}

	return playerRolls, poin
}
