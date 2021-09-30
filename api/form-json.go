// Reading From data and returning a JSON
package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type jsonResponse struct {
	OK        bool   `json:"ok"`
	Message   string `json:"Message"`
	Start     string `json:"start"`
	End       string `json:"end"`
	Sonstiges string `json:"sonstiges"`
}

func PostMyFormHandlerJson(w http.ResponseWriter, r *http.Request) {
	start := r.FormValue("start")
	end := r.FormValue("end")

	resp := jsonResponse{
		OK:        true,
		Message:   "Hallo aus der Serverless go function",
		Start:     start,
		End:       end,
		Sonstiges: "Eigenes Zeugs vom" + start,
	}

	out, err := json.MarshalIndent(resp, "", "   ")
	if err != nil {
		log.Println(err)
	}

	log.Println((string(out)))

	w.Header().Set("Content-Type", "aplication/json")
	w.Write(out)
}
