package crawler

import (
	"io"
	"net/http"
)

// type CrawledData struct {
// 	content string
// }

func CrawlData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// return &CrawledData{content: string(body)}, nil
	return body, nil

}