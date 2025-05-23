package main
import "fmt"
import "bufio"
import "strconv"
import "os"



func send_digit(number int64,dg chan int64){
   var digit int64
   for number!=0{
      digit=number%10
      dg<-digit
      number=number/int64(10)
   }
   close(dg)
}


func find_square_sum(number int64,square_sum chan int64){
   var sum int64
   sum=0
   ch:=make(chan int64)
   go send_digit(number,ch)
   for num:=range ch{
      sum=sum+(num*num)
   }
   square_sum<-sum
   close(square_sum)
} 



func find_cube_sum(number int64,cube_sum chan int64){
   var sum int64
   sum=0
   ch:=make(chan int64)
   go send_digit(number,ch)
   for num:=range ch{
      sum=sum+(num*num*num)
   }
   cube_sum<-sum
   close(cube_sum)
} 



func main(){
   ch1:=make(chan int64)
   ch2:=make(chan int64)
   scanner:=bufio.NewScanner(os.Stdin)
   fmt.Println("Enter the number:")
   scanner.Scan()
   number,err:=strconv.ParseInt(scanner.Text(),10,64)
   if err!=nil{
     fmt.Println(err)
     return
   }
   go find_cube_sum(number,ch1)
   go find_square_sum(number,ch2)
   cube_sum:=<-ch1
   square_sum:=<-ch2
   fmt.Println(cube_sum)
   fmt.Println(square_sum)
}