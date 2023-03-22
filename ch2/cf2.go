/*Exercise 2.2: Write a general-purpose unit-conversion program analogous to cf that reads numbers from its
command-line arguments or from the standard input if there are no arguments, and converts each number into units
like temperature in Celsius and Fahrenheit,length in feet and meters, weight in pounds and kilograms, and the like.*/

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"

	"gopl.io/ch2/tempconv"
)

func main() {
	input, _ := strconv.Atoi(os.Args[1])

	celsius := tempconv.Celsius(input)
	fahrenheit := tempconv.Fahrenheit(input)

	fmt.Printf("number -> celsius: %v\n", celsius)
	fmt.Printf("number -> farenheit: %v\n", fahrenheit)
	fmt.Printf("%v celsius -> fahrenheit == %v\n", celsius, tempconv.CToF(celsius))
	fmt.Printf("%v farenheit -> celsius: %v\n", fahrenheit, math.Round(float64(tempconv.FToC(fahrenheit))*100/100))

	feet := tempconv.Feet(input)
	meters := tempconv.Meters(input)

	fmt.Printf("%s -> meters = %v\n", feet, tempconv.FtoM(feet))
	fmt.Printf("%s -> feet = %v\n", meters, tempconv.MtoF(meters))

}
