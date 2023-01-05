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
	json.NewEncoder(w).Encode(string(search(key)))
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
	// req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

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
