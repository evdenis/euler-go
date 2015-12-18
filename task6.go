package main


func sum(num int) (sum int) {
   for i := 1; i <= num; i++ {
      sum += i
   }

   return
}

func sum2(num int) (sum int) {
   for i := 1; i <= num; i++ {
      sum += i * i
   }

   return
}


func main() {
   sum := sum(100)
   sum *= sum

  println(sum - sum2(100))

}

