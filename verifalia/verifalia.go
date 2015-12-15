// Package verifalia provides an API client for Verifalia API. For more details, see http://verifalia.com/developers
package verifalia

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
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
		log.Fatalln(err)
		return nil, err
	}
	url := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			log.Fatalln(err)
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url.String(), buf)
	if err != nil {
		log.Fatalln(err)
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

// Query the Email Validations API with an array of emails to validate.
// Response returned by this API is available in "Data" struct.
// POST: https://api.verifalia.com/v1.1/email-validations
// Emails to validate are passed as a slice of string.
func (c *Client) Validate(emails []string) (*Response, error) {
	if len(emails) <= 0 {
		err := errors.New("emails must not be empty")
		log.Fatalln(err)
		return nil, err
	}
	// create a request object to send in http request body
	params := Request{}
	for _, email := range emails {
		inp := inputEmail{email}
		params.Entries = append(params.Entries, inp)
	}
	// build request object for "email-validations" API
	req, err := c.NewRequest("POST", "email-validations", params)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	log.Println(req.URL)
	// send request to "email-validations" API with request params
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	// close request body after "Validate" method executes
	defer resp.Body.Close()

	// build a "Response" object from API response body
	r := Response{}
	err = json.NewDecoder(resp.Body).Decode(&r.Data)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	r.StatusCode = resp.StatusCode
	r.Status = http.StatusText(r.StatusCode)
	// next "Location" if any is returned in the "Location" header
	r.Location, _ = resp.Location()
	if r.Data != nil {
		r.UniqueID = r.Data.UniqueID
	}
	return &r, nil
}

// Email address is sent to email validations API as "inputData"
// [{inputData: "test@test.com"}]
type inputEmail struct {
	InputData string `json:"inputData"`
}

// Request body for email validations API is represented by "Request" struct
type Request struct {
	Entries []inputEmail `json:"entries"`
}

// All information about an email in a validation job is represented by an "Entry" struct
type Entry struct {
	InputData                   string     `json:"inputData"`
	Status                      string     `json:"status"`
	CompletedOn                 *time.Time `json:"completedOn"`
	EmailAddress                string     `json:"emailAddress"`
	AsciiEmailAddressDomainPart string     `json:"asciiEmailAddressDomainPart"`
	EmailAddressLocalPart       string     `json:"emailAddressLocalPart"`
	EmailAddressDomainPart      string     `json:"emailAddressDomainPart"`
	HasInternationalDomainName  bool       `json:"hasInternationalDomainName"`
	HasInternationalMailboxName bool       `json:"hasInternationalMailboxName"`
	IsDisposableEmailAddress    bool       `json:"isDisposableEmailAddress"`
	IsRoleAccount               bool       `json:"isRoleAccount"`
	SyntaxFailureIndex          int        `json:"syntaxFailureIndex"`
	IsCatchAllFailure           bool       `json:"isCatchAllFailure"`
	IsSuccess                   bool       `json:"isSuccess"`
	IsSyntaxFailure             bool       `json:"isSyntaxFailure"`
	IsDnsFailure                bool       `json:"isDnsFailure"`
	IsSmtpFailure               bool       `json:"isSmtpFailure"`
	IsMailboxFailure            bool       `json:"isMailboxFailure"`
	IsTimeoutFailure            bool       `json:"isTimeoutFailure"`
	IsNetworkFailure            bool       `json:"isNetworkFailure"`
}

// Data returned by Verifalia for an email validation job is represented by "Data" struct
type Data struct {
	UniqueID      string     `json:"uniqueID"`
	EngineVersion string     `json:"engineVersion"`
	SubmittedOn   *time.Time `json:"submittedOn"`
	CompletedOn   *time.Time `json:"completedOn"`
	Entries       []Entry    `json:"entries"`
	Progress      struct {
		NoOfTotalEntries     int `json:"noOfTotalEntries"`
		NoOfCompletedEntries int `json:"noOfCompletedEntries"`
	} `json:"progress"`
}

// All API response will be represented by general purpose "Response" struct
// Response returned after an email validation job is represented by "Data" struct
// Data is pointer so we can ignore it for DELETE job as it will be nil
// Status code represents if job was queued, executed or rejected by Verifalia
// Location stores the next API location after current request.
type Response struct {
	StatusCode int
	Status     string
	Location   *url.URL
	UniqueID   string
	*Data
}
