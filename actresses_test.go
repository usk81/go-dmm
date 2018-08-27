package dmm

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"
)

const testActressesRequest = `
{
	"request": {
		"parameters": {
			"api_id": "sample",
			"affiliate_id": "affiliate-990",
			"site": "DMM.R18",
			"hits": "10",
			"offset": "10",
			"output": "json"
		}
	},
	"result": {
		"status": "200",
		"result_count": 10,
		"total_count": "48122",
		"first_position": "10",
		"actress": [
			{
				"id": "26617",
				"name": "愛内あみ",
				"ruby": "あいうちあみ",
				"bust": "92",
				"cup": "E",
				"waist": "59",
				"hip": "88",
				"height": "152",
				"birthday": "1987-12-15",
				"blood_type": "B",
				"hobby": "音楽鑑賞",
				"prefectures": "静岡県",
				"imageURL": {
					"small": "http://pics.dmm.co.jp/mono/actjpgs/thumbnail/aiuti_ami.jpg",
					"large": "http://pics.dmm.co.jp/mono/actjpgs/aiuti_ami.jpg"
				},
				"listURL": {
					"digital": "http://www.dmm.co.jp/digital/videoa/-/list/=/article=actress/id=26617/affiliate-990",
					"monthly": "http://www.dmm.co.jp/monthly/premium/-/list/=/article=actress/id=26617/affiliate-990",
					"ppm": "http://www.dmm.co.jp/ppm/video/-/list/=/article=actress/id=26617/affiliate-990",
					"mono": "http://www.dmm.co.jp/mono/dvd/-/list/=/article=actress/id=26617/affiliate-990",
					"rental": "http://www.dmm.co.jp/rental/ppr/-/list/=/article=actress/id=26617/affiliate-990"
				}
			},
			{
				"id": "1038122",
				"name": "愛内陽菜",
				"ruby": "あいうちはるな",
				"bust": null,
				"waist": null,
				"hip": null,
				"height": null,
				"birthday": null,
				"blood_type": null,
				"hobby": null,
				"prefectures": null,
				"listURL": {
					"digital": "http://www.dmm.co.jp/digital/videoa/-/list/=/article=actress/id=1038122/affiliate-990",
					"monthly": "http://www.dmm.co.jp/monthly/premium/-/list/=/article=actress/id=1038122/affiliate-990",
					"ppm": "http://www.dmm.co.jp/ppm/video/-/list/=/article=actress/id=1038122/affiliate-990",
					"mono": "http://www.dmm.co.jp/mono/dvd/-/list/=/article=actress/id=1038122/affiliate-990",
					"rental": "http://www.dmm.co.jp/rental/ppr/-/list/=/article=actress/id=1038122/affiliate-990"
				}
			}
		]
	}
}`

func TestActresses_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(`/`+actressBasePath, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, testActressesRequest)
	})

	actual, r, err := client.Actresses.List(ctx, nil)
	if err != nil {
		t.Errorf("Actresses.List returned error: %v", err)
	}

	expected := []Actress{
		{
			ID:          `26617`,
			Name:        `愛内あみ`,
			Ruby:        `あいうちあみ`,
			Bust:        `92`,
			Cup:         `E`,
			Waist:       `59`,
			Hip:         `88`,
			Height:      `152`,
			Birthday:    `1987-12-15`,
			BloodType:   `B`,
			Hobby:       `音楽鑑賞`,
			Prefectures: `静岡県`,
			ImageURL: ImageURL{
				Small: `http://pics.dmm.co.jp/mono/actjpgs/thumbnail/aiuti_ami.jpg`,
				Large: `http://pics.dmm.co.jp/mono/actjpgs/aiuti_ami.jpg`,
			},
			ListURL: ListURL{
				Digital: `http://www.dmm.co.jp/digital/videoa/-/list/=/article=actress/id=26617/affiliate-990`,
				Mono:    `http://www.dmm.co.jp/mono/dvd/-/list/=/article=actress/id=26617/affiliate-990`,
				Monthly: `http://www.dmm.co.jp/monthly/premium/-/list/=/article=actress/id=26617/affiliate-990`,
				PPM:     `http://www.dmm.co.jp/ppm/video/-/list/=/article=actress/id=26617/affiliate-990`,
				Rental:  `http://www.dmm.co.jp/rental/ppr/-/list/=/article=actress/id=26617/affiliate-990`,
			},
		},
		{
			ID:   `1038122`,
			Name: `愛内陽菜`,
			Ruby: `あいうちはるな`,
			ListURL: ListURL{
				Digital: `http://www.dmm.co.jp/digital/videoa/-/list/=/article=actress/id=1038122/affiliate-990`,
				Mono:    `http://www.dmm.co.jp/mono/dvd/-/list/=/article=actress/id=1038122/affiliate-990`,
				Monthly: `http://www.dmm.co.jp/monthly/premium/-/list/=/article=actress/id=1038122/affiliate-990`,
				PPM:     `http://www.dmm.co.jp/ppm/video/-/list/=/article=actress/id=1038122/affiliate-990`,
				Rental:  `http://www.dmm.co.jp/rental/ppr/-/list/=/article=actress/id=1038122/affiliate-990`,
			},
		},
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Actresses.List returned %+v, expected %+v", actual, expected)
	}

	re := Response{
		Parameters: &ActressOptions{
			APIID:       "sample",
			AffiliateID: "affiliate-990",
			Initial:     ``,
			ActressID:   ``,
			Keyword:     ``,
			GteBust:     0,
			LteBust:     0,
			GteWaist:    0,
			LteWaist:    0,
			GteHip:      0,
			LteHip:      0,
			GteHeight:   0,
			LteHeight:   0,
			GteBirthday: ``,
			LteBirthday: ``,
			Hits:        10,
			Offset:      10,
			Output:      `json`,
			Callback:    ``,
		},
		ResultCount:   10,
		TotalCount:    48122,
		FirstPosition: 10,
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

func TestActresses_First(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(`/`+actressBasePath, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, testActressesRequest)
	})

	actual, r, err := client.Actresses.First(ctx, nil)
	if err != nil {
		t.Errorf("Actresses.First returned error: %v", err)
	}

	expected := Actress{
		ID:          `26617`,
		Name:        `愛内あみ`,
		Ruby:        `あいうちあみ`,
		Bust:        `92`,
		Cup:         `E`,
		Waist:       `59`,
		Hip:         `88`,
		Height:      `152`,
		Birthday:    `1987-12-15`,
		BloodType:   `B`,
		Hobby:       `音楽鑑賞`,
		Prefectures: `静岡県`,
		ImageURL: ImageURL{
			Small: `http://pics.dmm.co.jp/mono/actjpgs/thumbnail/aiuti_ami.jpg`,
			Large: `http://pics.dmm.co.jp/mono/actjpgs/aiuti_ami.jpg`,
		},
		ListURL: ListURL{
			Digital: `http://www.dmm.co.jp/digital/videoa/-/list/=/article=actress/id=26617/affiliate-990`,
			Mono:    `http://www.dmm.co.jp/mono/dvd/-/list/=/article=actress/id=26617/affiliate-990`,
			Monthly: `http://www.dmm.co.jp/monthly/premium/-/list/=/article=actress/id=26617/affiliate-990`,
			PPM:     `http://www.dmm.co.jp/ppm/video/-/list/=/article=actress/id=26617/affiliate-990`,
			Rental:  `http://www.dmm.co.jp/rental/ppr/-/list/=/article=actress/id=26617/affiliate-990`,
		},
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Actresses.First returned %+v, expected %+v", actual, expected)
	}

	re := Response{
		Parameters: &ActressOptions{
			APIID:       "sample",
			AffiliateID: "affiliate-990",
			Initial:     ``,
			ActressID:   ``,
			Keyword:     ``,
			GteBust:     0,
			LteBust:     0,
			GteWaist:    0,
			LteWaist:    0,
			GteHip:      0,
			LteHip:      0,
			GteHeight:   0,
			LteHeight:   0,
			GteBirthday: ``,
			LteBirthday: ``,
			Hits:        10,
			Offset:      10,
			Output:      `json`,
			Callback:    ``,
		},
		ResultCount:   10,
		TotalCount:    48122,
		FirstPosition: 10,
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
