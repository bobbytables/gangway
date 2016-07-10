package server

import "net/http"

func getRecipes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
