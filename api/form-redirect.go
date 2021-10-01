// Redirect
package handler

import (
	// "encoding/json"
	// "log"
	"net/http"
)

// type jsonResponse struct {
// 	OK        bool   `json:"ok"`
// 	Message   string `json:"message"`
// 	Start     string `json:"start"`
// 	End       string `json:"end"`
// 	Sonstiges string `json:"sonstiges"`
// }

func PostMyFormHandlerJson(w http.ResponseWriter, r *http.Request) {
	

	//create and send JSON response
	// resp := jsonResponse{
	// 	OK:        true,
	// 	Message:   "Hallo aus der Serverless go function",
	// 	Start:     "fake",
	// 	End:       "data",
	// 	Sonstiges: "Eigenes Zeugs in JSON.",
	// }

	// out, err := json.MarshalIndent(resp, "", "   ")
	// if err != nil {
	// 	log.Println(err)
	// }

	// log.Println((string(out)))

	// w.Header().Set("Content-Type", "aplication/json")
	// w.Write(out)

	http.Redirect(w, r, "/redirected", http.StatusSeeOther)
}

