package goinsta

import "net/http"

// Option func is used to update goinsta.Instagram fields
type Option func(*Instagram)

// WithHTTPClient is used to set the base http.Client for making requests
func WithHTTPClient(c *http.Client) Option {
	return func(insta *Instagram) {
		insta.c = c
	}
}
