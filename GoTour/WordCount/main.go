package main

// Implementar WordCount.
// Ele deve retornar um map das contagens de cada "word" na string s.
// A função wc.Test executa um conjunto de testes contra a função fornecida e imprime o sucesso ou falha.

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	count_word := make(map[string]int)

	var count int
	for _, item := range words {
		_, ok := count_word[item]
		if ok {
			continue
		}
		count = 0
		for i := range words {
			if item == words[i] {
				count++
			}
		}
		count_word[item] = count
	}
	//fmt.Println(count_word)
	return count_word
}

func main() {
	wc.Test(WordCount)
}
