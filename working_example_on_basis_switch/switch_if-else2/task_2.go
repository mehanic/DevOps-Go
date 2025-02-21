package main
import "fmt"

func main(){

	
	var (
		s= 3214568
		i int=0
		sum float32= 0
	)
	
	for s>0{
		sum+=float32(s%10)
		s=(s-s%10)/10
		i++

	}
	fmt.Println(sum/float32(i))
}