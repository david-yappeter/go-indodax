package indodax

import (
	"encoding/json"
	"go-indodax/model"
)

func (c *Client) Summaries() (*model.Summaries, error) {
	c.newRequest(c.endpointSummaries(nil))
	body, err := c.do()
	if err != nil {
		return nil, err
	}
	var respModel model.Summaries
	if err = json.Unmarshal(body, &respModel); err != nil {
		return nil, err
	}
	return &respModel, nil
}
