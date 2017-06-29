package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func waitingHandler(waitingTime time.Duration) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusOK)
			time.Sleep(waitingTime)
			fmt.Fprint(w, string("Hello world"))
		})
}

func main() {
	waitingTime, err := time.ParseDuration(os.Getenv("WAITING_TIME"))
	if err != nil {
		fmt.Println(waitingTime)
		waitingTime = 1 * time.Minute
	}
	http.HandleFunc("/", waitingHandler(waitingTime))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
