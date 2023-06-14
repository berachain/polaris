package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

type RPCOutput struct {
	Method   string      `json:"method"`
	Response RPCResponse `json:"response"`
}

// Query loads prexisting JSON-RPC calls from a file and queries the chain
func Query(outputFile string) error {
	calls := make([]RPCRequest, 0)
	loadCalls(&calls)

	var output []RPCOutput
	for i := 0; i < len(calls); i++ {
		result, err := call(calls[i])
		if err != nil {
			return fmt.Errorf("Query: An error occurred %v when calling\n", err)
		}
		output = append(output, result)
	}

	// add the results to a file and format
	content, err := Marshal(output)
	if err != nil {
		return fmt.Errorf("Query: An error occurred %v when marshalling output\n", err)
	}

	if err = os.WriteFile("./"+outputFile, content, 0644); err != nil {
		return fmt.Errorf("call: An error occurred %v when writing output\n", err)
	}

	fmt.Println("finished querying")
	return nil
}

func loadCalls(calls *[]RPCRequest) error {
	jsonFile, err := os.Open(TESTS)
	if err != nil {
		return fmt.Errorf("loadCalls: An error occurred %v when opening TESTS\n", err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	if err := json.Unmarshal(byteValue, calls); err != nil {
		return fmt.Errorf("loadCalls: An error occurred %v when unmarshalling TESTS\n", err)
	}

	return nil
}

// call makes a JSON-RPC call to the chain and saves the results
func call(postRequest RPCRequest) (RPCOutput, error) {
	postBody, _ := json.Marshal(postRequest)
	buffer := bytes.NewBuffer(postBody)

	body, err := makeRequest(POLARIS_RPC, buffer)
	if err != nil {
		return RPCOutput{}, fmt.Errorf("call: An error occurred %v when making the request\n", err)
	}
	var response RPCResponse
	json.Unmarshal([]byte(body), &response)

	return RPCOutput{Method: postRequest.Method, Response: response}, nil
}

// makeRequest makes the actual HTTP request to the chain
func makeRequest(rpc string, postBuffer *bytes.Buffer) (string, error) {
	response, err := http.Post(rpc, "application/json", postBuffer)
	if err != nil {
		return "", fmt.Errorf("makeRequest: An Error Occured %v when posting\n", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("makeRequest: An Error Occured %v when reading response\n", err)
	}
	return string(body), nil
}

// Marshal marshals the output slice to JSON
func Marshal(output []RPCOutput) ([]byte, error) {
	jsonOutput, err := json.MarshalIndent(output, "", "    ")
	if err != nil {
		return nil, fmt.Errorf("Marshal: An error occurred %v trying to marshal data\n", err)
	}
	return jsonOutput, nil
}
