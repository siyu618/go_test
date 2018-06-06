package main

import "fmt"

func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i ++ {
		if predicate((i)) {
			return i
		}
	}
	return -1

}

func Filter(limit int, predicate func(i int) bool, appender func(i int)) {
	for i := 0; i < limit; i ++ {
		if predicate(i) {
			appender(i)
		}
	}
}

type memorizeFunction func(int, ...int) interface{}

var Fibonacci memorizeFunction

func Memoize(function memorizeFunction) memorizeFunction {
	cache := make(map[string]interface{})
	return func(x int, xs ...int) interface{} {
		key := fmt.Sprint(x)
		for _, i := range xs {
			key += fmt.Sprintf(",%d", i)
		}
		if value, found := cache[key]; found {
			return value
		}
		value := function(x, xs...)
		cache[key] = value
		return value
	}

}
func init() {
	Fibonacci = Memoize(func(x int, xs ...int) interface{} {
		if x < 2 {
			return x
		}
		return Fibonacci(x - 1).(int) + Fibonacci(x - 2).(int)
	})
}

func main() {
	xs := []int{2, 4, 6, 7}
	ys := []string{"C", "B", "K", "A"}

	fmt.Println(
		SliceIndex(len(xs), func(i int) bool {
			return xs[i] == 5
		}),
		SliceIndex(len(xs), func(i int) bool {
			return xs[i] == 6
		}),
		SliceIndex(len(ys), func(i int) bool {
			return ys[i] == "Z"
		}),
		SliceIndex(len(ys), func(i int) bool {
			return ys[i] == "A"
		}),
	)

	readings := []int{4, -3, 2, -7, 8, 19, -11, 7, 18, 6}
	even := make([]int, 0, len(readings))
	Filter(len(readings),
		func(i int) bool {
			return readings[i] % 2 == 0
		},
		func(i int) {
			even = append(even, readings[i])
		})
	fmt.Println(even)

	fmt.Println(Fibonacci(45).(int))

}
