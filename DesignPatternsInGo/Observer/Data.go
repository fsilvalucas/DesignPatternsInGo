package Observer

import "fmt"

// from gopl.io ch7 The Go Programming Language - Alan A. A. Donovan · Brian W. Kernighan
// tempconv.go

type Temperature float64
type Humidity float64
type Pressure float64

// Celsius type
type Celsius float64

// Fahrenheit type
type Fahrenheit float64

// Kelvin type
type Kelvin float64

// CToF Function that converts Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9.0/5.0 + 32.0) }

// FToC Function that converts Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32.0) * 5.0 / 9.0) }

// CToK Function that converts Celsius to Kelvin
func CToK(c Celsius) Kelvin { return Kelvin(c - 273.5) }

// FToK Function that converts Fahrenheit to Kelvin
func FToK(f Fahrenheit) Kelvin { return Kelvin(FToC(f) - 273.5) }

// KToC Function that converts Kelvin to Celsius
func KToC(k Kelvin) Celsius { return Celsius(k - 273.5) }

// KToF Function that converts Kelvin to Fahrenheit
func KToF(k Kelvin) Fahrenheit { return Fahrenheit(CToF(Celsius(k + 273.5))) }

// String methods to implement Stringer in fmt package
func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }
func (t Temperature)String() string {return fmt.Sprintf("%g°C", t) }
func (h Humidity) String() string   {return fmt.Sprintf("%g%v", h, "%")}
func (p Pressure) String() string   {return fmt.Sprintf("%gPa", p)}