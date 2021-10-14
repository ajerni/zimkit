package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func MyUserHandler(w http.ResponseWriter, r *http.Request) {
	sampleUsers := map[string]bool{
		"Andi": true,
		"Max":  true,
	}

	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	sbs := string(bs)
	fmt.Println("USERNAME: ", sbs)

	fmt.Fprint(w, sampleUsers[sbs]) // retourniert true, wenn user vorhanden oder den default Wert false
}
