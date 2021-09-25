// function publishEvent(){
// 	let eventname = "alarmall";
// 	let daten = "";
// 	fetch("https://api.particle.io/v1/devices/events", {
// 		body: "name=" + eventname + "&data=" + daten + "&access_token=" + AUTH_TOKEN,
// 		headers: {"Content-Type": "application/x-www-form-urlencoded"},
// 		method: "POST"
// 	})
// }

//GO
package handler

import (
   "net/http"
   "net/url"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	//fmt.Fprintf(w, "reached at least now")
   //https://zimkit.vercel.app/api/get?key=nameOfEventToPublish alarmall (Parameter key auslesen) 

   keys := r.URL.Query()["key"]
   key := keys[0]

   params := url.Values{}
  	params.Add("name", key)
  	params.Add("access_token", "906d5e4a9041e4c0773cad80ccf23490fe83e76c")
  	params.Add("data", "")
  
   resp:= http.PostForm("https://api.particle.io/v1/devices/events", params) 
  
   defer resp.Body.Close()
   
}