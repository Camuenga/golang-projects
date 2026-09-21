// main project main.go
package main

import (
	"fmt"
)

type celsius float64
type kelvin float64

func celsiusToKelvin(c celsius) kelvin {
	return kelvin(c + 273.15)
}

func main() {

	var temparature celsius = 127
	kTemparature := celsiusToKelvin(temparature)
	fmt.Println(kTemparature)

}
