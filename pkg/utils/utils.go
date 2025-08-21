package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseRequestBody(r *http.Request, v interface{}) {

	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), v); err != nil {
			return
		}
	}
}
