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
func makeCalls(outputFile string) {
	calls := make([]RPCRequest, 0)
	loadCalls(&calls)

	var output string
	for i := 0; i < len(calls); i++ {
		output += formatOutput(calls[i].Method, call(calls[i]))
	}

	// add the results to a file and format
	err := os.WriteFile("./"+outputFile, []byte(output), 0644)
	if err != nil {
		log.Fatalf("call: An error occurred %v when writing output", err)
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
func call(postRequest RPCRequest) RPCResponse {
	postBody, _ := json.Marshal(postRequest)
	buffer := bytes.NewBuffer(postBody)

	body := makeRequest(POLARIS_RPC, buffer)
	response := RPCResponse{}
	json.Unmarshal([]byte(body), &response)

	return response
}

// makeRequest makes the actual HTTP request to the chain
func makeRequest(rpc string, postBuffer *bytes.Buffer) string {
	response, err := http.Post(rpc, "application/json", postBuffer)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(body)
}

func formatOutput(method string, result RPCResponse) string {
	return fmt.Sprintln("-------------------------\n" +
		"Method: " + method + "\n" +
		"Result: " + result.Result.(string) + "\n" +
		"Error: " + result.Err.Message + "\n" +
		"-------------------------\n")
}
