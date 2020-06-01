package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("testing go app")
	testMessage := os.Getenv("TEST_MESSAGE")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, testMessage)
	})

	log.Fatal(http.ListenAndServe(":9091", nil))
}
