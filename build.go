package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
)

const (
	templateInfo = `<?xml version="1.0" encoding="UTF-8"?>
  <!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
  <plist version="1.0">
  <dict>
    <key>CFBundleExecutable</key>
    <string>{{.Name}}</string>
    <key>CFBundleIconFile</key>
    <string>icon</string>
    <key>CFBundleIdentifier</key>
    <string>com.{{.Name}}.1</string>
    <key>NSHighResolutionCapable</key>
    <true/>
    <key>LSUIElement</key>
    <true/>
    <key>NSAppTransportSecurity</key>
    <dict>
        <key>NSAllowsLocalNetworking</key>
        <true/>
    </dict>
    <!-- <dict>
      <key>NSAllowsArbitraryLoads</key>
      <true/>
    </dict> -->
  </dict>
  </plist>
  `
)

// package json data
type Data struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
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
func parse(path string, tmpl string, data Data) {
	t, err := template.New("tmpl").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	err = t.Execute(f, data)
	if err != nil {
		panic(err)
	}
	f.Close()
}

// Spawn process command
func spawn(command string, msg string, passthroughArgs bool) {
	parts := strings.Split(command, " ")
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
	cmd.Run()
	// err := cmd.Run()
	// if err != nil {
	// 	panic(cmd.Stderr)
	// }
	var stdout string
	if out.Len() == 0 {
		stdout = msg
	} else {
		stdout = out.String()
	}

	fmt.Printf("%v\n%v", stdout, stderr.String())
}

// copy file
func copy(source string, destination string) {
	input, err := ioutil.ReadFile(source)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(destination, input, 0644)
	if err != nil {
		panic(err)
	}
}

// Read package
func readPackage() Data {
	jsonPackage, err := os.Open("package.json")
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

	spawn("go version", "undefined", false)

	// Read package.json
	data := readPackage()

	fmt.Println("ℹ Name: " + data.Name)
	fmt.Println("ℹ Description: " + data.Description)
	fmt.Println("ℹ Version: " + data.Version)

	var binaryPath string

	if runtime.GOOS == "darwin" {
		err := os.MkdirAll("bin/"+data.Name+".app/Contents/MacOS", 0755)
		if err != nil {
			panic(err)
		}

		err = os.MkdirAll("bin/"+data.Name+".app/Contents/Resources", 0755)
		if err != nil {
			panic(err)
		}

		parse(
			"bin/"+data.Name+".app/Contents/Info.plist",
			templateInfo,
			data)

		copy(
			"assets/icon.icns",
			"bin/"+data.Name+".app/Contents/Resources/icon.icns")

		binaryPath = "bin/" + data.Name + ".app/Contents/MacOS/" + data.Name
		// Build binary
		spawn(
			"go build -v -o "+binaryPath+" main.go",
			"✔ Go build "+data.Name+".app",
			false)
	} else if runtime.GOOS == "window" {
		binaryPath = "bin/" + data.Name + "-win-amd64.exe -ldflags=\"-H windowsgui\""
		// Build binary
		spawn(
			"go build -v -o "+binaryPath+" main.go",
			"✔ Go build your exe. \n ℹ Use with webview.dll WebView2Loader.dll",
			false)
	} else if runtime.GOOS == "linux" {
		spawn(
			"make build-deb",
			"✔ Go build "+data.Name+".app",
			false)
	} else {
		panic("⚠ Unsupported platform: " + runtime.GOOS)
	}

}
