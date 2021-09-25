//POST Request to Particle cloud publishing an event to all listeners
package handler

import (
   "net/http"
   "net/url"
   "log"
)

func Handler(w http.ResponseWriter, r *http.Request) {

   //use like this: https://zimkit.vercel.app/api/get?key=nameOfEventToPublish z.B alarmall

   keys := r.URL.Query()["key"] //hier wird der parameter key ausgelesen (eventname der published werden soll)
   key := keys[0]

   params := url.Values{} //definiert den Body des POST request
  	params.Add("name", key)
  	params.Add("access_token", "906d5e4a9041e4c0773cad80ccf23490fe83e76c")
  	params.Add("data", "")
  
   resp, err := http.PostForm("https://api.particle.io/v1/devices/events", params) //Enthält automatisch den richtigen Header mit Content Type

   if err != nil {   
      log.Printf("Request Failed: %s", err)
      return
   }
  
   defer resp.Body.Close()
   
}