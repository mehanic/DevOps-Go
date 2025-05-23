package main 

import "fmt"

func main  () {

     fmt.Println("1.yanvar \n2.fevral \n3.mart \n4.aprel \n5.may \n6.iyun \n7.iyul \n8.avgust \n9.sentabr \n10.oktabr \n11.noyabr \n12dekabr")
     var choice i
     fmt.Println("Chooce number: ")
     fmt.Println("%d", &choice)

     switch choice {
 case 1:
    fmt.Println("yanvar")
    break  // break or NOT NEEDED
 case 2:
	fmt.Println("fevral")
	break	
 case 3:
	fmt.Println("mart")
	break
 case 4:
	fmt.Println("aprel")
	break
 case 5:
	fmt.Println("may")
	break
 case 6:
	fmt.Println("iyun")
	break
 case 7:
	fmt.Println("iyul")
	break
 case 8:
	fmt.Println("avgust")
	break
 case 9:
	fmt.Println("sentabr")
	break
 case 10:
	fmt.Println("oktabr")
	break 
 case 11:
	fmt.Println("noyabr")
    break 
 case 12:
	fmt.Println("dekabr")
   break
 default:
	  fmt.Println("this number is NO! DANG ")
       }
}

