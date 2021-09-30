//Reading From data
package handler

import (
	"fmt"
	"net/http"
)

func PostMyFormHandler(w http.ResponseWriter, r *http.Request) {

   //use like this: https://zimkit.vercel.app/api/form (als action link im form)

   //f := r.URL.Query().Get("f") // um Query Vars auszulesen
   start := r.Form.Get("start")
   end := r.Form.Get("end")

   w.Write([]byte(fmt.Sprintf("The start date is %s and the end date is %s", start, end)))
   
}