package fetcher

import (
	"io"
	"net/http"
)

type RSSFetcher struct {
	URL string
}

func (r *RSSFetcher) Fetch() ([]byte, error) {
	resp, err := http.Get(r.URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
