package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"example.com/go-demo-1/Documents/GitHub/counting-requests/src/counter"
	"example.com/go-demo-1/Documents/GitHub/counting-requests/src/storage"
)

func main() {
	store := storage.NewFileStorage("counter_data.json")
	cnt := counter.NewCounter(60*time.Second, store)

	cnt.Load()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cnt.Increment()
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(cnt.Count()))
	})

	go func() {
		log.Println("Starting server on :8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("Server failed: %s\n", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
}
