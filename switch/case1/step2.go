package main 

import "fmt"

func main () {

	fmt.Println ("1.Dushanba \n2.seshanba\n3.chorshanba\n4.payshanba\n5.juma\n6.shanba\n7.yakshanba")
	var choice int
	fmt.Println("Chooce number: ")
	fmt.Scanf("%d", &choice)

	switch choice {
	case 1:
		fmt.Println("Dushanba")
		break  // --> break or NOT NEEDED
	case 2:
		fmt.Println("seshanba")
		break
	case 3:
		fmt.Println("chorshanba")
		break
	case 4:		
		fmt.Println("payshanba")
		break
	case 5:
		fmt.Println("juma")
		break
	case 6:
		fmt.Println("shanba")
		break
	case 7:
		fmt.Println("yakshanba")
		break	
	default:
		fmt.Println("This number is No! DANG ")
	}
}
