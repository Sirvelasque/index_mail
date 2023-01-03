package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"strings"
)

// ** Use a map instead of struct, then make the To values to be a slice to get all the posible remitents
// ** Use an Emails proccess library like mail or go-imap to optimize the data reading
// ** Paralize the the process with goroutines
// ** Analyze data with boyer-moore algorithm
type Email struct {
	MessageID string
	From      string
	To        string
	Subject   string
	Content   string
}

var Mails []Email

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	fmt.Println("working")
	root := os.Args[1]
	err := filepath.Walk(root, visit)
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", root, err)
	}
	// **Rebuild the way to call bulk in order to bulk info every 64 items and clean the Mails slice
	bulkEmails()

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
	// **Use strings.splitafter(`\r\n`) to optimize the splitting proccess
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

// ** Use a map structure to avoid using switch
// ** Use stringsTrimSpace to take out the blank spaces
// ** Use strings.Index() to find the `:` and the use the index to exteract the values
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
	return content[pos:]
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

func bulkEmails() {
	bulkUrl := "http://localhost:4080/api/_bulkv2"
	type JSONObject struct {
		Index   string  `json:"index"`
		Records []Email `json:"records"`
	}
	toBulk := JSONObject{
		Index:   "emails",
		Records: Mails,
	}
	jsonMails, err := json.Marshal(toBulk)
	if err != nil {
		fmt.Println(err)
		return
	}
	req, erreq := http.NewRequest("POST", bulkUrl, bytes.NewBuffer(jsonMails))
	if erreq != nil {
		fmt.Println(erreq)
	}
	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("res error: ", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
}
