package dmm

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/usk81/generic/v2"
)

const itemBasePath = `affiliate/v3/ItemList`

// ItemsService is an interface for interfacing with the Item
// endpoints of the DMM Affiliate API
// See: https://affiliate.dmm.com/api/v3/itemlist.html
type ItemsService interface {
	First(context.Context, *ItemOptions) (Item, *Response, error)
	List(context.Context, *ItemOptions) ([]Item, *Response, error)
	Unmarshal(context.Context, *ItemOptions, interface{}) (*Response, error)
}

// ItemsServiceOp handles communication with the Item related methods of
// the DMM Affiliate API.
type ItemsServiceOp struct {
	client *Client
}

var _ ItemsService = &ItemsServiceOp{}

type itemRoot struct {
	Request struct {
		Parameters *internalItemOptions `json:"parameters"`
	} `json:"request"`
	Result itemResult `json:"result"`
}

type itemResult struct {
	Status        generic.Int     `json:"status"`
	ResultCount   generic.Int     `json:"result_count"`
	TotalCount    generic.Int     `json:"total_count"`
	FirstPosition generic.Int     `json:"first_position"`
	Items         json.RawMessage `json:"items"`
}

// Item represents a DMM product
type Item struct {
	AffiliateURL       string                     `json:"affiliateURL"`
	AffiliateURLMobile string                     `json:"affiliateURLsp"`
	BandaiInformation  BandaiInformation          `json:"bandaiinfo"`
	CategoryName       string                     `json:"category_name"`
	CdInformation      CdInformation              `json:"cdinfo"`
	Comment            string                     `json:"comment"`
	ContentID          string                     `json:"content_id"`
	Date               string                     `json:"date"`
	FloorCode          string                     `json:"floor_code"`
	FloorName          string                     `json:"floor_name"`
	ImageURL           ImageURL                   `json:"imageURL"`
	ISBN               string                     `json:"isbn,omitempty"`
	ItemInfo           map[string][]ItemComponent `json:"iteminfo"`
	JANCode            string                     `json:"jancode,omitempty"`
	MakerProduct       string                     `json:"maker_product"`
	Prices             Prices                     `json:"prices"`
	ProductID          string                     `json:"product_id"`
	Review             Review                     `json:"review"`
	SampleImageURL     SampleImage                `json:"sampleImageURL,omitempty"`
	SampleMovieURL     SampleMovie                `mapstructure:"sampleMovieURL"`
	ServiceCode        string                     `json:"service_code"`
	ServiceName        string                     `json:"service_name"`
	Stock              string                     `json:"stock"`
	Title              string                     `json:"title"`
	URL                string                     `json:"URL"`
	URLMobile          string                     `json:"URLsp"`
	Volume             string                     `json:"volume"`
}

// ImageURL is image URLs for a product
type ImageURL struct {
	List  string `json:"list"`
	Small string `json:"small"`
	Large string `json:"large"`
}

// ItemInformation is a DMM product information
// type ItemInformation struct {
// 	Actress  []ItemComponent `json:"actress"`
// 	Director []ItemComponent `json:"director"`
// 	Genre    []ItemComponent `json:"genre"`
// 	Label    []ItemComponent `json:"label"`
// 	Maker    []ItemComponent `json:"maker"`
// 	Series   []ItemComponent `json:"series"`
// }

// ItemComponent is a product detail
type ItemComponent struct {
	ID   generic.String `json:"id"`
	Name string         `json:"name"`
}

// Prices is a price information
type Prices struct {
	Price      string     `json:"price"`
	ListPrice  string     `json:"list_price"`
	Deliveries Deliveries `json:"deliveries"`
}

// Deliveries is a collection that streaming prices
type Deliveries struct {
	Delivery []Delivery `json:"delivery"`
}

// Delivery is a price information for streaming
type Delivery struct {
	Type  string `json:"type"`
	Price string `json:"price"`
}

// Review is a review for a product
type Review struct {
	Count   int    `json:"count"`
	Average string `json:"average"`
}

// BandaiInformation is a information of bandai for a product
type BandaiInformation struct {
	TitleCode string `json:"titlecode"`
}

// CdInformation is a information of cd for a product
type CdInformation struct {
	Kind string `json:"kind"`
}

// SampleImage is a collection of sample image URL for a product
type SampleImage struct {
	SampleS SampleImageURLs `json:"sample_s"`
}

// SampleImageURLs is sample image URLs for a product
type SampleImageURLs struct {
	Image []string `json:"image"`
}

// SampleMovie is sample movie URLs for a product
type SampleMovie struct {
	Size476_306 string `json:"size_476_306"`
	Size560_360 string `json:"size_560_360"`
	Size644_414 string `json:"size_644_414"`
	Size720_480 string `json:"size_720_480"`
	PCFlag      int    `json:"pc_flag"`
	SPFlag      int    `json:"sp_flag"`
}

// ItemOptions specifies the optional parameters to various List methods
type ItemOptions struct {
	APIID       string `json:"api_id" url:"api_id"`
	AffiliateID string `json:"affiliate_id" url:"affiliate_id"`
	Site        string `json:"site" url:"site"`
	Service     string `json:"service" url:"service,omitempty"`
	Floor       string `json:"floor" url:"floor,omitempty"`
	Sort        string `json:"sort" url:"sort,omitempty"`
	Keyword     string `json:"keyword" url:"keyword,omitempty"`
	ContentID   string `json:"cid" url:"cid,omitempty"`
	Article     string `json:"article" url:"article,omitempty"`
	ArticleID   string `json:"article_id" url:"article_id,omitempty"`
	GteDate     string `json:"gte_date" url:"gte_date,omitempty"`
	LteDate     string `json:"lte_date" url:"lte_date,omitempty"`
	Stock       string `json:"mono_stock" url:"mono_stock,omitempty"`
	Hits        int    `json:"hits" url:"hits,omitempty"`
	Offset      int    `json:"offset" url:"offset,omitempty"`
	Output      string `json:"output" url:"output,omitempty"`
	Callback    string `json:"callback,omitempty" url:"callback,omitempty"`
}

type internalItemOptions struct {
	APIID       string      `json:"api_id" url:"api_id"`
	AffiliateID string      `json:"affiliate_id" url:"affiliate_id"`
	Site        string      `json:"site" url:"site"`
	Service     string      `json:"service" url:"service,omitempty"`
	Floor       string      `json:"floor" url:"floor,omitempty"`
	Sort        string      `json:"sort" url:"sort,omitempty"`
	Keyword     string      `json:"keyword" url:"keyword,omitempty"`
	ContentID   string      `json:"cid" url:"cid,omitempty"`
	Article     string      `json:"article" url:"article,omitempty"`
	ArticleID   string      `json:"article_id" url:"article_id,omitempty"`
	GteDate     string      `json:"gte_date" url:"gte_date,omitempty"`
	LteDate     string      `json:"lte_date" url:"lte_date,omitempty"`
	Stock       string      `json:"mono_stock" url:"mono_stock,omitempty"`
	Hits        generic.Int `json:"hits" url:"hits,omitempty"`
	Offset      generic.Int `json:"offset" url:"offset,omitempty"`
	Output      string      `json:"output" url:"output,omitempty"`
	Callback    string      `json:"callback,omitempty" url:"callback,omitempty"`
}

func (r *itemRoot) populatePageValues(res *Response) {
	res.FirstPosition = r.Result.FirstPosition.Int()
	res.ResultCount = r.Result.ResultCount.Int()
	res.TotalCount = r.Result.TotalCount.Int()
	res.Parameters = r.Request.Parameters.Convert()
}

func (s *ItemsServiceOp) list(ctx context.Context, path string) (itemResult, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return itemResult{}, nil, err
	}

	var root itemRoot
	resp, err := s.client.Do(ctx, req, &root)
	if err != nil {
		return itemResult{}, resp, err
	}
	return root.Result, resp, err
}

// First gets first item
func (s *ItemsServiceOp) First(ctx context.Context, opt *ItemOptions) (Item, *Response, error) {
	is, r, err := s.List(ctx, opt)
	if err != nil || len(is) == 0 {
		return Item{}, r, err
	}
	return is[0], r, err
}

// List gets all items
func (s *ItemsServiceOp) List(ctx context.Context, opt *ItemOptions) ([]Item, *Response, error) {
	var is []Item
	r, err := s.Unmarshal(ctx, opt, &is)
	return is, r, err
}

// Unmarshal parses item API response
func (s *ItemsServiceOp) Unmarshal(ctx context.Context, opt *ItemOptions, out interface{}) (*Response, error) {
	path := itemBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, err
	}
	var res itemResult
	var r *Response
	res, r, err = s.list(ctx, path)
	if err != nil {
		return r, err
	}
	if res.Items == nil {
		return r, nil
	}
	if err = json.Unmarshal(res.Items, out); err != nil {
		return r, err
	}
	return r, err
}

// Next updates offset
func (o *ItemOptions) Next() (err error) {
	o.Offset, err = nextOffset(o.Hits, o.Offset)
	return
}

// GetHits gets request hits parameter
func (o *ItemOptions) GetHits() int {
	return o.Hits
}

// GetOffset gets request offset parameter
func (o *ItemOptions) GetOffset() int {
	return o.Offset
}

func (i *internalItemOptions) Convert() *ItemOptions {
	return &ItemOptions{
		APIID:       i.APIID,
		AffiliateID: i.AffiliateID,
		Site:        i.Site,
		Service:     i.Service,
		Floor:       i.Floor,
		Sort:        i.Sort,
		Keyword:     i.Keyword,
		ContentID:   i.ContentID,
		Article:     i.Article,
		ArticleID:   i.ArticleID,
		GteDate:     i.GteDate,
		LteDate:     i.LteDate,
		Stock:       i.Stock,
		Hits:        i.Hits.Int(),
		Offset:      i.Offset.Int(),
		Output:      i.Output,
		Callback:    i.Callback,
	}
}
