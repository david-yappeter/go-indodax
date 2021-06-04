package api

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (c *Client) endpointServerTime(data url.Values) (string, string, io.Reader) {
	return http.MethodGet, fmt.Sprintf("%s/%s/%s", c.BaseEndpoint, "api", "server_time"), nil
}

func (c *Client) endpointPairs(data url.Values) (string, string, io.Reader) {
	return http.MethodGet, fmt.Sprintf("%s/%s/%s", c.BaseEndpoint, "api", "pairs"), nil
}

func (c *Client) endpointPriceIncrements(data url.Values) (string, string, io.Reader) {
	return http.MethodGet, fmt.Sprintf("%s/%s/%s", c.BaseEndpoint, "api", "price_increments"), nil
}

func (c *Client) endpointSummaries(data url.Values) (string, string, io.Reader) {
	return http.MethodGet, fmt.Sprintf("%s/%s/%s", c.BaseEndpoint, "api", "summaries"), nil
}

func (c *Client) endpointTickerAll(data url.Values) (string, string, io.Reader) {
	return http.MethodGet, fmt.Sprintf("%s/%s/%s", c.BaseEndpoint, "api", "ticker_all"), nil
}

func (c *Client) endpointTickerByPair(data url.Values, pair string) (string, string, io.Reader) {
	return http.MethodGet, fmt.Sprintf("%s/%s/%s/%s", c.BaseEndpoint, "api", "ticker", pair), nil
}

func (c *Client) endpointTradesByPair(data url.Values, pair string) (string, string, io.Reader) {
	return http.MethodGet, fmt.Sprintf("%s/%s/%s/%s", c.BaseEndpoint, "api", "trades", pair), nil
}

func (c *Client) endpointDepthsByPair(data url.Values, pair string) (string, string, io.Reader) {
	return http.MethodGet, fmt.Sprintf("%s/%s/%s/%s", c.BaseEndpoint, "api", "depth", pair), nil
}

