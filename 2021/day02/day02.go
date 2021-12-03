package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)


func navigate(route []string) (int){
    var r int
    hPosition := 0
    depth := 0
    aim := 0
    
    for r = 0; r <= len(route)-1; r++ {
        routeRawValues := strings.Fields(route[r])
        
        direction := routeRawValues[0]
        unit, _ := strconv.Atoi(routeRawValues[1])

        if direction == "forward" {
            hPosition += unit
            depth += aim * unit
        } else if direction == "down" {
            // depth += unit
            aim += unit
        } else if direction == "up" {
            // depth -= unit
            aim -= unit
        }
        
		fmt.Printf("Direction: %s Units: %d\n", direction, unit)
	}

	fmt.Printf("Direction: %d Postition: %d Aim: %d\n", depth, hPosition, aim)

    return depth * hPosition
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
    // var commands = []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
    var commands []string
	absPath, _ := filepath.Abs("../2021/util/data/day02Input.txt")
	lines, _ := readFile(absPath)

	commands = append(commands, lines...)
    
    fmt.Printf("Horizontal position and depth after following the planned course: %d\n", navigate(commands))
}