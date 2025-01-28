package main 

import  "fmt" 

func main () {

     var i int = 0

	 for {

		if 7 == i {
			break
		}
		
		if i <= 10 {
			 fmt.Println("loop", i)

	    }
		i++
		
	}	

}
	

// loop 0
// loop 1
// loop 2
// loop 3
// loop 4
// loop 5
// loop 6
