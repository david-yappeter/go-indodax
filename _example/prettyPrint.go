package example

import (
	"encoding/json"
	"fmt"
)

func prettyPrint(value interface{}) {
	jsoned, err := json.MarshalIndent(value, " ", "  ")
	if err != nil {
		fmt.Printf("Pretty Printing Log: %+v\n", err)
		return
	}
	fmt.Println(string(jsoned))
}
