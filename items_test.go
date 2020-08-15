package dmm

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"github.com/usk81/generic/v2"
)

const testItemsRequest = `
{
  "request": {
    "parameters": {
      "api_id": "sample",
      "affiliate_id": "affiliate-990",
      "site": "DMM.R18",
      "service": "mono",
      "floor": "dvd",
      "hits": "2",
      "offset": "1000"
    }
  },
  "result": {
    "status": 200,
    "result_count": 2,
    "total_count": 50000,
    "first_position": 1000,
    "items": [
      {
        "service_code": "mono",
        "service_name": "通販",
        "floor_code": "dvd",
        "floor_name": "DVD",
        "category_name": "DVD通販",
        "content_id": "juy553",
        "product_id": "juy553",
        "title": "元タレント人妻 マドンナ専属 第2弾！！ 絶倫だとウワサの、娘の彼氏が泊りに来て…。 壇えみ",
        "volume": "120",
        "URL": "http://www.dmm.co.jp/mono/dvd/-/detail/=/cid=juy553/",
        "URLsp": "http://sp.dmm.co.jp/mono/detail/index/shop/dvd/cid/juy553/",
        "affiliateURL": "http://www.dmm.co.jp/mono/dvd/-/detail/=/cid=juy553/affiliate-990",
        "affiliateURLsp": "http://sp.dmm.co.jp/mono/detail/index/shop/dvd/cid/juy553/affiliate-990",
        "imageURL": {
          "list": "https://pics.dmm.co.jp/mono/movie/adult/juy553/juy553pt.jpg",
          "small": "https://pics.dmm.co.jp/mono/movie/adult/juy553/juy553ps.jpg",
          "large": "https://pics.dmm.co.jp/mono/movie/adult/juy553/juy553pl.jpg"
        },
        "sampleImageURL": {
          "sample_s": {
            "image": [
              "https://pics.dmm.co.jp/digital/video/juy00553/juy00553-1.jpg",
              "https://pics.dmm.co.jp/digital/video/juy00553/juy00553-2.jpg",
              "https://pics.dmm.co.jp/digital/video/juy00553/juy00553-3.jpg",
              "https://pics.dmm.co.jp/digital/video/juy00553/juy00553-4.jpg",
              "https://pics.dmm.co.jp/digital/video/juy00553/juy00553-5.jpg",
              "https://pics.dmm.co.jp/digital/video/juy00553/juy00553-6.jpg",
              "https://pics.dmm.co.jp/digital/video/juy00553/juy00553-7.jpg",
              "https://pics.dmm.co.jp/digital/video/juy00553/juy00553-8.jpg",
              "https://pics.dmm.co.jp/digital/video/juy00553/juy00553-9.jpg",
              "https://pics.dmm.co.jp/digital/video/juy00553/juy00553-10.jpg"
            ]
          }
        },
        "prices": {
          "price": "2381",
          "list_price": "3218"
        },
        "date": "2018-07-25 10:00:00",
        "iteminfo": {
          "genre": [
            {
              "id": 6102,
              "name": "サンプル動画"
            },
            {
              "id": 6004,
              "name": "デジモ"
            },
            {
              "id": 4025,
              "name": "単体作品"
            },
            {
              "id": 1032,
              "name": "お母さん"
            },
            {
              "id": 2001,
              "name": "巨乳"
            },
            {
              "id": 1039,
              "name": "人妻"
            },
            {
              "id": 4111,
              "name": "寝取り・寝取られ"
            },
            {
              "id": 1014,
              "name": "熟女"
            }
          ],
          "maker": [
            {
              "id": 2661,
              "name": "マドンナ"
            }
          ],
          "actress": [
            {
              "id": 1046150,
              "name": "壇えみ"
            },
            {
              "id": "1046150_ruby",
              "name": "だんえみ"
            },
            {
              "id": "1046150_classify",
              "name": "av"
            }
          ],
          "director": [
            {
              "id": 107546,
              "name": "豆沢豆太郎"
            },
            {
              "id": "107546_ruby",
              "name": "まめざわまめたろう"
            }
          ],
          "label": [
            {
              "id": 2931,
              "name": "Madonna"
            }
          ]
        },
        "jancode": "4549831291424",
        "maker_product": "JUY-553",
        "stock": "reserve"
      },
      {
        "service_code": "digital",
        "service_name": "動画",
        "floor_code": "videoa",
        "floor_name": "ビデオ",
        "category_name": "ビデオ (動画)",
        "content_id": "1hawa00124",
        "product_id": "1hawa00124",
        "title": "夫に内緒で他人棒SEX 30歳すぎて初めての精飲特別編 想定外の中出し懇願 10，000人に1人の全身性感帯保育士妻 みひなさん30歳",
        "volume": "244",
        "review": {
          "count": 5,
          "average": "3.60"
        },
        "URL": "http://www.dmm.co.jp/digital/videoa/-/detail/=/cid=1hawa00124/",
        "URLsp": "http://www.dmm.co.jp/digital/videoa/-/detail/=/cid=1hawa00124/",
        "affiliateURL": "http://www.dmm.co.jp/digital/videoa/-/detail/=/cid=1hawa00124/affiliate-990",
        "affiliateURLsp": "http://www.dmm.co.jp/digital/videoa/-/detail/=/cid=1hawa00124/affiliate-990",
        "imageURL": {
          "list": "https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124pt.jpg",
          "small": "https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124ps.jpg",
          "large": "https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124pl.jpg"
        },
        "sampleImageURL": {
          "sample_s": {
            "image": [
              "https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-1.jpg",
              "https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-2.jpg",
              "https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-3.jpg",
              "https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-4.jpg",
              "https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-5.jpg",
              "https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-6.jpg",
              "https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-7.jpg",
              "https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-8.jpg",
              "https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-9.jpg",
              "https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-10.jpg",
              "https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-11.jpg",
              "https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-12.jpg",
              "https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-13.jpg",
              "https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-14.jpg",
              "https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-15.jpg",
              "https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-16.jpg",
              "https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-17.jpg",
              "https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-18.jpg",
              "https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-19.jpg",
              "https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-20.jpg"
            ]
          }
        },
        "sampleMovieURL": {
          "size_476_306": "http://www.dmm.co.jp/litevideo/-/part/=/cid=1hawa00124/size=476_306/",
          "size_560_360": "http://www.dmm.co.jp/litevideo/-/part/=/cid=1hawa00124/size=560_360/",
          "size_644_414": "http://www.dmm.co.jp/litevideo/-/part/=/cid=1hawa00124/size=644_414/",
          "size_720_480": "http://www.dmm.co.jp/litevideo/-/part/=/cid=1hawa00124/size=720_480/",
          "pc_flag": 1,
          "sp_flag": 1
        },
        "prices": {
          "price": "500~",
          "deliveries": {
            "delivery": [
              {
                "type": "stream",
                "price": "500"
              },
              {
                "type": "download",
                "price": "980"
              },
              {
                "type": "hd",
                "price": "1480"
              },
              {
                "type": "androiddl",
                "price": "980"
              },
              {
                "type": "iosdl",
                "price": "980"
              }
            ]
          }
        },
        "date": "2017-12-07 10:00:54",
        "iteminfo": {
          "genre": [
            {
              "id": 4007,
              "name": "企画"
            },
            {
              "id": 1039,
              "name": "人妻"
            },
            {
              "id": 4024,
              "name": "素人"
            },
            {
              "id": 5009,
              "name": "ごっくん"
            },
            {
              "id": 5001,
              "name": "中出し"
            },
            {
              "id": 5068,
              "name": "イラマチオ"
            },
            {
              "id": 6012,
              "name": "4時間以上作品"
            },
            {
              "id": 6533,
              "name": "ハイビジョン"
            }
          ],
          "series": [
            {
              "id": 211038,
              "name": "夫に内緒で他人棒SEX"
            }
          ],
          "maker": [
            {
              "id": 46115,
              "name": "コスモス映像"
            }
          ],
          "director": [
            {
              "id": 102777,
              "name": "長瀬ハワイ"
            },
            {
              "id": "102777_ruby",
              "name": "ながせはわい"
            }
          ],
          "label": [
            {
              "id": 24342,
              "name": "コスモス映像（ソフトオンデマンド）"
            }
          ]
        }
      }
    ]
  }
}`

func TestItems_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(`/`+itemBasePath, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, testItemsRequest)
	})

	actual, r, err := client.Items.List(ctx, nil)
	if err != nil {
		t.Errorf("Items.List returned error: %v", err)
	}

	expected := []Item{
		{
			AffiliateURL:       `http://www.dmm.co.jp/mono/dvd/-/detail/=/cid=juy553/affiliate-990`,
			AffiliateURLMobile: `http://sp.dmm.co.jp/mono/detail/index/shop/dvd/cid/juy553/affiliate-990`,
			CategoryName:       `DVD通販`,
			ContentID:          `juy553`,
			Date:               `2018-07-25 10:00:00`,
			FloorCode:          `dvd`,
			FloorName:          `DVD`,
			ImageURL: ImageURL{
				List:  `https://pics.dmm.co.jp/mono/movie/adult/juy553/juy553pt.jpg`,
				Large: `https://pics.dmm.co.jp/mono/movie/adult/juy553/juy553pl.jpg`,
				Small: `https://pics.dmm.co.jp/mono/movie/adult/juy553/juy553ps.jpg`,
			},
			ItemInfo: map[string][]ItemComponent{
				"actress": {
					{
						ID:   generic.MustString(1046150),
						Name: `壇えみ`,
					},
					{
						ID:   generic.MustString("1046150_ruby"),
						Name: `だんえみ`,
					},
					{
						ID:   generic.MustString("1046150_classify"),
						Name: `av`,
					},
				},
				"genre": {
					{
						ID:   generic.MustString("6102"),
						Name: `サンプル動画`,
					},
					{
						ID:   generic.MustString("6004"),
						Name: `デジモ`,
					},
					{
						ID:   generic.MustString("4025"),
						Name: `単体作品`,
					},
					{
						ID:   generic.MustString("1032"),
						Name: `お母さん`,
					},
					{
						ID:   generic.MustString("2001"),
						Name: `巨乳`,
					},
					{
						ID:   generic.MustString("1039"),
						Name: `人妻`,
					},
					{
						ID:   generic.MustString("4111"),
						Name: `寝取り・寝取られ`,
					},
					{
						ID:   generic.MustString("1014"),
						Name: `熟女`,
					},
				},
				"director": {
					{
						ID:   generic.MustString("107546"),
						Name: `豆沢豆太郎`,
					},
					{
						ID:   generic.MustString("107546_ruby"),
						Name: `まめざわまめたろう`,
					},
				},
				"label": {
					{
						ID:   generic.MustString("2931"),
						Name: `Madonna`,
					},
				},
				"maker": {
					{
						ID:   generic.MustString("2661"),
						Name: `マドンナ`,
					},
				},
			},
			JANCode:      `4549831291424`,
			MakerProduct: `JUY-553`,
			Prices: Prices{
				Price:     `2381`,
				ListPrice: `3218`,
			},
			ProductID: `juy553`,
			SampleImageURL: SampleImage{
				SampleS: SampleImageURLs{
					Image: []string{
						`https://pics.dmm.co.jp/digital/video/juy00553/juy00553-1.jpg`,
						`https://pics.dmm.co.jp/digital/video/juy00553/juy00553-2.jpg`,
						`https://pics.dmm.co.jp/digital/video/juy00553/juy00553-3.jpg`,
						`https://pics.dmm.co.jp/digital/video/juy00553/juy00553-4.jpg`,
						`https://pics.dmm.co.jp/digital/video/juy00553/juy00553-5.jpg`,
						`https://pics.dmm.co.jp/digital/video/juy00553/juy00553-6.jpg`,
						`https://pics.dmm.co.jp/digital/video/juy00553/juy00553-7.jpg`,
						`https://pics.dmm.co.jp/digital/video/juy00553/juy00553-8.jpg`,
						`https://pics.dmm.co.jp/digital/video/juy00553/juy00553-9.jpg`,
						`https://pics.dmm.co.jp/digital/video/juy00553/juy00553-10.jpg`,
					},
				},
			},
			ServiceCode: `mono`,
			ServiceName: `通販`,
			Stock:       `reserve`,
			Title:       `元タレント人妻 マドンナ専属 第2弾！！ 絶倫だとウワサの、娘の彼氏が泊りに来て…。 壇えみ`,
			URL:         `http://www.dmm.co.jp/mono/dvd/-/detail/=/cid=juy553/`,
			URLMobile:   `http://sp.dmm.co.jp/mono/detail/index/shop/dvd/cid/juy553/`,
			Volume:      `120`,
		},
		{
			AffiliateURL:       `http://www.dmm.co.jp/digital/videoa/-/detail/=/cid=1hawa00124/affiliate-990`,
			AffiliateURLMobile: `http://www.dmm.co.jp/digital/videoa/-/detail/=/cid=1hawa00124/affiliate-990`,
			CategoryName:       `ビデオ (動画)`,
			Comment:            ``,
			ContentID:          `1hawa00124`,
			Date:               `2017-12-07 10:00:54`,
			FloorCode:          `videoa`,
			FloorName:          `ビデオ`,
			ImageURL: ImageURL{
				List:  `https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124pt.jpg`,
				Small: `https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124ps.jpg`,
				Large: `https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124pl.jpg`,
			},
			ISBN: ``,
			ItemInfo: map[string][]ItemComponent{
				"director": {
					{
						ID:   generic.MustString(102777),
						Name: `長瀬ハワイ`,
					},
					{
						ID:   generic.MustString(`102777_ruby`),
						Name: `ながせはわい`,
					},
				},
				"genre": {
					{
						ID:   generic.MustString(4007),
						Name: `企画`,
					},
					{
						ID:   generic.MustString(1039),
						Name: `人妻`,
					},
					{
						ID:   generic.MustString(4024),
						Name: `素人`,
					},
					{
						ID:   generic.MustString(5009),
						Name: `ごっくん`,
					},
					{
						ID:   generic.MustString(5001),
						Name: `中出し`,
					},
					{
						ID:   generic.MustString(5068),
						Name: `イラマチオ`,
					},
					{
						ID:   generic.MustString(6012),
						Name: `4時間以上作品`,
					},
					{
						ID:   generic.MustString(6533),
						Name: `ハイビジョン`,
					},
				},
				"label": {
					{
						ID:   generic.MustString(24342),
						Name: `コスモス映像（ソフトオンデマンド）`,
					},
				},
				"maker": {
					{
						ID:   generic.MustString(46115),
						Name: `コスモス映像`,
					},
				},
				"series": {
					{
						ID:   generic.MustString(211038),
						Name: `夫に内緒で他人棒SEX`,
					},
				},
			},
			Prices: Prices{
				Price: `500~`,
				Deliveries: Deliveries{
					Delivery: []Delivery{
						{
							Type:  `stream`,
							Price: `500`,
						},
						{
							Type:  `download`,
							Price: `980`,
						},
						{
							Type:  `hd`,
							Price: `1480`,
						},
						{
							Type:  `androiddl`,
							Price: `980`,
						},
						{
							Type:  `iosdl`,
							Price: `980`,
						},
					},
				},
			},
			ProductID: `1hawa00124`,
			Review: Review{
				Count:   5,
				Average: "3.60",
			},
			SampleImageURL: SampleImage{
				SampleS: SampleImageURLs{
					Image: []string{
						`https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-1.jpg`,
						`https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-2.jpg`,
						`https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-3.jpg`,
						`https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-4.jpg`,
						`https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-5.jpg`,
						`https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-6.jpg`,
						`https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-7.jpg`,
						`https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-8.jpg`,
						`https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-9.jpg`,
						`https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-10.jpg`,
						`https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-11.jpg`,
						`https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-12.jpg`,
						`https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-13.jpg`,
						`https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-14.jpg`,
						`https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-15.jpg`,
						`https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-16.jpg`,
						`https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-17.jpg`,
						`https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-18.jpg`,
						`https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-19.jpg`,
						`https://pics.dmm.co.jp/digital/video/1hawa00124/1hawa00124-20.jpg`,
					},
				},
			},
			SampleMovieURL: SampleMovie{
				Size476_306: `http://www.dmm.co.jp/litevideo/-/part/=/cid=1hawa00124/size=476_306/`,
				Size560_360: `http://www.dmm.co.jp/litevideo/-/part/=/cid=1hawa00124/size=560_360/`,
				Size644_414: `http://www.dmm.co.jp/litevideo/-/part/=/cid=1hawa00124/size=644_414/`,
				Size720_480: `http://www.dmm.co.jp/litevideo/-/part/=/cid=1hawa00124/size=720_480/`,
				PCFlag:      1,
				SPFlag:      1,
			},
			ServiceCode: `digital`,
			ServiceName: `動画`,
			Title:       `夫に内緒で他人棒SEX 30歳すぎて初めての精飲特別編 想定外の中出し懇願 10，000人に1人の全身性感帯保育士妻 みひなさん30歳`,
			URL:         `http://www.dmm.co.jp/digital/videoa/-/detail/=/cid=1hawa00124/`,
			URLMobile:   `http://www.dmm.co.jp/digital/videoa/-/detail/=/cid=1hawa00124/`,
			Volume:      `244`,
		},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Items.List is not correct; %s", pretty.Compare(actual, expected))
	}

	re := Response{
		Parameters: &ItemOptions{
			APIID:       `sample`,
			AffiliateID: `affiliate-990`,
			Site:        `DMM.R18`,
			Floor:       `dvd`,
			Service:     `mono`,
			Hits:        2,
			Offset:      1000,
			Callback:    ``,
		},
		ResultCount:   2,
		TotalCount:    50000,
		FirstPosition: 1000,
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

func TestItems_First(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(`/`+itemBasePath, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, testItemsRequest)
	})

	actual, r, err := client.Items.First(ctx, nil)
	if err != nil {
		t.Errorf("Items.List returned error: %v", err)
	}

	expected := Item{
		AffiliateURL:       `http://www.dmm.co.jp/mono/dvd/-/detail/=/cid=juy553/affiliate-990`,
		AffiliateURLMobile: `http://sp.dmm.co.jp/mono/detail/index/shop/dvd/cid/juy553/affiliate-990`,
		CategoryName:       `DVD通販`,
		ContentID:          `juy553`,
		Date:               `2018-07-25 10:00:00`,
		FloorCode:          `dvd`,
		FloorName:          `DVD`,
		ImageURL: ImageURL{
			List:  `https://pics.dmm.co.jp/mono/movie/adult/juy553/juy553pt.jpg`,
			Large: `https://pics.dmm.co.jp/mono/movie/adult/juy553/juy553pl.jpg`,
			Small: `https://pics.dmm.co.jp/mono/movie/adult/juy553/juy553ps.jpg`,
		},
		ItemInfo: map[string][]ItemComponent{
			"actress": {
				{
					ID:   generic.MustString(1046150),
					Name: `壇えみ`,
				},
				{
					ID:   generic.MustString("1046150_ruby"),
					Name: `だんえみ`,
				},
				{
					ID:   generic.MustString("1046150_classify"),
					Name: `av`,
				},
			},
			"genre": {
				{
					ID:   generic.MustString("6102"),
					Name: `サンプル動画`,
				},
				{
					ID:   generic.MustString("6004"),
					Name: `デジモ`,
				},
				{
					ID:   generic.MustString("4025"),
					Name: `単体作品`,
				},
				{
					ID:   generic.MustString("1032"),
					Name: `お母さん`,
				},
				{
					ID:   generic.MustString("2001"),
					Name: `巨乳`,
				},
				{
					ID:   generic.MustString("1039"),
					Name: `人妻`,
				},
				{
					ID:   generic.MustString("4111"),
					Name: `寝取り・寝取られ`,
				},
				{
					ID:   generic.MustString("1014"),
					Name: `熟女`,
				},
			},
			"director": {
				{
					ID:   generic.MustString("107546"),
					Name: `豆沢豆太郎`,
				},
				{
					ID:   generic.MustString("107546_ruby"),
					Name: `まめざわまめたろう`,
				},
			},
			"label": {
				{
					ID:   generic.MustString("2931"),
					Name: `Madonna`,
				},
			},
			"maker": {
				{
					ID:   generic.MustString("2661"),
					Name: `マドンナ`,
				},
			},
		},
		JANCode:      `4549831291424`,
		MakerProduct: `JUY-553`,
		Prices: Prices{
			Price:     `2381`,
			ListPrice: `3218`,
		},
		ProductID: `juy553`,
		SampleImageURL: SampleImage{
			SampleS: SampleImageURLs{
				Image: []string{
					`https://pics.dmm.co.jp/digital/video/juy00553/juy00553-1.jpg`,
					`https://pics.dmm.co.jp/digital/video/juy00553/juy00553-2.jpg`,
					`https://pics.dmm.co.jp/digital/video/juy00553/juy00553-3.jpg`,
					`https://pics.dmm.co.jp/digital/video/juy00553/juy00553-4.jpg`,
					`https://pics.dmm.co.jp/digital/video/juy00553/juy00553-5.jpg`,
					`https://pics.dmm.co.jp/digital/video/juy00553/juy00553-6.jpg`,
					`https://pics.dmm.co.jp/digital/video/juy00553/juy00553-7.jpg`,
					`https://pics.dmm.co.jp/digital/video/juy00553/juy00553-8.jpg`,
					`https://pics.dmm.co.jp/digital/video/juy00553/juy00553-9.jpg`,
					`https://pics.dmm.co.jp/digital/video/juy00553/juy00553-10.jpg`,
				},
			},
		},
		ServiceCode: `mono`,
		ServiceName: `通販`,
		Stock:       `reserve`,
		Title:       `元タレント人妻 マドンナ専属 第2弾！！ 絶倫だとウワサの、娘の彼氏が泊りに来て…。 壇えみ`,
		URL:         `http://www.dmm.co.jp/mono/dvd/-/detail/=/cid=juy553/`,
		URLMobile:   `http://sp.dmm.co.jp/mono/detail/index/shop/dvd/cid/juy553/`,
		Volume:      `120`,
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Items.First returned %+v, expected %+v", actual, expected)
	}

	re := Response{
		Parameters: &ItemOptions{
			APIID:       `sample`,
			AffiliateID: `affiliate-990`,
			Site:        `DMM.R18`,
			Floor:       `dvd`,
			Service:     `mono`,
			Hits:        2,
			Offset:      1000,
			Callback:    ``,
		},
		ResultCount:   2,
		TotalCount:    50000,
		FirstPosition: 1000,
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

func TestItemOptions_Next(t *testing.T) {
	op := &ItemOptions{
		APIID:       `sample`,
		AffiliateID: `affiliate-990`,
		Site:        `DMM.R18`,
		Floor:       `dvd`,
		Service:     `mono`,
		Hits:        2,
		Offset:      1000,
		Callback:    ``,
	}
	err := op.Next()
	if err != nil {
		t.Errorf("Next not expected error; %+v", err)
	}
	if op.Offset != 1002 {
		t.Errorf("Offset returned %d", op.Offset)
	}
}

func TestItemOptions_GetHits(t *testing.T) {
	op := &ItemOptions{
		APIID:       `sample`,
		AffiliateID: `affiliate-990`,
		Site:        `DMM.R18`,
		Floor:       `dvd`,
		Service:     `mono`,
		Hits:        2,
		Offset:      1000,
		Callback:    ``,
	}
	hits := op.GetHits()
	if hits != 2 {
		t.Errorf("GetHits returned %d", hits)
	}
}

func TestItemOptions_GetOffset(t *testing.T) {
	op := &ItemOptions{
		APIID:       `sample`,
		AffiliateID: `affiliate-990`,
		Site:        `DMM.R18`,
		Floor:       `dvd`,
		Service:     `mono`,
		Hits:        2,
		Offset:      1000,
		Callback:    ``,
	}
	offset := op.GetOffset()
	if offset != 1000 {
		t.Errorf("GetHits returned %d", offset)
	}
}
