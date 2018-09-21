package main

import (
	"fmt"
	"strings"
)

func map1() {
	results := map[string]int{}
	for i := 0; i < 100000; i++ {
		m := make(map[int]string)
		m[1] = "A"
		m[2] = "B"
		m[3] = "C"
		m[4] = "D"
		m[5] = "E"
		m[6] = "F"
		m[7] = "G"
		m[8] = "H"
		m[9] = "I"
		m[10] = "J"
		m[11] = "K"
		m[12] = "L"
		m[13] = "M"
		m[14] = "N"
		m[15] = "O"
		m[16] = "P"
		m[17] = "Q"
		m[18] = "R"
		m[19] = "S"
		m[20] = "T"
		m[21] = "U"
		m[22] = "V"
		m[23] = "W"
		m[24] = "X"
		m[25] = "Y"
		m[26] = "Z"

		str := ""

		for _, v := range m {
			str += v
		}
		results[str]++
	}
	fmt.Println(len(results))
}

func map2() {
	m := map[string]int{
		"G": 7, "A": 1, "C": 3, "E": 5,
		"D": 4, "B": 2, "F": 6, "I": 9, "H": 8,
	}
	counts := map[string]int{}
	for i := 0; i < 1000000; i++ {
		var order strings.Builder
		for k, _ := range m {
			order.WriteString(k)
		}
		counts[order.String()]++
	}
	fmt.Println(len(counts))
}

func main() {
	map1()
	map2()
}
