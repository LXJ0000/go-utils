package httpx

import (
	"context"
	"net/http"
)

func ExampleRequest() {
	var (
		ctx    = context.Background()
		url    = ""
		method = ""
		client = http.DefaultClient
	)
	var resp response
	req := request{
		name: "name",
		age:  20,
	}
	if err := NewRequest(ctx, method, url).BodyWithJSON(req).
		Client(client).
		Header("key", "value").
		Do().ScanJSON(&resp); err != nil {
		// handle error
	}
}

type request struct {
	name string
	age  int
}

type response struct {
}
