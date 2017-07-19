package main

import (
	"net/http"
	"fmt"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	rStr := "<a href='/public'>Hello World!</a>"
	w.Header().Set("Content-Length", fmt.Sprint(len(rStr)))
	fmt.Fprint(w, string(rStr))
}

func main() {
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	http.HandleFunc("/", helloHandler)

	fmt.Println("Server running on port 3000.")
	http.ListenAndServe(":3000", nil)

}