package dmm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

const (
	// SiteGeneral is the code as DMM.com
	SiteGeneral = "DMM.com"
	// SiteAdult is the code as DMM.co.jp (FANZA)
	SiteAdult      = "FANZA"
	libraryVersion = "0.0.1"
	defaultBaseURL = "https://api.dmm.com/"
	userAgent      = "go-dmm/" + libraryVersion
	mediaType      = "application/json"
)

// ListOptions is interface that specifies the optional parameters to various List methods
type ListOptions interface {
	Next() error
	GetOffset() int
	GetHits() int
}

type searchResult interface {
	populatePageValues(*Response)
}

// Client manages communication with DMM Affiliate V3 API.
type Client struct {
	// HTTP client used to communicate with the DO API.
	client *http.Client

	// Base URL for API requests.
	BaseURL *url.URL

	// User agent for client
	UserAgent string

	// Services used for communicating with the API
	Actresses ActressesService
	Authors   AuthorsService
	Floors    FloorsService
	Genres    GenresService
	Items     ItemsService
	Makers    MakersService
	Series    SeriesService
}

// NewClient returns a new DMM API client.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent}
	c.Actresses = &ActressesServiceOp{client: c}
	c.Authors = &AuthorsServiceOp{client: c}
	c.Floors = &FloorsServiceOp{client: c}
	c.Genres = &GenresServiceOp{client: c}
	c.Items = &ItemsServiceOp{client: c}
	c.Makers = &MakersServiceOp{client: c}
	c.Series = &SeriesServiceOp{client: c}

	return c
}

// ClientOpt are options for New.
type ClientOpt func(*Client) error

// New returns a new DIgitalOcean API client instance.
func New(httpClient *http.Client, opts ...ClientOpt) (*Client, error) {
	c := NewClient(httpClient)
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

// SetBaseURL is a client option for setting the base URL.
func SetBaseURL(bu string) ClientOpt {
	return func(c *Client) error {
		u, err := url.Parse(bu)
		if err != nil {
			return err
		}

		c.BaseURL = u
		return nil
	}
}

// SetUserAgent is a client option for setting the user agent.
func SetUserAgent(ua string) ClientOpt {
	return func(c *Client) error {
		c.UserAgent = fmt.Sprintf("%s %s", ua, c.UserAgent)
		return nil
	}
}

// Response is a DMM API response. This wraps the standard http.Response returned from API.
type Response struct {
	*http.Response
	Parameters ListOptions

	ResultCount   int
	TotalCount    int
	FirstPosition int
}

// An ErrorResponse reports the error caused by an API request
type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response

	Parameters *ListOptions
	Result     ErrResult `json:"result"`
}

// An ErrResult reports the error caused by an API request
type ErrResult struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors,omitempty"`
}

func (r *ErrorResponse) Error() string {
	if r.Result.Errors != nil {
		return fmt.Sprintf("%v %v: %d %s (%v)",
			r.Response.Request.Method,
			r.Response.Request.URL,
			r.Response.StatusCode,
			r.Result.Message,
			r.Result.Errors,
		)
	}
	return fmt.Sprintf("%v %v: %d %s",
		r.Response.Request.Method,
		r.Response.Request.URL,
		r.Response.StatusCode,
		r.Result.Message,
	)
}

// newResponse creates a new Response for the provided http.Response
func newResponse(r *http.Response, v interface{}) *Response {
	resp := &Response{Response: r}
	resp.populatePageValues(v)
	return resp
}

// Sets paging values if response json was parsed to searchResult type
// (can be extended with other types if they also need paging info)
func (r *Response) populatePageValues(v interface{}) {
	switch value := v.(type) {
	case searchResult:
		value.populatePageValues(r)
	}
}

// IsLast returns true if the current request is the last
func (r *Response) IsLast() bool {
	if r.ResultCount == 0 {
		return true
	}
	return r.ResultCount < r.Parameters.GetHits()
}

// NewRequest creates an API request. A relative URL can be provided in urlStr, which will be resolved to the
// BaseURL of the Client. Relative URLS should always be specified without a preceding slash. If specified, the
// value pointed to by body is JSON encoded and included in as the request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	buf := new(bytes.Buffer)
	if body != nil {
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", mediaType)
	req.Header.Add("Accept", mediaType)
	req.Header.Add("User-Agent", c.UserAgent)
	return req, nil
}

// Do sends an API request and returns the API response. The API response is JSON decoded and stored in the value
// pointed to by v, or returned as an error if an API error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (response *Response, err error) {
	var resp *http.Response
	if resp, err = DoRequestWithClient(c.client, req); err != nil {
		return nil, err
	}
	defer func() {
		if rerr := resp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if err = CheckResponse(resp); err != nil {
		return newResponse(resp, nil), err
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
		if err != nil {
			return nil, err
		}
	}

	response = newResponse(resp, v)
	return
}

// DoRequestWithClient submits an HTTP request using the specified client.
func DoRequestWithClient(client *http.Client, req *http.Request) (*http.Response, error) {
	return client.Do(req)
}

// CheckResponse checks the API response for errors, and returns them if present. A response is considered an
// error if it has a status code outside the 200 range. API error responses are expected to have either no response
// body, or a JSON response body that maps to ErrorResponse. Any other response body will be silently ignored.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && len(data) > 0 {
		if err = json.Unmarshal(data, errorResponse); err != nil {
			return err
		}
	}
	return errorResponse
}

func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)

	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	origURL, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	origValues := origURL.Query()

	newValues, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	for k, v := range newValues {
		origValues[k] = v
	}

	origURL.RawQuery = origValues.Encode()
	return origURL.String(), nil
}

func nextOffset(l, o int) (result int, err error) {
	if l == 0 && o == 0 {
		return 0, fmt.Errorf("limit and offset are empty")
	}
	if l == 0 {
		return 0, fmt.Errorf("limit is empty")
	}
	return o + l, nil
}
