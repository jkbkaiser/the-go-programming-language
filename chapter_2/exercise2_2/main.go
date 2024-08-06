package main

import (
	"fmt"
    "os"
    "strconv"

    "exercise2-2/conv"
)

func main() {
   for _, arg := range os.Args[1:] {
      t, err := strconv.ParseFloat(arg, 64)
      if err != nil {
         fmt.Fprintf(os.Stderr, "cf: %v\n", err)
         os.Exit(1)
      }

      c := conv.Celsius(t)
      f := conv.Fahrenheit(t)
      l := conv.Feet(t)
      m := conv.Metre(t)
      p := conv.Pound(t)
      k := conv.Kilogram(t)

      fmt.Println("Temperature:")
      fmt.Printf("If it is temperature of %s degrees Celsius then it's %s degrees Fahrenheit\n", c, conv.CToF(c))
      fmt.Printf("If it is temperature of %s degrees Fahrenheit then it's %s degrees Celsius\n", f, conv.FToC(f))

      fmt.Println("Length:")
      fmt.Printf("Test first %v and second %v \n", m, conv.MToF(m))
      fmt.Printf("If it is length of %v metres then it's %v feets\n", m, conv.MToF(m))
      fmt.Printf("If it is length of %v feets then it's %v metres \n", l, conv.FToM(l))

      fmt.Println("Weight:")
      fmt.Printf("If it is weight of %v kilograms then it's %v pounds\n", k, conv.KToP(k))
      fmt.Printf("If it is weight of %v pounds then it's %v kilograms\n", p, conv.PToK(p))
   }

}
