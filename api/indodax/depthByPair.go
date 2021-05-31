package indodax

import (
	"encoding/json"
	"myapp/model"
)

func (c *Client) DepthByPair(pair string) (*model.DepthByPair, error) {
	c.newRequest(c.endpointDepthsByPair(nil, pair))
	body, err := c.do()
	if err != nil {
		return nil, err
	}
	var respModel model.DepthByPair
	var respModelNodes model.DepthByPairChild
	if err = json.Unmarshal(body, &respModelNodes); err != nil {
		if err = json.Unmarshal(body, &respModel); err != nil {
			return nil, err
		}
	}
	respModel.Depths = &respModelNodes
	return &respModel, nil
}
