package main
import(
  "net/http"
  "log"
  "io/ioutil"
  "os"
)

func main(){
  resp, err:=http.Get("http://google.com")
  if err!=nil{
    log.Fatal(err)
  }

  defer resp.Body.Close()
  body,err:= ioutil.ReadAll(resp.Body)
  if err!=nil{
    log.Fatal(err)
  }
  os.Stdout.Write(body)

  //or

  /*_,err=os.Stdout.Write(body)

  if err!=nil{
    log.Fatal(err)
  }*/
}
