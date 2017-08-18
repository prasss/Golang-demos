package main
import "fmt"

type person struct{
  name string
  age int
  prof string
}

func main(){
  fmt.Println(person{"bob",20,"doctor"})
  fmt.Println(person{name:"pika"})
  fmt.Println(person{name:"prasss",age:21,prof:"engineer"})
  s:=person{"yolo",123,"aaaa"}
  fmt.Println(s.name)
  sp:=&s
  fmt.Println(sp.age)
  sp.age=432
  fmt.Println(sp.age)

}
