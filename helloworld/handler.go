package function

import (
	"fmt"
	"net/http"
	"github.com/openfaas-incubator/go-function-sdk"
)

// Handle a function invocation
func Handle(req handler.Request) (handler.Response, error) {
	var err error

	a := make(map[string]int)

	a["Matt"] = 28
	a["Jordan"] = 28
	a["Trevor"] = 27
	a["Tom"] = 31

	name := string(req.Body)

	message := fmt.Sprintf("%s is %d years old", name, a[name])

	return handler.Response{
		Body:       []byte(message),
		StatusCode: http.StatusOK,
	}, err
}
