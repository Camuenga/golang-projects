// pointer project main.go
package main

import (
	"fmt"
	"math/rand"
)

type operation int

func (aritimetic operation) sum() operation {
	return aritimetic + operation(rand.Int())
}

func main() {

	timy := operation(32)
	x := timy.sum()

	fmt.Printf("%d", x)

}
