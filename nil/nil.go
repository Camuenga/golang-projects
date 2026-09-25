package main

import (
	"fmt"
	"sort"
)

// nil with structs and pointer function---------------------

type Person struct {
	age int
}

func (arg *Person) inc() {
	if arg == nil {
		return
	}
	arg.age++
}

// nil with structs and pointer function----------------------

// nil with variable function -----------------------
func sortStrings(s []string, less func(i, j int) bool) {
	if less == nil {
		less = func(i, j int) bool { return s[i] > s[j] }
	}
	sort.Slice(s, less)
}

// nil with variable function ------------------------

// alternative to nil ---------------------------------
type number struct {
	value int
	valid bool
}

func newNumber(arg int) number {
	return number{value: arg, valid: true}
}

func (arg number) String() string {
	if !arg.valid {
		return "not set"
	}
	return fmt.Sprintf("%d", arg.value)
}

// alternative to nil--------------------------------

func main() {

	// nil with integers and pointers
	var nowhere *int
	fmt.Println(nowhere)

	if nowhere != nil {
		fmt.Println(*nowhere)
	}
	// nil with integers and pointers

	// nil with struct and pointers --------------------
	var p *Person
	fmt.Println(p)
	p.inc()
	// nil with struct and pointers --------------------
	// nil with variables functions --------------------
	var sum func(arg1, arg2 int) int
	sum = func(arg1, arg2 int) int {
		return arg1 + arg2
	}
	fmt.Println(sum == nil)
	fmt.Println(sum(1, 2))
	// nil with variables functions --------------------

	// nil with slices ---------------------------------
	food := []string{"onion", "avocato", "banana"}

	sortStrings(food, nil)
	fmt.Println(food)

	var soupe []string
	fmt.Println(soupe == nil)

	for _, ingredients := range soupe {
		fmt.Println(ingredients)
	}

	fmt.Println(len(soupe))
	soupe = append(soupe, "onion", "tomato", "celery")
	fmt.Println(soupe)
	// nil with slices ---------------------------------

	// nil with Map -------------------------------------
	var soup map[string]int
	fmt.Println(soup == nil)
	measurement, ok := soup["onion"]

	//soup["onion"] = 1 // writing to a nil map will panic

	if ok {
		fmt.Println(measurement)
	}

	for ingredient, measurement := range soup {
		fmt.Println(ingredient, measurement)

	}

	// nil with Map--------------------------------------

	// nil with interface---------------------------------
	var v interface{}

	fmt.Printf("%T %v\n", v, v == nil)

	var n *int
	v = n
	fmt.Printf("%#v", v)
	var s fmt.Stringer
	fmt.Printf("%#v", s)
	// nil with interface---------------------------------

	// Alternative to nil--------------------------------
	s1 := newNumber(100075)
	fmt.Println(s1)
	s2 := newNumber(751)
	result := s2.String()
	fmt.Println(result)
	// Alternative to nil--------------------------------

}
