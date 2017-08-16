package main
import "fmt"
func main(){
 var a [10] string
 a[0]= "prasanna"
 a[1]="is"
 a[2]="awsm!"
 fmt.Println(a)
 fmt.Println("length:",len(a))
 b := [6] int {1,2,3,4,5,6}
 fmt.Println("numbers:",b)
 var twoD [5][5]int
 for i:=0;i<5;i++ {
  for j:=0;j<5;j++ {
   fmt.Print("2D one by one:")
   twoD[i][j]=i+j
   fmt.Println("->",twoD[i][j])
}}
fmt.Println("2D all:",twoD)
}


