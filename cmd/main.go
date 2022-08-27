package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/Aksh-Bansal-dev/wordle/internal/color"
)

const loopCount = 6

var mode = flag.String("mode", "hard", "Specify game mode (easy/hard)")

func main() {
	flag.Parse()
	if *mode != "hard" && *mode != "easy" {
		fmt.Println("Invalid game mode")
		return
	}
	reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().UnixNano())
	reqWord := words[rand.Intn(len(words))]
	wordmap := map[string]bool{}
	for _, e := range words {
		wordmap[e] = true
	}
	for i := 0; i < loopCount; i++ {
		fmt.Print(strings.Repeat("-", len(reqWord)), "\r")
		word, _ := reader.ReadString('\n')
		word = word[:len(word)-1]
		if len(word) != len(reqWord) {
			fmt.Printf("The word must have %d letters\n", len(reqWord))
			i--
			continue
		}
		if _, present := wordmap[word]; *mode == "hard" && !present {
			fmt.Println("Invalid word!")
			i--
			continue
		}
		notUsed := map[byte]int{}
		for _, c := range reqWord {
			_, present := notUsed[byte(c)]
			if present {
				notUsed[byte(c)]++
			} else {
				notUsed[byte(c)] = 1
			}
		}
		var allCorrect = true
		for i := range word {
			if reqWord[i] == word[i] {
				notUsed[word[i]]--
			} else {
				allCorrect = false
			}
		}
		res := strings.Builder{}
		for i := range word {
			if reqWord[i] == word[i] {
				res.WriteString(color.Green(string(word[i])))
			} else {
				if _, ok := notUsed[word[i]]; ok && notUsed[word[i]] > 0 {
					notUsed[word[i]]--
					res.WriteString(color.Yellow(string(word[i])))
				} else {
					res.WriteString(color.Grey(string(word[i])))
				}
			}
		}
		fmt.Printf("#%d %s\n", i+1, res.String())
		if allCorrect {
			fmt.Println("Correct!")
			return
		}
	}
	fmt.Println("GAME OVER!")
	fmt.Println("The word was", reqWord)
	return
}
