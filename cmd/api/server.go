package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	mw "schoolREST/internal/api/middlewares"
)

type user struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	City string `json:"city"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello From The Root Route"))
	fmt.Println("Hello From The Root Route")
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("This is the GET Method for teachers routes"))
		fmt.Println("This is the GET Method for teachers routes")
	case http.MethodPost:
		w.Write([]byte("This is the PUT Method for teachers routes"))
		fmt.Println("This is the PUT Method for teachers routes")
	case http.MethodDelete:
		w.Write([]byte("This is the DELETE Method for teachers routes"))
		fmt.Println("This is the DELETE Method for teachers routes")
	case http.MethodPatch:
		w.Write([]byte("This is the PATCH Method for teachers routes"))
		fmt.Println("This is the PATCH Method for teachers routes")
	}

}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("This is the GET Method for students routes"))
		fmt.Println("This is the GET Method for students routes")
	case http.MethodPost:
		w.Write([]byte("This is the PUT Method for students routes"))
		fmt.Println("This is the PUT Method for students routes")
	case http.MethodDelete:
		w.Write([]byte("This is the DELETE Method for students routes"))
		fmt.Println("This is the DELETE Method for students routes")
	case http.MethodPatch:
		w.Write([]byte("This is the PATCH Method for students routes"))
		fmt.Println("This is the PATCH Method for students routes")
	}

}

func execsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("This is the GET Method for Executives routes"))
		fmt.Println("This is the GET Method for Executives routes")
	case http.MethodPost:
		w.Write([]byte("This is the PUT Method for Executives routes"))
		fmt.Println("This is the PUT Method for Executives routes")
	case http.MethodDelete:
		w.Write([]byte("This is the DELETE Method for Executives routes"))
		fmt.Println("This is the DELETE Method for Executives routes")
	case http.MethodPatch:
		w.Write([]byte("This is the PATCH Method for Executives routes"))
		fmt.Println("This is the PATCH Method for Executives routes")
	}

}

func main() {

	port := ":3000"

	cert := "cert.pem"
	key := "key.pem"

	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)

	mux.HandleFunc("/teachers", teachersHandler)

	mux.HandleFunc("/students", studentsHandler)

	mux.HandleFunc("/execs", execsHandler)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	server := &http.Server{
		Addr:      port,
		Handler:   mw.Security_headers(mw.Cors(mux)),
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server starting on port ", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln(err.Error())
	}

}
