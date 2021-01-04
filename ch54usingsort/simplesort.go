package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{10, 5, 3, 7, 6}
	b := []float64{4.2, 7.6, 5.5, 1.3, 9.9}
	c := []string{"Maria", "Andrew", "John"}

	fmt.Println("before sort : ", a, ", ", b, ", ", c)

	sort.Sort(sort.IntSlice(a))
	sort.Sort(sort.Float64Slice(b))
	sort.Sort(sort.StringSlice(c))
	fmt.Println("after sort : ", a, ", ", b, ", ", c)

	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	sort.Sort(sort.Reverse(sort.Float64Slice(b)))
	sort.Sort(sort.Reverse(sort.StringSlice(c)))
	fmt.Println("after reverse : ", a, ", ", b, ", ", c)
}