# Go-DMM
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/usk81/go-dmm)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://github.com/usk81/go-dmm/blob/master/LICENSE)
![](https://github.com/usk81/go-dmm/workflows/preliminary%20review/badge.svg)
[![codecov](https://codecov.io/gh/usk81/go-dmm/branch/master/graph/badge.svg)](https://codecov.io/gh/usk81/go-dmm)
[![Go Report Card](https://goreportcard.com/badge/github.com/usk81/go-dmm)](https://goreportcard.com/report/github.com/usk81/go-dmm)

## Install

```
go get -u github.com/usk81/go-dmm
```

enhanced package

```
go get -u github.com/usk81/go-dmm-enhanced
```

## Usage

e.g. Request to Actress search API:

```go
import (
  "github.com/usk81/go-dmm"
)

ctx := context.TODO()
cli := dmm.NewClient(nil)
result, _, _ := cli.Actresses.First(ctx, nil)
// Actress{
//   ID:          `26617`,
//   Name:        `愛内あみ`,
//   Ruby:        `あいうちあみ`,
//   Bust:        `92`,
//   Cup:         `E`,
//   Waist:       `59`,
//   Hip:         `88`,
//   Height:      `152`,
//   Birthday:    `1987-12-15`,
//   BloodType:   `B`,
//   Hobby:       `音楽鑑賞`,
//   Prefectures: `静岡県`,
//   ImageURL: ImageURL{
//     Small: `http://pics.dmm.co.jp/mono/actjpgs/thumbnail/aiuti_ami.jpg`,
//     Large: `http://pics.dmm.co.jp/mono/actjpgs/aiuti_ami.jpg`,
//   },
//   ListURL: ListURL{
//     Digital: `http://www.dmm.co.jp/digital/videoa/-/list/=/article=actress/id=26617/affiliate-990`,
//     Mono:    `http://www.dmm.co.jp/mono/dvd/-/list/=/article=actress/id=26617/affiliate-990`,
//     Monthly: `http://www.dmm.co.jp/monthly/premium/-/list/=/article=actress/id=26617/affiliate-990`,
//     PPM:     `http://www.dmm.co.jp/ppm/video/-/list/=/article=actress/id=26617/affiliate-990`,
//     Rental:  `http://www.dmm.co.jp/rental/ppr/-/list/=/article=actress/id=26617/affiliate-990`,
//   },
// }
```

use enhanced package:

```go
import (
  "github.com/usk81/go-dmm"
  "github.com/usk81/go-dmm-enhanced"
)

ctx := context.TODO()
cli := dmm.NewClient(nil)
r, _, _ := cli.Actresses.First(ctx, nil)
result, _ := enhanced.ConvertActress(r)
// Actress{
//   ID:          `26617`,
//   Name:        `愛内あみ`,
//   Ruby:        `あいうちあみ`,
//   Bust:        92,
//   Cup:         `E`,
//   Waist:       59,
//   Hip:         88,
//   Height:      152,
//   Birthday:    `1987-12-15`,
//   BloodType:   `B`,
//   Hobby:       `音楽鑑賞`,
//   Prefectures: `静岡県`,
//   ImageURL: ImageURL{
//     Small: `http://pics.dmm.co.jp/mono/actjpgs/thumbnail/aiuti_ami.jpg`,
//     Large: `http://pics.dmm.co.jp/mono/actjpgs/aiuti_ami.jpg`,
//   },
//   ListURL: ListURL{
//     Digital: `http://www.dmm.co.jp/digital/videoa/-/list/=/article=actress/id=26617/affiliate-990`,
//     Mono:    `http://www.dmm.co.jp/mono/dvd/-/list/=/article=actress/id=26617/affiliate-990`,
//     Monthly: `http://www.dmm.co.jp/monthly/premium/-/list/=/article=actress/id=26617/affiliate-990`,
//     PPM:     `http://www.dmm.co.jp/ppm/video/-/list/=/article=actress/id=26617/affiliate-990`,
//     Rental:  `http://www.dmm.co.jp/rental/ppr/-/list/=/article=actress/id=26617/affiliate-990`,
//   },
// }
```


## Licence

[MIT](https://github.com/usk81/go-dmm/blob/master/LICENSE)

## Author

[Yusuke Komatsu](https://github.com/usk81)
