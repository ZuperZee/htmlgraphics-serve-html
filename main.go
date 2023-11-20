package main

import (
	"log"
	"net/http"
	"os"
)

func getSomethingHtml(w http.ResponseWriter, r *http.Request) {
	content, err := os.ReadFile("./html/something.html")
	if err != nil {
		w.WriteHeader(500)
		log.Print(err)
		return
	}
	w.WriteHeader(200)
	w.Write(content)
}

func main() {
	http.HandleFunc("/something.html", getSomethingHtml)

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
