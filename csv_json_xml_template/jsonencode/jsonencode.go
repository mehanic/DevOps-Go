package main
import "fmt"
import "encoding/json"

type Hero interface{
   describe() string
}

type SuperHero struct{
  Name string   `json:"name"`
  Weapons [5]string `json:"weapons"`
  Energy int64   `json:"energy"`
  IsHero bool    `json:"ishero"`
}

func (s *SuperHero)add_name(name string){
   s.Name=name
}

func (s *SuperHero)add_weapon(weapons ...string){
   for index,wpn:=range weapons{
      s.Weapons[index]=wpn
   }
}

func (s *SuperHero)add_energy(energy int64){
   s.Energy=energy
}

func (s *SuperHero)set_hero(ishero bool){
   s.IsHero=ishero
}

func (s *SuperHero)describe()string{
   myweapons:=""
   for _,weapon:=range s.Weapons{
      myweapons+=weapon
      myweapons+=" "
   }
   return s.Name+"  "+myweapons
}

func myheros(t Hero){
   fmt.Println("Inside the interface")
   fmt.Println(t.describe())
}

func main(){
  superhero1:=&SuperHero{"myhero",[5]string{"","","","",""},10,false}
  superhero2:=&SuperHero{"mynewhero",[5]string{"knife","dagger","sword","stick","chain"},20,false}
  superhero1.add_name("Superman")
  superhero1.add_weapon("axe","saw","knife","dagger","sword")
  superhero1.add_energy(100)
  superhero1.set_hero(true)
  fmt.Println(*superhero1)
  fmt.Println(superhero1.describe())
  myheros(superhero1)
  myheros(superhero2)
  data1,_:=json.Marshal(superhero1)
  data2,_:=json.Marshal(superhero2)
  fmt.Println(string(data2))
  fmt.Println(string(data1))
}