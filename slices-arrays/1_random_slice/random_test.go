package main

import (
	"reflect"
	"testing"
)

func TestRandom(t *testing.T) {
	testingNums := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	resultNums := random(testingNums)

	if len(testingNums) != len(resultNums) {
		t.Errorf("Expected an equal length of numbers but got %v", len(resultNums))
	}
	if reflect.DeepEqual(testingNums, resultNums) == true {
		t.Errorf("Expected the slices to not be the same but they are. testingNums: %v resultNums: %v", testingNums, resultNums)

	}

}