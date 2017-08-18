package main
import "fmt"

func main(){
  i:=3
  j:=4
  p:=&i  //pointer
  fmt.Println(*p)
  *p=20
  fmt.Println(i) //i value is affect by change in *p value
  p=&j
  *p= *p/2
    fmt.Println(j)
    
}
