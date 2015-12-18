package main

import "fmt"

func main() {
O:   for a := 1; a < 999; a++ {
      for b := 1; b < 999; b++ {
         for c := 1; c < 999; c++ {
            if (a + b + c == 1000) {
               if (a*a + b*b - c*c == 0) {
                  fmt.Printf("%v %v %v: %v\n", a, b, c, a * b * c);
                  break O;
               }
            }
         }
      }
   }
}

