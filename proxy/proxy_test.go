package proxy

import (
	"net/http"
	"testing"
)

func createEmptyHeaders(headerKeys []string) http.Header {
	headers := make(http.Header, len(headerKeys))

	for _, key := range headerKeys {
		headers.Set(key, "")
	}
	return headers
}

func TestAllowedHeadersWithNoIgnoredHeader(t *testing.T) {
	headerKeys := []string{
		"Content-Type",
		"Content-Encoding",
		"Content-Length",
	}
	headers := AllowedHeaders(createEmptyHeaders(headerKeys))

	if len(headerKeys) != len(headers) {
		t.Errorf(
			"expected length of old headers equals new headers but found old(%v) != new(%v)",
			len(headerKeys),
			len(headers),
		)
	}
}

func TestAllowedHeadersWithIgnoredHeader(t *testing.T) {
	headerKeys := []string{
		"Content-Type",
		"Set-Cookie",
		"Content-Length",
		"Keep-Alive",
	}
	headers := AllowedHeaders(createEmptyHeaders(headerKeys))

	for k := range headers {
		if headerExists(k, IgnoredHeaders) {
			t.Errorf("expected %s to be removed, but found", k)
		}
	}
}
