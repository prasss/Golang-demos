package main
import "fmt"
import "time"

func main(){

for i:=1;i<=5;i++ {
switch i{
case 1:
 fmt.Println("one")

case 2:
 fmt.Println("two")

case 3:
 fmt.Println("three")

case 4:
 fmt.Println("four")


case 5:
 fmt.Println("five")
}
}

switch time.Now().Weekday(){
case time.Saturday, time.Sunday:
 fmt.Println("Weekend!!!")
default:
 fmt.Println("awww")
 }
t:=time.Now().Hour()
m:=time.Now().Minute()
s:=time.Now().Second()
 fmt.Println("time:",t,":",m,":",s)
switch {
case t>12:
  fmt.Println("aaa")
default:
 fmt.Println("bbbb")
}}
