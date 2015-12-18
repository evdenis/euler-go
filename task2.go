package main


func main() {
   sum := int64(0)

   for a, b := 1, 2; b < 4000000; a, b = b, a + b {
      if (b % 2 == 0) {
         sum += int64(b)
      }
   }

   println(sum)
}

