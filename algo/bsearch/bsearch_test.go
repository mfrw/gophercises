package main

import "testing"

func TestSearch(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	idx := Search(3, s)
	if idx != 2 {
		t.Fail()
	}
}

func TestLeftMost(t *testing.T) {
	s := []int{1, 2, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 7}
	idx := LeftMost(4, s)
	if idx != 3 {
		t.Fail()
	}
}

func TestRightMost(t *testing.T) {
	s := []int{1, 2, 3, 3, 3, 3, 3, 3, 4}
	idx := RightMost(3, s)
	if idx != 7 {
		t.Fail()
	}
}
