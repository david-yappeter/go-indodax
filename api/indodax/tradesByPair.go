package indodax

import (
	"encoding/json"
	"myapp/model"
)

func (c *Client) TradesByPair(pair string) (*model.TradesByPair, error) {
	c.newRequest(c.endpointTradesByPair(nil, pair))
	body, err := c.do()
	if err != nil {
		return nil, err
	}
	var respModel model.TradesByPair
	var respModelNodes []*model.TradesByPairChild
	if err = json.Unmarshal(body, &respModelNodes); err != nil {
		if err = json.Unmarshal(body, &respModel); err != nil {
			return nil, err
		}
	}
	respModel.Trades = respModelNodes
	return &respModel, nil
}
