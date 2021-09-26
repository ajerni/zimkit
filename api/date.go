package handler

import (
    "fmt"
    "net/http"
    "time"
)

func MyDateHandler(w http.ResponseWriter, r *http.Request) {
    currentTime := time.Now().Format(time.RFC850)
    fmt.Fprintf(w, currentTime)
}