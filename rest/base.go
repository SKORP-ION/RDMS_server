package rest

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status string `json:"status"`
	Data interface{} `json:"data"`
}

func SendResponse(status int, w *http.ResponseWriter, data interface{}) {
	//wp = Writer_Pointer
	wp := *w
	wp.Header().Add("Content-type", "application/json")
	switch {
	case status >= 400 && status < 500:
		{
			wp.WriteHeader(status)
			json.NewEncoder(wp).Encode(Response{"Client error", data})
		}
	case status >= 500:
		{
			wp.WriteHeader(status)
			json.NewEncoder(wp).Encode(Response{"Server error", data})
		}
	case status >= 300 && status < 400:
		{
			wp.WriteHeader(status)
			json.NewEncoder(wp).Encode(Response{"Redirection", data})
		}
	case status >= 200 && status < 300:
		{
			wp.WriteHeader(status)
			json.NewEncoder(wp).Encode(Response{"Success", data})
		}
	case status >= 100 && status < 200:
		{
			wp.WriteHeader(status)
			json.NewEncoder(wp).Encode(Response{"Informational", data})
		}
	}
}
