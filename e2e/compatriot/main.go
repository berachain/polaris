// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
)

const GETH_RPC = "http://localhost:8555"
const OTHER_RPC = "http://localhost:8545"
const TESTS = "./tests.json"

var supportedMethods []string
var possiblySupportedMethods []string
var unsupportedMethods []string

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

func main() {
	calls := make([]RPCRequest, 0)
	loadCalls(&calls)
	for i := 0; i < len(calls); i++ {
		call(calls[i])
	}

	color.Set(color.FgGreen)
	fmt.Println("The following JSON-RPC methods are likely supported in your EVM chain:")
	for _, val := range supportedMethods {
		fmt.Println(val)
	}
	fmt.Println()

	color.Set(color.FgYellow)
	fmt.Println("The following JSON-RPC methods may or may not be supported in your EVM chain:")
	for _, val := range possiblySupportedMethods {
		fmt.Println(val)
	}
	fmt.Println()

	color.Set(color.FgRed)
	fmt.Println("The following JSON-RPC methods are likely unsupported in your EVM chain:")
	for _, val := range unsupportedMethods {
		fmt.Println(val)
	}
	fmt.Println()

}

func loadCalls(calls *[]RPCRequest) {
	jsonFile, err := os.Open(TESTS)
	if err != nil {
		log.Fatalf("An error occurred %v", err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, calls)
}

func call(postRequest RPCRequest) {
	postBody, _ := json.Marshal(postRequest)
	gethBuffer := bytes.NewBuffer(postBody)
	otherBuffer := bytes.NewBuffer(postBody)

	gethBody := makeRequest(GETH_RPC, gethBuffer)
	otherBody := makeRequest(OTHER_RPC, otherBuffer)

	fmt.Println(postRequest.Method)
	otherResp := RPCResponse{}
	json.Unmarshal([]byte(otherBody), &otherResp)

	if gethBody != otherBody {
		if otherResp.Err != (RPCResponse{}).Err {
			color.Set(color.FgRed)
			fmt.Printf("ERROR: %v\n", otherResp.Err.Message)
			unsupportedMethods = append(unsupportedMethods, postRequest.Method)
		} else {
			possiblySupportedMethods = append(possiblySupportedMethods, postRequest.Method)
		}

		color.Set(color.FgYellow)
	} else {
		supportedMethods = append(supportedMethods, postRequest.Method)
		color.Set(color.FgGreen)
	}
	fmt.Printf("Geth returned: %vOther returned: %v\n\n", gethBody, otherBody)
	color.Unset()
}

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
