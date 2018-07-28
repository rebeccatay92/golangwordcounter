package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Please enter, or paste input text body. To end input, hit Esc on a new line and then Enter.\n")

	// var lines []string
	var textbody string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 1 {
			// ascii hexadecimal for Esc
			if line[0] == '\x1B' {
				break
			}
		}
		// lines = append(lines, line)
		textbody += line
	}

	// if len(lines) > 0 {
	// 	fmt.Println()
	// 	fmt.Println("Result:")
	// 	for _, line := range lines {
	// 		fmt.Println("line", line)
	// 	}
	// 	fmt.Println()
	// }

	if len(textbody) > 0 {
		fmt.Println("Textbody:")
		fmt.Println(textbody)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading std input:", err)
	}

	// words := strings.Fields(textbody)
	// fmt.Println(words, len(words))

	// refactor to use Replacer instead, to target many chars.
	textbody = strings.ToLower(textbody)
	textbody = strings.Replace(textbody, ".", "", -1)
	textbody = strings.Replace(textbody, ",", "", -1)
	fmt.Println("Cleaned:")
	fmt.Println(textbody)
}
