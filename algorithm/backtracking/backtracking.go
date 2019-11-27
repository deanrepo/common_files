package main

/*
Backtracking algorithm - Combinational Sum
Given an array: [1,2] and a target: 4
Find the solution set that adds up to the target
in this case:
[1,1,1,1]
[1,1,2]
[2,2]
*/

import (
	"fmt"
	"sort"
)

func main() {
	res := combinationSum([]int{1, 7, 3}, 13)
	for _, v := range res {
		fmt.Println(v)
	}
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	return combine(0, target, []int{}, candidates)
}

func combine(addSum, target int, curCombinations []int, candidates []int) [][]int {
	var result [][]int
	if addSum == target {
		tempRes := make([]int, len(curCombinations))
		copy(tempRes, curCombinations)
		return [][]int{tempRes}
	} else if addSum < target {
		for i, v := range candidates {
			tempCombs := append(curCombinations, v)
			temp := combine(addSum+v, target, tempCombs, candidates[i:])
			result = append(result, temp...)
		}
	}
	return result
}
