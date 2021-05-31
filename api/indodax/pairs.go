package indodax

import (
	"encoding/json"
	"myapp/model"
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
