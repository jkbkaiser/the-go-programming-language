package conv

import (
   "fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64
type Metre float64
type Feet float64
type Kilogram float64
type Pound float64

const (
   AbsoluteZeroC Celsius = -273.15
   FreezingC     Celsius = 0
   BoilingC      Celsius = 100
   AbsoluteZeroK Kelvin  = 0
   FreezingK     Kelvin  = 273.15
   BoilingK      Kelvin  = 373.15
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }
func (m Metre) String() string      { return fmt.Sprintf("%g m", m) }
func (l Feet) String() string       { return fmt.Sprintf("%g ft", l) }
func (k Kilogram) String() string   { return fmt.Sprintf("%g kg", k) }
func (p Pound) String() string      { return fmt.Sprintf("%g lb", p) }

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) Celsius { return Celsius(k - 279.15) }

// CToK converts a Celsius temperature to Kelvin.
func CToK(c Celsius) Kelvin { return Kelvin(c + 279.15) }

// FToK converts a Fahrenheit temperature to Kelvin.
func FToK(f Fahrenheit) Kelvin { return Kelvin(((f - 32) * 5 / 9) + 279.15) }

// KToF converts a Kelvin temperature to Fahrenheit.
func KToF(k Kelvin) Fahrenheit { return Fahrenheit((k-279.15)*9/5 + 32) }

// FToM converts a Feets length to Meters.
func FToM(f Feet) Metre { return Metre(f * 0.3048) }

// MToF converts a Metre length to Feets.
func MToF(m Metre) Feet { return Feet(m / 0.3048) }

// PToK converts a Pound weight to Kilograms.
func PToK(p Pound) Kilogram { return Kilogram(p * 0.45359237) }

// KToP converts a Kilogram weight to Pounds.
func KToP(k Kilogram) Pound { return Pound(k / 0.45359237) }
