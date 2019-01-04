package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"
)

func pandocHandler(w http.ResponseWriter, r *http.Request) {
	arg := r.URL.Path[len("/pandoc/"):]
	if !(arg == "version" || arg == "help") {
		log.Printf("404 %s", r.URL.Path)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	start := time.Now()
	out, err := exec.Command("pandoc", "--"+arg).Output()
	elapsed := time.Since(start)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Printf("%d %s", http.StatusInternalServerError, r.URL.Path)
		return
	}
	log.Printf("200 %s", r.URL.Path)
	log.Printf("CMD \"pandoc --%s\" took %s", arg, elapsed)
	fmt.Fprintf(w, "<!DOCTYPE html><html lang=\"en\"><body><pre>%s</pre></body></html>", out)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("404 %s", r.URL.Path)
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func main() {
	http.HandleFunc("/pandoc/", pandocHandler)
	http.HandleFunc("/", defaultHandler)
	log.Print("Listening on <http://localhost:8080>")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
