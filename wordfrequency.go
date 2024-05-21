package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter the block of code: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	// fmt.Println(input)
	input= strings.TrimSpace(input)
	s := strings.Split(input," ")


	// fmt.Println(s[0])
	completedWords := make(map[string]int)
	for _, word := range s {
		completedWords[word]++
	}

	for word, count := range completedWords {
		fmt.Println("Word: ", word, "Count:", count)
	}
}