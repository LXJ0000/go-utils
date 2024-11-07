package httpx

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type Request struct {
	request *http.Request
	client  *http.Client
	err     error
}

func NewRequest(ctx context.Context, method, url string) *Request {
	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	return &Request{request: req, err: err, client: http.DefaultClient}
}

func (r *Request) Client(client *http.Client) *Request {
	r.client = client
	return r
}

func (r *Request) BodyWithJSON(jsonData any) *Request {
	data, err := json.Marshal(jsonData)
	if err != nil {
		r.err = err
		return r
	}
	reader := bytes.NewReader(data)
	r.request.Body = io.NopCloser(reader)
	r.request.Header.Set("Content-Type", "application/json")
	return r
}

func (r *Request) Param(key, value string) *Request {
	q := r.request.URL.Query()
	q.Add(key, value)
	r.request.URL.RawQuery = q.Encode()
	return r
}

func (r *Request) Header(key, value string) *Request {
	r.request.Header.Add(key, value)
	return r
}

func (r *Request) Do() *Response {
	if r.err != nil {
		return &Response{err: r.err}
	}
	resp, err := r.client.Do(r.request)
	return &Response{Response: resp, err: err}
}
