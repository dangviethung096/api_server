package core

import "net/http"

func init() {
	httpClient = &http.Client{}
}
