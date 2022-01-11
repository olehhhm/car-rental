package routes

import (
	"net/http"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Luke, I am Your Server."))
}
