package main

import (
	"fmt"
	"io"
	"strings"
	"unicode"
)

type LowerCaser interface {
	LowerCase()
}

type UpperCaser interface {
	UpperCase()
}

type LowerUpperCaser interface {
	LowerCaser
	UpperCaser
}

type FixCaser interface {
	FixCase()
}

type ChangeCaser interface {
	LowerUpperCaser
	FixCaser
}

type Exchanger interface {
	Exchange()
}

type StringPair struct {
	first, second string
}

func (pair *StringPair) Exchange() {
	pair.first, pair.second = pair.second, pair.first
}
func (part *StringPair) UpperCase () {
	part.first = strings.ToUpper(part.first)
	part.second = strings.ToUpper(part.second)
}

func (pair *StringPair) FixCase() {
	pair.first = fixCase(pair.first)
	pair.second = fixCase(pair.second)
}

type Point [2]int
func (point *Point) Exchange() {
	point[0], point[1] = point[1], point[0]
}

func (pair StringPair) String() string {
	return fmt.Sprintf("%q+%q", pair.first, pair.second)
}

func exchangeThese(exchangers...Exchanger) {
	for _, exchanger := range exchangers {
		exchanger.Exchange()
	}
}

func test() {
	jekyll := StringPair{"Henry", "Jekyll"}
	hyde := StringPair{"Edward", "Hyde"}
	point := Point{5, -3}
	fmt.Println("Before :", jekyll, hyde, point)
	jekyll.Exchange()
	hyde.Exchange()
	point.Exchange()
	fmt.Println("After #1 :", jekyll, hyde, point)
	exchangeThese(&jekyll, &hyde, &point)
	fmt.Println("After #2 :", jekyll, hyde, point)

}

func (pair *StringPair) Read(data[] byte) (n int, err error) {
	if pair.first == "" && pair.second == ""  {
		return 0, io.EOF
	}
	if pair.first != "" {
		n = copy(data, pair.first)
		pair.first = pair.first[n:]
	}
	if n < len(data) && pair.second != "" {
		m := copy(data[n:], pair.second)
		pair.second = pair.second[m:]
		n += m
	}
	return n, nil
}

func test2() {
	const size = 16
	robert := &StringPair{"Robert L.", "Setveson"}
	david := StringPair{"David", "Balfour"}
	for _, reader := range []io.Reader{robert, &david} {
		raw, err := ToBytes(reader, size)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%q\n", raw)
	}
}

func ToBytes(reader io.Reader, size int) ([]byte, error) {
	data := make([]byte, size)
	n, err := reader.Read(data)
	if err != nil {
		return data, err
	}
	return data[:n], nil

}




func fixCase(str string) string {
	var chars []rune
	upper := true
	for _, char := range str {
		if upper {
			char = unicode.ToUpper(char)
		} else {
			char = unicode.ToLower(char)
		}
		chars = append(chars, char)
		upper = unicode.IsSpace(char) || unicode.Is(unicode.Hyphen, char)
	}
	return string(chars)
}



func test3() {
	lobelia := StringPair{"LOBELIA", "SACKVILLE-BGGINS"}
	lobelia.FixCase()
}

func main() {
	test()
	test2()
}
