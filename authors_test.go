package dmm

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"
)

const testAuthorsRequest = `{
	"request": {
		"parameters": {
			"Author": null,
			"api_id": "sample",
			"affiliate_id": "affiliate-990",
			"floor_id": "27",
			"hits": "10",
			"offset": "5000",
			"output": "json"
		}
	},
	"result": {
		"status": "200",
		"result_count": 10,
		"total_count": "99852",
		"first_position": 5000,
		"site_name": "DMM.com（一般）",
		"site_code": "DMM.com",
		"service_name": "通販",
		"service_code": "mono",
		"floor_id": "27",
		"floor_name": "本・コミック",
		"floor_code": "book",
		"author": [
			{
				"author_id": "217780",
				"name": "安藤美華代",
				"ruby": "あんどうみかよ"
			},
			{
				"author_id": "217781",
				"name": "安東みきえ",
				"ruby": "あんどうみきえ"
			},
			{
				"author_id": "180054",
				"name": "安東実",
				"ruby": "あんどうみのる"
			}
		]
	}
}`

func TestAuthors_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(`/`+AuthorBasePath, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, testAuthorsRequest)
	})

	actual, r, err := client.Authors.List(ctx, nil)
	if err != nil {
		t.Fatalf("Authors.List returned error: %v; response: %#v", err, r)
	}

	expected := []Author{
		Author{
			AuthorID:    "217780",
			Name:        "安藤美華代",
			Ruby:        "あんどうみかよ",
			SiteName:    "DMM.com（一般）",
			SiteCode:    "DMM.com",
			ServiceName: "通販",
			ServiceCode: "mono",
			FloorID:     "27",
			FloorName:   "本・コミック",
			FloorCode:   "book",
		},
		Author{
			AuthorID:    "217781",
			Name:        "安東みきえ",
			Ruby:        "あんどうみきえ",
			SiteName:    "DMM.com（一般）",
			SiteCode:    "DMM.com",
			ServiceName: "通販",
			ServiceCode: "mono",
			FloorID:     "27",
			FloorName:   "本・コミック",
			FloorCode:   "book",
		},
		Author{
			AuthorID:    "180054",
			Name:        "安東実",
			Ruby:        "あんどうみのる",
			SiteName:    "DMM.com（一般）",
			SiteCode:    "DMM.com",
			ServiceName: "通販",
			ServiceCode: "mono",
			FloorID:     "27",
			FloorName:   "本・コミック",
			FloorCode:   "book",
		},
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Authors.List returned %+v, expected %+v", actual, expected)
	}

	re := Response{
		Parameters: &AuthorOptions{
			APIID:       "sample",
			AffiliateID: "affiliate-990",
			FloorID:     `27`,
			Initial:     ``,
			Hits:        10,
			Offset:      5000,
			Output:      `json`,
			Callback:    ``,
		},
		ResultCount:   10,
		TotalCount:    99852,
		FirstPosition: 5000,
	}
	if !reflect.DeepEqual(re.Parameters, r.Parameters) {
		t.Errorf("Response.Parameters is not correct; %s", pretty.Compare(r.Parameters, re.Parameters))
	}
	if r.ResultCount != re.ResultCount {
		t.Errorf("Response.ResultCount returned %+v, expected %+v", r.ResultCount, re.ResultCount)
	}
	if r.TotalCount != re.TotalCount {
		t.Errorf("Response.TotalCount returned %+v, expected %+v", r.TotalCount, re.TotalCount)
	}
	if r.FirstPosition != re.FirstPosition {
		t.Errorf("Response.FirstPosition returned %+v, expected %+v", r.FirstPosition, re.FirstPosition)
	}
}

func TestAuthors_First(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(`/`+AuthorBasePath, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, testAuthorsRequest)
	})

	actual, r, err := client.Authors.First(ctx, nil)
	if err != nil {
		t.Errorf("Authors.First returned error: %v", err)
	}

	expected := Author{
		AuthorID:    "217780",
		Name:        "安藤美華代",
		Ruby:        "あんどうみかよ",
		SiteName:    "DMM.com（一般）",
		SiteCode:    "DMM.com",
		ServiceName: "通販",
		ServiceCode: "mono",
		FloorID:     "27",
		FloorName:   "本・コミック",
		FloorCode:   "book",
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Authors.First returned %+v, expected %+v", actual, expected)
	}

	re := Response{
		Parameters: &AuthorOptions{
			APIID:       "sample",
			AffiliateID: "affiliate-990",
			FloorID:     `27`,
			Initial:     ``,
			Hits:        10,
			Offset:      5000,
			Output:      `json`,
			Callback:    ``,
		},
		ResultCount:   10,
		TotalCount:    99852,
		FirstPosition: 5000,
	}
	if !reflect.DeepEqual(re.Parameters, r.Parameters) {
		t.Errorf("Response.Parameters is not correct; %s", pretty.Compare(r.Parameters, re.Parameters))
	}
	if r.ResultCount != re.ResultCount {
		t.Errorf("Response.ResultCount returned %+v, expected %+v", r.ResultCount, re.ResultCount)
	}
	if r.TotalCount != re.TotalCount {
		t.Errorf("Response.TotalCount returned %+v, expected %+v", r.TotalCount, re.TotalCount)
	}
	if r.FirstPosition != re.FirstPosition {
		t.Errorf("Response.FirstPosition returned %+v, expected %+v", r.FirstPosition, re.FirstPosition)
	}
}
