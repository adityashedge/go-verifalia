// Package verifalia provides an API client for Verifalia API. For more details, see http://verifalia.com/developers
package verifalia

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const (
	libraryVersion = "0.1"
	defaultBaseUrl = "https://api.verifalia.com/v1.1/"
	userAgent      = "go-verifalia/" + libraryVersion
	contentType    = "application/json"
)

// A Client manages communication with the Verifalia REST API.
type Client struct {
	// Authentication to the API occurs via HTTP Basic Auth
	// using the sub-account SID as the username and the auth token as password
	AccountSID string
	AuthToken  string

	// Base URL for communicating with the API..
	BaseURL *url.URL

	// User agent used when communicating with the API.
	UserAgent string
}

// Returns a new Verifalia API client.
// It requires account SID and auth token which are used for basic http authentication
func NewClient(accountSID, authToken string) *Client {
	if accountSID == "" || authToken == "" {
		return nil
	}

	baseUrl, _ := url.Parse(defaultBaseUrl)

	c := &Client{
		AccountSID: accountSID,
		AuthToken:  authToken,
		UserAgent:  userAgent,
		BaseURL:    baseUrl,
	}
	return c
}

// NewRequest creates an API request.
// method is the HTTP VERB
// path is the relative URL resolved relative to the BaseURL of the Client (eg. "email-validations").
// Relative URLs should always be specified without a preceding slash.
// It can also be an absolute URL.
// If specified, the value pointed to by body is JSON encoded and included as the request body.
func (c *Client) NewRequest(method, path string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	url := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url.String(), buf)
	if err != nil {
		return nil, err
	}

	// Verifalia uses basic auth
	req.SetBasicAuth(c.AccountSID, c.AuthToken)
	req.Header.Add("Content-Type", contentType)
	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	}
	return req, nil
}
