package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Pizza is the best vegetable"

	fmt.Println(SpinWords(str))
}

func SpinWords(str string) string {
	words := strings.Fields(str)
	var anew string

	for i := 0; i != len(words); i++ {
		if len(words[i]) >= 5 {
			runes := []rune(words[i])
			for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
				runes[i], runes[j] = runes[j], runes[i]
			}
			anew += (string(runes) + " ") //rehtona
		} else {
			anew += words[i] + " "
		}
	}

	return anew[:len(anew)-1]
}
