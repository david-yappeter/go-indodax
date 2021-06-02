package example

import "go-indodax/api/indodax"

func ExampleServerTime() {
	c := indodax.NewDefaultClient()
	resp, _ := c.ServerTime()

	prettyPrint(resp)
	/*
		    {
			   "timezone": "UTC",
			   "server_time": 1622454122414
		    }
	*/
}

func ExamplePairs() {
	c := indodax.NewDefaultClient()
	resp, _ := c.Pairs()

	prettyPrint(resp)
}

func ExamplePriceIncrements() {
	c := indodax.NewDefaultClient()
	resp, _ := c.PriceIncrements()

	prettyPrint(resp)
}

func ExampleSummaries() {
	c := indodax.NewDefaultClient()
	resp, _ := c.Summaries()

	prettyPrint(resp)
}

func ExampleTickerAll() {
	c := indodax.NewDefaultClient()
	resp, _ := c.TickerAll()

	prettyPrint(resp)
}

func ExampleTickerByPairs() {
	c := indodax.NewDefaultClient()
	resp, _ := c.TickerByPair("btcidr")

	prettyPrint(resp)
	/*
	   {
	       "error": null,
	       "error_description": null,
	       "ticker": {
	           "high": "546768000",
	           "low": "522889000",
	           "vol_btc": "118.16536439",
	           "vol_idr": "63145588982",
	           "last": "525600000",
	           "buy": "525602000",
	           "sell": "526413000",
	           "server_time": 1622597186
	       }
	   }
	*/
}

func ExampleTickerByPairsError() {
	c := indodax.NewDefaultClient()
	resp, _ := c.TickerByPair("asd")

	prettyPrint(resp)
	/*
	   {
	       "error": "invalid_pair",
	       "error_description": "Invalid Pair",
	       "ticker": null
	   }
	*/
}

func ExampleDepthByPairs() {
	c := indodax.NewDefaultClient()
	resp, _ := c.DepthByPair("btcidr")

	prettyPrint(resp)
}

func ExampleDepthByPairsError() {
	c := indodax.NewDefaultClient()
	resp, _ := c.DepthByPair("asd")

	prettyPrint(resp)

	/*
	   {
	       "error": "invalid_pair",
	       "error_description": "Invalid Pair",
	       "depths": null
	   }
	*/
}
func TradesByPair() {
	c := indodax.NewDefaultClient()
	resp, _ := c.TradesByPair("btcidr")

	prettyPrint(resp)
	/*
	   {
	       "error": "invalid_pair",
	       "error_description": "Invalid Pair",
	       "depths": {
	           "buy": null,
	           "sell": null
	       }
	   }
	*/
}

func TradesByPairError() {
	c := indodax.NewDefaultClient()
	resp, _ := c.TradesByPair("asd")

	prettyPrint(resp)
	/*
	   {
	       "error": "invalid_pair",
	       "error_description": "Invalid Pair",
	       "trades": null
	   }
	*/
}
