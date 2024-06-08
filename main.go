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

	slices.Sort[[]int](hand)

	end := len(hand)

	for range totalGroups {

		current := hand[0]
		hand = hand[1:]
		end -= 1

		for range groupSize - 1 {

			idx, found := slices.BinarySearch(hand[:end], current+1)

			if !found {
				return false
			}

			slices.Delete(hand, idx, idx+1)
			end -= 1
			current += 1
		}
	}

	return true
}

func main() {
	nums := []int{1, 2, 3, 4, 7, 8}
	result := isNStraightHand(nums, 2)
	fmt.Println(result)
}
