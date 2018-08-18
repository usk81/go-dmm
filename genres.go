package dmm

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/usk81/generic"
)

const genreBasePath = `affiliate/v3/GenreSearch`

// GenresService is an interface for interfacing with the Genre
// endpoints of the DMM Affiliate API
// See: https://affiliate.dmm.com/api/v3/genresearch.html
type GenresService interface {
	First(context.Context, *GenreOptions) (Genre, *Response, error)
	List(context.Context, *GenreOptions) ([]Genre, *Response, error)
	Unmarshal(context.Context, *GenreOptions, interface{}) (*Response, error)
}

// GenresServiceOp handles communication with the genre related methods of
// the DMM Affiliate API.
type GenresServiceOp struct {
	client *Client
}

var _ GenresService = &GenresServiceOp{}

type genreRoot struct {
	Request struct {
		Parameters *internalGenreOptions `json:"parameters"`
	} `json:"request"`
	Result genreResult `json:"result"`
}

type genreResult struct {
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
	Genre         json.RawMessage `json:"genre"`
}

// Genre represents a DMM genre
type Genre struct {
	GenreID     string `json:"genre_id"`
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

// GenreOptions specifies the optional parameters to various List methods
type GenreOptions struct {
	APIID       string `json:"api_id" url:"api_id"`
	AffiliateID string `json:"affiliate_id" url:"affiliate_id"`
	FloorID     string `json:"floor_id" url:"floor_id"`
	Initial     string `json:"initial,omitempty" url:"initial,omitempty"`
	Hits        int    `json:"hits" url:"hits,omitempty"`
	Offset      int    `json:"offset" url:"offset,omitempty"`
	Output      string `json:"output" url:"output,omitempty"`
	Callback    string `json:"callback,omitempty" url:"callback,omitempty"`
}

type internalGenreOptions struct {
	APIID       string      `json:"api_id" url:"api_id"`
	AffiliateID string      `json:"affiliate_id" url:"affiliate_id"`
	FloorID     string      `json:"floor_id" url:"floor_id"`
	Initial     string      `json:"initial,omitempty" url:"initial,omitempty"`
	Hits        generic.Int `json:"hits" url:"hits,omitempty"`
	Offset      generic.Int `json:"offset" url:"offset,omitempty"`
	Output      string      `json:"output" url:"output,omitempty"`
	Callback    string      `json:"callback,omitempty" url:"callback,omitempty"`
}

func (r *genreRoot) populatePageValues(res *Response) {
	res.FirstPosition = r.Result.FirstPosition.Int()
	res.ResultCount = r.Result.ResultCount.Int()
	res.TotalCount = r.Result.TotalCount.Int()
	res.Parameters = r.Request.Parameters.Convert()
}

func (s *GenresServiceOp) list(ctx context.Context, path string) (genreResult, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return genreResult{}, nil, err
	}

	var root genreRoot
	resp, err := s.client.Do(ctx, req, &root)
	if err != nil {
		return genreResult{}, resp, err
	}
	return root.Result, resp, err
}

// First gets first genre
func (s *GenresServiceOp) First(ctx context.Context, opt *GenreOptions) (Genre, *Response, error) {
	gs, r, err := s.List(ctx, opt)
	if err != nil || len(gs) == 0 {
		return Genre{}, r, err
	}
	return gs[0], r, err
}

// List gets all genres
func (s *GenresServiceOp) List(ctx context.Context, opt *GenreOptions) ([]Genre, *Response, error) {
	path := genreBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}
	var res genreResult
	var r *Response
	res, r, err = s.list(ctx, path)
	if err != nil {
		return nil, r, err
	}
	var gs []Genre
	if err = json.Unmarshal(res.Genre, &gs); err != nil {
		return nil, r, err
	}
	for i := range gs {
		gs[i].SiteName = res.SiteName
		gs[i].SiteCode = res.SiteCode
		gs[i].ServiceName = res.ServiceName
		gs[i].ServiceCode = res.ServiceCode
		gs[i].FloorID = res.FloorID
		gs[i].FloorName = res.FloorName
		gs[i].FloorCode = res.FloorCode
	}
	return gs, r, err
}

// Unmarshal parses genre API response
func (s *GenresServiceOp) Unmarshal(ctx context.Context, opt *GenreOptions, out interface{}) (*Response, error) {
	path := genreBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, err
	}
	var res genreResult
	var r *Response
	res, r, err = s.list(ctx, path)
	if err != nil {
		return r, err
	}
	if err = json.Unmarshal(res.Genre, out); err != nil {
		return r, err
	}
	return r, err
}

// Next updates offset
func (o *GenreOptions) Next() (err error) {
	o.Offset, err = nextOffset(o.Hits, o.Offset)
	return
}

// GetHits gets request hits parameter
func (o *GenreOptions) GetHits() int {
	return o.Hits
}

// GetOffset gets request offset parameter
func (o *GenreOptions) GetOffset() int {
	return o.Offset
}

func (i *internalGenreOptions) Convert() *GenreOptions {
	return &GenreOptions{
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
