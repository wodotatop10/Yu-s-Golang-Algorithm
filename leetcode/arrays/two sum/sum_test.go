package main

import (
	"fmt"
	"testing"
)

var i int
var nums []int

func init() {
	max := 10000000
	i = 677787
	nums = RandomIntArray(max)
}
func BenchmarkSumArray(t *testing.B) {
	result := twoSumBadWay(nums, i)
	// fmt.Println("result tag=", result)
	if result != nil {
		fmt.Println("result num=", nums[result[0]], nums[result[1]])
	} else {
		fmt.Println("result[] is nil")
	}
}

func BenchmarkSumMap(t *testing.B) {
	result := towSumMapFindBest(nums, i)
	fmt.Println("result tag=", result)
}
