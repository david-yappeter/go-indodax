package indodax

import (
	"encoding/json"
	"go-indodax/model"
)

func (c *Client) ServerTime() (*model.ServerTime, error) {
	c.newRequest(c.endpointServerTime(nil))
	body, err := c.do()
	if err != nil {
		return nil, err
	}
	var respModel model.ServerTime
	if err = json.Unmarshal(body, &respModel); err != nil {
		return nil, err
	}
	return &respModel, nil
}
