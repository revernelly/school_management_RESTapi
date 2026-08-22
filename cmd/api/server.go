package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	mw "restapi/internal/api/middlewares"
	"time"
)

// The struct is lowercase we want to keep it private. No other package can access it.
type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Root Route"))
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	// teachers/?key=value&query=value2&sortby=email&sortorder=ASC
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET Method on Teachers Route"))
		return
	case http.MethodPost:
		w.Write([]byte("Hello POST Method on Teachers Route"))
		return
	case http.MethodPut:
		w.Write([]byte("Hello PUT Method on Teachers Route"))
		return
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH Method on Teachers Route"))
		return
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE Method on Teachers Route"))
		return
	}
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET Method on Students Route"))
		return
	case http.MethodPost:
		w.Write([]byte("Hello POST Method on Students Route"))
		return
	case http.MethodPut:
		w.Write([]byte("Hello PUT Method on Students Route"))
		return
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH Method on Students Route"))
		return
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE Method on Students Route"))
		return
	}
}

func execsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET Method on Execs Route"))
		return
	case http.MethodPost:
		w.Write([]byte("Hello POST Method on Execs Route"))
		return
	case http.MethodPut:
		w.Write([]byte("Hello PUT Method on Execs Route"))
		return
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH Method on Execs Route"))
		return
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE Method on Execs Route"))
		return
	}
}

func main() {

	port := ":3000"
	cert := "cert.pem"
	key := "key.pem"

	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)

	mux.HandleFunc("/teachers/", teachersHandler)

	mux.HandleFunc("/students/", studentsHandler)

	mux.HandleFunc("/execs/", execsHandler)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	rl := mw.NewRateLimiter(5, time.Minute)

	hppOptions := mw.HPPOptions{
		CheckQuery:                  true,
		CheckBody:                   true,
		CheckBodyOnlyForContentType: "application/x-www-form-urlencoded",
		Whitelist:                   []string{"sortBy", "sortOrder", "name", "age", "class"},
	}

	secureMux := mw.Hpp(hppOptions)(rl.Middleware(mw.Compression(mw.ResponseTimeMiddleware(mw.SecurityHeaders(mw.Cors(mux))))))

	// Create custom server
	server := &http.Server{
		Addr: port,
		// Handler: mux
		Handler: secureMux,
		// Handler:   middlewares.Cors(mux),
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port:", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}
}
