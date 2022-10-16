package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func PrettyPrintData(data interface{}) {
	dataBytes, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Printf("error : could not MarshalIndent json : %v", err.Error())
		return
	}
	fmt.Printf("\n%v\n\n", string(dataBytes))
}
