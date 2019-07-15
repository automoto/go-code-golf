package main

import (
	"math/rand"
	"time"
)

func random(nums []int) []int {
	newNums := make([]int, len(nums))
	copy(newNums, nums)

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i:= range nums {
		newPosition := r.Intn(len(newNums))

		newNums[i], newNums[newPosition] = newNums[newPosition], newNums[i]
	}
	return newNums
}