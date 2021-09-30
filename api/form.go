//Reading From data
package handler

import (
	"fmt"
	"net/http"
)

func MyFormHandler(w http.ResponseWriter, r *http.Request) {

   //use like this: https://zimkit.vercel.app/api/form (als action link im form)

   //f := r.URL.Query().Get("f") // um Query Vars auszulesen
   from := r.Form.Get("date-from")
   to := r.Form.Get("date-to")

   w.Write([]byte(fmt.Sprintf("The start date is %s and the end date is %s", from, to)))
   
}