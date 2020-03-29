package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

func main() {
	log.Println("start server")

	if err := run(); err != nil {
		log.Println(err.Error())
		return
	}
}

func run() error {

	server, err := NewServer()
	if err != nil {
		return errors.Wrap(err, "error NewServer")
	}

	if err := http.ListenAndServe(":5000", server); err != nil {
		return errors.Wrap(err, "error ListenAndServe")
	}

	return nil
}

// NewServer is XXX
func NewServer() (*http.ServeMux, error) {
	server := http.NewServeMux()

	server.HandleFunc("/api/login", LoginHandler)
	server.HandleFunc("/api/me", MeHandler)

	return server, nil

}

// LoginHandler is XXX
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer r.Body.Close()

	req := &LoginRequest{}
	if err := json.Unmarshal(buf, req); err != nil {
		log.Println(err.Error())
		return
	}

	if req.Name != "yumaito" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	log.Printf("requestt: %+v", req)

	res := &LoginResponse{
		Token: "hoge-token",
	}
	j, err := json.Marshal(res)
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Fprint(w, string(j))
	return
}

// MeHandler is XXX
func MeHandler(w http.ResponseWriter, r *http.Request) {
	res := &MeResponse{
		ID:   123,
		Name: "yumaito",
		Age:  31,
	}
	j, err := json.Marshal(res)
	if err != nil {
		log.Println(err.Error())
		return
	}

	fmt.Fprint(w, string(j))
	return
}

// LoginRequest is XXX
type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// LoginResponse is XXX
type LoginResponse struct {
	Token string `json:"token"`
}

// MeResponse is XXX
type MeResponse struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
	Age  uint32 `json:"age"`
}
