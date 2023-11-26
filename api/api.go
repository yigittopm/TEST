package api

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	connStr string
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		_, err := w.Write([]byte("Hello World"))
		if err != nil {
			//TODO: Implement me
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (s *Server) NewServer(connStr string) *Server {
	return &Server{connStr: connStr}
}

func (s *Server) Start() {
	http.HandleFunc("/hello", HelloWorld)
	fmt.Printf("Server is running on port: %s", s.connStr)
	log.Fatal(http.ListenAndServe(s.connStr, nil))
}
