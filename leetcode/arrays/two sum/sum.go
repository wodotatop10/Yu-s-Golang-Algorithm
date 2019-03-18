package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	max := 100
	i := 23
	nums := RandomIntArray(max)
	result := twoSumBadWay(nums, i)
	fmt.Println("result tag=", result)
	if result != nil {
		fmt.Println("result num=", nums[result[0]], nums[result[1]])
	} else {
		fmt.Println("result[] is nil")
	}

	result = towSumMapFindBest(nums, i)
	fmt.Println("result tag=", result)
	// fmt.Println("result num=", nums[result[0]], nums[result[1]])
}

//生成随机数
func RandomIntArray(max int) []int {
	rand.Seed(time.Now().Unix())
	nums := make([]int, max-1)
	for i, _ := range nums {
		nums[i] = rand.Intn(max)
	}
	return nums
}

//暴力破解
func twoSumBadWay(nums []int, target int) []int {
	l := len(nums)
	for i, _ := range nums {
		for j := i + 1; j < l; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

//hashmap 查找
//要求输入数组不能有重复的
func towSumMapFind(nums []int, target int) []int {
	numMap := map[int]int{}
	for index, value := range nums {
		numMap[value] = index
	}

	for index, value := range nums {
		com := target - value
		_, exists := numMap[com]
		if numMap[com] != index && exists {
			return []int{value, com}
		}
	}
	return nil
}

//map find with only one looper
func towSumMapFindBest(nums []int, target int) []int {
	numMap := map[int]int{}
	for index, value := range nums {
		component := target - value
		_, exists := numMap[component]

		if exists {
			return []int{value, component}
		} else {
			numMap[value] = index
		}
	}
	return nil
}
