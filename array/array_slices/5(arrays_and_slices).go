package main

import "fmt"

func main(){
    todoList := [4]string{
        "купить молоко",
        "купить хлеб",
        "купить пиво",
        "погулять",
    }

    tasks:=todoList[1:4]
    //fmt.Println(tasks==nil)

    for i:= range tasks{
        fmt.Println(tasks[i])
    }
    fmt.Println("\n----После выполнения changeTask()----\n")
    changeTask(tasks)

    for i:= range tasks{
            fmt.Println(tasks[i])
    }
    fmt.Println()
    for i:= range todoList{
                fmt.Println(todoList[i])
        }

    fmt.Println("Number of existed elements: \n",len(todoList))
    fmt.Println("Total number of elements: \n",cap(todoList))

    for index,item := range todoList{
        fmt.Printf("%d. %s\n",index+1,item)
    }
    todoList=append(todoList,"почесать мне спинку")
    fmt.Println("Number of existed elements: \n",len(todoList))
    fmt.Println("Total number of elements: \n",cap(todoList))
    var arr [3] int
    fillArray(arr)
    fmt.Println(arr)

}
func changeTask(tasks []string){
    tasks[0]="прибраться в комнате"
    tasks[1]="погладить кота"
}
func fillArray(arr [3] int){
    for i:=0;i<len(arr);i++{
        arr[i]=i
    }
    fmt.Println("fillArray():",arr)
}