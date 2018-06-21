package main

import (
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
	"strings"
	"io/ioutil"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	log.Fatal(http.ListenAndServe(":8080", proxy))

	proxy.OnResponse().DoFunc(InjectField)
}

func InjectField(response *http.Response, proxyCtx *goproxy.ProxyCtx) *http.Response {

	contentType := response.Header.Get("Content-Type")

	// NOTE: Should also check for successful 2xx status code
	if strings.Contains(contentType, "json") {

		defer response.Body.Close()

		bodyBytes, err := ioutil.ReadAll(response.Body)

		if err == nil {

			bodyString := string(bodyBytes)

			bodyStringTampered := TamperJsonString(bodyString)

			response.Body = bodyStringTampered

		} else {

			// NOTE: would want to record this somewhere to be debugged by Chaos team

			// NOTE: since we read, or attempted to read, the response body if we continue from here
			// is it empty now?
		}
	}

	return response
}


