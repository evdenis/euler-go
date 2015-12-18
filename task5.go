package main


func divisible(num int) bool {
   for i := 2; i <= 20; i++ {
      if (num %i != 0) {
         return false
      }
   }

   return true
}


func main() {
   for i := 1; ; i++ {
      if (divisible(i)) {
         println(i)
         break
      }
   }
}

