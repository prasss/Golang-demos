package main
import "fmt"
func main(){
  m:=make(map[string]int)
  m["k1"]=1
  m["k2"]=2
  m["k3"]=3
  fmt.Println(m)
  v1:= m["k1"]
  fmt.Println("v1:",v1)
  delete(m,"k2")
  fmt.Println(m)

  _,prs:=m["k3"]   //value ignored
  fmt.Println(prs) //presence

  v,prs:=m["k1"]
  fmt.Println(prs)
  fmt.Println("k1:",v) //value


  n:=map[string]int{"p":1,"i":2,"k":3,"a":4}
  fmt.Println(n)
}
