package slicex

import "fmt"

func ExampleMap() {
	ints := []int{1, 2, 3, 4, 5}
	strings := Map(ints, func(i int) string {
		return fmt.Sprintf("Number: %d", i)
	})

	fmt.Println(strings)
	// Output: [Number: 1 Number: 2 Number: 3 Number: 4 Number: 5]
}

func ExampleFilterMap() {
	ints := []int{1, 2, 3, 4, 5}
	strings := FilterMap(ints, func(i int) (string, bool) {
		return fmt.Sprintf("Number: %d", i), i%2 == 0
	})

	fmt.Println(strings)
	// Output: [Number: 2 Number: 4]
}