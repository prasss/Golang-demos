//right angled triangle of perimeter 24, sides less than 10

package main
import "fmt"
import "math"
func main(){
for i:=1;i<=10;i++ {
 for j:=1;j<=10;j++{
  for k:=1;k<=10;k++{
 // fmt.Println("loops done")
   if k<=10 && j<k && i<j {
   //fmt.Println("1")   
 if j+k+i == 24 {
   //fmt.Println("sum")
 //  fmt.Println(i,j,k)
    g:= math.Pow(float64(j),2)
//fmt.Println(j)
//fmt.Println(g)
    h:=math.Pow(float64(i),2)
    l:=math.Pow(float64(k),2)
     if g+h == l {
     // fmt.Println("squares")
      fmt.Print("triplet: ",i,j,k)
}}}}}}}
