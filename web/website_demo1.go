package main
import(
//"fmt"
"net/http"
"log"
//"fmt"
)

/*func run(w http.ResponseWriter,r *http.Request){
  //http.ServeFile(w,r,r.URL.Path[1:])
  http.FileServer(http.Dir("static"))
}*/ //not reqd for websites

/*type web int

func (m web) ServeHTTP(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w,"yoooo")
}
func main(){
  //http.Handle("/",http.FileServer(http.Dir("./static")))
  //log.Fatal(http.ListenAndServe(":8081",nil))
var d web
http.ListenAndServe(":8080", d)

}
*/
func main(){
  http.Handle("/",http.FileServer(http.Dir("./static")))
  log.Fatal(http.ListenAndServe(":8081",nil))
}
