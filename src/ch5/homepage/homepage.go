package main

import (
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", homepage)

}

func logPanics(function func(http.ResponseWriter, request *http.Request)) func(http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic: %v", request.RemoteAddr, x)
			}
		}()
		function(writer, request)
	}
}

func homepage(write http.ResponseWriter, request *http.Request) {

	//dummy
}

