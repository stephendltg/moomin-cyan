package main

import (
  "encoding/json"
  "io/ioutil"
	"fmt"
	"runtime"
  "strconv"
  "path/filepath"
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
	if err != nil {
		panic(err)
	}
	exe, err = filepath.EvalSymlinks(exe)
	if err != nil {
		panic(err)
	}
	dir := filepath.Dir(exe)
	return dir
}

// Parse template
func parse(path string, tmpl string) {
  t, err := template.New("tmpl").Parse(tmpl)
  if err != nil { panic(err) }

	f, err := os.Create(path)
  if err != nil { panic(err) }

	err = t.Execute(f, data)
  if err != nil { panic(err) }
	f.Close()
}


func main() {
  
  // Read package.json
  jsonPackage, err := os.Open(filepath.Join(abspath(), "package.json"))
  if err != nil {
    panic(err)
  }
  defer jsonPackage.Close()
  byteValue, _ := ioutil.ReadAll(jsonPackage)
  var data Data
  if err := json.Unmarshal(byteValue, &data); err != nil {
	 	panic(err)
	}
                                 
  fmt.Println("Data name: " + data.name)
  
  fmt.Println( os.Args[1:] )
  
  err := os.MkdirAll(filepath.Join(abspath(), "testing/test"), 0755)
  if err != nil { panic(err) }
  
  parse("test.txt", templateDeb)

	if runtime.GOOS == "darwin" {
		fmt.Println("mac")
	} else if runtime.GOOS == "window" {
		fmt.Println("window")
	} else if runtime.GOOS == "linux" {
		fmt.Println("linux")
	}
  
  /*
  parts := strings.Split(`pushnotify.exe --to=gqukgkJyLtchaLE41WUEJ2qFM7Q3tb --title=New Plex content --message={showname} season {showseasonnumber} episode {showepisodenumber} just landed on Plex --sound=magic`, " -")
	head := parts[0]
	args := parts[1:len(parts)]
	//	Format the command
	cmd := exec.Command(head, args...)
	//	Sanity check -- capture stdout and stderr:
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	//	Run the command
	cmd.Run()
	//	Output our results
	fmt.Printf("Result: %v / %v", out.String(), stderr.String())
  */
}
