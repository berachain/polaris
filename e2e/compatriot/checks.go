package main

import "fmt"

func sanityCheck(results []RPCOutput) error {
	for i := 0; i < len(results); i++ {
		if err := checkNull(results[i]); err != nil {
			return err
		}
	}
	return nil
}

func checkNull(result RPCOutput) error {
	if result.Response.Result == nil {
		return fmt.Errorf("checkNull: %v returns null result %v\n",
			result.Request.Method, result.Response.Result)
	}
	return nil
}
