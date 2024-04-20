package send_json

import (
	"encoding/json"
	"net/http"
)

func SendJson(w *http.ResponseWriter, data interface{}) {
	(*w).Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(data)
	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = (*w).Write(b)
	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		return
	}
}
