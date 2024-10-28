package slice

import "github.com/LXJ0000/go-utils/setx"

// Shrink 函数用于根据一定的规则缩减切片的容量。
func Shrink[T comparable](nums []T) []T {
	calCapacity := func(c, l int) (int, bool) {
		if c <= 64 { // TODO const
			return c, false
		}
		if c <= 2048 && (c/l >= 4) {
			return int(float32(c) * 0.5), true
		}
		if c > 2048 && (c/l >= 2) {
			return int(float32(c) * 0.625), true
		}
		return c, false
	}
	c, l := cap(nums), len(nums)
	n, isChang := calCapacity(c, l)
	if !isChang {
		return nums
	}
	newNums := make([]T, 0, n)
	newNums = append(newNums, nums...)
	return newNums
}

func UniqueByLoop[T comparable](nums []T) []T {
	newNums := make([]T, 0, len(nums))
	for _, num := range nums {
		flag := true
		for _, newNum := range newNums {
			if num == newNum {
				flag = false
			}
		}
		if flag {
			newNums = append(newNums, num)
		}
	}
	return newNums
}

func UniqueBySet[T comparable](nums []T) []T {
	newNums := make([]T, 0, len(nums))
	s := setx.NewMapSet[T](len(nums))
	for _, num := range nums {
		if !s.Exists(num) {
			newNums = append(newNums, num)
		}
		s.Add(num)
	}
	return newNums
}

// Contains 函数用于检查一个或多个元素是否存在于切片中。
func Contains[T comparable](nums []T, num ...T) bool {
	set := setx.NewMapSet[T](0).WithSlice(nums)
	for _, n := range num {
		if !set.Exists(n) {
			return false
		}
	}
	return true
}
