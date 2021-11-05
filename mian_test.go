package main

import (
	"testing"
)

type pair struct {
	inputValue     int
	expectedResult int
}

type pairArrays struct {
	inputArray    []int
	expectedArray []int
}

var testingPair = []pair{
	{1000, 1024},
	{0, 0},
	{-123, 0},
	// {10000000000000000, 18014398509481984},
	{7, 8},
	{32, 64},
	{128, 256},
}

var testingArray = []pairArrays{
	{[]int{1, 2, 3, 4, 5}, []int{120, 24, 6, 2, 1}},
	{[]int{0, 0, 0, 0, -1, 5, 0}, []int{0, 0, 0, 0}},
	{[]int{}, []int{}},
	{[]int{-1, 2, 3, 4, 5}, []int{}},
	{[]int{100, 1, 1, -10, 1}, []int{100, 100, 100}},
}

func TestTask108(t *testing.T) {
	for _, pair := range testingPair {
		n := task108(pair.inputValue)
		if n != pair.expectedResult {
			t.Error(
				"For", pair.inputValue,
				"expected", pair.expectedResult,
				"got", n,
			)
		}

	}
}

// повільно працює з великими числами
// не працює з від'ємними (згідно з умовою задачі)
func TestTask331a(t *testing.T) {
	for _, item := range testingPair {
		a, b, c := task331a(item.inputValue)
		if !(((a*a)+(b*b)+(c*c) == item.inputValue) || (a+b+c == 0)) {
			t.Error("err for value", item.inputValue)
		}
	}
}

func TestTask165d(t *testing.T) {
	for _, item := range testingArray {
		res := task165d(item.inputArray)
		for i, val := range item.expectedArray {
			if val != res[i] {
				t.Error(
					"for", item.inputArray,
					"expected", item.expectedArray,
					"got", res,
				)
			}
		}
	}
}
