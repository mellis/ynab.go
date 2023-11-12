package ynab

import "net/http"

func HTTPClient(hc *http.Client) func(*client) {
	return func(c *client) {
		c.client = hc
	}
}
