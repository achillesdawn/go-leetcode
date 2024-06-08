package main

import (
	"fmt"
	"slices"
)

func isNStraightHand(hand []int, groupSize int) bool {

	totalGroups := len(hand) / groupSize

	if totalGroups < 1 {
		return false
	} else if totalGroups*groupSize != len(hand) {
		return false
	}

	var numMap map[int]int = make(map[int]int)

	for _, item := range hand {
		value, exists := numMap[item]
		if exists {
			numMap[item] = value + 1
		} else {
			numMap[item] = 1
		}
	}

	keys := make([]int, 0, len(numMap))
	for k := range numMap {
		keys = append(keys, k)
	}

	slices.Sort(keys)

	for _, key := range keys {

		count := numMap[key]

		if count == 0 {
			continue
		} else if count < 0 {
			return false
		}

		for range count {
			check := key
			for range groupSize {
				value, exists := numMap[check]

				if !exists || value < 0 {
					return false
				}

				numMap[check] = value - 1
				check += 1
			}
		}
	}
	return true
}

func main() {
	nums := []int{5, 1, 0, 6, 4, 5, 3, 0, 8, 9}
	result := isNStraightHand(nums, 2)
	fmt.Println(result)
}
