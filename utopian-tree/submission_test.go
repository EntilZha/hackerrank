package main

import "testing"

func TestCalculateHeight(t *testing.T) {
	if calculateHeight(0) != 1 {
		Fail(t)
	}
	if calculateHeight(1) != 2 {
		Fail(t)
	}
	if calculateHeight(3) != 6 {
		Fail(t)
	}
	if calculateHeight(4) != 7 {
		Fail(t)
	}
}

func Fail(t *testing.T) {
	t.Errorf("Wrong answer")
}
