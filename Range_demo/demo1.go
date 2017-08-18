package main
import "fmt"
func main(){
  nums:= []int{1,2,3,4,5,6,7}
  sum:=0
  for i, n:= range nums{
    fmt.Println("number",i,"is:",n)
    sum+=n
    if(n==5){
      break
    }
  }
    fmt.Println("sum:",sum)

kvs:=map[string]string{"a":"apple","b":"ball","c":"cat","d":"dog"}
for k,v:= range kvs{
    //fmt.Println(k,"for",v)
      fmt.Printf("%s for %s\n",k,v)
    }
      for i,c:=range "pikaaa"{
          fmt.Printf("%d - %c\n",i,c)
      }
}
