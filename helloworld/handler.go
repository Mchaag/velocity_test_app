package function

import (
	"fmt"
	"net/http"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

// Handle a function invocation
func Handle(req handler.Request) (handler.Response, error) {
	var err error

	name := string(req.Body)

	age, errMsg := getAge(name)

	if errMsg != "" {
		return handler.Response{
			Body:       []byte(errMsg),
			StatusCode: http.StatusOK,
		}, err

	}
	message := fmt.Sprintf("%s is %d years old", name, age)

	return handler.Response{
		Body:       []byte(message),
		StatusCode: http.StatusOK,
	}, err
}

func getAge(name string) (int, string) {

	a := make(map[string]int)

	a["Matt"] = 28
	a["Jordan"] = 28
	a["Trevor"] = 27
	a["Tom"] = 31

	if val, ok := a[name]; ok {
		return val, ""
	}

	return 0, fmt.Sprintf("No entry in our file for the name: %s", name)

}
