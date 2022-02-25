package main

import "fmt"

var calledNumbers []int
var yelledBingo = false

var winningRowCount = make(map[int]int)
var winningCard = 0

var markedSpotsOnCard = make(map[int][]int)

func bingoSubsystem(drawnBingoNumbers []int) {
	for bs := 0; bs <= len(drawnBingoNumbers)-1; bs++ {
		submarineBingoSystems(drawnBingoNumbers[bs])

		if yelledBingo {
			break
		}
	}
}

func incremenetWinningRows(calledNumber int, cardNumber int, numberLocationOnCard int) {
	// TODO: Check if this is in a winning row.
	// Loop through each winning row and increment

	var winningRows = make(map[int][]int)
	winningRows[0] = []int{1, 2, 3, 4, 5}      // Straight accross any row
	winningRows[1] = []int{6, 7, 8, 9, 10}     // Straight accross any row
	winningRows[2] = []int{11, 12, 13, 14, 15} // Straight accross any row
	winningRows[3] = []int{16, 17, 18, 19, 20} // Straight accross any row
	winningRows[4] = []int{21, 22, 23, 24, 25} // Straight accross any row

	winningRows[5] = []int{1, 6, 11, 16, 21}  // Straigt down
	winningRows[6] = []int{2, 7, 12, 17, 22}  // Straigt down
	winningRows[7] = []int{3, 8, 13, 18, 23}  // Straigt down
	winningRows[8] = []int{4, 9, 14, 19, 24}  // Straigt down
	winningRows[9] = []int{5, 10, 15, 20, 25} // Straigt down1

	winningRows[10] = []int{1, 7, 13, 19, 25} // Top left to bottom right
	winningRows[11] = []int{21, 17, 13, 9, 5} // Bottom left to bottom right

	for wr := 0; wr <= len(winningRows)-1; wr++ {
		var winningRow = winningRows[wr]
		// var sumForRows []int
		//Checking if card has number.
		for r := 0; r <= len(winningRow)-1; r++ {
			// fmt.Printf("Location on card: %d  calledNumberRow: %d\n", numberLocationOnCard, winningRow[r])
			if numberLocationOnCard == winningRow[r] {

				winningRowCount[r] = winningRowCount[r] + 1

				fmt.Printf("The row is: %d and row count is: %d \n", winningRow[r], winningRowCount[r])

				// bitZeroCountByRow[r] = bitZeroCountByRow[r] + 1

				if winningRowCount[r] == 5 {
					winningCard = calledNumber
					yelledBingo = true
					break
				}

				// fmt.Printf("The incremented count is: %d \n", checkForBingo)
			}
		}

		if yelledBingo {
			break
		}

		// sumForRows = append(sumForRows, inc)
		// if sumForRows[wr] == 5 {
		// 	yelledBingo = true
		// 	break
		// }
	}
}

func checkBingoCards(calledNumber int) {
	var cards = make(map[int][]int)
	// var sumOfUnmarkedNumbers = 0
	var finalScore = 0

	// Read card without the 0 -- +1 always
	// Probably will needed to read this differntly

	cards[0] = []int{22, 13, 17, 11, 0, 8, 2, 23, 4, 24, 21, 9, 14, 16, 7, 6, 10, 3, 18, 5, 1, 12, 20, 15, 19}
	cards[1] = []int{3, 15, 0, 2, 22, 9, 18, 13, 17, 5, 19, 8, 7, 25, 23, 20, 11, 10, 24, 4, 14, 21, 16, 12, 6}
	cards[2] = []int{14, 21, 17, 24, 4, 10, 16, 15, 9, 19, 18, 8, 23, 26, 20, 22, 11, 13, 6, 5, 2, 0, 12, 3, 7}

	//TODO: Loop through cards and check for called numbers
	//Count up if for each winning row.

	for c := 0; c <= len(cards)-1; c++ {
		var cardNumbers = cards[c]
		// winningRowCount = make(map[int]int)

		// Resetting on each card -- total of 3
		// winningRowCount0 = 0
		// winningRowCount1 = 0
		// winningRowCount2 = 0
		// winningRowCount3 = 0

		fmt.Printf("The card numbers are: %d \n", cardNumbers)
		//Checking if card has number.
		for c1 := 0; c1 <= len(cardNumbers)-1; c1++ {
			if cardNumbers[c1] == calledNumber {
				// fmt.Printf("The number row on the card is: %d \n", c1+1)
				incremenetWinningRows(calledNumber, c, c1+1)
			}

			// if winningRowCount0 == 5 || winningRowCount1 == 5 || winningRowCount2 == 5 || winningRowCount3 == 5 {
			if yelledBingo {
				break
			}
		}

		if yelledBingo {
			break
		}

		// if winningRowCount0 == 5 || winningRowCount1 == 5 || winningRowCount2 == 5 || winningRowCount3 == 5 {
		// 	yelledBingo = true
		// 	winningCard = c
		// 	break
		// }
	}

	if yelledBingo {
		fmt.Print("=========\n======\n YELL BINGO \n")
		fmt.Printf("The last numnber called was: %d \n", calledNumber)
		fmt.Printf("The final score is: %d \n", finalScore)
		fmt.Printf("Winning row count: %d \n", winningRowCount)
		// var numbersToRemove []int
		// var cardNumbers = cards[winningCard]

		// for n := 0; n <= len(cardNumbers)-1; n++ {
		// 	for cn := 0; cn <= len(calledNumbers)-1; cn++ {
		// 		if cardNumbers[n] == calledNumbers[cn] {
		// 			numbersToRemove = append(numbersToRemove, calledNumbers[cn])
		// 		}
		// 	}
		// }

		// Rework below
		// for nr := 0; nr <= len(numbersToRemove)-1; nr++ {
		// 	cardNumbers = remove(cardNumbers, getBinaryIndex(cardNumbers, numbersToRemove[nr]))
		// }

		// fmt.Printf("The card number after removals: %d\n", len(cardNumbers))

		// for c := 0; c <= len(cardNumbers)-1; c++ {
		// 	sumOfUnmarkedNumbers += cardNumbers[c]
		// }

		// finalScore = sumOfUnmarkedNumbers * calledNumbers[len(calledNumbers)]

	} else {
		fmt.Printf("The last numnber called was: %d \n", calledNumber)
	}
}

func checkCardPaths(cardNumber int) {
	var winnablePaths = make(map[int][]int)
	winnablePaths[0] = []int{1, 2, 3, 4, 5}      // Straight accross any row
	winnablePaths[1] = []int{6, 7, 8, 9, 10}     // Straight accross any row
	winnablePaths[2] = []int{11, 12, 13, 14, 15} // Straight accross any row
	winnablePaths[3] = []int{16, 17, 18, 19, 20} // Straight accross any row
	winnablePaths[4] = []int{21, 22, 23, 24, 25} // Straight accross any row
	winnablePaths[5] = []int{1, 6, 11, 16, 21}   // Straigt down
	winnablePaths[6] = []int{2, 7, 12, 17, 22}   // Straigt down
	winnablePaths[7] = []int{3, 8, 13, 18, 23}   // Straigt down
	winnablePaths[8] = []int{4, 9, 14, 19, 24}   // Straigt down
	winnablePaths[9] = []int{5, 10, 15, 20, 25}  // Straigt down
	winnablePaths[10] = []int{1, 7, 13, 19, 25}  // Top left to bottom right
	winnablePaths[11] = []int{21, 17, 13, 9, 5}  // Bottom left to top right

	for msoc := 0; msoc <= len(markedSpotsOnCard)-1; msoc++ {

		for soc := 0; soc <= len(markedSpotsOnCard[msoc])-1; soc++ {

			//TODO: Get winnable path
			var countWinPath = 0

			for wp := 0; wp <= len(winnablePaths)-1; wp++ {

				if contains(winnablePaths[wp], markedSpotsOnCard[msoc][soc]) {
					countWinPath++
				}
				if countWinPath == 3 {
					// fmt.Printf("Marked spots on card: %d\n", markedSpotsOnCard)
					yelledBingo = true
					break
				}

				// for p := 0; p <= len(winnablePaths[wp])-1; p++ {

				// 	if markedSpotsOnCard[msoc][soc] == winnablePaths[wp][p] {
				// 		fmt.Printf("We are checking if marked spot -> %d on card is equal to -> %d for path -> %d\n", markedSpotsOnCard[msoc][soc], winnablePaths[wp][p], wp)

				// 		// if cardNumber == 2 {
				// 		countWinPath++
				// 		// fmt.Printf("The count: %d \n", countWinPath)
				// 		// }
				// 	}

				// 	if countWinPath == 5 {
				// 		// fmt.Printf("Marked spots on card: %d\n", markedSpotsOnCard)
				// 		yelledBingo = true
				// 		break
				// 	}

				// }
				fmt.Printf("Path %d -- count them up: %d \n", wp, countWinPath)
				if yelledBingo {
					break
				}
			}

			if yelledBingo {
				break
			}

		}

		if yelledBingo {
			break
		}
	}
}

func submarineBingoSystems(drawnNumber int) {
	var cards = make(map[int][]int)
	cards[0] = []int{22, 13, 17, 11, 0, 8, 2, 23, 4, 24, 21, 9, 14, 16, 7, 6, 10, 3, 18, 5, 1, 12, 20, 15, 19}
	cards[1] = []int{3, 15, 0, 2, 22, 9, 18, 13, 17, 5, 19, 8, 7, 25, 23, 20, 11, 10, 24, 4, 14, 21, 16, 12, 6}
	cards[2] = []int{14, 21, 17, 24, 4, 10, 16, 15, 9, 19, 18, 8, 23, 26, 20, 22, 11, 13, 6, 5, 2, 0, 12, 3, 7}

	// DONE: A number is called.
	// DONE: We take that number and iterate the cards
	// DONE: IF the number exist tag the card spot
	// DONE: Each card should iterate through the mutliple winnables paths
	// DONE: Means this should be stored globally
	// When a winnable path reaches 5 we yell bingo.
	// We will tag [card][spots]

	for c := 0; c <= len(cards)-1; c++ {
		var cardNumbers = cards[c]
		for cn := 0; cn <= len(cardNumbers)-1; cn++ {
			if drawnNumber == cardNumbers[cn] {
				markedSpotsOnCard[c] = append(markedSpotsOnCard[c], cn+1)
			}
		}

		// fmt.Printf("Marked spots: %d \n", markedSpotsOnCard)
		checkCardPaths(c) // Move this into the if statment

		if yelledBingo {
			fmt.Print("=========\n======\n YELL BINGO \n")
			fmt.Printf("The last numnber called was: %d and card number is: %d \n", drawnNumber, c)
			break
			// fmt.Printf("The final score is: %d \n", finalScore)
			// fmt.Printf("Winning row count: %d \n", winningRowCount)
		}

		fmt.Printf("Mark sport on card: %d \n", markedSpotsOnCard)
	}
}

func getBinaryIndex(slice []int, value int) int {
	for s := 0; s <= len(slice)-1; s++ {
		if slice[s] == value {
			return s
		}
	}
	return 0
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func contains(slice []int, item int) bool {
	set := make(map[int]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func main() {
	fmt.Print("BINGO TIME \n")

	bingoNumbers := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}

	bingoSubsystem(bingoNumbers)
}
