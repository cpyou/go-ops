// 完成 Gocn 每日新闻抓取
package crawler

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"go-ops/pkg/apirequest"
	"io"
	"strings"
)

var headers = map[string]string{
	"Host":       "gocn.vip",
	"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36",
}

type GocnScrawler struct {
	Url string
}

func (gs *GocnScrawler) Request(url string) io.ReadCloser {
	_, body := apirequest.Get(url, headers, nil)
	return body
}

func (gs *GocnScrawler) Parse(closer io.ReadCloser) interface{} {
	doc, err := goquery.NewDocumentFromReader(closer)
	res := make([]string, 0, 32)
	if err != nil {
		panic(err)
	} else {
		h4 := doc.Find("div[class=aw-question-content]").Find("h4")
		h4.Find("a").Each(func(i int, selection *goquery.Selection) {
			u, _ := selection.Attr("href")
			if u != "" {
				res = append(res, u)
			}
		})
	}
	_, d := apirequest.Get(res[0], headers, nil)
	news, _ := goquery.NewDocumentFromReader(d)
	r := make([]string, 0, 32)
	news.Find("div[class=mod-body]").Each(func(i int, selection *goquery.Selection) {
		r = append(r, strings.TrimSpace(selection.Text()))
	})
	fmt.Println(r)
	return r
}

func (gs *GocnScrawler) Save(interface{}) error {
	return nil
}
