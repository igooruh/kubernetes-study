package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var startedAt = time.Now()

func main() {
	http.HandleFunc("/healthz", Healthz)
	http.HandleFunc("/configmap", ConfigMap)
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":80", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World! Tudo bem?"))
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("go/myfamily/family.txt")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(string(data))
}

func Healthz(w http.ResponseWriter, r *http.Request) {

	duration := time.Since(startedAt)

	if duration.Seconds() < 10 {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
	} else {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	}
}
