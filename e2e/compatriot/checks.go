package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func sanityCheck(file string) error {
	var results []RPCOutput

	jsonFile, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("sanityCheck: An error occurred %v when opening the file\n", err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	if err := json.Unmarshal(byteValue, &results); err != nil {
		return fmt.Errorf("sanityCheck: An error occurred %v when unmarshalling the file\n", err)
	}

	return checkNull(results)
}

func checkNull(results []RPCOutput) error {
	for i := 0; i < len(results); i++ {
		if results[i].Response.Result == nil {
			return fmt.Errorf("checkNull: %v returns null result %v\n",
				results[i].Method, results[i].Response.Result)
		}
	}
	return nil
}
