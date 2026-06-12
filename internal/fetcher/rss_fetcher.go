package fetcher

import (
	"io"
	"net/http"
	"time"
)

type RSSFetcher struct {
	URL string
}

func (r *RSSFetcher) Fetch() ([]byte, error) {

	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	req, err := http.NewRequest("GET", r.URL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (writeup-fetcher)")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
