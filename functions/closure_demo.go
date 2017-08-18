package main
import "fmt"

func maincall() func() int{
  i:=0
  return func() int{
    i+=1
    return i
  }
}

func main(){
  j:=maincall()
  fmt.Println(j())
    fmt.Println(j())
      fmt.Println(j())

  k:=maincall()
    fmt.Println(k())
      fmt.Println(k())

}
