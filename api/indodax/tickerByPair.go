package indodax

import (
	"encoding/json"

	"github.com/david-yappeter/go-indodax/model"
)

func (c *Client) TickerByPair(pair string) (*model.TickerByPair, error) {
	c.newRequest(c.endpointTickerByPair(nil, pair))
	body, err := c.do()
	if err != nil {
		return nil, err
	}
	var respModel model.TickerByPair
	if err = json.Unmarshal(body, &respModel); err != nil {
		return nil, err
	}
	return &respModel, nil
}
