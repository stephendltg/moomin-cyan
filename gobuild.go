package main

import (
  "encoding/json"
  "io/ioutil"
	"fmt"
	"runtime"
  "strconv"
)

// package json data
type Data struct {
    Name   string `json:"name"`
    Description   string `json:"description"`
    Version    string    `json:"version"`
}

func main() {
  
  // Read package.json
  jsonPackage, err := os.Open("package.json")
  if err != nil {
     panic(err)
  }
  defer jsonPackage.Close()
  byteValue, _ := ioutil.ReadAll(jsonPackage)
  var data Data
  json.Unmarshal(byteValue, &data)
                                 
  fmt.Println("Data name: " + data.name)

	if runtime.GOOS == "darwin" {
		fmt.Println("mac")
	} else if runtime.GOOS == "window" {
		fmt.Println("window")
	} else if runtime.GOOS == "linux" {
		fmt.Println("linux")
	}
}
