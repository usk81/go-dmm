package dmm

import (
	"context"
	"encoding/json"
	"net/http"
)

const floorBasePath = `affiliate/v3/FloorList`

// FloorsService is an interface for interfacing with the Floor
// endpoints of the DMM Affiliate API
// See: https://affiliate.dmm.com/api/v3/floorlist.html
type FloorsService interface {
	First(context.Context, *FloorOptions) (Site, *Response, error)
	List(context.Context, *FloorOptions) ([]Site, *Response, error)
	Unmarshal(context.Context, *FloorOptions, interface{}) (*Response, error)
}

// FloorsServiceOp handles communication with the Floor related methods of
// the DMM Affiliate API.
type FloorsServiceOp struct {
	client *Client
}

var _ FloorsService = &FloorsServiceOp{}

type floorRoot struct {
	Request struct {
		Parameters *FloorOptions `json:"parameters"`
	} `json:"request"`
	Result floorResult `json:"result"`
}

type floorResult struct {
	Site json.RawMessage `json:"site"`
}

// Site represents a DMM site
type Site struct {
	Name     string    `json:"name"`
	Code     string    `json:"code"`
	Services []Service `json:"service"`
}

// Service represents a DMM service
type Service struct {
	Name  string  `json:"name"`
	Code  string  `json:"code"`
	Floor []Floor `json:"floor"`
}

// Floor represents a DMM floor
type Floor struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	SiteName    string
	SiteCode    string
	ServiceName string
	ServiceCode string
}

// FloorOptions specifies the optional parameters to various List methods
type FloorOptions struct {
	APIID       string `json:"api_id" url:"api_id"`
	AffiliateID string `json:"affiliate_id" url:"affiliate_id"`
	Output      string `json:"output" url:"output,omitempty"`
	Callback    string `json:"callback" url:"callback,omitempty"`
}

func (r *floorRoot) populatePageValues(res *Response) {
	res.Parameters = r.Request.Parameters
}

func (s *FloorsServiceOp) list(ctx context.Context, path string) (floorResult, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return floorResult{}, nil, err
	}

	var root floorRoot
	resp, err := s.client.Do(ctx, req, &root)
	if err != nil {
		return floorResult{}, resp, err
	}
	return root.Result, resp, err
}

// First gets first site and floors
func (s *FloorsServiceOp) First(ctx context.Context, opt *FloorOptions) (Site, *Response, error) {
	ss, r, err := s.List(ctx, opt)
	if err != nil || len(ss) == 0 {
		return Site{}, r, err
	}
	return ss[0], r, nil
}

// List gets all floors
func (s *FloorsServiceOp) List(ctx context.Context, opt *FloorOptions) ([]Site, *Response, error) {
	var ss []Site
	r, err := s.Unmarshal(ctx, opt, &ss)
	return ss, r, err
}

// Unmarshal parses floor API response
func (s *FloorsServiceOp) Unmarshal(ctx context.Context, opt *FloorOptions, out interface{}) (*Response, error) {
	path := floorBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, err
	}
	var res floorResult
	var r *Response
	res, r, err = s.list(ctx, path)
	if err != nil {
		return r, err
	}
	if err = json.Unmarshal(res.Site, out); err != nil {
		return r, err
	}
	return r, err
}

// Next updates offset
func (o *FloorOptions) Next() error {
	// no action
	return nil
}

// GetHits gets request hits parameter
func (o *FloorOptions) GetHits() int {
	return 0
}

// GetOffset gets request offset parameter
func (o *FloorOptions) GetOffset() int {
	return 0
}
