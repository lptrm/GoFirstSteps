package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	sum := 0
	regexPattern := `\b\d{1,2}\s\w+\b`
	reader := bufio.NewReader(os.Stdin)

	greenCubesConstraint := 13
	redCubesConstraint := 12
	blueCubesConstraint := 14

	gameCounter := 0

	for {
		valid := true
		line, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		} else {
			if strings.TrimSpace(line) == "end" {
				break
			}
			if strings.TrimSpace(line) == "" {
				continue
			}
			gameCounter += 1
			regex, err := regexp.Compile(regexPattern)
			if err != nil {
				fmt.Println("Error compiling regex:", err)
				return
			}
			if err != nil {
				fmt.Println("Error compiling regex:", err)
				return
			}
			cubeSet := strings.Split(line, ";")
			for _, cube := range cubeSet {
				greenCubes := 0
				redCubes := 0
				blueCubes := 0
				matches := regex.FindAllString(cube, -1)
				for _, match := range matches {
					number, err := strconv.Atoi(strings.Split(match, " ")[0])
					if err != nil {
						fmt.Println("Error converting string to int:", err)
						return
					}
					color := strings.Split(match, " ")[1]
					switch color {
					case "green":
						greenCubes += number
					case "red":
						redCubes += number
					case "blue":
						blueCubes += number
					}
				}
				if greenCubes > greenCubesConstraint || redCubes > redCubesConstraint || blueCubes > blueCubesConstraint {
					valid = false
				}
			}
			if valid {
				sum += gameCounter
			}

		}
	}
	fmt.Println("Sum:", sum)
}
