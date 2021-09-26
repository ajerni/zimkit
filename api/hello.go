package handler

import (
   "net/http"
)

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	  nam := r.URL.Query().Get("name")
	  w.Write([]byte("Hallo " + nam))
  }