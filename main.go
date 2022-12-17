package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Email struct {
	MessageID string
	From      string
	To        string
	Subject   string
	Content   string
}

var Mails []Email

func main() {
	fmt.Println("working")
	root := os.Args[1]
	err := filepath.Walk(root, visit)
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", root, err)
	}
	fmt.Println(Mails)

}

func visit(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if !info.IsDir() {
		appendInfo(path)
	}
	return nil
}

func appendInfo(path string) {

	data := Email{}
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("error reading file: %v\n", err)
		return
	}

	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		key, obj := asignLine(line)
		if key != "" {
			// fmt.Println(line)
			switch key {
			case "MessageID":
				data.MessageID = obj
			case "From":
				data.From = obj
			case "To":
				data.To = obj
			case "Subject":
				data.Subject = obj
			default:
				continue
			}

		}
	}
	fmt.Println("DATA : \n", data)
	{
		pushData(data)
	}
	fmt.Println("MAILS \n", Mails)
}

func asignLine(line string) (string, string) {
	var key string
	var obj string
	switch getKey(line) {
	case "Message-ID:":
		key = "MessageID"
		obj = strings.Split(line, " ")[1]
	case "From:":
		key = "From"
		obj = strings.Split(line, " ")[1]
	case "To:":
		key = "To"
		obj = strings.Split(line, " ")[1]
	case "Subject:":
		key = "Subject"
		obj = strings.Join(strings.Split(line, " ")[1:], " ")
	default:
		key = ""
		obj = ""
	}
	return key, obj
}

func getKey(line string) string {
	return strings.Split(line, " ")[0]
}

func pushData(data Email) {
	Mails = append(Mails, data)
}
