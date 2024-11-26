package main

import (
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("<h1>Hello Kubernetes from cefalo Bangladesh</h1>"))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("<h1>Hello from Argo CD from cefalo Bangladesh</h1>"))
	})

	log.Fatal(http.ListenAndServe(":9090", nil))
}