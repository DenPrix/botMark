package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	text, _ := readText("book.txt")
	states := markov.train(text)

	for {
		var testWord string
		print("Input: ")
		reader := bufio.NewReader(os.Stdin)
		testWord , _ = reader.ReadString('\n')


		words := textToWords(testWord)
		last := strings.TrimRight(words[len(words)-1], "\n")
		generatedText := markov.generateText(states, last, 20)
		println("Output: " + generatedText + "\n")
	}
}
