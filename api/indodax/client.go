package indodax

import (
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	BaseEndpoint        string        `json:"base_endpoint"`
	PrivateBaseEndpoint string        `json:"private_base_endpoint"`
	HttpClient          *http.Client  `json:"http_client"`
	Request             *http.Request `json:"request"`
}

type Option struct {
	HttpClient *http.Client `json:"http_client"`
}

func NewClient(base string, privateBase string, option ...Option) *Client {
	client := Client{
		BaseEndpoint:        base,
		PrivateBaseEndpoint: privateBase,
	}
	if len(option) >= 1 {
		client.HttpClient = option[0].HttpClient
	}
	return &client
}

func NewDefaultClient(option ...Option) *Client {
	client := Client{
		BaseEndpoint:        "https://indodax.com",
		PrivateBaseEndpoint: "https://indodax.com/tapi",
	}
	if len(option) >= 1 {
		client.HttpClient = option[0].HttpClient
	} else {
		client.HttpClient = &http.Client{Timeout: time.Second * 10}
	}
	return &client
}

func (c *Client) newRequest(method string, url string, body io.Reader) error {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	c.Request = request
	return nil
}

func (c *Client) do() ([]byte, error) {
	resp, err := c.HttpClient.Do(c.Request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
