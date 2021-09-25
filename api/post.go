//POST Request to invoke Particle.function
package handler

import (
   "net/http"
   "net/url"
   "log"
)

func Handler(w http.ResponseWriter, r *http.Request) {

   //use like this: https://zimkit.vercel.app/api/post?f=particleFunction&a=newMessageText
   // https://zimkit.vercel.app/api/post?f=setmesssage&a=neu

   funcs := r.URL.Query()["f"] //hier wird der parameter key ausgelesen (eventname der published werden soll)
   f := funcs[0] // die Particle.Function

   args := r.URL.Query()["a"]
   a := args[] // das Argument (neuer Text)

   params := url.Values{} //definiert den Body des POST request
  	params.Add("args", a)
  	params.Add("access_token", "906d5e4a9041e4c0773cad80ccf23490fe83e76c")
  
   resp, err := http.PostForm("https://api.particle.io/v1/devices/2c0030000447343337373739/" + f, params) //Enthält automatisch den richtigen Header mit {"Content-Type": "application/x-www-form-urlencoded"},

   if err != nil {   
      log.Printf("Request Failed: %s", err)
      return
   }
  
   defer resp.Body.Close()
   
}