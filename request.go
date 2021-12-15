package be

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	defaultRequestTimeout = 10 * time.Second
	defaultHttpClient     = &http.Client{
		Timeout: defaultRequestTimeout,
	}
)

func request(client *Client, method, uri string, headers map[string]string, body []byte) (*http.Response, error) {
	return realRequest(client, method, uri, headers, body)
}

// request sends a request to alibaba cloud Log Service.
// @note if error is nil, you must call http.Response.Body.Close() to finalize reader
func realRequest(client *Client, method, uri string, headers map[string]string,
	body []byte) (*http.Response, error) {

	headers["Host"] = client.Endpoint

	if body != nil {
		// TODO set body
	}

	digest, err := signature(client)
	if err != nil {
		return nil, NewClientError(err)
	}
	auth := fmt.Sprintf("Basic %v", digest)
	headers["Authorization"] = auth

	// Initialize http request
	reader := bytes.NewReader(body)

	// Handle the endpoint
	urlStr := fmt.Sprintf("%s/%s", client.Endpoint, uri)
	req, err := http.NewRequest(method, urlStr, reader)
	if err != nil {
		return nil, NewClientError(err)
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	// Get ready to do request
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Parse the sls error from body.
	if resp.StatusCode != http.StatusOK {
		err := &Error{}
		err.HTTPCode = (int32)(resp.StatusCode)
		defer resp.Body.Close()
		buf, ioErr := ioutil.ReadAll(resp.Body)
		if ioErr != nil {
			return nil, NewBadResponseError(ioErr.Error(), resp.Header, resp.StatusCode)
		}
		err.Code = "StatusCodeError"
		err.Message = "Content:" + string(buf)
		return nil, err
	}
	return resp, nil
}

func signature(client *Client) (string, error) {
	if client.UserName == "" || client.PassWord == "" {
		return "", NewClientError(fmt.Errorf("Empty userName or passWord"))
	}
	auth := client.UserName + ":" + client.PassWord
	return base64.StdEncoding.EncodeToString([]byte(auth)), nil
}
