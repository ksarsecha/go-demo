package utils

import (
	"encoding/json"
	"errors"
	"net/http"
)

func ParseRequest(req *http.Request, data interface{}) error {
	if req == nil {
		return errors.New("request can not be nil")
	}
	if req.Body == nil {
		return errors.New("request body can not be nil")
	}

	decoder := json.NewDecoder(req.Body)

	return decoder.Decode(data)
}
