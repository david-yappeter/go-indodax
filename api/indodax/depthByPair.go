package indodax

import (
	"encoding/json"
	"go-indodax/model"
)

func (c *Client) DepthByPair(pair string) (*model.DepthByPair, error) {
	c.newRequest(c.endpointDepthsByPair(nil, pair))
	body, err := c.do()
	if err != nil {
		return nil, err
	}
	var respModel model.DepthByPair
	var respModelNodes *model.DepthByPairChild
	if err = json.Unmarshal(body, &respModelNodes); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, &respModel); err != nil {
		return nil, err
	}
	if respModelNodes.Buy == nil || respModelNodes.Sell == nil {
		respModelNodes = nil
	}
	respModel.Depths = respModelNodes
	return &respModel, nil
}
