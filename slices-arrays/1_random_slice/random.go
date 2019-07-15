package main

import (
	"math/rand"
	"time"
)

func random(nums []int) []int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i:= range nums {
		newPosition := r.Intn(len(nums))

		nums[i], nums[newPosition] = nums[newPosition], nums[i]
	}
	return nums
}