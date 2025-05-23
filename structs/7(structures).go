package main

import "fmt"

type employee struct{
        name string
        sex string
        age int
        salary int
}

//конструктор
func newEmployee(name,sex string, age,salary int)employee{
    return employee{
        name:name,
        sex:sex,
        age:age,
        salary:salary,
    }
}

func (e employee) getInfo() string{
    return fmt.Sprintf("Сотрудник: %s\nВозраст: %d\nЗарплата: %d\n",e.name, e.age, e.salary)
}

func (e *employee) setName(name string){
    e.name=name
}

func main(){
    emp1:=newEmployee("Anek","F",20,100000000)
    emp2:=newEmployee("Tem","M",22,320000)

    emp1.setName("Kira")

    fmt.Printf("%+v\n",emp1.getInfo())
    fmt.Printf("%+v\n",emp2.getInfo())
}

