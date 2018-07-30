package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	fmt.Printf("Please enter, or paste input text body. To end input, hit Esc and then Enter.\n")

	var textbody string
	scanner := bufio.NewScanner(os.Stdin)

	// for each line that is scanned
	for scanner.Scan() {
		line := scanner.Text()

		// concatenate the strings, inserting a space to separate lines
		textbody += line + " "

		// check if last char is Esc key, and stop scanning
		if len(line) >= 1 && line[len(line)-1] == '\x1B' {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading stdin:", err)
	}

	// convert to lower case, remove esc key
	textbody = strings.ToLower(textbody)
	textbody = strings.Replace(textbody, "\x1b", " ", -1)

	// strip off punctuation
	regex, _ := regexp.Compile("[^a-z0-9]")
	textbody = regex.ReplaceAllString(textbody, " ")

	wordsSlice := strings.Fields(textbody)

	if len(wordsSlice) < 1 {
		fmt.Printf("There are no words found.\n")
		return
	}

	frequencyMap := make(map[string]int)

	for _, word := range wordsSlice {
		_, found := frequencyMap[word]
		if found {
			frequencyMap[word]++
		} else if len(word) >= 1 {
			frequencyMap[word] = 1
		}
	}

	// turn word, count pairs into sortable arr
	type Pair struct {
		word  string
		count int
	}
	var keyValueArr []Pair
	for key, value := range frequencyMap {
		keyValueArr = append(keyValueArr, Pair{key, value})
	}

	sort.Slice(keyValueArr, func(i, j int) bool {
		return keyValueArr[i].count > keyValueArr[j].count
	})

	if len(keyValueArr) < 10 {
		fmt.Println("There are less than 10 unique words")
		for _, Pair := range keyValueArr {
			fmt.Printf("%s appears %d times\n", Pair.word, Pair.count)
		}
	} else if len(keyValueArr) >= 10 {
		fmt.Println("The 10 most frequent words are")
		for i := 1; i <= 10; i++ {
			fmt.Printf("%s appears %d times\n", keyValueArr[i].word, keyValueArr[i].count)
		}
	}

}
