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
    position := 0
    depth := 0
    
    for r = 0; r <= len(route)-1; r++ {
        routeRawValues := strings.Fields(route[r])
        
        direction := routeRawValues[0]
        units, _ := strconv.Atoi(routeRawValues[1])

        if direction == "forward" {
            position = position + units
        } else if direction == "down" {
            depth = depth + units
        } else if direction == "up" {
            depth = depth - units
        }
        
		fmt.Printf("Direction: %s Units: %d\n", direction, units)
	}

    return depth * position
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
    var commands []string
	absPath, _ := filepath.Abs("../2021/util/data/day02Input.txt")
	lines, _ := readFile(absPath)

	commands = append(commands, lines...)
    
    fmt.Printf("Horizontal position and depth after following the planned course: %d\n", navigate(commands))
}