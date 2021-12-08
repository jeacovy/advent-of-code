package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func diagnosticReport(binary []string) {
	bitZeroCountByRow := make(map[int]int)
	bitOneCountByRow := make(map[int]int)

	var bitByPositionSize = 0

	for c := 0; c <= len(binary)-1; c++ { // Column
		bytes := binary[c]
		bitByPosition := strings.Split(bytes, "")
		bitByPositionSize = len(bitByPosition)

		for r := 0; r <= bitByPositionSize-1; r++ { // Row
			if bitByPosition[r] == "0" {
				bitZeroCountByRow[r] = bitZeroCountByRow[r] + 1
			} else if bitByPosition[r] == "1" {
				bitOneCountByRow[r] = bitOneCountByRow[r] + 1
			}
		}
	}

	findBit(bitOneCountByRow, bitZeroCountByRow, bitByPositionSize)

	// Change boolean value to reflect needed rating.
	// True for oxygen generator rating
	// False for CO2 scrubber rating
	lifeSupportRating(bitByPositionSize, binary, false)
}

func findBit(bitOneCount map[int]int, bitZeroCount map[int]int, count int) {
	var gammaRateBinary string
	var epsilonRateBinary string

	for c := 0; c <= count-1; c++ {
		if bitOneCount[c] > bitZeroCount[c] {
			gammaRateBinary = gammaRateBinary + string("0")
			epsilonRateBinary = epsilonRateBinary + string("1")
		} else {
			gammaRateBinary = gammaRateBinary + string("1")
			epsilonRateBinary = epsilonRateBinary + string("0")
		}
	}

	gammaRate, _ := strconv.ParseInt(gammaRateBinary, 2, 64)
	epsilonRate, _ := strconv.ParseInt(epsilonRateBinary, 2, 64)
	powerConsumtion := epsilonRate * gammaRate

	fmt.Printf("Gamma rate binary: %s \n", gammaRateBinary)
	fmt.Printf("Gamma rate as int: %d \n", gammaRate)
	fmt.Printf("Epsilon rate binary: %s \n", epsilonRateBinary)
	fmt.Printf("Epsilon rate as in: %d \n", epsilonRate)
	fmt.Printf("Power consumption: %d \n", powerConsumtion)
}

func lifeSupportRating(sizeOfBinary int, binary []string, isOxygenGeneratorRating bool) {
	if isOxygenGeneratorRating {
		fmt.Printf("==========================\nOxygen generator rating\n==========================\n")
	} else {
		fmt.Printf("==========================\nCO2 scrubber rating\n==========================\n")
	}

	for lsr := 0; lsr <= sizeOfBinary-1; lsr++ {
		var oneCount = 0
		var zeroCount = 0

		for b := 0; b <= len(binary)-1; b++ {
			binaryTopPortion := strings.Split(binary[b], "")

			if binaryTopPortion[lsr] == "1" {
				oneCount++
			} else {
				zeroCount++
			}
		}

		var binaryToDelete []string
		for b := 0; b <= len(binary)-1; b++ {
			binaryTopPortion := strings.Split(binary[b], "")

			if len(binary) == 1 {
				break
			}

			if zeroCount == oneCount {
				oneCount++
			}

			if zeroCount >= oneCount && binaryTopPortion[lsr] != "0" && isOxygenGeneratorRating {
				binaryToDelete = append(binaryToDelete, binary[b])
			}

			if zeroCount <= oneCount && binaryTopPortion[lsr] != "1" && isOxygenGeneratorRating {
				binaryToDelete = append(binaryToDelete, binary[b])
			}

			if zeroCount >= oneCount && binaryTopPortion[lsr] != "1" && !isOxygenGeneratorRating {
				binaryToDelete = append(binaryToDelete, binary[b])
			}

			if zeroCount <= oneCount && binaryTopPortion[lsr] != "0" && !isOxygenGeneratorRating {
				binaryToDelete = append(binaryToDelete, binary[b])
			}

			if len(binary) == 2 && len(binaryToDelete) == 1 {
				break
			}
		}

		for b := 0; b <= len(binaryToDelete)-1; b++ {
			binary = remove(binary, getBinaryIndex(binary, binaryToDelete[b]))
		}

		if len(binary) == 1 {
			break
		}
	}

	// if len(binary) == 1 {
	rating, _ := strconv.ParseInt(binary[0], 2, 64)
	fmt.Printf("Binary: %s \n", binary[0])
	fmt.Printf("Rating: %d \n", rating)
	// return rating

	// return rating
}
func getBinaryIndex(binary []string, value string) int {
	for b := 0; b <= len(binary)-1; b++ {
		if binary[b] == value {
			return b
		}
	}
	return 0
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func readFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	// var binaries = []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	var binaries []string
	absPath, _ := filepath.Abs("/Users/jeacovygayle/workstation/playground/advent-of-code/2021/util/data/day03Input.txt")
	lines, _ := readFile(absPath)

	binaries = append(binaries, lines...)

	diagnosticReport(binaries)
}
