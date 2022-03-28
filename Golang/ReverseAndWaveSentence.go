package main

import (
	"fmt"
	"strings"
)

//Function reverse of the sentence
func Reverse(s string) string {
	arr := strings.Split(s, " ")
	res := ""
	for a := 0; a < len(arr); a++ {
		runes := []rune(arr[a])
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		res += string(runes)
		if a != len(arr)-1 {
			res += " "
		}
	}
	return res
}

//make wave function
func wave(words string) []string {
	result := []string{}
	for i := 0; i < len(words); i++ {
		runes := []rune(words)
		res := ""
		for j := 0; j < len(words); j++ {
			if i == j {
				res += strings.ToUpper(string(runes[j]))
			} else {
				res += string(runes[j])
			}
		}
		if string(runes[i]) == " " {
			continue
		} else {
			result = append(result, res)
		}
	}
	return result
}

func main() {
	//try to make wave
	fmt.Println(wave(" x yz"))
	//try to reverse sentence
	fmt.Println(Reverse("fsafas fsafas h"))
}
