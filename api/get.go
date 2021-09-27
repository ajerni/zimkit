//this is a server less function running on vercel (Acces token etc. kann direkt im Http request sein, da diese Funktion nur auf dem Sever ist!)
//GET Request auf Particle Argon's teperature variable (Particle.variable)

package handler

import (
   "io/ioutil"
   "log"
   "net/http"
   "fmt"
   "os"
)

func MyGetHandler(w http.ResponseWriter, r *http.Request) {

   //https://zimkit.vercel.app/api/get?key=nameOfParicleVariable (Parameter key auslesen) / wenn mehrere Parameter siehe post.go
   keys := r.URL.Query()["key"]
   key := keys[0]

   url := "https://api.particle.io/v1/devices/e00fce6839e00714b083442d/" + key + "?access_token=" + os.Getenv("MY_API_KEY")

   resp, err := http.Get(url)
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