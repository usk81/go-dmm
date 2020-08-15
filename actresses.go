package dmm

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/usk81/generic"
)

const actressBasePath = `affiliate/v3/ActressSearch`

// ActressesService is an interface for interfacing with the Actress
// endpoints of the DMM Affiliate API
// See: https://affiliate.dmm.com/api/v3/actresssearch.html
type ActressesService interface {
	First(context.Context, *ActressOptions) (Actress, *Response, error)
	List(context.Context, *ActressOptions) ([]Actress, *Response, error)
	Unmarshal(context.Context, *ActressOptions, interface{}) (*Response, error)
}

// ActressesServiceOp handles communication with the Actress related methods of
// the DMM Affiliate API.
type ActressesServiceOp struct {
	client *Client
}

var _ ActressesService = &ActressesServiceOp{}

type actressRoot struct {
	Request struct {
		Parameters *internalActressOptions `json:"parameters"`
	} `json:"request"`
	Result actressResult `json:"result"`
}

type actressResult struct {
	Status        generic.Int     `json:"status"`
	ResultCount   generic.Int     `json:"result_count"`
	TotalCount    generic.Int     `json:"total_count"`
	FirstPosition generic.Int     `json:"first_position"`
	Actress       json.RawMessage `json:"actress"`
}

// Actress represents a actress data
type Actress struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Ruby        string   `json:"ruby"`
	Bust        string   `json:"bust"`
	Cup         string   `json:"cup,omitempty"`
	Waist       string   `json:"waist"`
	Hip         string   `json:"hip"`
	Height      string   `json:"height"`
	Birthday    string   `json:"birthday"`
	BloodType   string   `json:"blood_type"`
	Hobby       string   `json:"hobby"`
	Prefectures string   `json:"prefectures"`
	ImageURL    ImageURL `json:"imageURL,omitempty"`
	ListURL     ListURL  `json:"listURL"`
}

// ListURL is list page URLs are specified actresses
type ListURL struct {
	Digital string `json:"digital"`
	Monthly string `json:"monthly"`
	PPM     string `json:"ppm"`
	Mono    string `json:"mono"`
	Rental  string `json:"rental"`
}

// ActressOptions specifies the optional parameters to various List methods
type ActressOptions struct {
	APIID       string `json:"api_id" url:"api_id"`
	AffiliateID string `json:"affiliate_id" url:"affiliate_id"`
	Initial     string `json:"initial" url:"initial,omitempty"`
	ActressID   string `json:"actress_id" url:"actress_id,omitempty"`
	Keyword     string `json:"keyword" url:"keyword,omitempty"`
	GteBust     int    `json:"gte_bust" url:"gte_bust,omitempty"`
	LteBust     int    `json:"lte_bust" url:"lte_bust,omitempty"`
	GteWaist    int    `json:"gte_waist" url:"gte_waist,omitempty"`
	LteWaist    int    `json:"lte_waist" url:"lte_waist,omitempty"`
	GteHip      int    `json:"gte_hip" url:"gte_hip,omitempty"`
	LteHip      int    `json:"lte_hip" url:"lte_hip,omitempty"`
	GteHeight   int    `json:"gte_height" url:"gte_height,omitempty"`
	LteHeight   int    `json:"lte_height" url:"lte_height,omitempty"`
	GteBirthday string `json:"gte_birthday" url:"gte_birthday,omitempty"`
	LteBirthday string `json:"lte_birthday" url:"lte_birthday,omitempty"`
	Sort        string `json:"sort" url:"sort,omitempty"`
	Hits        int    `json:"hits" url:"hits,omitempty"`
	Offset      int    `json:"offset" url:"offset,omitempty"`
	Output      string `json:"output" url:"output,omitempty"`
	Callback    string `json:"callback" url:"callback,omitempty"`
}

type internalActressOptions struct {
	APIID       string      `json:"api_id"`
	AffiliateID string      `json:"affiliate_id"`
	Initial     string      `json:"initial"`
	ActressID   string      `json:"actress_id"`
	Keyword     string      `json:"keyword"`
	GteBust     generic.Int `json:"gte_bust"`
	LteBust     generic.Int `json:"lte_bust"`
	GteWaist    generic.Int `json:"gte_waist"`
	LteWaist    generic.Int `json:"lte_waist"`
	GteHip      generic.Int `json:"gte_hip"`
	LteHip      generic.Int `json:"lte_hip"`
	GteHeight   generic.Int `json:"gte_height"`
	LteHeight   generic.Int `json:"lte_height"`
	GteBirthday string      `json:"gte_birthday"`
	LteBirthday string      `json:"lte_birthday"`
	Sort        string      `json:"sort"`
	Hits        generic.Int `json:"hits"`
	Offset      generic.Int `json:"offset"`
	Output      string      `json:"output"`
	Callback    string      `json:"callback"`
}

func (r *actressRoot) populatePageValues(res *Response) {
	res.FirstPosition = r.Result.FirstPosition.Int()
	res.ResultCount = r.Result.ResultCount.Int()
	res.TotalCount = r.Result.TotalCount.Int()
	res.Parameters = r.Request.Parameters.Convert()
}

func (s *ActressesServiceOp) list(ctx context.Context, path string) (actressResult, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return actressResult{}, nil, err
	}

	var root actressRoot
	resp, err := s.client.Do(ctx, req, &root)
	if err != nil {
		return actressResult{}, resp, err
	}
	return root.Result, resp, err
}

// First gets first actress
func (s *ActressesServiceOp) First(ctx context.Context, opt *ActressOptions) (Actress, *Response, error) {
	as, r, err := s.List(ctx, opt)
	if err != nil || len(as) == 0 {
		return Actress{}, r, err
	}
	return as[0], r, err
}

// List gets all actresses
func (s *ActressesServiceOp) List(ctx context.Context, opt *ActressOptions) ([]Actress, *Response, error) {
	var as []Actress
	r, err := s.Unmarshal(ctx, opt, &as)
	return as, r, err
}

// Unmarshal parses actress API response
func (s *ActressesServiceOp) Unmarshal(ctx context.Context, opt *ActressOptions, out interface{}) (*Response, error) {
	path := actressBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, err
	}
	var res actressResult
	var r *Response
	res, r, err = s.list(ctx, path)
	if err != nil {
		return r, err
	}
	if err = json.Unmarshal(res.Actress, out); err != nil {
		return r, err
	}
	return r, err
}

// Next updates offset
func (o *ActressOptions) Next() (err error) {
	o.Offset, err = nextOffset(o.Hits, o.Offset)
	return err
}

// GetHits gets request hits parameter
func (o *ActressOptions) GetHits() int {
	return o.Hits
}

// GetOffset gets request offset parameter
func (o *ActressOptions) GetOffset() int {
	return o.Offset
}

func (i *internalActressOptions) Convert() *ActressOptions {
	return &ActressOptions{
		APIID:       i.APIID,
		AffiliateID: i.AffiliateID,
		Initial:     i.Initial,
		ActressID:   i.ActressID,
		Keyword:     i.Keyword,
		GteBust:     i.GteBust.Int(),
		LteBust:     i.LteBust.Int(),
		GteWaist:    i.GteWaist.Int(),
		LteWaist:    i.LteWaist.Int(),
		GteHip:      i.GteHip.Int(),
		LteHip:      i.LteHip.Int(),
		GteHeight:   i.GteHeight.Int(),
		LteHeight:   i.LteHeight.Int(),
		GteBirthday: i.GteBirthday,
		LteBirthday: i.LteBirthday,
		Hits:        i.Hits.Int(),
		Offset:      i.Offset.Int(),
		Output:      i.Output,
		Callback:    i.Callback,
	}
}
