package main

import (
	"fmt"
	"strings"
)

func main() {
	personages := map[string]int{
		"Dead Duck": 22,
		"Bob Cuspe": 27,
	}

	// OR
	// personages := make(map[string]int)
	// personages["Dead Duck"] = 22
	// personages["Bob Cuspe"] = 27

	personages["Rê Bordosa"] = 31

	for name, age := range personages {
		fmt.Printf("Name: %s ===>>> Age: %d\n", name, age)
	}
	fmt.Println(strings.Repeat("<>", 40))

	p := "Dead Duc"
	if personage, ok := personages[p]; !ok {
		fmt.Printf(`"%s" não é uma chave nesse map!`+"\n", p)
	} else {
		fmt.Printf("Dead Duck age is %d\n", personage)
	}
	fmt.Println(strings.Repeat("+-", 40))

	fmt.Printf("Total elements: %d\n", len(personages))
	fmt.Println(strings.Repeat("><", 40))

	keys := make([]string, 0, len(personages))
	for k := range personages {
		keys = append(keys, k)
	}
	fmt.Println(strings.Join(keys, "\n"))
}
