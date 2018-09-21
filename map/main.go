package main

import "fmt"

func main() {
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
