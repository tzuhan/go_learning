package main

import "fmt"

func main28() {
	haystack, needle := "hello", "ll"
	pmt := make([]int, len(needle)+1)
	pmt = makePMT(needle, pmt)
	fmt.Println(haystack)
	fmt.Println(needle)
	fmt.Println(pmt)
	firstIndex := checkKMP(haystack, needle, pmt)
	fmt.Println("first Index", firstIndex)
}
func checkKMP(haystack string, needle string, pmt []int) int {
	i, j := 0, 0
	for i < len(haystack) && j < len(needle) {
		fmt.Println("i:%d j:%d", i, j)
		if j == -1 || haystack[i] == needle[j] {
			i++
			j++
		} else {
			j = pmt[j]
		}
	}
	if j == len(needle) {
		return i - j
	} else {
		return -1
	}
}
func makePMT(needle string, table []int) []int {
	table[0] = -1
	//the value of pmt[j] is the max length of substring from the intersection of prefix and postfix substring union of needle[0 ~ j-1]
	//so if haystack[i] is not equal to needle[j], we know that haystack[i-j ~ i-1] is equal to needle[0 ~ j-1]
	//we'll move j to pmt[j] to start comparing the next character after the max substring
	for i, j := 0, -1; i < len(needle); {
		if j == -1 || needle[i] == needle[j] {
			i++
			j++
			table[i] = j
		} else {
			//when not equal
			j = table[j]
		}
	}
	return table
}
