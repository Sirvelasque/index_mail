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

	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("error reading file: %v\n", err)
		return
	}

	lines := strings.Split(string(content), "\n")
	var MessageID string
	var From string
	var To string
	var Subject string
	var Content string
	for _, line := range lines {
		key, obj := asignLine(line)
		if key != "" {
			switch key {
			case "MessageID":
				MessageID = obj
			case "From":
				From = obj
			case "To":
				To = obj
			case "Subject":
				Subject = obj
			case "Content":
				Content = getMessage(string(content), obj)
				continue
			default:

			}

		}
	}
	pushData(createMail(MessageID, From, To, Subject, Content))
}

func asignLine(line string) (string, string) {
	var key string
	var obj string
	switch getKey(line) {
	case "Message-ID:":
		key = "MessageID"
		if len(strings.Split(line, " ")) > 1 {
			obj = strings.Split(line, " ")[1]
		} else {
			obj = ""
		}
	case "From:":
		key = "From"
		if len(strings.Split(line, " ")) > 1 {
			obj = strings.Split(line, " ")[1]
		} else {
			obj = ""
		}
	case "To:":
		key = "To"
		if len(strings.Split(line, " ")) > 1 {
			obj = strings.Split(line, " ")[1]
		} else {
			obj = ""
		}
	case "Subject:":
		key = "Subject"
		if len(strings.Split(line, " ")) > 2 {
			obj = line[9:]
		} else {
			obj = ""
		}

	case "X-FileName:":
		key = "Content"
		words := strings.Split(line, " ")
		obj = words[len(words)-1]
	default:
		key = ""
		obj = ""
	}
	return key, obj
}

func getKey(line string) string {
	return strings.Split(line, " ")[0]
}

func getMessage(content string, obj string) string {
	pos := strings.Index(content, obj)
	if pos == -1 {
		return "Oops! something went wrong getting the message content"
	}
	return content[pos+21:]
}

func pushData(data Email) {
	Mails = append(Mails, data)
}

func createMail(MessageID string, From string, To string, Subject string, Content string) Email {
	return Email{
		MessageID: MessageID,
		From:      From,
		To:        To,
		Subject:   Subject,
		Content:   Content,
	}
}
