package main

import (
	"fmt"
    "errors"
)

type employee struct{
    id int
    name string
    age string
    salary int
}

type storage interface{
    insert (e employee) error
    get (id int) (employee,error)
    delete (id int) error
}

type memoryStorage struct {
    data map[int]employee
}

func newMemoryStorage() *memoryStorage{
    return &memoryStorage{
        data: make(map[int]employee),
    }
}

func (s *memoryStorage) insert(e employee) error{
    s.data[e.id]=e
    return nil
}

func (s *memoryStorage) get(id int) (employee,error){
    e,exists:=s.data[id]
    if !exists{
        return employee{}, errors.New("This employee doesn't exist")
    }

    return e,nil
}

func (s *memoryStorage) delete (id int) error{
    delete(s.data,id)
    return nil
}

type dumbStorage struct {}

func newDumbStorage()*dumbStorage{
    return &dumbStorage{}
}

func (s *dumbStorage) insert(e employee) error{
    fmt.Printf("вставка пользователя с id=%d прошла успешно\n",e.id)
    return nil
}

func (s *dumbStorage) get(id int) (employee,error){
    e:=employee{
        id:id,
    }

    return e,nil
}

func (s *dumbStorage) delete (id int) error{
    fmt.Printf("удаление пользователя с id=%d прошло успешно\n",id)
    return nil
}

func main(){
    var s storage

    fmt.Println("s==nil",s==nil)
    fmt.Printf("type of s:%T\n",s)

    s=newMemoryStorage()

    fmt.Println("s==nil",s==nil)
    fmt.Printf("type of s:%T\n",s)

    s=newDumbStorage()

    fmt.Println("s==nil",s==nil)
    fmt.Printf("type of s:%T\n",s)

    s=nil

    fmt.Println("s==nil",s==nil)
    fmt.Printf("type of s:%T\n",s)
}