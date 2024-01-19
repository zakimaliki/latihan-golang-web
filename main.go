package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// blue Print
type Response struct {
	Message string
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// validasi method
	if r.Method == "GET" {
		// implement blue print
		obj := Response{
			Message: "Hello World",
		}
		// konversi ke json
		res, err := json.Marshal(obj)
		// validasi konversi json
		if err != nil {
			http.Error(w, "Gagal Konversi ke json", http.StatusInternalServerError)
		}
		// set header
		w.Write(res)
		w.Header().Set("Content-Type", "application/json")
	} else {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
	}

}

func main() {
	http.HandleFunc("/hello", HelloHandler)
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
