package api

import (
	"encoding/json"

	"github.com/david-yappeter/go-indodax/model"
)

func (c *Client) Pairs() (*model.Pairs, error) {
	c.newRequest(c.endpointPairs(nil))
	body, err := c.do()
	if err != nil {
		return nil, err
	}
	var respModel model.Pairs
	if err = json.Unmarshal(body, &respModel); err != nil {
		return nil, err
	}
	return &respModel, nil
}
