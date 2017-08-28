package main
import(
  "fmt"
  "os"
  "encoding/json"
  "io/ioutil"
)

type pokeobj struct{
  pokemon_type []string
}

type pokemon_type struct{
  Name string `json:"name"`
  Attack string `json:"attack"`
}

func main(){
  file, err:=ioutil.ReadFile("./data1.json")
  if err!=nil {
    fmt.Println("error:",err)
    os.Exit(1)
  }
  fmt.Printf("%s\n",string(file))
 pika:=pokeobj{}
 json.Unmarshal(file, &pika)
 fmt.Printf("%v\n",pika)
  fmt.Printf("%v\n",pika.pokemon_type)
}
