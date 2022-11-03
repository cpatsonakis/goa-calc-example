package openapiui

import (
	"fmt"
	"net/url"
	"strings"
)

func validateExternalURLString(urlString string) (*url.URL, error) {
	url, err := url.Parse(urlString)
	if err != nil {
		return nil, fmt.Errorf("url.Parse() returned error: %w", err)
	}
	if url.Scheme != "http" && url.Scheme != "https" {
		return nil, fmt.Errorf("invalid scheme %s specified, must be http or https", url.Scheme)
	}
	if !strings.HasPrefix(urlString, url.Scheme+"://"+url.Host) {
		return nil, fmt.Errorf("url string %s does not have a valid http(s) prefix, must be of the form <scheme>://<host>", urlString)
	}
	return url, nil
}
