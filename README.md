# configuration-go


The configuration-go package intended for normal work with configuration files.
For now we support HJSON format three-like configs.
Library is intended to ease getting configuration file values if you know path to them

# Install

Make sure you have a working Go environment. See the [install instructions](http://golang.org/doc/install.html).

$ go get -u github.com/ilya1st/configuration-go

# Usage as command line tool


Sample:

For example: we have file test.hjson


# Usage as a GO library

```go

package main
import(
  "fmt"
  "github.com/ilya1st/configuration-go"
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

```

See examples directory.