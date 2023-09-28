package core

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func HttpRequest(url string, method string, headers map[string]string, body []byte) (*HttpResponse, Error) {
	if body == nil {
		body = []byte{}
	}
	fmt.Printf("Request: \n%s\n", string(body))
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, NewError(HttpErrorCodeInitRequestFail, err)
	}

	for key, val := range headers {
		req.Header.Add(key, val)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, NewError(HttpErrorCodeRequestHttpFail, err)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, NewError(HttpErrorCodeParseResponseFail, err)
	}

	response := &HttpResponse{
		Body:       resBody,
		StatusCode: int32(res.StatusCode),
	}

	return response, nil
}
