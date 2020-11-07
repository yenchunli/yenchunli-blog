package main

import (
	"fmt"
	"os"
	"path"
	"log"
	"io/ioutil"
)

func main() {
	var location string
	location = "content/posts"
	fullPath := path.Join("..", location)
	info, err := os.Stat(fullPath)

	if err != nil {
		log.Fatal(err)
	}

	if !info.IsDir() {
		log.Fatal("Path doesn't exist: %s", fullPath)
	}
	files, _ := ioutil.ReadDir(fullPath)

	
	for _, f := range files {
		fmt.Println(f.Name())
	}
	

	//fmt.Println(files[1])
	//file, err := ioutil.ReadFile(path.Join(fullPath, files[1].Name()))

	//fmt.Println(string(file))
	
}