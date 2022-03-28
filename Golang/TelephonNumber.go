package main

import (
	"fmt"
	"strconv"
)

func CreatePhoneNumber(numbers [10]uint) string {
	res := "("
	for i := 0; i < 3; i++ {
		res += strconv.FormatInt(int64(numbers[i]), 10)
	}
	res += ") "
	for i := 3; i < 6; i++ {
		res += strconv.FormatInt(int64(numbers[i]), 10)
	}
	res += "-"
	for i := 6; i < 10; i++ {
		res += strconv.FormatInt(int64(numbers[i]), 10)
	}
	return res
}

func main() {
	numb := [10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(CreatePhoneNumber(numb))
}
