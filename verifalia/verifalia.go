// Package verifalia provides an API client for Verifalia API. For more details, see http://verifalia.com/developers
package verifalia

import "net/url"

const (
	libraryVersion = "0.1"
	defaultBaseUrl = "https://api.verifalia.com/v1.1/"
	userAgent      = "go-verifalia/" + libraryVersion
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
