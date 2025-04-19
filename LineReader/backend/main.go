package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
)

func FileHandler(w http.ResponseWriter, r *http.Request) {
	// CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*") 
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	
	r.ParseMultipartForm(10 << 20) // 10 MB MAX
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "File read error", http.StatusBadRequest)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var lines, words, chars int

	for {
		line, err := reader.ReadString('\n')
		if line != "" {
			lines++
			// Count actual words by splitting the line
			words += len(strings.Fields(line))
			chars += len(line)
		}
		if err != nil {
			break
		}
	}
	
	result := fmt.Sprintf(`{
	"lines": %d,
	"words": %d,
	"characters": %d
	}`, lines, words, chars)
	
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(result))
}

func main() {
	http.HandleFunc("/upload", FileHandler)
	fmt.Println("Server is running in http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}