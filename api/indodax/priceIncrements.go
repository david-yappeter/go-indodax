package indodax

import (
	"encoding/json"
	"myapp/model"
)

func (c *Client) PriceIncrements() (*model.PriceIncrements, error) {
	c.newRequest(c.endpointPriceIncrements(nil))
	body, err := c.do()
	if err != nil {
		return nil, err
	}
	var respModel model.PriceIncrements
	if err = json.Unmarshal(body, &respModel); err != nil {
		return nil, err
	}
	return &respModel, nil
}
