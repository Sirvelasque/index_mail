package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Email struct {
	MessageID string `json:"message_id"`
	From      string `json:"from"`
	To        string `json:"to"`
	Subject   string `json:"subject"`
	Content   string `json:"content"`
}

func main() {
	fmt.Println("starting server...")
	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)
	router.Get("/search", searchHandler)
	http.ListenAndServe(":8080", router)

}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	data := search(key)
	var result map[string]interface{}
	json.Unmarshal(data, &result)

	emails := make([]Email, 0)
	hits := result["hits"].(map[string]interface{})
	for _, hit := range hits["hits"].([]interface{}) {
		hitMap := hit.(map[string]interface{})
		source := hitMap["_source"].(map[string]interface{})
		emails = append(emails, Email{
			MessageID: source["MessageID"].(string),
			From:      source["From"].(string),
			To:        source["To"].(string),
			Subject:   source["Subject"].(string),
			Content:   source["Content"].(string),
		})
	}
	fmt.Println("hits----------------")
	fmt.Print(hits)
	fmt.Println("hits-----------------")
	fmt.Println("emails----------------")
	fmt.Print(emails)
	fmt.Println("emails-----------------")
	json.NewEncoder(w).Encode(emails)
}

func search(key string) []byte {
	query := fmt.Sprintf(`{
    "search_type": "match",
    "query": {
        "term": "%s",
        "start_time": "2023-01-01T14:28:31.894Z",
        "end_time": "2023-01-03T23:28:31.894Z"
    },
    "from": 0,
    "max_results": 20,
    "_source": []
}`, key)
	req, err := http.NewRequest("POST", "http://localhost:4080/api/emails/_search", strings.NewReader(query))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}
