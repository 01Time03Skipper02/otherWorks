package main

import (
	"fmt"
	"strings"
)

func DuplicateEncode(word string) string {
	var res string
	res = ""
	runes := []rune(word)
	for i := 0; i < len(word); i++ {
		//fmt.Println(string(runes[i]), word)
		if strings.Count(word, string(runes[i])) > 1 || strings.Count(word, strings.ToLower(string(runes[i]))) > 1 || strings.Count(word, strings.ToUpper(string(runes[i]))) > 1 {
			res += ")"
		} else {
			res += "("
		}
	}
	return res
}

func main() {
	a := "Aaara"
	fmt.Println(DuplicateEncode(a))
}
