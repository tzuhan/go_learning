package main

import "fmt"

func main76() {
	s, t := "ADOBECODEBANC", "ABC"
	fmt.Printf("s:%s t:%s\n", s, t)
	r := minWindow(s, t)
	fmt.Printf("result: %s\n", r)
}
func minWindow(s string, t string) string {
	//first create map[char]int of string to store character-count info
	m := make(map[byte]int)
	for _, val := range t {
		m[byte(val)]++
	}
	fmt.Println(m)

	/*We defined two pointers(start & end of the window) to walk through string s.
	  First we move end forward till the window contains all characters of string t (use a variable 'count' to check that), if we encounter charcters of string t, we subtract 1 from the number of the character in map

	*/
	start, end, count, minWindowLen, minStart := 0, 0, len(t), len(s)+1, 0
	for end < len(s) {
		fmt.Printf("substring S: %s count:%d\n", s[start:end+1], count)
		_, ok1 := m[s[end]]
		if ok1 {
			if m[s[end]] > 0 {
				count--
			}
			m[s[end]]--
			fmt.Printf("end: %d, ", end)
			fmt.Println(m)
		}
		end++
		for count == 0 {
			if end-start < minWindowLen {
				minStart, minWindowLen = start, end-start
				fmt.Printf("minStart, minLen update: %d %d\n", minStart, minWindowLen)
			}
			_, ok := m[s[start]]
			if ok {
				m[s[start]]++
				if m[s[start]] > 0 {
					count++
				}
				fmt.Printf("start: %d, count:%d\n", start, count)
				fmt.Println(m)
			}
			start++
		}
	}
	if minWindowLen != len(s)+1 {
		return s[minStart : minStart+minWindowLen]
	} else {
		return ""
	}
}
