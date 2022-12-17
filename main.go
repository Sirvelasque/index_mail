package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("working")
	root := os.Args[1]
	err := filepath.Walk(root, visit)
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", root, err)
	}

}

func visit(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if !info.IsDir() {
		fmt.Println(path)
	}
	return nil
}
