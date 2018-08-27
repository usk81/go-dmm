package dmm

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"
)

const testGenresRequest = `
{
	"request": {
		"parameters": {
			"Genre": null,
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
		"total_count": "284",
		"first_position": 100,
		"site_name": "FANZA（アダルト）",
		"site_code": "FANZA",
		"service_name": "動画",
		"service_code": "digital",
		"floor_id": "43",
		"floor_name": "ビデオ",
		"floor_code": "videoa",
		"genre": [
			{
				"genre_id": "1034",
				"name": "ギャル",
				"ruby": "ぎゃる"
			},
			{
				"genre_id": "6017",
				"name": "ギリモザ",
				"ruby": "ぎりもざ"
			},
			{
				"genre_id": "5069",
				"name": "くすぐり",
				"ruby": "くすぐり"
			},
			{
				"genre_id": "5007",
				"name": "クスコ",
				"ruby": "くすこ"
			},
			{
				"genre_id": "1075",
				"name": "くノ一",
				"ruby": "くのいち"
			},
			{
				"genre_id": "4033",
				"name": "クラシック",
				"ruby": "くらしっく"
			},
			{
				"genre_id": "38",
				"name": "クンニ",
				"ruby": "くんに"
			},
			{
				"genre_id": "4060",
				"name": "ゲイ・ホモ",
				"ruby": "げいほも"
			},
			{
				"genre_id": "6151",
				"name": "ゲロ",
				"ruby": "げろ"
			},
			{
				"genre_id": "4138",
				"name": "原作コラボ",
				"ruby": "げんさくこらぼ"
			}
		]
	}
}`

func TestGenres_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(`/`+genreBasePath, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, testGenresRequest)
	})

	actual, r, err := client.Genres.List(ctx, nil)
	if err != nil {
		t.Errorf("Genres.List returned error: %v", err)
	}

	expected := []Genre{
		{
			GenreID:     `1034`,
			Name:        `ギャル`,
			Ruby:        `ぎゃる`,
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
			GenreID:     `6017`,
			Name:        `ギリモザ`,
			Ruby:        `ぎりもざ`,
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
			GenreID:     `5069`,
			Name:        `くすぐり`,
			Ruby:        `くすぐり`,
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
			GenreID:     `5007`,
			Name:        `クスコ`,
			Ruby:        `くすこ`,
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
			GenreID:     `1075`,
			Name:        `くノ一`,
			Ruby:        `くのいち`,
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
			GenreID:     `4033`,
			Name:        `クラシック`,
			Ruby:        `くらしっく`,
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
			GenreID:     `38`,
			Name:        `クンニ`,
			Ruby:        `くんに`,
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
			GenreID:     `4060`,
			Name:        `ゲイ・ホモ`,
			Ruby:        `げいほも`,
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
			GenreID:     `6151`,
			Name:        `ゲロ`,
			Ruby:        `げろ`,
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
			GenreID:     `4138`,
			Name:        `原作コラボ`,
			Ruby:        `げんさくこらぼ`,
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
		t.Errorf("Genres.List is not correct; %s", pretty.Compare(actual, expected))
	}

	re := Response{
		Parameters: &GenreOptions{
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
		TotalCount:    284,
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

func TestGenres_First(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(`/`+genreBasePath, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, testGenresRequest)
	})

	actual, r, err := client.Genres.First(ctx, nil)
	if err != nil {
		t.Errorf("Genres.First returned error: %v", err)
	}

	expected := Genre{
		GenreID:     `1034`,
		Name:        `ギャル`,
		Ruby:        `ぎゃる`,
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
		t.Errorf("Genres.First returned %+v, expected %+v", actual, expected)
	}

	re := Response{
		Parameters: &GenreOptions{
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
		TotalCount:    284,
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
