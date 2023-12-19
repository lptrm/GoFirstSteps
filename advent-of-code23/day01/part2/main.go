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
	var sum int

	regexPattern := "(one|two|three|four|five|six|seven|eight|nine|[1-9])"
	regexPatternRev := "(eno|owt|eerht|ruof|evif|xis|neves|thgie|enin|[1-9])"
	reader := bufio.NewReader(os.Stdin)

	for {
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

			// Compile the regex pattern
			regex, err := regexp.Compile(regexPattern)
			if err != nil {
				fmt.Println("Error compiling regex:", err)
				return
			}
			regexRev, err := regexp.Compile(regexPatternRev)
			if err != nil {
				fmt.Println("Error compiling regex:", err)
				return
			}

			matchesFirst := regex.FindString(line)
			lineRev := reverse(line)
			matchesLastRev := regexRev.FindString(lineRev)
			matchesLast := reverse(matchesLastRev)
			resultString := parse(matchesFirst) + parse(matchesLast)
			result, err := strconv.Atoi(resultString)

			if err != nil {
				fmt.Println("Error", err)
				return
			}
			sum += result
		}
	}
	fmt.Println("Sum:", sum)
}
func reverse(input string) string {
	result := ""
	for _, char := range input {
		result = string(char) + result
	}
	return result
}
func parse(input string) string {
	switch input {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"

	default:
		return input
	}
}
