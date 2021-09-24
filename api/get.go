//this is a server less function running on vercel (Acces token etc. kann direkt im Http request sein, da diese Funktion nur auf dem Sever ist!)
//GET Request auf Particle Argon's teperature variable

package handler

import (
   "io/ioutil"
   "log"
   "net/http"
   "fmt"
)

func Handler(w http.ResponseWriter, r *http.Request) {
   resp, err := http.Get("https://api.particle.io/v1/devices/e00fce6839e00714b083442d/temperature?access_token=906d5e4a9041e4c0773cad80ccf23490fe83e76c")
   if err != nil {
      log.Fatalln(err)
   }
//We Read the response body on the line below.
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      log.Fatalln(err)
   }
//Convert the body to type string
   sb := string(body)
   //log.Printf(sb)
   fmt.Fprintf(w, sb)
}
