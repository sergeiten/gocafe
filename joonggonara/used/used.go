package used

import (
	"fmt"
	"log"

	"bufio"
	"io"
	"net/http"

	"strings"

	"net/url"

	"sync"

	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/htmlindex"
	iconv "gopkg.in/iconv.v1"
)

const (
	listURL = "http://cafe.naver.com/ArticleSearchList.nhn?search.clubid=10050146&search.searchdate=all&search.searchBy=0&search.query=%s&search.defaultValue=1&search.sortBy=date&userDisplay=50&search.media=0&search.option=0&search.page=%d"
)

type item struct {
	No      string
	Title   string
	Content string
	Name    string
	Views   string
	Likes   string
}

func Fetch(query string, pages int) [][]string {
	var result [][]string

	query = convert(query, "utf-8", "euc-kr")

	var wg sync.WaitGroup

	wg.Add(pages)

	for i := 1; i <= pages; i++ {
		go func(page int) {
			time.Sleep(time.Duration(page*100) * time.Millisecond)

			listURL := getListURL(query, page)

			list, err := getList(&wg, listURL)
			if err != nil {
				log.Printf("failed to get listURL: %v", err)
			}

			for _, l := range list {
				result = append(result, []string{
					l.No,
					l.Title,
					l.Content,
					l.Name,
					l.Views,
					l.Likes,
				})
			}
		}(i)
	}

	wg.Wait()

	return result
}

func getListURL(query string, page int) string {
	return fmt.Sprintf(listURL, url.QueryEscape(query), page)
}

func getList(wg *sync.WaitGroup, url string) ([]item, error) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response code, url: %s, code: %d", url, resp.StatusCode)
	}

	reader, err := decodeHTMLBody(resp.Body, "KSC5601")
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}

	var output []item

	doc.Find(`form[name="ArticleList"] table.board-box tr[align="center"]`).Each(func(i int, sel *goquery.Selection) {
		var no string
		var title string
		var name string
		var views string
		var likes string
		var pageUrl string

		sel.Find("td").Each(func(i int, sel *goquery.Selection) {
			switch i {
			case 0:
				no = strings.TrimSpace(sel.Text())
			case 1:
				title = getTitle(sel)
				pageUrl = getPageURL(sel)
			case 2:
				name = getName(sel)
			case 6:
				views = strings.TrimSpace(sel.Text())
			case 7:
				likes = strings.TrimSpace(sel.Text())
			}
		})

		rowItem := item{
			no,
			title,
			getContent(pageUrl),
			name,
			views,
			likes,
		}

		output = append(output, rowItem)
	})

	return output, nil
}

func getContent(url string) string {
	if url == "" {
		return ""
	}

	resp, err := http.Get("http://cafe.naver.com" + url)
	if err != nil {
		log.Fatal("failed to get resp")
		return ""
	}
	defer resp.Body.Close()

	reader, err := decodeHTMLBody(resp.Body, "KSC5601")
	if err != nil {
		log.Fatal("failed to get reader")
		return ""
	}

	doc, err := goquery.NewDocumentFromReader(reader)

	if err != nil {
		log.Printf("failed to get url: %v", err)
	}

	var html string
	doc.Find("#tbody p").Each(func(i int, selection *goquery.Selection) {
		html = html + selection.Text()
	})

	return html
}

func getTitle(sel *goquery.Selection) string {
	return strings.TrimSpace(sel.Find("a").Text())
}

func getName(sel *goquery.Selection) string {
	return strings.TrimSpace(sel.Find("span.wordbreak").Text())
}

func getPageURL(sel *goquery.Selection) string {
	url, exist := sel.Find("a").Attr("href")
	if exist {
		return url
	}
	return ""
}

func detectContentCharset(body io.Reader) string {
	r := bufio.NewReader(body)
	if data, err := r.Peek(1024); err == nil {
		if _, name, ok := charset.DetermineEncoding(data, ""); ok {
			return name
		}
	}
	return "utf-8"
}

// DecodeHTMLBody returns an decoding reader of the html Body for the specified `charset`
// If `charset` is empty, DecodeHTMLBody tries to guess the encoding from the content
func decodeHTMLBody(body io.Reader, charset string) (io.Reader, error) {
	if charset == "" {
		charset = detectContentCharset(body)
	}
	e, err := htmlindex.Get(charset)
	if err != nil {
		return nil, err
	}
	if name, _ := htmlindex.Name(e); name != "utf-8" {
		body = e.NewDecoder().Reader(body)
	}
	return body, nil
}

func convert(str string, from string, to string) string {
	converter, _ := iconv.Open(to, from)
	defer converter.Close()

	o := converter.ConvString(str)
	return o
}
