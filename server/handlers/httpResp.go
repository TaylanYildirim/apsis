package handlers

import (
	"apsis/util/errorUtil"
	"encoding/json"
	"net/http"
)

func SetHTTPStatus(w http.ResponseWriter, status int, message string, data int) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	body := map[string]interface{}{
		"message": message,
		"data":    data,
	}
	jsonResp, err := json.Marshal(body)
	errorUtil.Check(err)
	w.Write(jsonResp)
}

func SetHTTPResponse(w http.ResponseWriter, status int, data *[]map[string]int) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	jsonResp, err := json.Marshal(data)
	errorUtil.Check(err)
	w.Write(jsonResp)
}

func SetHTTP(w http.ResponseWriter, status int, data *[]map[string]string) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	jsonResp, err := json.Marshal(data)
	errorUtil.Check(err)
	w.Write(jsonResp)
}
