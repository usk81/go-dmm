package dmm

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/usk81/generic/v2"
)

const makerBasePath = `affiliate/v3/MakerSearch`

// MakersService is an interface for interfacing with the Maker
// endpoints of the DMM Affiliate API
// See: https://affiliate.dmm.com/api/v3/makersearch.html
type MakersService interface {
	First(context.Context, *MakerOptions) (Maker, *Response, error)
	List(context.Context, *MakerOptions) ([]Maker, *Response, error)
	Unmarshal(context.Context, *MakerOptions, interface{}) (*Response, error)
}

// MakersServiceOp handles communication with the Maker related methods of
// the DMM Affiliate API.
type MakersServiceOp struct {
	client *Client
}

var _ MakersService = &MakersServiceOp{}

type makerRoot struct {
	Request struct {
		Parameters *internalMakerOptions `json:"parameters"`
	} `json:"request"`
	Result makerResult `json:"result"`
}

type makerResult struct {
	Status        generic.Int     `json:"status"`
	FirstPosition generic.Int     `json:"first_position"`
	ResultCount   generic.Int     `json:"result_count"`
	TotalCount    generic.Int     `json:"total_count"`
	SiteName      string          `json:"site_name"`
	SiteCode      string          `json:"site_code"`
	ServiceName   string          `json:"service_name"`
	ServiceCode   string          `json:"service_code"`
	FloorID       string          `json:"floor_id"`
	FloorName     string          `json:"floor_name"`
	FloorCode     string          `json:"floor_code"`
	Maker         json.RawMessage `json:"maker"`
}

// Maker represents a DMM maker
type Maker struct {
	MakerID     string `json:"maker_id"`
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

// MakerOptions specifies the optional parameters to various List methods
type MakerOptions struct {
	APIID       string `json:"api_id" url:"api_id"`
	AffiliateID string `json:"affiliate_id" url:"affiliate_id"`
	FloorID     string `json:"floor_id" url:"floor_id,omitempty"`
	Initial     string `json:"initial,omitempty" url:"initial,omitempty"`
	Hits        int    `json:"hits" url:"hits,omitempty"`
	Offset      int    `json:"offset" url:"offset,omitempty"`
	Output      string `json:"output" url:"output,omitempty"`
	Callback    string `json:"callback,omitempty" url:"callback,omitempty"`
}

type internalMakerOptions struct {
	APIID       string      `json:"api_id" url:"api_id"`
	AffiliateID string      `json:"affiliate_id" url:"affiliate_id"`
	FloorID     string      `json:"floor_id" url:"floor_id,omitempty"`
	Initial     string      `json:"initial,omitempty" url:"initial,omitempty"`
	Hits        generic.Int `json:"hits" url:"hits,omitempty"`
	Offset      generic.Int `json:"offset" url:"offset,omitempty"`
	Output      string      `json:"output" url:"output,omitempty"`
	Callback    string      `json:"callback,omitempty" url:"callback,omitempty"`
}

func (r *makerRoot) populatePageValues(res *Response) {
	res.FirstPosition = r.Result.FirstPosition.Int()
	res.ResultCount = r.Result.ResultCount.Int()
	res.TotalCount = r.Result.TotalCount.Int()
	res.Parameters = r.Request.Parameters.Convert()
}

func (s *MakersServiceOp) list(ctx context.Context, path string) (makerResult, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return makerResult{}, nil, err
	}

	var root makerRoot
	resp, err := s.client.Do(ctx, req, &root)
	if err != nil {
		return makerResult{}, resp, err
	}
	return root.Result, resp, err
}

// First gets first maker
func (s *MakersServiceOp) First(ctx context.Context, opt *MakerOptions) (Maker, *Response, error) {
	ms, r, err := s.List(ctx, opt)
	if err != nil || len(ms) == 0 {
		return Maker{}, r, err
	}
	return ms[0], r, err
}

// List gets all makers
func (s *MakersServiceOp) List(ctx context.Context, opt *MakerOptions) ([]Maker, *Response, error) {
	path := makerBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	var res makerResult
	var r *Response
	res, r, err = s.list(ctx, path)
	if err != nil {
		return nil, r, err
	}
	var ms []Maker
	if err = json.Unmarshal(res.Maker, &ms); err != nil {
		return nil, r, err
	}
	for i := range ms {
		ms[i].SiteName = res.SiteName
		ms[i].SiteCode = res.SiteCode
		ms[i].ServiceName = res.ServiceName
		ms[i].ServiceCode = res.ServiceCode
		ms[i].FloorID = res.FloorID
		ms[i].FloorName = res.FloorName
		ms[i].FloorCode = res.FloorCode
	}
	return ms, r, err
}

// Unmarshal parses maker API response
func (s *MakersServiceOp) Unmarshal(ctx context.Context, opt *MakerOptions, out interface{}) (*Response, error) {
	path := makerBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, err
	}
	var res makerResult
	var r *Response
	res, r, err = s.list(ctx, path)
	if err != nil {
		return r, err
	}
	if res.Maker == nil {
		return r, nil
	}
	if err = json.Unmarshal(res.Maker, out); err != nil {
		return r, err
	}
	return r, err
}

// Next updates offset
func (o *MakerOptions) Next() (err error) {
	o.Offset, err = nextOffset(o.Hits, o.Offset)
	return
}

// GetHits gets request hits parameter
func (o *MakerOptions) GetHits() int {
	return o.Hits
}

// GetOffset gets request offset parameter
func (o *MakerOptions) GetOffset() int {
	return o.Offset
}

func (i *internalMakerOptions) Convert() *MakerOptions {
	return &MakerOptions{
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
