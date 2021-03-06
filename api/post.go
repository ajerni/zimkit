//POST Request to invoke Particle.function
package handler

import (
   "net/http"
   "net/url"
   "log"
   "os"
)

func MyPostFunctionHandler(w http.ResponseWriter, r *http.Request) {

   //use like this: https://zimkit.vercel.app/api/post?f=fuktionName&a=argumetName (f=setmessage&a=neuerWert)

   f := r.URL.Query().Get("f") // die Particle.function die aufgerufen werden soll
   a := r.URL.Query().Get("a")

   params := url.Values{} //definiert den Body des POST request
  	params.Add("args", a)
  	params.Add("access_token", os.Getenv("MY_API_KEY"))
  
   resp, err := http.PostForm("https://api.particle.io/v1/devices/2c0030000447343337373739/" + f, params) //Enthält automatisch den richtigen Header mit {"Content-Type": "application/x-www-form-urlencoded"},

   if err != nil {   
      log.Printf("Request Failed: %s", err)
      return
   }
  
   defer resp.Body.Close()
   
}