package util

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// Response Payload
type Payload struct {
	Result  interface{} `json:"result,omitempty"`
	Success string      `json:"success,omitempty"`
	Error   string      `json:"error,omitempty"`
	Count   int         `json:"count,omitempty"`
}

// Write JSON Response
func WriteResponse(res http.ResponseWriter, payload Payload, statusCode int) error {
	content, err := json.Marshal(payload)
	if err == nil {
		res.WriteHeader(statusCode)
		res.Header().Set("Content-Type", "application/json")
		res.Write(content)
	}
	return err
}

// Decode JSON intro struct
func DecodeReqBody(reqBody io.ReadCloser, v interface{}) error {
	body, _ := ioutil.ReadAll(reqBody)
	err := json.Unmarshal(body, v)
	if err != nil {
		log.Printf("Error unmarshaling request body: %s", err)
	}

	reqBody.Close()
	return err
}

// Get Query String
func GetQS(req *http.Request, key string) string {
	return req.URL.Query().Get(key)
}
