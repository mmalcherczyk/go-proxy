package cmd

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", ProxyHandler)
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
