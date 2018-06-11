package main

import (
	"fmt"
	"strings"
)

type Count int

func (count *Count) Increment()   { *count ++ }
func (count *Count) Decrement()   { *count -- }
func (count *Count) IsZero() bool { return *count == 0 }

func test() {
	var count Count
	i := int(count)
	count.Increment()
	j := int(count)
	count.Decrement()
	fmt.Println(i, j, count, count.IsZero())
}


type Part struct {
	Id int
	Name string
}

func (part *Part) LowerCase() {
	part.Name = strings.ToLower(part.Name)
}

func (part *Part) UpperCase() {
	part.Name = strings.ToUpper(part.Name)
}
func (part *Part) HasPrefix(prefix string) bool {
	return strings.HasPrefix(part.Name, prefix)
}

func (part Part) String() string {
	return fmt.Sprintf("<<%d %q>>", part.Id, part.Name)
}

func test2() {
	part := Part{5, "wrench"}
	part.UpperCase()
	part.Id += 11
	fmt.Println(part, part.HasPrefix("w"))
}


type Item struct {
	id string
	price float64
	quantity int
}

func (item *Item) Cost() float64 {
	return item.price * float64(item.quantity)
}

type SpecialItem struct {
	Item
	catalogId int
}

func test3() {
	special := SpecialItem{Item{"Green", 3, 5}, 207}
	fmt.Println(special.id, special.price, special.quantity, special.catalogId)
	fmt.Println(special.Cost())
}

type LuxuryITem struct {
	Item
	markup float64
}

func (item *LuxuryITem) Cost() float64 {
	return item.Item.Cost() * item.markup
}

func test4() {
	item := LuxuryITem{Item{"aaa", 5, 9}, 3.5}
	fmt.Println(item, item.Cost())
}

type Place struct {
	latitude, longitude float64
	Name string
}

//(place Place)
func  New(latitude, longitude float64, name string) *Place {
	return &Place{saneAngle(0, latitude), saneAngle(0, longitude), name}
}
func saneAngle(i float64, f float64) float64 {
	return f
}

func (place *Place) Latitude() float64 {
	return place.latitude
}
func (place *Place) SetLatitude(latitude float64)  {
	place.latitude = saneAngle(place.latitude, latitude)
}
func (place *Place) Longitude() float64 {
	return place.longitude
}
func (place *Place) setLongitude(longitude float64)  {
	place.longitude = saneAngle(place.longitude, longitude)
}

func (place *Place) String() string {
	return fmt.Sprintf("(%.3f, %.3f) %q", place.latitude, place.longitude, place.Name )
}

func (original *Place) Copy() *Place {
	return &Place{original.latitude, original.longitude, original.Name}
}

func test5()  {
	newYork := New(40.71667, -74, "New York")
	fmt.Println(newYork)
	baltimore := newYork.Copy()
	baltimore.SetLatitude(newYork.Latitude() - 1.4333)
	baltimore.setLongitude(newYork.Longitude() - 2.61667)
	baltimore.Name = "Baltimore"
	fmt.Println(baltimore)

}




func main() {
	test()
	test2()
	test3()
	test4()
	test5()

}
