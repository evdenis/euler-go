package main

func palindrome(num int) bool {
   rev := 0
   for i := num; i > 0; i /= 10 {
      rev = rev * 10 + (i % 10)
   }

   return num == rev
}

func factorize3x3(num int) bool {
   for i := 100; i <= 999; i++ {
      for j := i; j <= 999; j++ {
         if (i * j == num) {
            return true
         }
      }
   }

   return false
}


func main() {
   for i := 999 * 999; i >= 100 * 100; i-- {
      if (palindrome(i)) {
         if (factorize3x3(i)) {
            println(i)
            break
         }
      }
   }
}

