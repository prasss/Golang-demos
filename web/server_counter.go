package main
import(
  "fmt"
  "log"
  "net/http"
  //"strconv"
  "sync"
)

var counter int
var mutex=&sync.Mutex{}

func echo(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "hello")
}

func increment(w http.ResponseWriter, r *http.Request){
  mutex.Lock()
  counter++
  fmt.Fprintf(w, "counter incremented %d",counter)
  mutex.Unlock()
}
func main(){
  http.HandleFunc("/",echo)
  http.HandleFunc("/c",increment)

log.Fatal(http.ListenAndServe(":8081", nil))
}
