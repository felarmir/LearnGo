package lesson4

import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	io.WriteString(res, "<doctype html><html><head><title>Hello from Go Server</title></head><body>Hello GO!</body></html>")
	//static files
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
}

func HttpServerGO() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
