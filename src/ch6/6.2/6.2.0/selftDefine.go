package main

import (
	"fmt"
	"unicode"
	"strings"
)

type Count int
type StringMap map[string]string
type FloatChan chan float64

type RuneForRuneFunc func(rune) rune

func test2() {
	var removePunctuation RuneForRuneFunc
	phrases := []string{"Day", "long"}
	removePunctuation = func(char rune) rune {
		if unicode.Is(unicode.Terminal_Punctuation, char) {
			return -1
		}
		return char
	}

	for _, pharse := range phrases {
		fmt.Println(strings.Map(removePunctuation, pharse))
	}
}

func test() {
	var i Count = 7
	i ++
	fmt.Println(i)
	sm := make(StringMap)
	sm["key1"] = "value1"
	sm["key2"] = "value2"
	fmt.Println(sm)

	fc := make(FloatChan, 1)
	fc <- 2.29558714938
	fmt.Println(<-fc)
}

func main()  {
	test()
	test2()
}
