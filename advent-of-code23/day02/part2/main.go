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

	for {
		greenCubes := 0
		redCubes := 0
		blueCubes := 0
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
			regex, err := regexp.Compile(regexPattern)
			if err != nil {
				fmt.Println("Error compiling regex:", err)
				return
			}
			if err != nil {
				fmt.Println("Error compiling regex:", err)
				return
			}
			cubeSets := strings.Split(line, ";")
			for _, cubeSet := range cubeSets {
				matches := regex.FindAllString(cubeSet, -1)
				for _, match := range matches {
					number, err := strconv.Atoi(strings.Split(match, " ")[0])
					if err != nil {
						fmt.Println("Error converting string to int:", err)
						return
					}
					color := strings.Split(match, " ")[1]
					switch color {
					case "green":
						if number > greenCubes {
							greenCubes = number
						}
					case "red":
						if number > redCubes {
							redCubes = number
						}
					case "blue":
						if number > blueCubes {
							blueCubes = number
						}
					}
				}

			}
			sum += greenCubes * redCubes * blueCubes
		}
	}
	fmt.Println("Sum:", sum)
}
