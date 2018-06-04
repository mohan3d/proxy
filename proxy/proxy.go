package proxy

import (
	"net/http"
	"strings"
)

// IgnoredHeaders returns a list of not-allowed headers.
var IgnoredHeaders = []string{
	"set-cookie",
	"expires",
	"cache-control",
	"connection",
	"keep-alive",
	"proxy-authenticate",
	"proxy-authorization",
	"te",
	"trailers",
	"transfer-encoding",
	"upgrade",
}

func headerExists(headerKey string, headers []string) bool {
	key := strings.ToLower(headerKey)
	for _, v := range headers {
		if key == v {
			return true
		}
	}
	return false
}

// AllowedHeaders returns only non-ignored headers.
func AllowedHeaders(headers http.Header) http.Header {
	h2 := make(http.Header)
	for k, v := range headers {
		if !headerExists(k, IgnoredHeaders) {
			v2 := make([]string, len(v))
			copy(v2, v)
			h2[k] = v2
		}
	}
	return h2
}

// Fetcher interface for all url fetchers
type Fetcher interface {
	Fetch(url string) (*http.Response, error)
}

type URLFetcher struct{}

func (f URLFetcher) Fetch(url string) (*http.Response, error) {
	return &http.Response{}, nil
}
