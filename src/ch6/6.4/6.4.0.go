package main

import "fmt"


func test1() {
	points := [][2]int{{6,4}, {}, {3,5}}
	for _, point := range points {
		fmt.Println(point[0], point[1])
	}

	points2 := []struct {x, y int} {{4,6}, {4,5}, {}, {0,9}}
	for _, point := range points2 {
		fmt.Println(point.x, point.y)
	}
}

type Person struct {
	Title string
	Forenames []string
	Surname string
}

type Author1 struct {
	Names Person
	Title []string
	YearBorn int
}

func test2() {
	author1 := Author1{
		Person{"Mr", []string {"Robert", "Louis", "Balfour"}, "Steveson"}, []string{"Kidnapped", "Treasure Island"}, 1850}
	fmt.Println(author1)
	author1.Names.Title = ""
	author1.Names.Forenames = []string{"abc", "def"}
	author1.Names.Surname = "Widle"
	author1.Title = []string{"The picture of xxx"}
	author1.YearBorn += 4
	fmt.Println(author1)
}


type Author2 struct {
	Person
	Title []string
	YearBorn int
}

func test3() {
	author2 := Author2{
		Person{"Mr", []string {"Robert", "Louis", "Balfour"}, "Steveson"}, []string{"Kidnapped", "Treasure Island"}, 1850}
	fmt.Println(author2)
	author2.Title = []string{"The pic"}
	author2.Person.Title = ""
	author2.Forenames = []string {"Oscar"}
	author2.Surname = "Wilde"
	author2.YearBorn += 4
	fmt.Println(author2)
}


type Count int

func (count *Count) Increment()   { *count ++ }
func (count *Count) Decrement()   { *count -- }
func (count *Count) IsZero() bool { return *count == 0 }

type Tasks struct {
	slice []string
	Count
}

func (tasks *Tasks) Add(task string) {
	tasks.slice = append(tasks.slice, task)
	tasks.Increment()
}

func (tasks *Tasks) Tally() int {
	return int(tasks.Count)
}

func test4() {
	tasks := Tasks{}
	fmt.Println(tasks.IsZero(), tasks.Tally(), tasks)
	tasks.Add("one")
	tasks.Add("two")
	fmt.Println(tasks.IsZero(), tasks.Tally(), tasks)

}

type Optioner interface {
	Name() string
	IsValid() bool
}

type OptionCommon struct {
	ShortName string "short option name"
	LongName string "long option name"
}

type IntOption struct {
	OptionCommon
	Value, Min, Max int
}

func (option IntOption) Name() string {
	return name(option.ShortName, option.LongName)
}

func (option IntOption) IsValid() bool {
	return option.Min <= option.Value && option.Value <= option.Max
}
func name(shortName string, longName string) string {
	if longName == "" {
		return shortName
	}
	return longName
}

type FloatOption struct {
	Optioner
	Value float64
}

type GenericOption struct {
	OptionCommon
}

func (option GenericOption) Name() string  {
	return name(option.ShortName, option.LongName)
}

func (option GenericOption) IsValid() bool  {
	return true
}

func test5() {
}

func main()  {
	test1()
	test2()
	test3()
	test4()

}
