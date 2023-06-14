package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type RPCRequest struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  []any  `json:"params"`
	Id      int64  `json:"id"`
}

type ResponseErr struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type RPCResponse struct {
	Jsonrpc string      `json:"jsonrpc"`
	Id      int64       `json:"id"`
	Result  any         `json:"result"`
	Err     ResponseErr `json:"error"`
}

// makeCalls loads prexisting JSON-RPC calls from a file and queries the chain
func makeCalls(output string) {
	calls := make([]RPCRequest, 0)
	loadCalls(&calls)

	for i := 0; i < len(calls); i++ {
		call(calls[i])
	}
}

func loadCalls(calls *[]RPCRequest) {
	jsonFile, err := os.Open(TESTS)
	if err != nil {
		log.Fatalf("An error occurred %v", err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, calls)
}

// call makes a JSON-RPC call to the chain and saves the results
func call(postRequest RPCRequest) {
	postBody, _ := json.Marshal(postRequest)
	otherBuffer := bytes.NewBuffer(postBody)

	otherBody := makeRequest(POLARIS_RPC, otherBuffer)

	fmt.Println(postRequest.Method)
	otherResp := RPCResponse{}
	json.Unmarshal([]byte(otherBody), &otherResp)

	// add the results to a file and format
}

// makeRequest makes the actual HTTP request to the chain
func makeRequest(rpc string, postBuffer *bytes.Buffer) string {
	resp, err := http.Post(rpc, "application/json", postBuffer)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(body)
}
