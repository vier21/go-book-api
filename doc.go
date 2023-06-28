package main

import "fmt"

type myInt int

type Comparable interface {
	LessThan(c Comparable) bool
}

func (this myInt) LessThan(other Comparable) (r bool) {
	return this < other.(myInt)
}

func test(c []Comparable) {
	if c[0].LessThan(c[1]) {
		fmt.Print("lessThan\n")
	} else {
		fmt.Print("greaterThan\n")
	}
}

func main() {
	slc := []Comparable{myInt(4), myInt(2)}
	test(slc)
}
