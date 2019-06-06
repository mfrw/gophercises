package main

func Search(n int, s []int) int {
	r := len(s)
	l := 0
	for l <= r {
		m := l + (r-l)/2
		if s[m] == n {
			return m
		} else if s[m] < n {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}

func LeftMost(n int, s []int) int {
	idx := Search(n, s)
	for idx > 0 {
		if s[idx-1] == n {
			idx--
		} else {
			break
		}
	}
	return idx
}

func RightMost(n int, s []int) int {
	idx := Search(n, s)
	for idx < len(s)-1 {
		if s[idx+1] == n {
			idx++
		} else {
			break
		}
	}
	return idx
}
