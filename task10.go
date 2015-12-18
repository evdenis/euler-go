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
   sum := 2
   for j := 3; j < 2000000; j += 2 {
      if (prime(j)) {
         sum += j
      }
   }

   println(sum)
}
