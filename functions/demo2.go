package main
import "fmt"

func vals()(int,int,int){
  return 1, 2, 3
}

func main(){
  a,b,c:=vals()
  fmt.Println(a,b,c)
  _,_,d:=vals()
  fmt.Println(d)



  func vals(a,b,c int)(int,int,int){ //LOL
    return a,b,c
  }

  func main(){
    a,b,c:=vals(1,2,3)
    fmt.Println(a,b,c)
}
