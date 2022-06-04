package response

import (
	"encoding/json"
	"net/http"
	"ransmart_auth/app/models/responseModel"
)

func Response(w http.ResponseWriter, code int, msg string, data interface{}) {
	receiver := responseModel.Response{}

	receiver.Status = cekStatus(code)
	receiver.Code = code
	receiver.Message = msg
	receiver.Data = data

	jadiJson, err := json.Marshal(receiver) // nge convert jadi json

	w.Header().Set("Content-Type", "application/json") // return type data nya

	if err != nil {
		w.WriteHeader(500) // status code
		w.Write([]byte("Error to marshall"))
	}

	// success
	w.WriteHeader(code) // status code
	w.Write(jadiJson)   // return datanya
}

func ResponseRunningService(w http.ResponseWriter) {
	Response(w, http.StatusOK, service_running, nil)
}
