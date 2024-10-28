package slicex

import (
	"github.com/LXJ0000/go-utils/internal/slice"
)

// Map 函数将一个切片中的每个元素通过一个映射函数转换为另一个类型的元素，并返回一个新的切片。
func Map[Src, Dst any](slice []Src, mapper func(Src) Dst) []Dst {
	result := make([]Dst, 0, len(slice))
	for _, v := range slice {
		result = append(result, mapper(v))
	}
	return result
}

// FilterMap 函数将一个切片中的每个元素通过一个映射函数转换为另一个类型的元素，并过滤掉不满足条件的元素。
func FilterMap[Src, Dst any](slice []Src, mapper func(Src) (Dst, bool)) []Dst {
	result := make([]Dst, 0, len(slice))
	for _, v := range slice {
		if dst, ok := mapper(v); ok {
			result = append(result, dst)
		}
	}
	return result
}

// Unique 函数用于对切片进行去重。
func Unique[T comparable](nums []T) []T {
	newNums := make([]T, 0, len(nums))
	if len(nums) < 1024 { // TODO const
		slice.UniqueByLoop(nums)
	} else {
		slice.UniqueBySet(nums)
	}
	return slice.Shrink(newNums)
}

func Contains[T comparable](nums []T, num ...T) bool {
	return slice.Contains(nums, num...)
}
