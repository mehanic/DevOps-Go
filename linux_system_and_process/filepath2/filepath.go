package main
import "os"
import "path/filepath"
import "fmt"


func main(){
   count_file:=0
   count_dir:=0
   wd,err:=os.Getwd()
   if err!=nil{
     fmt.Println(err)
   }
   fmt.Println(wd)
   filepath.Walk(wd,func(path string,info os.FileInfo,err error)error{
      if info.IsDir(){
        count_dir=count_dir+1
        wd,err:=filepath.Abs(path)
        if err!=nil{
           return err
        }
        fmt.Println(wd)
      }else{
        count_file=count_file+1
        wd,err:=filepath.Abs(path)
        if err!=nil{
           return err
        }
        fmt.Println(wd)
      }
      return nil
   })
   fmt.Println("The no of files is:")
   fmt.Println(count_file)
   fmt.Println("The no of directory is:")
   fmt.Println(count_dir)
}