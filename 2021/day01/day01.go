package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func depthMeasurements(measurments []int){
	var lastMeasuredValue int
	var totalMeasurementsLargerThanPrevious int
	var totalMeasurementsSmallerThanPrevious int
	var r int
	
	for r = 0; r <= len(measurments)-1; r++ {
		if(r != 0){
			if(lastMeasuredValue < measurments[r]){
				fmt.Printf("%d (increased)\n", measurments[r])
				totalMeasurementsLargerThanPrevious++
			}else{
				fmt.Printf("%d (decreased)\n", measurments[r])
				totalMeasurementsSmallerThanPrevious++
			}
			lastMeasuredValue = measurments[r]
		}else{
			fmt.Printf("%d (N/A - no previous measurement)\n", measurments[r])
		}
	}

	fmt.Printf("%d measurements are larger than the previous measurement.\n", totalMeasurementsLargerThanPrevious)
	fmt.Printf("%d measurements are larger than the previous measurement.\n", totalMeasurementsSmallerThanPrevious)
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
	var measurments []int
	absPath, _ := filepath.Abs("../advent-of-code/2021/util/data/day01Input.txt")
	lines, _ := readFile(absPath)

    for _, line := range lines {
		intVar, _ := strconv.Atoi(line)
		measurments = append(measurments, intVar)
    }

	depthMeasurements(measurments)
}