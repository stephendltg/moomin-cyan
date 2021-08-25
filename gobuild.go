package main

import (
  "encoding/json"
  "io/ioutil"
	"fmt"
	"runtime"
  "strconv"
  "path/filepath"
  "exec"
)

// Templates
const templateDeb := "{{.name}} items are made of {{.description}}"

// package json data
type Data struct {
    Name   string `json:"name"`
    Description   string `json:"description"`
    Version    string    `json:"version"`
}

// Absolu path exec
func abspath() string {
	exe, err := os.Executable()
	if err != nil { panic(err) }
  
	exe, err = filepath.EvalSymlinks(exe)
	if err != nil {	panic(err) }

	dir := filepath.Dir(exe)
	return dir
}

// Parse template
func parse(path string, tmpl string, data Data) {
  t, err := template.New("tmpl").Parse(tmpl)
  if err != nil { panic(err) }

	f, err = os.Create(path)
  if err != nil { panic(err) }

	err = t.Execute(f, data)
  if err != nil { panic(err) }
	f.Close()
}

// Spawn process command
func spawn(cmd string, passthroughArgs bool) {
  parts := strings.Split(cmd)
	head := parts[0]
	args := parts[1:len(parts)]
  if passthroughArgs {
     args = append(args, os.Args[1:]...)
  }
  
	cmd := exec.Command(head, args...)
  
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
  
  err := cmd.Run()
  if err != nil { panic(err) }
  
	fmt.Printf("Result: %v / %v", out.String(), stderr.String())
}

// copy file
func copy(source string, destination string) {  
  input, err := ioutil.ReadFile(source)
  if err != nil { panic(err) }

	err = ioutil.WriteFile(destination, input, 0644)
  if err != nil { panic(err) }
}

// Read package
func readPackage() Data {
  jsonPackage, err := os.Open(filepath.Join(abspath(), "package.json"))
  if err != nil {
    panic(err)
  }
  defer jsonPackage.Close()
  byteValue, _ := ioutil.ReadAll(jsonPackage)
  var data Data
  if err = json.Unmarshal(byteValue, &data); err != nil {
	 	panic(err)
	}
  return data
}

func main() {
  
  // Read package.json
  data := readPackage()
                                 
  fmt.Println("Data name: " + data.name)
  
  fmt.Println( os.Args[1:] )
  
  err = os.MkdirAll(filepath.Join(abspath(), "testing/test"), 0755)
  if err != nil { panic(err) }
  
  parse("test.txt", templateDeb, data)
  
  spawn("go version", false)

	if runtime.GOOS == "darwin" {
		fmt.Println("mac")
	} else if runtime.GOOS == "window" {
		fmt.Println("window")
	} else if runtime.GOOS == "linux" {
		fmt.Println("linux")
	}
  
}
