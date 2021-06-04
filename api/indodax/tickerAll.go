package indodax

import (
	"encoding/json"

	"github.com/david-yappeter/go-indodax/model"
)

func (c *Client) TickerAll() (*model.TickerAll, error) {
	c.newRequest(c.endpointTickerAll(nil))
	body, err := c.do()
	if err != nil {
		return nil, err
	}
	var respModel model.TickerAll
	if err = json.Unmarshal(body, &respModel); err != nil {
		return nil, err
	}
	return &respModel, nil
}
