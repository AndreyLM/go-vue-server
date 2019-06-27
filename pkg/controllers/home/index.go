package home

import "net/http"

// Index - index action
func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}
