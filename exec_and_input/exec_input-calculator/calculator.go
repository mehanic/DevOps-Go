package main
import "fmt"
import "os"
import "bufio"
import "strconv"


func error_handler_int(err error)int64{
  if err!=nil{
     fmt.Println(err)
     return int64(-1)
  }
  return int64(0)
}

func error_handler_float(err error)float64{
   if err!=nil{
      fmt.Println(err)
      return float64(-1)
   }
   return float64(0)
}

func add_two_number()int64{
   scanner:=bufio.NewScanner(os.Stdin)
   fmt.Println("Enter the first number")
   scanner.Scan()
   a,err:=strconv.ParseInt(scanner.Text(),10,64)
   error_handler_int(err)
   fmt.Println("Enter the second number")
   scanner.Scan()
   b,err:=strconv.ParseInt(scanner.Text(),10,64)
   error_handler_int(err)
   return a+b
}

func divide_two_number()float64{
   scanner:=bufio.NewScanner(os.Stdin)
   fmt.Println("Enter the first number")
   scanner.Scan()
   a,err:=strconv.ParseInt(scanner.Text(),10,64)
   error_handler_float(err)
   fmt.Println("Enter the second number")
   scanner.Scan()
   b,err:=strconv.ParseInt(scanner.Text(),10,64)
   error_handler_float(err)
   return float64(a)/float64(b)
}

func multiply_two_number()int64{
   scanner:=bufio.NewScanner(os.Stdin)
   fmt.Println("Enter the first number")
   scanner.Scan()
   a,err:=strconv.ParseInt(scanner.Text(),10,64)
   error_handler_int(err)
   fmt.Println("Enter the second number:")
   scanner.Scan()
   b,err:=strconv.ParseInt(scanner.Text(),10,64)
   error_handler_int(err)
   return a*b
}

func main(){
  var num int64
  var num1 float64
  num=0
  num1=0
  fmt.Println("Calculator app")
  fmt.Println("1:Add operation")
  fmt.Println("2:Substract operation")
  fmt.Println("3:Multiply operation")
  scanner1:=bufio.NewScanner(os.Stdin)
  fmt.Println("Enter your choice:")
  scanner1.Scan()
  op,_:=strconv.ParseInt(scanner1.Text(),10,64)
  switch op{
     case 1:
       num=add_two_number()
       fmt.Println(num)
     case 2:
       num1=divide_two_number()
       fmt.Println(num1)
     case 3:
       num=multiply_two_number()
       fmt.Println(num)
  }
}