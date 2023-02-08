package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed version.txt
var version string

//go:embed logo.jpg
var logo []byte

//go:embed files/*.txt
var path embed.FS

func main() {
	// Output embed version.txt
	fmt.Println(version)

	// Output embed logo.jpg
	err := ioutil.WriteFile("logo_new.jpg", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	// Output embed files/*.txt
	dirEntry, _ := path.ReadDir("files")
	for _, entry := range dirEntry {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}

}
