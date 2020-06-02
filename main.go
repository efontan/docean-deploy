package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("starting go app")
	testMessage := os.Getenv("TEST_MESSAGE")
	log.Printf("TEST_MESSAGE=%s\n", testMessage)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("handle request")
		fmt.Fprintf(w, testMessage)
	})

	log.Fatal(http.ListenAndServe(":9091", nil))
}
