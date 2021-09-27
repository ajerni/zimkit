//POST Request to invoke Particle.function //generische version mit Paremeter f und a --> siehe post.go
package handler

import (
	"log"
	"net/http"
	"net/url"
   "os"
)

func MyMessageHandler(w http.ResponseWriter, r *http.Request) {

   //use like this: https://zimkit.vercel.app/api/postmessage?key=neuerWert

   keys := r.URL.Query()["key"] //hier wird der parameter key ausgelesen (eventname der published werden soll)
   key := keys[0] //Argument, das an die Particle.function setmessage übergeben wird (also der neue Wert)

   f := "setmessage" //die Particle.function die aufgerufen werden soll

   params := url.Values{} //definiert den Body des POST request
  	params.Add("args", key)
  	params.Add("access_token", os.Getenv("MY_API_KEY"))
   //params.Add("access_token", "906d5e4a9041e4c0773cad80ccf23490fe83e76c")
  
   resp, err := http.PostForm("https://api.particle.io/v1/devices/2c0030000447343337373739/" + f, params) //Enthält automatisch den richtigen Header mit {"Content-Type": "application/x-www-form-urlencoded"},

   if err != nil {   
      log.Printf("Request Failed: %s", err)
      return
   }
  
   defer resp.Body.Close()
   
}