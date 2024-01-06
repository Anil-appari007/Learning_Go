package main

import (
	"testing"
)

type addTest struct {
	arg1, arg2, exp int
}

func TestAddTwoNum(t *testing.T) {
	got := addTwoNum(1, 4)
	exp := 5

	if got != exp {
		t.Errorf("Got: %d. Expected: %d", got, exp)
	}

	TestData := []addTest{
		addTest{1, 2, 3},
		addTest{0, 9, 9},
		addTest{9, 8, 17},
	}

	for _, each := range TestData {
		if res := addTwoNum(each.arg1, each.arg2); res != each.exp {
			t.Errorf("Result: %d. Expected: %d", res, each.exp)
		}
	}
}

func TestSubTwoNum(t *testing.T) {
	got := subTwoNum(9, 8)
	exp := 1
	if got != exp {
		t.Errorf("Got: %d, Expected: %d", got, exp)
	}
}
