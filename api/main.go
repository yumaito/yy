package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("start server")

	res := &Response{
		ID: 12345,
	}
	j, err := json.Marshal(res)
	if err != nil {
		log.Println(err.Error())
		return
	}
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, string(j))
		return
	})

	http.ListenAndServe(":5000", nil)
}

// Response is XXX
type Response struct {
	ID uint64 `json:"id,string"`
}
