package main

import (
	"net/http"
	"fmt"
	"strconv"
	"io/ioutil"
	"encoding/binary"
)

// TODO move to properties or parameter
const hostname = "localhost"
const port = "8080"
const protocol = "http"
const defaultError = "Ошибка"

func main()  {
	result := SomePublicMethods(32.09)
	println(result)
}

func SomePublicMethods(sum float64) string {
	response, err := http.Get(fmt.Sprintf("%s://%s:%s/request/%s", protocol, hostname, port,  FloatToString(sum)))

	var result = defaultError
	if err == nil {
		defer response.Body.Close()
		responseData, err := ioutil.ReadAll(response.Body)
		if(err == nil && binary.Size(responseData) > 0) {
			result = string(responseData)
		}
	}

	return result
}

func FloatToString(input_num float64) string {
	// to convert a float number to a string 2 decimal places
	return strconv.FormatFloat(input_num, 'f', 2, 64)
}