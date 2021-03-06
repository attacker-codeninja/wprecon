package gohttp

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"
)

// HttpRequest :: This function will be used for any request that is made.
func HttpRequest(httpStructs Http) (Response, error) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	if !strings.HasSuffix(httpStructs.URL, "/") {
		httpStructs.URL = fmt.Sprintf("%s/", httpStructs.URL)
	}

	httpStructs.URLFULL = httpStructs.URL+httpStructs.Dir

	if httpStructs.Method == "" {
		httpStructs.Method = "GET"
	}

	request, err := http.NewRequest(httpStructs.Method, httpStructs.URL + httpStructs.Dir, nil)

	if err != nil {
		return Response{}, err
	}

	response, err := client.Do(request)

	if err != nil {
		return Response{}, err
	}

	request.Header.Set("User-Agent", "wprecon - Wordpress Recon (Vulnerability Scanner) (GoHttp 0.0.0.1)")
	if httpStructs.RandomUserAgent == true {
		userAgent := RandomUserAgent()
		request.Header.Set("User-Agent", userAgent)
	}

	httpResponse := Response{
		URL:        httpStructs.URL,
		URLFULL:    httpStructs.URLFULL,
		Dir:		httpStructs.Dir,
		StatusCode: response.StatusCode,
		UserAgent:  request.UserAgent(),
		Body:       response.Body,
	}

	return httpResponse, nil
}
