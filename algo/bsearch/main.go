package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const MAX = 1024 * 1024 * 128

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	s := make([]int, MAX)
	for i := 0; i < MAX; i++ {
		s[i] = rand.Intn(100)
	}
	sort.Ints(s)
	r := rand.Intn(100)
	fmt.Printf("Linear: %d found %d times\n", r, CountIntLinear(r, s))
	fmt.Printf("Binary: %d found %d times\n", r, binary(r, s))
}

func binary(n int, s []int) int {
	l := LeftMost(n, s)
	r := RightMost(n, s)
	if l < 0 {
		return -1
	}
	return r - l
}

func CountIntLinear(n int, s []int) int {
	count := 0
	for _, v := range s {
		if n == v {
			count++
		}
	}
	return count
}
