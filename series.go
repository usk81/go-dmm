package dmm

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/usk81/generic/v2"
)

const seriesBasePath = `affiliate/v3/SeriesSearch`

// SeriesService is an interface for interfacing with the Serie
// endpoints of the DMM Affiliate API
// See: https://affiliate.dmm.com/api/v3/seriessearch.html
type SeriesService interface {
	First(context.Context, *SeriesOptions) (Series, *Response, error)
	List(context.Context, *SeriesOptions) ([]Series, *Response, error)
	Unmarshal(context.Context, *SeriesOptions, interface{}) (*Response, error)
}

// SeriesServiceOp handles communication with the Series related methods of
// the DMM Affiliate API.
type SeriesServiceOp struct {
	client *Client
}

var _ SeriesService = &SeriesServiceOp{}

type seriesRoot struct {
	Request struct {
		Parameters *internalSeriesOptions `json:"parameters"`
	} `json:"request"`
	Result seriesResult `json:"result"`
}

type seriesResult struct {
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
	Series        json.RawMessage `json:"series"`
}

// Series represents a DMM series
type Series struct {
	SeriesID    string `json:"series_id"`
	Name        string `json:"name"`
	Ruby        string `json:"ruby"`
	ListURL     string `json:"list_url"`
	SiteName    string
	SiteCode    string
	ServiceName string
	ServiceCode string
	FloorID     string
	FloorName   string
	FloorCode   string
}

// SeriesOptions specifies the optional parameters to various List methods
type SeriesOptions struct {
	APIID       string `json:"api_id" url:"api_id"`
	AffiliateID string `json:"affiliate_id" url:"affiliate_id"`
	FloorID     string `json:"floor_id" url:"floor_id,omitempty"`
	Initial     string `json:"initial,omitempty" url:"initial,omitempty"`
	Hits        int    `json:"hits" url:"hits,omitempty"`
	Offset      int    `json:"offset" url:"offset,omitempty"`
	Output      string `json:"output" url:"output,omitempty"`
	Callback    string `json:"callback,omitempty" url:"callback,omitempty"`
}

type internalSeriesOptions struct {
	APIID       string      `json:"api_id" url:"api_id"`
	AffiliateID string      `json:"affiliate_id" url:"affiliate_id"`
	FloorID     string      `json:"floor_id" url:"floor_id,omitempty"`
	Initial     string      `json:"initial,omitempty" url:"initial,omitempty"`
	Hits        generic.Int `json:"hits" url:"hits,omitempty"`
	Offset      generic.Int `json:"offset" url:"offset,omitempty"`
	Output      string      `json:"output" url:"output,omitempty"`
	Callback    string      `json:"callback,omitempty" url:"callback,omitempty"`
}

func (r *seriesRoot) populatePageValues(res *Response) {
	res.FirstPosition = r.Result.FirstPosition.Int()
	res.ResultCount = r.Result.ResultCount.Int()
	res.TotalCount = r.Result.TotalCount.Int()
	res.Parameters = r.Request.Parameters.Convert()
}

func (s *SeriesServiceOp) list(ctx context.Context, path string) (seriesResult, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return seriesResult{}, nil, err
	}

	var root seriesRoot
	resp, err := s.client.Do(ctx, req, &root)
	if err != nil {
		return seriesResult{}, resp, err
	}
	return root.Result, resp, err
}

// First gets first series
func (s *SeriesServiceOp) First(ctx context.Context, opt *SeriesOptions) (Series, *Response, error) {
	ss, r, err := s.List(ctx, opt)
	if err != nil || len(ss) == 0 {
		return Series{}, r, err
	}
	return ss[0], r, err
}

// List gets all series
func (s *SeriesServiceOp) List(ctx context.Context, opt *SeriesOptions) ([]Series, *Response, error) {
	path := seriesBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	var res seriesResult
	var r *Response
	res, r, err = s.list(ctx, path)
	if err != nil {
		return nil, r, err
	}
	var ss []Series
	if err = json.Unmarshal(res.Series, &ss); err != nil {
		return nil, r, err
	}
	for i := range ss {
		ss[i].SiteName = res.SiteName
		ss[i].SiteCode = res.SiteCode
		ss[i].ServiceName = res.ServiceName
		ss[i].ServiceCode = res.ServiceCode
		ss[i].FloorID = res.FloorID
		ss[i].FloorName = res.FloorName
		ss[i].FloorCode = res.FloorCode
	}
	return ss, r, err
}

// Unmarshal parses series API response
func (s *SeriesServiceOp) Unmarshal(ctx context.Context, opt *SeriesOptions, out interface{}) (*Response, error) {
	path := seriesBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, err
	}
	var res seriesResult
	var r *Response
	res, r, err = s.list(ctx, path)
	if err != nil {
		return r, err
	}
	if res.Series == nil {
		return r, nil
	}
	if err = json.Unmarshal(res.Series, out); err != nil {
		return r, err
	}
	return r, err
}

// Next updates offset
func (o *SeriesOptions) Next() (err error) {
	o.Offset, err = nextOffset(o.Hits, o.Offset)
	return
}

// GetHits gets request hits parameter
func (o *SeriesOptions) GetHits() int {
	return o.Hits
}

// GetOffset gets request offset parameter
func (o *SeriesOptions) GetOffset() int {
	return o.Offset
}

func (i *internalSeriesOptions) Convert() *SeriesOptions {
	return &SeriesOptions{
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
