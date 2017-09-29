package main
import(
  "fmt"
  "github.com/Ilya1st/configuration-go"
)

func main(){
  config, err:=configuration.GetConfigInstance("mainconfig", "HJSON", "test.hjson")
  if err != nil {
    fmt.Printf("Error occured %v\n", err);
  }
  val, err:=config.GetValue("section1", "subsection2", "value")
  fmt.Printf("getting section1/subsection2/value: val=%v, error=%v\n", val, err)
  val, err=config.GetValue("test")
  fmt.Printf("getting test: val=%v, error=%v\n", val, err)
}
