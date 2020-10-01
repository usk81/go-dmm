package dmm

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"
)

const (
	testSeriesRequest = `
{
	"request": {
		"parameters": {
			"Series": null,
			"api_id": "sample",
			"affiliate_id": "affiliate-990",
			"floor_id": "43",
			"hits": "10",
			"offset": "100",
			"output": "json"
		}
	},
	"result": {
		"status": "200",
		"result_count": 10,
		"total_count": "30828",
		"first_position": 100,
		"site_name": "FANZA（アダルト）",
		"site_code": "FANZA",
		"service_name": "動画",
		"service_code": "digital",
		"floor_id": "43",
		"floor_name": "ビデオ",
		"floor_code": "videoa",
		"series": [
			{
				"series_id": "205129",
				"name": "アイドル候補生騙し撮り過激映像",
				"ruby": "あいどるこうほせいだましとりかげきえいぞう"
			},
			{
				"series_id": "441",
				"name": "アイドルザーメン",
				"ruby": "あいどるざーめん"
			},
			{
				"series_id": "208920",
				"name": "IDOL SEX 8時間",
				"ruby": "あいどるせっくす8じかん"
			},
			{
				"series_id": "203306",
				"name": "アイドル潜入捜査官",
				"ruby": "あいどるせんにゅうそうさかん"
			},
			{
				"series_id": "202904",
				"name": "アイドルソープ",
				"ruby": "あいどるそーぷ"
			},
			{
				"series_id": "4039",
				"name": "アイドル宅配便",
				"ruby": "あいどるたっきゅうびん"
			},
			{
				"series_id": "61358",
				"name": "アイドル魂",
				"ruby": "あいどるだましい"
			},
			{
				"series_id": "219168",
				"name": "アイドルになるのを夢見る女子校生",
				"ruby": "あいどるになるのをゆめみるじょしこうせい"
			},
			{
				"series_id": "75005",
				"name": "アイドル寝起き襲撃",
				"ruby": "あいどるねおきしゅうげき"
			},
			{
				"series_id": "78651",
				"name": "アイドル○○の童貞筆おろし",
				"ruby": "あいどるまるまるのどうていふでおろし"
			}
		]
	}
}`

	testSeriesNilRequest = `
{
	"request": {
		"parameters": {
			"Series": null,
			"api_id": "sample",
			"affiliate_id": "affiliate-990",
			"floor_id": "43",
			"hits": "10",
			"offset": "100",
			"output": "json"
		}
	},
	"result": {
		"status": "200",
		"result_count": 0,
		"total_count": "0",
		"first_position": 100,
		"site_name": "FANZA（アダルト）",
		"site_code": "FANZA",
		"service_name": "動画",
		"service_code": "digital",
		"floor_id": "43",
		"floor_name": "ビデオ",
		"floor_code": "videoa"
	}
}`
)

func TestSeries_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(`/`+seriesBasePath, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, testSeriesRequest)
	})

	actual, r, err := client.Series.List(ctx, nil)
	if err != nil {
		t.Errorf("Series.List returned error: %v", err)
	}

	expected := []Series{
		{
			SeriesID:    `205129`,
			Name:        `アイドル候補生騙し撮り過激映像`,
			Ruby:        `あいどるこうほせいだましとりかげきえいぞう`,
			ListURL:     ``,
			SiteName:    `FANZA（アダルト）`,
			SiteCode:    `FANZA`,
			ServiceName: `動画`,
			ServiceCode: `digital`,
			FloorID:     `43`,
			FloorName:   `ビデオ`,
			FloorCode:   `videoa`,
		},
		{
			SeriesID:    `441`,
			Name:        `アイドルザーメン`,
			Ruby:        `あいどるざーめん`,
			ListURL:     ``,
			SiteName:    `FANZA（アダルト）`,
			SiteCode:    `FANZA`,
			ServiceName: `動画`,
			ServiceCode: `digital`,
			FloorID:     `43`,
			FloorName:   `ビデオ`,
			FloorCode:   `videoa`,
		},
		{
			SeriesID:    `208920`,
			Name:        `IDOL SEX 8時間`,
			Ruby:        `あいどるせっくす8じかん`,
			ListURL:     ``,
			SiteName:    `FANZA（アダルト）`,
			SiteCode:    `FANZA`,
			ServiceName: `動画`,
			ServiceCode: `digital`,
			FloorID:     `43`,
			FloorName:   `ビデオ`,
			FloorCode:   `videoa`,
		},
		{
			SeriesID:    `203306`,
			Name:        `アイドル潜入捜査官`,
			Ruby:        `あいどるせんにゅうそうさかん`,
			ListURL:     ``,
			SiteName:    `FANZA（アダルト）`,
			SiteCode:    `FANZA`,
			ServiceName: `動画`,
			ServiceCode: `digital`,
			FloorID:     `43`,
			FloorName:   `ビデオ`,
			FloorCode:   `videoa`,
		},
		{
			SeriesID:    `202904`,
			Name:        `アイドルソープ`,
			Ruby:        `あいどるそーぷ`,
			ListURL:     ``,
			SiteName:    `FANZA（アダルト）`,
			SiteCode:    `FANZA`,
			ServiceName: `動画`,
			ServiceCode: `digital`,
			FloorID:     `43`,
			FloorName:   `ビデオ`,
			FloorCode:   `videoa`,
		},
		{
			SeriesID:    `4039`,
			Name:        `アイドル宅配便`,
			Ruby:        `あいどるたっきゅうびん`,
			ListURL:     ``,
			SiteName:    `FANZA（アダルト）`,
			SiteCode:    `FANZA`,
			ServiceName: `動画`,
			ServiceCode: `digital`,
			FloorID:     `43`,
			FloorName:   `ビデオ`,
			FloorCode:   `videoa`,
		},
		{
			SeriesID:    `61358`,
			Name:        `アイドル魂`,
			Ruby:        `あいどるだましい`,
			ListURL:     ``,
			SiteName:    `FANZA（アダルト）`,
			SiteCode:    `FANZA`,
			ServiceName: `動画`,
			ServiceCode: `digital`,
			FloorID:     `43`,
			FloorName:   `ビデオ`,
			FloorCode:   `videoa`,
		},
		{
			SeriesID:    `219168`,
			Name:        `アイドルになるのを夢見る女子校生`,
			Ruby:        `あいどるになるのをゆめみるじょしこうせい`,
			ListURL:     ``,
			SiteName:    `FANZA（アダルト）`,
			SiteCode:    `FANZA`,
			ServiceName: `動画`,
			ServiceCode: `digital`,
			FloorID:     `43`,
			FloorName:   `ビデオ`,
			FloorCode:   `videoa`,
		},
		{
			SeriesID:    `75005`,
			Name:        `アイドル寝起き襲撃`,
			Ruby:        `あいどるねおきしゅうげき`,
			ListURL:     ``,
			SiteName:    `FANZA（アダルト）`,
			SiteCode:    `FANZA`,
			ServiceName: `動画`,
			ServiceCode: `digital`,
			FloorID:     `43`,
			FloorName:   `ビデオ`,
			FloorCode:   `videoa`,
		},
		{
			SeriesID:    `78651`,
			Name:        `アイドル○○の童貞筆おろし`,
			Ruby:        `あいどるまるまるのどうていふでおろし`,
			ListURL:     ``,
			SiteName:    `FANZA（アダルト）`,
			SiteCode:    `FANZA`,
			ServiceName: `動画`,
			ServiceCode: `digital`,
			FloorID:     `43`,
			FloorName:   `ビデオ`,
			FloorCode:   `videoa`,
		},
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Series.List is not correct; %s", pretty.Compare(actual, expected))
	}

	re := Response{
		Parameters: &SeriesOptions{
			APIID:       `sample`,
			AffiliateID: `affiliate-990`,
			FloorID:     `43`,
			Initial:     ``,
			Hits:        10,
			Offset:      100,
			Output:      `json`,
			Callback:    ``,
		},
		ResultCount:   10,
		TotalCount:    30828,
		FirstPosition: 100,
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

func TestSeries_List_Nil(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(`/`+seriesBasePath, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, testSeriesNilRequest)
	})

	actual, r, err := client.Series.List(ctx, nil)
	if err != nil {
		t.Errorf("Series.List returned error: %v", err)
	}
	if actual != nil {
		t.Errorf("Series.List is not correct; %s", pretty.Compare(actual, nil))
	}

	re := Response{
		Parameters: &SeriesOptions{
			APIID:       `sample`,
			AffiliateID: `affiliate-990`,
			FloorID:     `43`,
			Initial:     ``,
			Hits:        10,
			Offset:      100,
			Output:      `json`,
			Callback:    ``,
		},
		ResultCount:   0,
		TotalCount:    0,
		FirstPosition: 100,
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

func TestSeries_First(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(`/`+seriesBasePath, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, testSeriesRequest)
	})

	actual, r, err := client.Series.First(ctx, nil)
	if err != nil {
		t.Errorf("Series.First returned error: %v", err)
	}

	expected := Series{
		SeriesID:    `205129`,
		Name:        `アイドル候補生騙し撮り過激映像`,
		Ruby:        `あいどるこうほせいだましとりかげきえいぞう`,
		ListURL:     ``,
		SiteName:    `FANZA（アダルト）`,
		SiteCode:    `FANZA`,
		ServiceName: `動画`,
		ServiceCode: `digital`,
		FloorID:     `43`,
		FloorName:   `ビデオ`,
		FloorCode:   `videoa`,
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Series.First returned %+v, expected %+v", actual, expected)
	}

	re := Response{
		Parameters: &SeriesOptions{
			APIID:       `sample`,
			AffiliateID: `affiliate-990`,
			FloorID:     `43`,
			Initial:     ``,
			Hits:        10,
			Offset:      100,
			Output:      `json`,
			Callback:    ``,
		},
		ResultCount:   10,
		TotalCount:    30828,
		FirstPosition: 100,
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

func TestSeriesOptions_Next(t *testing.T) {
	op := &SeriesOptions{
		APIID:       `sample`,
		AffiliateID: `affiliate-990`,
		FloorID:     `43`,
		Initial:     ``,
		Hits:        10,
		Offset:      100,
		Output:      `json`,
		Callback:    ``,
	}
	err := op.Next()
	if err != nil {
		t.Errorf("Next not expected error; %+v", err)
	}
	if op.Offset != 110 {
		t.Errorf("Offset returned %d", op.Offset)
	}
}

func TestSeriesOptions_GetHits(t *testing.T) {
	op := &SeriesOptions{
		APIID:       `sample`,
		AffiliateID: `affiliate-990`,
		FloorID:     `43`,
		Initial:     ``,
		Hits:        10,
		Offset:      100,
		Output:      `json`,
		Callback:    ``,
	}
	hits := op.GetHits()
	if hits != 10 {
		t.Errorf("GetHits returned %d", hits)
	}
}

func TestSeriesOptions_GetOffset(t *testing.T) {
	op := &SeriesOptions{
		APIID:       `sample`,
		AffiliateID: `affiliate-990`,
		FloorID:     `43`,
		Initial:     ``,
		Hits:        10,
		Offset:      100,
		Output:      `json`,
		Callback:    ``,
	}
	offset := op.GetOffset()
	if offset != 100 {
		t.Errorf("GetHits returned %d", offset)
	}
}
