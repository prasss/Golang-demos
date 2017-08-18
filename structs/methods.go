package main
import "fmt"

type rect struct{
  w,h int
}

func (r *rect) area1() int{
  return r.w * r.h
}

func (r rect) area2() int{
  return r.w * r.h
}

func main(){
  r:=rect{w:2,h:3}
  fmt.Println(r.area1())
  fmt.Println(r.area2())

  s:=&r

  fmt.Println(s.area1())
  fmt.Println(s.area2())

}
