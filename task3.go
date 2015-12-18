package main


func main() {
   num := int(600851475143)
   max := 1

   if (num % 2 == 0) {
      for ; num % 2 == 0; num /= 2 {}
      max = 2
   }

   for i := 3; num > 1; i+=2 {
      if (num % i == 0) {
         for ; num % i == 0; num /= i {}
         max = i
      }
   }

   println(max)
}

