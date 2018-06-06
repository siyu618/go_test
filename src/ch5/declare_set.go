package ch5

import (
	"fmt"
	"bytes"
	"encoding/json"
	"strings"
	"math/rand"
	"math"
)

func jsonObjectAsString(jsonObject map[string]interface{}) string {
	var buffer bytes.Buffer
	buffer.WriteString("{")
	comma := ""
	for key, value := range jsonObject {
		buffer.WriteString(comma)
		switch value := value.(type) {
		case nil:
			fmt.Fprintf(&buffer, "%q: null", key)
		case bool:
			fmt.Fprintf(&buffer, "%q: %t", key, value)
		case float64:
			fmt.Fprintf(&buffer, "%q %f", key, value)
		case string:
			fmt.Fprintf(&buffer, "%q: %q", key, value)
		case []interface{}:
			fmt.Fprintf(&buffer, "%q: [", key)
			innerComma := ""
			for _, s := range value {
				if s, ok := s.(string); ok {
					fmt.Fprintf(&buffer, "%s%q", innerComma, s)
					innerComma = ", "
				}
			}
			buffer.WriteString("]")

		}
		comma = ", "
	}
	buffer.WriteString("}")
	return buffer.String()
}

func classfier(items...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("param #%d is a bool\n", i)
		case float64:
			fmt.Printf("param #%d is a float\n", i)
		case int, int8, int16, int32, int64:
			fmt.Printf("param #%d is a int\n", i)
		case uint, uint8, uint16, uint32, uint64:
			fmt.Printf("param #%d is a unsinged int\n", i)
		case nil:
			fmt.Printf("param #%d is a nill\n", i)
		case string:
			fmt.Printf("param #%d is a string\n", i)
		default:
			fmt.Printf("param #%d is a unknown\n", i)
		}
	}
}

type State struct {
	Name     string
	Senators []string
	Water    float64
	Area     int
}

func (state State) String() string {
	var senators []string
	for _, senator := range state.Senators {
		senators = append(senators, fmt.Sprintf("%q", senator))
	}
	return fmt.Sprintf(
		"{\"name\": %q, \"area\": %d, \"water\": %f, \"senators\":[%s]}",
		state.Name, state.Area, state.Water, strings.Join(senators, ", "))
}

func couldFindMatch(table [][]int, target int) bool {
	found := false
	FOUND:
	for _, row := range table {
		for _, num := range row {
			if num == target {
				found = true
				break FOUND
			}
		}
	}
	return found
}

func createCounter(start int) chan int {
	next := make(chan int)
	go func(i int) {
		for {
			next <- i
			i ++
		}
	} (start)
	return next
}
func testGoRoutine() {
	counterA := createCounter(2)
	counterB := createCounter(102)
	for idx := range []int{1,2,3,4,5} {
		a := <- counterA
		fmt.Printf(" %d: (A->%d, B->%d)", idx, a, <-counterB)
	}
	fmt.Println()
}

func testSelect() {
	channelNum := 6
	channels := make([] chan bool, channelNum)
	for i := range channels {
		channels[i] = make(chan bool)
	}
	go func() {
		for {
			channels[rand.Intn(channelNum)] <- true
		}
	} ()

	for i:= 0; i < 36; i ++ {
		var x int
		select {
		case <- channels[0]: x  = 1
		case <- channels[1]: x  = 2
		case <- channels[2]: x  = 3
		case <- channels[3]: x  = 4
		case <- channels[4]: x  = 5
		case <- channels[5]: x  = 6
		}
		fmt.Printf("%d ", x)
	}
	fmt.Println()

}

func expensiveComputation(data int, answer chan int, done chan bool) {
	finished := false
	idx := 0
	for !finished {
		idx ++
		if idx == data {
			finished = true
		}
		answer <- idx
	}
	done <-true
}

func testSelect2() {
	const allDone = 2
	doneCount := 0
	answerA := make(chan int)
	answerB := make(chan int)
	defer func() {
		close(answerA)
		close(answerB)
	}()
	done := make(chan bool)
	defer func() { close(done)}()

	go expensiveComputation(5, answerA, done)
	go expensiveComputation(20, answerB, done)

	for doneCount != allDone {
		var which, result int
		select {
		case result = <- answerA:
			which = 'A'
		case result = <- answerB:
			which = 'B'
		case <-done:
			doneCount ++
		}
		if which != 0 {
			fmt.Printf("%c -> %d,", which, result)
		}
	}
	fmt.Println()
}

func ConvertInt64ToInt(x int64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 {
		return int(x)
	}
	panic(fmt.Sprint("%d is out of int range", x))
}

func IntFromInt64(x int64) (i int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	i = ConvertInt64ToInt(x)
	return i, nil

}



func main() {
	a, b, c := 2, 3, 5
	for a := 7; a < 8; a++ {
		b := 11
		c = 13
		fmt.Printf("inner: a->%d b->%d c->%d\n", a, b, c)
	}
	fmt.Printf("inner: a->%d b->%d c->%d\n", a, b, c)

	type StringSlice []string;

	fancy := StringSlice{"Lithium", "Sodium"}
	fmt.Println(fancy)

	var i interface{} = 99
	var s interface{} = []string{"left", "right"}
	j := i.(int)
	fmt.Printf("%T->%d\n", j, j)

	if i, ok := i.(int); ok {
		fmt.Printf("%T->%d\n", i, j)
	}
	if s, ok := s.([]string); ok {
		fmt.Printf("%T->%q\n", s, s)
	}

	classfier(5, -17.9, "ZIP", nil, true, complex(1, 1))

	jsonStr := "{\"name\":\"Massachusetts\", \"area\":27336, \"water\":25.7, \"senators\":[\"John Kerry\", \"Scott Brown\"]}"
	MA := []byte(jsonStr)
	var object interface{}
	if err := json.Unmarshal(MA, &object); err != nil {
		fmt.Println(err)
	} else {
		jsonObject := object.(map[string]interface{})
		fmt.Println(jsonObjectAsString(jsonObject))
	}

	var state State
	if err := json.Unmarshal(MA, &state); err != nil {
		fmt.Println(err)
	}
	fmt.Println(state)

	table := [][]int{{1,2,3,4}, {6,7,8,9}, {10,23,47}}
	target:=10
	fmt.Printf("find %d in table : %t\n", target, couldFindMatch(table, target))

	testGoRoutine()
	testSelect()
	testSelect2()
}
