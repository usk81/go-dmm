package dmm

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

const testFloorRequest = `
{
	"request": {
		"parameters": {
			"api_id": "sample",
			"affiliate_id": "affiliate-990",
			"output": "json"
		}
	},
	"result": {
		"site": [
			{
				"name": "DMM.com（一般）",
				"code": "DMM.com",
				"service": [
					{
						"name": "AKB48グループ",
						"code": "lod",
						"floor": [
							{
								"id": "1",
								"name": "AKB48",
								"code": "akb48"
							},
							{
								"id": "2",
								"name": "SKE48",
								"code": "ske48"
							},
							{
								"id": "3",
								"name": "NMB48",
								"code": "nmb48"
							},
							{
								"id": "4",
								"name": "HKT48",
								"code": "hkt48"
							},
							{
								"id": "5",
								"name": "NGT48",
								"code": "ngt48"
							},
							{
								"id": "6",
								"name": "REVIVAL!! ON DEMAND",
								"code": "rod"
							}
						]
					},
					{
						"name": "動画",
						"code": "digital",
						"floor": [
							{
								"id": "90",
								"name": "一般動画",
								"code": "videomarket"
							},
							{
								"id": "9",
								"name": "アイドル",
								"code": "idol"
							},
							{
								"id": "10",
								"name": "舞台",
								"code": "cinema"
							},
							{
								"id": "12",
								"name": "VR",
								"code": "video"
							}
						]
					}
				]
			},
			{
				"name": "DMM.R18（アダルト）",
				"code": "DMM.R18",
				"service": [
					{
						"name": "動画",
						"code": "digital",
						"floor": [
							{
								"id": "43",
								"name": "ビデオ",
								"code": "videoa"
							},
							{
								"id": "44",
								"name": "素人",
								"code": "videoc"
							},
							{
								"id": "45",
								"name": "成人映画",
								"code": "nikkatsu"
							},
							{
								"id": "46",
								"name": "アニメ動画",
								"code": "anime"
							}
						]
					},
					{
						"name": "通販",
						"code": "mono",
						"floor": [
							{
								"id": "74",
								"name": "DVD",
								"code": "dvd"
							},
							{
								"id": "75",
								"name": "大人のおもちゃ",
								"code": "goods"
							},
							{
								"id": "76",
								"name": "アニメ",
								"code": "anime"
							},
							{
								"id": "77",
								"name": "PCゲーム",
								"code": "pcgame"
							},
							{
								"id": "78",
								"name": "ブック",
								"code": "book"
							},
							{
								"id": "79",
								"name": "同人",
								"code": "doujin"
							}
						]
					}
				]
			}
		]
	}
}`

func TestFloors_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(`/`+floorBasePath, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, testFloorRequest)
	})

	actual, _, err := client.Floors.List(ctx, nil)
	if err != nil {
		t.Errorf("Floors.List returned error: %v", err)
	}

	expected := []Site{
		Site{
			Name: `DMM.com（一般）`,
			Code: `DMM.com`,
			Services: []Service{
				Service{
					Name: `AKB48グループ`,
					Code: `lod`,
					Floor: []Floor{
						Floor{
							ID:   `1`,
							Name: `AKB48`,
							Code: `akb48`,
						},
						Floor{
							ID:   `2`,
							Name: `SKE48`,
							Code: `ske48`,
						},
						Floor{
							ID:   `3`,
							Name: `NMB48`,
							Code: `nmb48`,
						},
						Floor{
							ID:   `4`,
							Name: `HKT48`,
							Code: `hkt48`,
						},
						Floor{
							ID:   `5`,
							Name: `NGT48`,
							Code: `ngt48`,
						},
						Floor{
							ID:   `6`,
							Name: `REVIVAL!! ON DEMAND`,
							Code: `rod`,
						},
					},
				},
				Service{
					Name: `動画`,
					Code: `digital`,
					Floor: []Floor{
						Floor{
							ID:   `90`,
							Name: `一般動画`,
							Code: `videomarket`,
						},
						Floor{
							ID:   `9`,
							Name: `アイドル`,
							Code: `idol`,
						},
						Floor{
							ID:   `10`,
							Name: `舞台`,
							Code: `cinema`,
						},
						Floor{
							ID:   `12`,
							Name: `VR`,
							Code: `video`,
						},
					},
				},
			},
		},
		Site{
			Name: `DMM.R18（アダルト）`,
			Code: `DMM.R18`,
			Services: []Service{
				Service{
					Name: `動画`,
					Code: `digital`,
					Floor: []Floor{
						Floor{
							ID:   `43`,
							Name: `ビデオ`,
							Code: `videoa`,
						},
						Floor{
							ID:   `44`,
							Name: `素人`,
							Code: `videoc`,
						},
						Floor{
							ID:   `45`,
							Name: `成人映画`,
							Code: `nikkatsu`,
						},
						Floor{
							ID:   `46`,
							Name: `アニメ動画`,
							Code: `anime`,
						},
					},
				},
				Service{
					Name: `通販`,
					Code: `mono`,
					Floor: []Floor{
						Floor{
							ID:   `74`,
							Name: `DVD`,
							Code: `dvd`,
						},
						Floor{
							ID:   `75`,
							Name: `大人のおもちゃ`,
							Code: `goods`,
						},
						Floor{
							ID:   `76`,
							Name: `アニメ`,
							Code: `anime`,
						},
						Floor{
							ID:   `77`,
							Name: `PCゲーム`,
							Code: `pcgame`,
						},
						Floor{
							ID:   `78`,
							Name: `ブック`,
							Code: `book`,
						},
						Floor{
							ID:   `79`,
							Name: `同人`,
							Code: `doujin`,
						},
					},
				},
			},
		},
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Floors.List returned %+v, expected %+v", actual, expected)
	}
}

func TestFloors_First(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(`/`+floorBasePath, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, testFloorRequest)
	})

	actual, _, err := client.Floors.First(ctx, nil)
	if err != nil {
		t.Errorf("Floors.List returned error: %v", err)
	}

	expected := Site{
		Name: `DMM.com（一般）`,
		Code: `DMM.com`,
		Services: []Service{
			Service{
				Name: `AKB48グループ`,
				Code: `lod`,
				Floor: []Floor{
					Floor{
						ID:   `1`,
						Name: `AKB48`,
						Code: `akb48`,
					},
					Floor{
						ID:   `2`,
						Name: `SKE48`,
						Code: `ske48`,
					},
					Floor{
						ID:   `3`,
						Name: `NMB48`,
						Code: `nmb48`,
					},
					Floor{
						ID:   `4`,
						Name: `HKT48`,
						Code: `hkt48`,
					},
					Floor{
						ID:   `5`,
						Name: `NGT48`,
						Code: `ngt48`,
					},
					Floor{
						ID:   `6`,
						Name: `REVIVAL!! ON DEMAND`,
						Code: `rod`,
					},
				},
			},
			Service{
				Name: `動画`,
				Code: `digital`,
				Floor: []Floor{
					Floor{
						ID:   `90`,
						Name: `一般動画`,
						Code: `videomarket`,
					},
					Floor{
						ID:   `9`,
						Name: `アイドル`,
						Code: `idol`,
					},
					Floor{
						ID:   `10`,
						Name: `舞台`,
						Code: `cinema`,
					},
					Floor{
						ID:   `12`,
						Name: `VR`,
						Code: `video`,
					},
				},
			},
		},
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Floors.First returned %+v, expected %+v", actual, expected)
	}
}
