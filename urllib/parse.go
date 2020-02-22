package urllib

import (
	"net/url"
)

func Quote(s string) string {
	return url.QueryEscape(s)
}

func UnQuote(s string) string {
	s, _ = url.QueryUnescape(s)
	return s
}
