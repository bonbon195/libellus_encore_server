package get_update_date

import (
	"encore.app/firebasesdk"
	"net/http"
)

//encore:api public raw method=GET path=/getUpdateDate
func getUpdateDate(w http.ResponseWriter, r *http.Request) {
	var date string
	err := firebasesdk.GetUpdateDate(&date)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write([]byte(date))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
