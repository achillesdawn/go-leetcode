package main

import (
	"fmt"
	"slices"
)

func isNStraightHand(hand []int, groupSize int) bool {

	fmt.Println("hand", hand, "size", groupSize)

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
		}
		fmt.Println("checking key", key, "for", count, "times", numMap)

		for range count {
			check := key
			for range groupSize {
				fmt.Println("checking key", check)
				value, exists := numMap[check]

				if !exists {
					fmt.Println(check, "does not exist")
					return false
				}

				if value < 0 {
					fmt.Println("less than zero")
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
	nums := []int{1, 1, 2, 2, 3, 3, 3, 4}
	result := isNStraightHand(nums, 2)
	fmt.Println(result)
}
