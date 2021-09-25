//https://zimkit.vercel.app/api/hello?name=name
// package handler

// import (
//    "net/http"
//    "fmt"
// )

// func Handler(w http.ResponseWriter, r *http.Request) {

//    names := r.URL.Query()["name"]
//    name := names[0]

//    fmt.Fprintf(w, "Hello " + name)
// }