// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// )

// // The struct is lowercase we want to keep it private. No other package can access it.
// type user struct {
// 	Name string `json:"name"`
// 	Age  int    `json:"age"`
// 	City string `json:"city"`
// }

// func main() {

// 	port := ":3000"

// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Hello Root Route"))
// 		fmt.Println("Hello Root Route")
// 	})

// 	http.HandleFunc("/teachers", func(w http.ResponseWriter, r *http.Request) {
// 		switch r.Method {
// 		case http.MethodGet:
// 			w.Write([]byte("Hello GET Method on Teachers Route"))
// 			fmt.Println("Hello GET Method on Teachers Route")
// 			return
// 		case http.MethodPost:
// 			// Parse form data (necessary for x-www-form-urlencoded)
// 			err := r.ParseForm()
// 			if err != nil {
// 				http.Error(w, "Error parsing form", http.StatusBadRequest)
// 				return
// 			}

// 			fmt.Println("Form:", r.Form)

// 			// Prepare response data
// 			response := make(map[string]interface{})
// 			for key, value := range r.Form {
// 				response[key] = value[0]
// 			}

// 			fmt.Println("Processed response map:", response)

// 			// RAW BODY
// 			body, err := io.ReadAll(r.Body)
// 			if err != nil {
// 				return
// 			}
// 			defer r.Body.Close()
// 			fmt.Println("RAW Body:", body)
// 			fmt.Println("RAW Body:", string(body))

// 			// If you expect json data, then Unmarshall it
// 			var userInstance user
// 			err = json.Unmarshal(body, &userInstance)
// 			if err != nil {
// 				return
// 			}
// 			fmt.Println("Unmarshaled JSON into an instance of user struct:", userInstance)

// 			response1 := make(map[string]interface{})
// 			for key, value := range r.Form {
// 				response[key] = value[0]
// 			}

// 			err = json.Unmarshal(body, &response1)
// 			if err != nil {
// 				return
// 			}

// 			fmt.Println("Unmarshalled JSON into a map:", response1)

// 			// Access the request details
// 			fmt.Println("Body:", r.Body)
// 			fmt.Println("Form:", r.Form)
// 			fmt.Println("Header:", r.Header)
// 			fmt.Println("Context:", r.Context())
// 			fmt.Println("ContentLength:", r.ContentLength)
// 			fmt.Println("Host:", r.Host)
// 			fmt.Println("Method:", r.Method)
// 			fmt.Println("Protocol:", r.Proto)
// 			fmt.Println("Remote Address:", r.RemoteAddr)
// 			fmt.Println("Request URI:", r.RequestURI)
// 			fmt.Println("TLS:", r.TLS)
// 			fmt.Println("Trailer:", r.Trailer)
// 			fmt.Println("Transfer Encoding:", r.TransferEncoding)
// 			fmt.Println("URL:", r.URL)
// 			fmt.Println("User Agent:", r.UserAgent())
// 			fmt.Println("Port:", r.URL.Port())
// 			fmt.Println("Scheme:", r.URL.Scheme)

// 			w.Write([]byte("Hello POST Method on Teachers Route"))
// 			fmt.Println("Hello POST Method on Teachers Route")
// 			return
// 		case http.MethodPut:
// 			w.Write([]byte("Hello PUT Method on Teachers Route"))
// 			fmt.Println("Hello PUT Method on Teachers Route")
// 			return
// 		case http.MethodPatch:
// 			w.Write([]byte("Hello PATCH Method on Teachers Route"))
// 			fmt.Println("Hello PATCH Method on Teachers Route")
// 			return
// 		case http.MethodDelete:
// 			w.Write([]byte("Hello DELETE Method on Teachers Route"))
// 			fmt.Println("Hello DELETE Method on Teachers Route")
// 			return
// 		}

// 		w.Write([]byte("Hello Teachers Route"))
// 		fmt.Println("Hello Teachers Route")
// 	})

// 	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Hello Students Route"))
// 		fmt.Println("Hello Students Route")
// 	})

// 	http.HandleFunc("/execs", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Hello Execs Route"))
// 		fmt.Println("Hello Execs Route")
// 	})

// 	fmt.Println("Server is running on port:", port)
// 	err := http.ListenAndServe(port, nil)
// 	if err != nil {
// 		log.Fatalln("Error starting the server", err)
// 	}
// }
