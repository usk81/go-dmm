package dmm

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/usk81/generic/v2"
)

// AuthorBasePath is used as the base url path to request DMM author API
const AuthorBasePath = `affiliate/v3/AuthorSearch`

// AuthorsService is an interface for interfacing with the Author
// endpoints of the DMM Affiliate API
// See: https://affiliate.dmm.com/api/v3/authorsearch.html
type AuthorsService interface {
	First(context.Context, *AuthorOptions) (Author, *Response, error)
	List(context.Context, *AuthorOptions) ([]Author, *Response, error)
	Unmarshal(context.Context, *AuthorOptions, interface{}) (*Response, error)
}

// AuthorsServiceOp handles communication with the Author related methods of
// the DMM Affiliate API.
type AuthorsServiceOp struct {
	client *Client
}

var _ AuthorsService = &AuthorsServiceOp{}

type authorRoot struct {
	Request struct {
		Parameters *internalAuthorOptions `json:"parameters"`
	} `json:"request"`
	Result authorResult `json:"result"`
}

type authorResult struct {
	Status        generic.Int     `json:"status"`
	ResultCount   generic.Int     `json:"result_count"`
	TotalCount    generic.Int     `json:"total_count"`
	FirstPosition generic.Int     `json:"first_position"`
	SiteName      string          `json:"site_name"`
	SiteCode      string          `json:"site_code"`
	ServiceName   string          `json:"service_name"`
	ServiceCode   string          `json:"service_code"`
	FloorID       string          `json:"floor_id"`
	FloorName     string          `json:"floor_name"`
	FloorCode     string          `json:"floor_code"`
	Author        json.RawMessage `json:"author"`
}

// Author represents a author data
type Author struct {
	AuthorID    string `json:"author_id"`
	Name        string `json:"name"`
	Ruby        string `json:"ruby"`
	SiteName    string
	SiteCode    string
	ServiceName string
	ServiceCode string
	FloorID     string
	FloorName   string
	FloorCode   string
}

// AuthorOptions specifies the optional parameters to various List methods
type AuthorOptions struct {
	APIID       string `json:"api_id" url:"api_id"`
	AffiliateID string `json:"affiliate_id" url:"affiliate_id"`
	FloorID     string `json:"floor_id" url:"floor_id,omitempty"`
	Initial     string `json:"initial,omitempty" url:"initial,omitempty"`
	Hits        int    `json:"hits" url:"hits,omitempty"`
	Offset      int    `json:"offset" url:"offset,omitempty"`
	Output      string `json:"output" url:"output,omitempty"`
	Callback    string `json:"callback,omitempty" url:"callback,omitempty"`
}

type internalAuthorOptions struct {
	APIID       string      `json:"api_id" url:"api_id"`
	AffiliateID string      `json:"affiliate_id" url:"affiliate_id"`
	FloorID     string      `json:"floor_id" url:"floor_id,omitempty"`
	Initial     string      `json:"initial,omitempty" url:"initial,omitempty"`
	Hits        generic.Int `json:"hits" url:"hits,omitempty"`
	Offset      generic.Int `json:"offset" url:"offset,omitempty"`
	Output      string      `json:"output" url:"output,omitempty"`
	Callback    string      `json:"callback,omitempty" url:"callback,omitempty"`
}

func (r *authorRoot) populatePageValues(res *Response) {
	res.FirstPosition = r.Result.FirstPosition.Int()
	res.ResultCount = r.Result.ResultCount.Int()
	res.TotalCount = r.Result.TotalCount.Int()
	res.Parameters = r.Request.Parameters.Convert()
}

func (s *AuthorsServiceOp) list(ctx context.Context, path string) (authorResult, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return authorResult{}, nil, err
	}

	var root authorRoot
	resp, err := s.client.Do(ctx, req, &root)
	if err != nil {
		return authorResult{}, resp, err
	}
	return root.Result, resp, err
}

// First gets first author
func (s *AuthorsServiceOp) First(ctx context.Context, opt *AuthorOptions) (Author, *Response, error) {
	as, r, err := s.List(ctx, opt)
	if err != nil || len(as) == 0 {
		return Author{}, r, err
	}
	return as[0], r, err
}

// List gets all authors
func (s *AuthorsServiceOp) List(ctx context.Context, opt *AuthorOptions) ([]Author, *Response, error) {
	path := AuthorBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	var res authorResult
	var r *Response
	res, r, err = s.list(ctx, path)
	if err != nil {
		return nil, r, err
	}
	var as []Author
	if err = json.Unmarshal(res.Author, &as); err != nil {
		return nil, r, err
	}
	for i := range as {
		as[i].SiteName = res.SiteName
		as[i].SiteCode = res.SiteCode
		as[i].ServiceName = res.ServiceName
		as[i].ServiceCode = res.ServiceCode
		as[i].FloorID = res.FloorID
		as[i].FloorName = res.FloorName
		as[i].FloorCode = res.FloorCode
	}
	return as, r, err
}

// Unmarshal parses author API response
func (s *AuthorsServiceOp) Unmarshal(ctx context.Context, opt *AuthorOptions, out interface{}) (*Response, error) {
	path := AuthorBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, err
	}
	var res authorResult
	var r *Response
	res, r, err = s.list(ctx, path)
	if err != nil {
		return r, err
	}
	if res.Author == nil {
		return r, nil
	}
	if err = json.Unmarshal(res.Author, out); err != nil {
		return r, err
	}
	return r, err
}

// Next updates offset
func (o *AuthorOptions) Next() (err error) {
	o.Offset, err = nextOffset(o.Hits, o.Offset)
	return
}

// GetHits gets request hits parameter
func (o *AuthorOptions) GetHits() int {
	return o.Hits
}

// GetOffset gets request offset parameter
func (o *AuthorOptions) GetOffset() int {
	return o.Offset
}

func (i *internalAuthorOptions) Convert() *AuthorOptions {
	return &AuthorOptions{
		APIID:       i.APIID,
		AffiliateID: i.AffiliateID,
		FloorID:     i.FloorID,
		Initial:     i.Initial,
		Hits:        i.Hits.Int(),
		Offset:      i.Offset.Int(),
		Output:      i.Output,
		Callback:    i.Callback,
	}
}
