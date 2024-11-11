package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var filename string = "wordlist.txt"
var actions []string = []string{"add", "list", "remove", "rnum"}
var words []string

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("wordlist <action>\nvalid actions: %s\n", actions)
		os.Exit(1)
	}

	var action string = os.Args[1]

	if !contains(actions, action) {
		fmt.Printf("action: \"%s\" is not present\n", action)
		fmt.Printf("valid actions: %s\n", actions)
		os.Exit(1)
	}

	if action == "add" {
		addWord()
		printWords()
	}

	if action == "remove" {
		removeWord()
		printWords()
	}

	if action == "rnum" {
		removeNum()
		printWords()
	}

	if action == "list" {
		printWords()
	}

	fmt.Println("End of program")
}

func removeNum() {
	if len(os.Args) < 3 || 3 < len(os.Args) {
		fmt.Println("wordlist rnum <number> to remove word by number")
		os.Exit(1)
	}

	num, err := strconv.Atoi(os.Args[2])
	if err != nil || num < 1 {
		fmt.Println("input an integer")
		os.Exit(1)
	}

	words := getWords()

	if len(words) < num {
		fmt.Println("number out of range")
		os.Exit(1)
	}

	file, err := os.OpenFile(filename, os.O_TRUNC|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	for i, word := range words {
		if i+1 != num {
			file.Write([]byte(";" + word))
		}
	}
}

func removeWord() {
	if len(os.Args) < 3 {
		fmt.Println("provide words to add")
		os.Exit(1)
	}

	wordToRemove := strings.Join(os.Args[2:], " ")

	words := getWords()

	if contains(words, wordToRemove) {

		file, err := os.OpenFile(filename, os.O_TRUNC|os.O_WRONLY, 0644)

		if err != nil {
			fmt.Printf("Error opening file: %s\n", err)
			os.Exit(1)
		}

		defer file.Close()

		for _, word := range words {
			file.Write([]byte(";" + word))
		}
	}
}

func addWord() {
	if len(os.Args) < 3 {
		fmt.Println("provide words to add")
		os.Exit(1)
	}

	wordToAdd := strings.Join(os.Args[2:], " ")

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	file.Write([]byte(";" + wordToAdd))
}

func printWords() {
	data, err := os.ReadFile(filename)

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No words, use add action")
		} else {
			fmt.Println("Error:", err)
		}
		os.Exit(1)
	}

	if len(data) == 0 {
		fmt.Println("No words added, use add action")
		os.Exit(1)
	}

	fmt.Println("Wordlist:")

	words := getWords()

	for i := 0; i < len(words); i++ {
		fmt.Printf("%v. %s\n", i+1, words[i])
	}
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func getWords() []string {
	if words != nil {
		return words
	}

	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)

	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		os.Exit(1)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	filesize := fileInfo.Size()

	if filesize == 0 {
		fmt.Println("File is empty")
		os.Exit(1)
	}

	buffer := make([]byte, filesize)

	_, err = file.Read(buffer)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	file.Close()

	content := string(buffer)
	words := strings.Split(strings.Trim(content, ";"), ";")

	return words
}
