package overseerr

import (
	"net/http"
	"net/url"
)

const (
	overseerrAPIPath = "/api/v1"
	userAgent        = "go-overseerr"
)

//go:generate mockery --name=Client --case=snake
// A Client manages communication with the API.
type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client
	// BaseURL base URL for API requests, should
	// always be specified with a trailing slash.
	BaseURL *url.URL
	// API Key for Overseerr API
	APIKey string
	// User agent used when communicating with the API.
	UserAgent string

	// Reuse a single struct instead of allocating one for each service on the heap.
	common service
}

type service struct {
	client *Client
}

// CommonParameters are the parameters that are common to all requests.
type CommonParameters struct {
	APIKey   string `url:"apikey"`   //
	Command  string `url:"cmd"`      //
	OutType  string `url:"out_type"` //
	Callback string `url:"callback"` //
	Debug    int    `url:"debug"`    //
}

// Client returns the http.Client used by this client.
func (c *Client) Client() *http.Client {
	clientCopy := *c.client
	return &clientCopy
}

// NewClient returns a new API client. If a nil httpClient is
// provided, a new http.Client will be used.
func NewClient(httpClient *http.Client, baseURL *url.URL, apiKey string) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	c := &Client{client: httpClient, BaseURL: baseURL, APIKey: apiKey, UserAgent: userAgent}
	c.common.client = c
	return c
}
