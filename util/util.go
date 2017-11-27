package util

import (
	"io/ioutil"
	"net/http"
	"github.com/djimenez/iconv-go"
)

func GetPageContent(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	utfBody, err := iconv.NewReader(resp.Body, "KSC5601", "utf-8")
	if err != nil {
		// handler error
	}

	dat, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(dat), nil
}
