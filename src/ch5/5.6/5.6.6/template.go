package main

import (
	"fmt"
	"reflect"
)

func Minimum(first interface{}, rest...interface{}) interface {} {
	minimum := first
	for _, x := range rest {
		switch x:= x.(type) {
		case int:
			if x < minimum.(int) {
				minimum = x
			}
		case float64:
			if x < minimum.(float64) {
				minimum = x
			}
		case string:
			if x < minimum.(string) {
				minimum = x
			}

		}
	}
	return minimum
}

func Index(xs interface{}, x interface{}) int {
	switch slice := xs.(type) {
	case []int:
		for i, y := range slice {
			if y == x.(int) {
				return i
			}
		}
	case []string:
		for i, y := range slice {
			if y == x.(string) {
				return i
			}
		}
	}
	return -1
}

func IndexReflectX(xs interface{}, x interface{}) int {
	if sliceX:= reflect.ValueOf(xs); sliceX.Kind() == reflect.Slice {
		for i:=0;  i <sliceX.Len(); i ++ {
			switch y:=sliceX.Index(i).Interface().(type) {
			case int:
				if x.(int) == y {
					return i
				}
			case string:
				if x.(string) == y {
					return i
				}

			}
		}
	}
	return -1
}

func IndexReflect(xs interface{}, x interface{}) int {
	if slice:=reflect.ValueOf(xs); slice.Kind() == reflect.Slice {
		for i := 0; i < slice.Len(); i ++ {
			if reflect.DeepEqual(x, slice.Index(i)) {
				return i
			}
		}
	}
	return -1
}


type Slicer interface {
	EqualTo(i int, x interface{}) bool
	Len() int
}

type IntSlice []int
func (slice IntSlice) EqualTo(i int, x interface{}) bool {
	return slice[i] == x.(int)
}
func (slice IntSlice) Len() int {
	return len(slice)
}

type StringSlice []string
func (slice StringSlice) EqualTo(i int, x interface{}) bool {
	return slice[i] == x.(string)
}

func (slice StringSlice) Len() int {
	return len(slice)
}

func IntIndexSlicer(ints []int, x int) int {
	return IndexSlicer(IntSlice(ints), x)
}
func IndexSlicer(slice Slicer, x interface{}) int {
	for i:=0; i < slice.Len(); i ++ {
		if slice.EqualTo(i, x) {
			return i
		}
	}
	return -1
}
func testSlice () {
	fmt.Println(IntIndexSlicer([]int{1,2,3,4,5,5}, 5))
}

func main()  {
	fmt.Println(Minimum(1, 2, 3, 5, -1, 5, 6))
	fmt.Println("5 @", Index([]int{2,4,56,8}, 5))
	fmt.Println("56 @", Index([]int{2,4,56,8}, 56))
	fmt.Println("A @", Index([]string{"A", "B", "C"}, "A"))
	fmt.Println("D @", Index([]string{"A", "B", "C"}, "D"))


	fmt.Println("D @", IndexReflectX([]string{"A", "B", "C"}, "D"))
	fmt.Println("56 @", IndexReflectX([]int{2,4,56,8}, 56))



	// ???? TODO
	fmt.Println("[]string{\"A\"} @", IndexReflect([][]string{{"A"}, {"B"}, {"C"}}, []string{"A"}))
	fmt.Println("56 @", IndexReflect([]int{2,4,56,8}, 56))

	testSlice()
}
