package main
import "fmt"

func add2(a int,b int) int{
  return a+b
}
func add4(a,b,c,d int) int{
  return a+b+c+d
}

func main(){

two:=add2(1,2)
  fmt.Println("add2:",two)
four:=add4(1,2,3,4)
  fmt.Println("add4:",four)
}
