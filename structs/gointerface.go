package main
import "fmt"

type Person interface{
   describe()string
}

type Police struct{
   Name string
   SocialId int64
   Rank string
}

type Teacher struct{
   Name string
   SocialId int64
   Rank string
}

func (p *Police)describe()string{
   return p.Name+" "+p.Rank
}

func (p *Police)report(){
    fmt.Println("Myself ",p.Name," Reporting")
    fmt.Println("My Rank ",p.Rank)
}

func (t *Teacher)describe()string{
  return t.Name+" "+t.Rank
}

func (t *Teacher)teach(){
  fmt.Println("I am ",t.Name)
  fmt.Println("I am ",t.Rank)
}

func myperson(p Person){
    switch p.(type){
         case *Police:
             fmt.Println(p.describe())
             p.(*Police).report()
         case *Teacher:
             fmt.Println(p.describe())
             p.(*Teacher).teach()
    }
}



func main(){
     pol:=&Police{"Bapan",121,"Inspector"}
     tch:=&Teacher{"Swami",123,"Professor"}
     fmt.Println(pol.describe())
     pol.report()
     fmt.Println(tch.describe())
     tch.teach()
     myperson(pol)
     myperson(tch)  
}