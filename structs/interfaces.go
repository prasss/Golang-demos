package main
import "fmt"
import "math"

type geometry interface{
  area() float64
  peri() float64
}

type rect struct{
  l,h float64
}

type circle struct{
  rad float64
}

func (r rect) area() float64{
  return (r.l)*(r.h)
}

func (r rect) peri() float64{
  return 2*(r.l+r.h)
}

func (c circle) peri() float64{
  return 2*math.Pi*c.rad
}

func (c circle) area() float64{
  return math.Pi*c.rad*c.rad
}

func measure(g geometry){
    fmt.Println(g)
  fmt.Println(g.area())
    fmt.Println(g.peri())
}

func main(){
rr:=rect{l:5,h:6}
cc:=circle{rad:5}

measure(rr)
  //fmt.Println(ar)
measure(cc)
}
