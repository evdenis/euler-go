package main

import "math"

func prime(num int) bool {

   if (num % 2 == 0) {
      return false
   }

   for i := 3; i <= int(math.Sqrt(float64(num))); i += 2 {
      if (num % i == 0) {
         return false
      }
   }

   return true
}


func main() {
   for i, j := 6, 13; i <= 10001; j += 2 {
      if (prime(j)) {
         println(j)
         i++
      }
   }
}
