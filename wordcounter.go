package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Please enter, or paste input text body. To end input, hit Esc and then Enter.\n")

	var textbody string
	scanner := bufio.NewScanner(os.Stdin)

	// for each line that is scanned
	for scanner.Scan() {
		line := scanner.Text()

		// concatenate the strings, inserting a space to seperate lines. else 'abc', 'def' will become 'abcdef'
		textbody += line + " "

		// check if last char is Esc key, and stop scanning
		if len(line) >= 1 && line[len(line)-1] == '\x1B' {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading stdin:", err)
	}

	if len(textbody) > 0 {
		fmt.Println("Textbody:")
		fmt.Println(textbody)
	} else {
		fmt.Println("No text input detected")
	}

	// refactor to use Replacer instead, to target many chars.
	textbody = strings.ToLower(textbody)
	textbody = strings.Replace(textbody, ".", " ", -1)
	textbody = strings.Replace(textbody, ",", " ", -1)
	fmt.Println("Cleaned:")
	fmt.Println(textbody)

	// words := strings.Fields(textbody)
	// fmt.Println(words, len(words))
}
