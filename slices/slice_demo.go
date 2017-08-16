package main
import "fmt"
func main(){
s := make([]string, 3)
fmt.Println("empty:",s)
s[0]="1"
s[1]="2"
s[2]="3"
fmt.Println("s:",s)
fmt.Println("length:",len(s))
s= append(s,"a")
s=append(s,"b")
fmt.Println("s:",s)
c:=make([]string,len(s))
copy(c,s)
fmt.Println("c:",c)
l:=s[2:3]
fmt.Println("l:",l)
m:=s[:3]
n:=s[2:]
fmt.Println("m:",m)
fmt.Println("n:",n)

twoD:=make([][]int,5)
for i:=0;i<5;i++ {
  jlen:=i+1
  twoD[i]=make([]int,jlen)
  for j:=0;j<jlen;j++ {
   twoD[i][j]=j
   fmt.Println("yolo:",twoD)
}
}
fmt.Println("2d:",twoD)
}
