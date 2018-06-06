package main

import (
	"math"
	"fmt"
)

func Heron(a, b, c int) float64 {
	fa, fb, fc := float64(a), float64(b), float64(c)
	s := (fa + fb + fc)/2
	return math.Sqrt(s * (s-fa) * (s-fb) * (s-fc))
}

func PythagoreanTriple(m, n int) (a , b , c int) {
	if m < n {
		m, n = n, m
	}
	return (m*m-n*n), 2*m*n, m*m + n*n
}

func minimumInt(first int, rest ... int) int {
	for _, x := range rest {
		if x < first {
			first = x
		}
	}
	return first
}
func main()  {
	fmt.Println(minimumInt(1, 3, 4, 5, 6, -1, 8))

}
