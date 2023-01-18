// Copyright (C) 2022, Berachain Foundation. All rights reserved.
// See the file LICENSE for licensing terms.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

//go:build test_unit

package contracts

import (
	"embed"
	"encoding/json"
	"fmt"

	evmtypes "github.com/evmos/ethermint/x/evm/types"
	"gopkg.in/yaml.v2"

	"github.com/berachain/berachain-node/pkg/dahlia/pkg/common/hexutil"
	"github.com/berachain/berachain-node/pkg/dahlia/pkg/types/abi"
)

const (
	contractsFile = "contracts.yaml"
	interfaceFile = "interfaces.yaml"
	contractsDir  = "out/"
)

var (
	//go:embed contracts.yaml
	contracts embed.FS

	//go:embed interfaces.yaml
	interfaces embed.FS

	//go:embed out
	out embed.FS
)

type ForgeOutputs struct {
	Name string `yaml:"name"`
	File string `yaml:"filename"`
}

type ForgeJSON struct {
	ABI      abi.ABI `json:"abi"`
	Bytecode struct {
		Object hexutil.Bytes `json:"object"`
	} `json:"deployedBytecode"`
}

type BerachainContractsYaml struct {
	ContractData []ForgeOutputs `yaml:"contracts"`
}

type BerachainInterfacesYaml struct {
	InterfaceData []ForgeOutputs `yaml:"interfaces"`
}

// GenerateContracts returns a map of contract names to CompiledContracts for all contracts
func GenerateContracts() (keys []string, contractData map[string]evmtypes.CompiledContract) {
	// Create a map with a capacity equal to the number of contracts.
	contractData = make(map[string]evmtypes.CompiledContract, len(contractData))

	// Read the contracts file and unmarshal the data into the contractsYaml variable.
	var contractsYaml BerachainContractsYaml
	data, err := contracts.ReadFile(contractsFile)
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(data, &contractsYaml); err != nil {
		panic(err)
	}

	// Loop through the contract data and build the CompiledContracts for each contract.
	for _, cData := range contractsYaml.ContractData {
		name := cData.Name
		file := cData.File
		contractData[name] = buildCompiledContract(file, name)
		keys = append(keys, name)
	}
	return
}

// GenerateABIs reads the contents of the interface file and unmarshals the data
// into a BerachainInterfacesYaml type. It then loops through the interface data
// and builds the ABI for each interface, storing the result in a map with the
// interface name as the key. It returns the map and a slice of the keys.
func GenerateABIs() (keys []string, abiData map[string]abi.ABI) {
	// Create a map with a capacity equal to the number of interfaces.
	abiData = make(map[string]abi.ABI, len(abiData))

	// Read the interface file and unmarshal the data into the interfacesYaml variable.
	var interfacesYaml BerachainInterfacesYaml
	data, err := interfaces.ReadFile(interfaceFile)
	// Panic if there was an error reading the file.
	if err != nil {
		panic(err)
	}

	// Panic if there was an error unmarshaling the data.
	if err := yaml.Unmarshal(data, &interfacesYaml); err != nil {
		fmt.Println(err)
	}

	// Loop through the interface data and build the ABIs for each interface.
	for _, iData := range interfacesYaml.InterfaceData {
		name := iData.Name
		file := iData.File
		abiData[name] = buildCompiledContract(file, name).ABI
		keys = append(keys, name)
	}
	return
}

// buildCompiledContract reads the contents of a JSON file located at
// contractsDir + folder + ".sol/" + file + ".json" and unmarshals it
// into a ForgeJSON type. It then returns a CompiledContract struct
// populated with the values from the ForgeJSON struct.
func buildCompiledContract(file, name string) evmtypes.CompiledContract {
	// Read the contents of the file at the specified location.
	data, err := out.ReadFile(contractsDir + file + ".sol/" + name + ".json")
	// Panic if there was an error reading the file.
	if err != nil {
		panic(err)
	}

	// Declare a variable to hold the unmarshaled JSON data.
	var forgeJSON ForgeJSON

	// Unmarshal the JSON data into the forgeJSON variable.
	if err := json.Unmarshal(data, &forgeJSON); err != nil {
		panic(err)
	}

	// Return a CompiledContract struct populated with the data loaded from the JSON file.
	return evmtypes.CompiledContract{
		ABI: forgeJSON.ABI,
		Bin: []byte(forgeJSON.Bytecode.Object),
	}
}
