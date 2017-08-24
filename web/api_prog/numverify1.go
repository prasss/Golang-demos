package main
import(
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "net/url"
)

type Numverify struct {
	Valid               bool   `json:"valid"`
	Number              string `json:"number"`
	LocalFormat         string `json:"local_format"`
	InternationalFormat string `json:"international_format"`
	CountryPrefix       string `json:"country_prefix"`
	CountryCode         string `json:"country_code"`
	CountryName         string `json:"country_name"`
	Location            string `json:"location"`
	Carrier             string `json:"carrier"`
	LineType            string `json:"line_type"`
}

func main(){
  no:="+919623856658"
  phone:= url.QueryEscape(no)
  url:= fmt.Sprintf("http://apilayer.net/api/validate?access_key=b33dfc6669c9902ea26080b7d05bb797&number=%s",phone)
  req, err:=http.NewRequest("GET",url,nil)
  if err!=nil{
    log.Fatal("NewRequest:",err)
    return
  }

client:=&http.Client{}
resp, err:=client.Do(req)
if err!=nil{
  log.Fatal("Do:",err)
  return
}

defer resp.Body.Close()
var record Numverify
if err:=json.NewDecoder(resp.Body).Decode(&record); err!=nil{
  log.Println(err)
}
fmt.Println("Phone no.=",record.LocalFormat)



}
