package core

import (
	"net/http"
	"time"
)

func init() {
	httpClient = &http.Client{
		Transport: http.DefaultTransport,
		Timeout:   time.Second * 60,
	}
}
