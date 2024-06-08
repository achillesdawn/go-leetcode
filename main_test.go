package main

import "testing"

func TestCase1(t *testing.T) {
	nums := []int{1, 2, 3, 6, 2, 3, 4, 7, 8}
	result := isNStraightHand(nums, 3)

	if !result {
		t.Errorf("should true")
	}
}

func TestCase2(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	result := isNStraightHand(nums, 4)

	if result {
		t.Errorf("should false")
	}
}

func TestCase3(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 1}
	result := isNStraightHand(nums, 2)

	if result {
		t.Errorf("should false")
	}
}

func TestCase4(t *testing.T) {
	nums := []int{1, 2, 3, 4, 7, 8}
	result := isNStraightHand(nums, 2)

	if !result {
		t.Errorf("should true")
	}
}

func TestCase5(t *testing.T) {
	nums := []int{1, 1, 2, 2, 3, 3}
	result := isNStraightHand(nums, 3)

	if !result {
		t.Errorf("should true")
	}
}

func TestCase6(t *testing.T) {
	nums := []int{0, 0}
	result := isNStraightHand(nums, 2)

	if result {
		t.Errorf("should false")
	}
}

func TestCase7(t *testing.T) {
	nums := []int{34, 80, 89, 15, 38, 69, 19, 17, 97, 98, 26, 77, 8, 31, 79, 70, 103, 3, 13, 21, 81, 53, 33, 14, 60, 68, 33, 59, 84, 23, 97, 90, 76, 82, 66, 83, 23, 22, 16, 18, 98, 25, 16, 61, 84, 100, 4, 68, 101, 25, 23, 9, 10, 55, 2, 67, 39, 52, 102, 99, 40, 11, 83, 24, 81, 53, 96, 23, 13, 24, 99, 67, 22, 51, 31, 58, 78, 88, 5, 15, 24, 32, 81, 91, 96, 16, 54, 22, 56, 69, 14, 82, 32, 34, 83, 24, 37, 82, 54, 21}
	result := isNStraightHand(nums, 4)

	if !result {
		t.Errorf("should be true")
	}
}

func TestCase8(t *testing.T) {
	nums := []int{1, 1, 2, 2, 3, 3}
	result := isNStraightHand(nums, 4)

	if result {
		t.Errorf("should be false")
	}
}

func TestCase9(t *testing.T) {
	nums := []int{5, 1, 0, 6, 4, 5, 3, 0, 8, 9}
	result := isNStraightHand(nums, 2)

	if result {
		t.Errorf("should be false")
	}
}
