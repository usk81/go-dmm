package dmm

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"
)

const testMakersRequest = `
{
  "request": {
    "parameters": {
      "Maker": null,
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
    "total_count": "4294",
    "first_position": 100,
    "site_name": "FANZA（アダルト）",
    "site_code": "FANZA",
    "service_name": "動画",
    "service_code": "digital",
    "floor_id": "43",
    "floor_name": "ビデオ",
    "floor_code": "videoa",
    "maker": [
      {
        "maker_id": "45313",
        "name": "アップス",
        "ruby": "あっぷす"
      },
      {
        "maker_id": "46068",
        "name": "アップソン",
        "ruby": "あっぷそん"
      },
      {
        "maker_id": "4616",
        "name": "アップダウン",
        "ruby": "あっぷだうん"
      },
      {
        "maker_id": "1229",
        "name": "アップ ユーピー企画",
        "ruby": "あっぷゆーぴーきかく"
      },
      {
        "maker_id": "4709",
        "name": "アップル企画",
        "ruby": "あっぷるきかく"
      },
      {
        "maker_id": "45664",
        "name": "apple Club",
        "ruby": "あっぷるくらぶ"
      },
      {
        "maker_id": "5090",
        "name": "あっぷるぷる",
        "ruby": "あっぷるぷる"
      },
      {
        "maker_id": "40160",
        "name": "アテナ映像",
        "ruby": "あてなえいぞう"
      },
      {
        "maker_id": "46487",
        "name": "アテナレジェンド",
        "ruby": "あてなれじぇんど"
      },
      {
        "maker_id": "1902",
        "name": "アディクト",
        "ruby": "あでぃくと"
      }
    ]
  }
}`

func TestMakers_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(`/`+makerBasePath, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, testMakersRequest)
	})

	actual, r, err := client.Makers.List(ctx, nil)
	if err != nil {
		t.Errorf("Makers.List returned error: %v", err)
	}

	expected := []Maker{
		{
			MakerID:     `45313`,
			Name:        `アップス`,
			Ruby:        `あっぷす`,
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
			MakerID:     `46068`,
			Name:        `アップソン`,
			Ruby:        `あっぷそん`,
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
			MakerID:     `4616`,
			Name:        `アップダウン`,
			Ruby:        `あっぷだうん`,
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
			MakerID:     `1229`,
			Name:        `アップ ユーピー企画`,
			Ruby:        `あっぷゆーぴーきかく`,
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
			MakerID:     `4709`,
			Name:        `アップル企画`,
			Ruby:        `あっぷるきかく`,
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
			MakerID:     `45664`,
			Name:        `apple Club`,
			Ruby:        `あっぷるくらぶ`,
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
			MakerID:     `5090`,
			Name:        `あっぷるぷる`,
			Ruby:        `あっぷるぷる`,
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
			MakerID:     `40160`,
			Name:        `アテナ映像`,
			Ruby:        `あてなえいぞう`,
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
			MakerID:     `46487`,
			Name:        `アテナレジェンド`,
			Ruby:        `あてなれじぇんど`,
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
			MakerID:     `1902`,
			Name:        `アディクト`,
			Ruby:        `あでぃくと`,
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
		t.Errorf("Makers.List returned %+v, expected %+v", actual, expected)
	}

	re := Response{
		Parameters: &MakerOptions{
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
		TotalCount:    4294,
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

func TestMakers_First(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(`/`+makerBasePath, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, testMakersRequest)
	})

	actual, r, err := client.Makers.First(ctx, nil)
	if err != nil {
		t.Errorf("Makers.First returned error: %v", err)
	}

	expected := Maker{
		MakerID:     `45313`,
		Name:        `アップス`,
		Ruby:        `あっぷす`,
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
		t.Errorf("Makers.First returned %+v, expected %+v", actual, expected)
	}

	re := Response{
		Parameters: &MakerOptions{
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
		TotalCount:    4294,
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

func TestMakerOptions_Next(t *testing.T) {
	op := &MakerOptions{
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

func TestMakerOptions_GetHits(t *testing.T) {
	op := &MakerOptions{
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

func TestMakerOptions_GetOffset(t *testing.T) {
	op := &MakerOptions{
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
