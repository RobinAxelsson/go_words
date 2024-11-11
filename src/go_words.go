package go_words

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var filename string = "wordlist.txt"
var actions []string = []string{"add", "list", "remove", "rnum"}

func Run(osArgs []string) {

	action := ""

	if len(osArgs) < 2 {
		fmt.Printf("wordlist <action>\nvalid actions: %s\n", actions)
	} else {
		action = osArgs[1]
	}

	if action != "" && !contains(actions, action) {
		fmt.Printf("action: \"%s\" is not present\n", action)
		fmt.Printf("valid actions: %s\n", actions)
	}

	if action != "" {
		words := getWords()
		actionArgs := osArgs[2:]

		if action == "add" {
			added := add(actionArgs)
			if added != "" {
				words = append(words, added)
			}
		}

		if action == "remove" {
			words = remove(actionArgs, words)
			saveWords(words)
		}

		if action == "rnum" {
			words = removeNum(actionArgs, words)
			saveWords(words)
		}

		//list & rest
		printWords(words)
	}

	fmt.Println("End of program")
}

func removeNum(args []string, words []string) []string {
	if len(args) != 1 {
		fmt.Println("wordlist rnum <number> to remove word by number")
		return words
	}

	if len(words) == 0 {
		return words
	}

	numArg := args[0]

	num, err := strconv.Atoi(numArg)
	if err != nil || num < 1 {
		fmt.Println("input an integer")
		return words
	}

	if len(words) < num {
		fmt.Println("number out of range")
		return words
	}

	return append(words[:num-1], words[num:]...)
}

func remove(args []string, words []string) []string {
	if len(args) < 1 {
		fmt.Println("provide word or phrase to remove")
		return words
	}

	if len(words) == 0 {
		return words
	}

	word := strings.Join(args, " ")

	index := -1

	for i, w := range words {
		if w == word {
			index = i
			break
		}
	}

	if index == -1 {
		return words
	}

	return append(words[:index], words[index+1:]...)
}

func add(args []string) string {
	if len(args) == 0 {
		fmt.Println("provide word or phrase to add")
		return ""
	}

	word := strings.Join(args, " ")

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	defer file.Close()

	_, err = file.Write([]byte(";" + word))

	if err != nil {
		fmt.Println("Error writing to file:", err)
	}

	return word
}

func printWords(words []string) {
	if len(words) == 0 {
		fmt.Println("No words saved, use add action")
		return
	}

	fmt.Println("Wordlist:")

	for i := 0; i < len(words); i++ {
		fmt.Printf("%v. %s\n", i+1, words[i])
	}
}

func contains(words []string, word string) bool {
	for _, v := range words {
		if v == word {
			return true
		}
	}
	return false
}

func saveWords(words []string) error {
	file, err := os.OpenFile(filename, os.O_TRUNC|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		return err
	}

	defer file.Close()

	for _, word := range words {
		file.Write([]byte(";" + word))
	}

	return err
}

func getWords() []string {
	file, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0644)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	filesize := fileInfo.Size()

	if filesize == 0 {
		return []string{}
	}

	buffer := make([]byte, filesize)

	_, err = file.Read(buffer)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	file.Close()

	content := string(buffer)

	return strings.Split(strings.Trim(content, ";"), ";")
}
