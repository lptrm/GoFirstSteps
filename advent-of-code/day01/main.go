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
	var input string
	var sum int

	regexPattern := "[A-Z,a-z]"
	numberPatterns := "[one,two,three,four,five,six,seven,eight,nine,ten,eleven]"
	numberSuffix := "[teen,ty]"

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
			regex, err := regexp.Compile(regexPattern)
			if err != nil {
				fmt.Println("Error compiling regex:", err)
				return
			}
			result := strings.TrimSpace(regex.ReplaceAllString(line, ""))
			if len(result) == 1 {
				result += result

			} else {
				result = string(result[0]) + string(result[len(result)-1])
			}
			number, err := strconv.Atoi(result)
			if err != nil {
				fmt.Println("Error", err)
				return
			}
			sum += number
			input += result + "\n"
		}
	}

	fmt.Println("You Entered:", input)
	fmt.Println("The Sum is:", sum)

}
