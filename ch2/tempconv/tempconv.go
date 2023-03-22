// Package tempconv performs Celsius and Fahrenheit conversions
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64
type Meters float64
type Feet float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%gC", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%gF", f) }
func (m Meters) String() string     { return fmt.Sprintf("%gm", m) }
func (f Feet) String() string       { return fmt.Sprintf("%gft", f) }
