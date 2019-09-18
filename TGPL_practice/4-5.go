package main

import (
	"fmt"
)

func delete_adjacent_redundent(slice []string) []string {
	if len(slice) <= 1 {
		return slice
	}

	i, tmpstr := 1, ""
	for idx, str := range slice {
		if idx == 0 {
			tmpstr = str
		} else {
			if str != tmpstr {
				slice[i] = str
				tmpstr = str
				i++
			}
		}
	}
	return slice[:i]
}

func main4_5() {
	input := []string{"ads", "ads", "", "b", "c", "c", "d"}
	fmt.Println(input)
	input = delete_adjacent_redundent(input)
	fmt.Println(input)
}
