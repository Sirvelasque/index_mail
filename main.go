package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("working")
	path := os.Args[1]
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("folder path", path)
	fmt.Println("file content", string(bytes))

}
