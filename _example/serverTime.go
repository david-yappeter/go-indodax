package example

import "go-indodax/api/indodax"

func ExampleServerTime() {
	c := indodax.NewDefaultClient()
	sTime, _ := c.ServerTime()

	prettyPrint(sTime)
	/*
		    {
			   "timezone": "UTC",
			   "server_time": 1622454122414
		    }
	*/
}
