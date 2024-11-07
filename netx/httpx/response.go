package httpx

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	*http.Response
	err error
}

func (r *Response) ScanJSON(data any)  error {
	if r.err != nil {
		return r.err
	}
	return json.NewDecoder(r.Body).Decode(data)
}