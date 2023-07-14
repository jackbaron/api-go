package helpers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func BindingDataBody(r *http.Request, data interface{}) {
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		// fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, data)
}
