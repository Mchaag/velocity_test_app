package function

import (
	"fmt"
	"net/http"
)

// Handle a function invocation
func Handle(req handler.Request) (handler.Response, error) {
	var err error

	name, age := determineAge(string(req.Body))
	message := fmt.Sprintf("%s is %d years old", string(req.Body))

	return handler.Response{
		Body:       []byte(message),
		StatusCode: http.StatusOK,
	}, err
}

func determineAge(name string) (string, int) {

	a := make(map[string]int)

	a["Matt"] = 28
	a["Jordan"] = 28
	a["Trevor"] = 27
	a["Tom"] = 31

	return name, a[name]
}
