// Package tempconv performs Celsius and Fahrenheit temperature computations.
package tempconv

import (
	"fmt"
)

type Fahrenheit float64

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%.2f°F", f)
}

type Celcius float64

func (c Celcius) String() string {
	return fmt.Sprintf("%.2f°C", c)
}

const (
	AbsoluteZeroC Celcius = -273.15
	FreezingC     Celcius = 0
	BoilingC      Celcius = 100
)

func CToF(c Celcius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celcius {
	return Celcius((f - 32) * 5 / 9)
}

