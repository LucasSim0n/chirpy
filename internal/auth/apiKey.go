package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(headers http.Header) (string, error) {

	fullAuth, ok := headers[http.CanonicalHeaderKey("Authorization")]
	if !ok {
		return "", errors.New("Api key not provided")
	}
	auth := strings.Split(fullAuth[0], " ")
	if auth[0] != "ApiKey" {
		return "", errors.New("Api key not provided")
	}
	return auth[1], nil
}
