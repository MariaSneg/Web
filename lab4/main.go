package main

import (
	"net/http"
)

func main() {
	const port = ":3000"

	mux := http.NewServeMux()
	//mux.HandleFunc("/home", index)

	//aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
	http.ListenAndServe(port, mux) //что то дополнительное нужно до этого, т к на этом оменте программа замирает и дальше не делает
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}
