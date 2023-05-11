// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cosmos

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// PolarisERC20MetaData contains all meta data concerning the PolarisERC20 contract.
var PolarisERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_denom\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denom\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b50604051620026e3380380620026e3833981810160405281019062000037919062000292565b80600090816200004891906200052e565b504660808181525050620000616200006f60201b60201c565b60a0818152505050620007ab565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6000604051620000a39190620006c4565b60405180910390207fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc64630604051602001620000e49594939291906200074e565b60405160208183030381529060405280519060200120905090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b62000168826200011d565b810181811067ffffffffffffffff821117156200018a57620001896200012e565b5b80604052505050565b60006200019f620000ff565b9050620001ad82826200015d565b919050565b600067ffffffffffffffff821115620001d057620001cf6200012e565b5b620001db826200011d565b9050602081019050919050565b60005b8381101562000208578082015181840152602081019050620001eb565b60008484015250505050565b60006200022b6200022584620001b2565b62000193565b9050828152602081018484840111156200024a576200024962000118565b5b62000257848285620001e8565b509392505050565b600082601f83011262000277576200027662000113565b5b81516200028984826020860162000214565b91505092915050565b600060208284031215620002ab57620002aa62000109565b5b600082015167ffffffffffffffff811115620002cc57620002cb6200010e565b5b620002da848285016200025f565b91505092915050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200033657607f821691505b6020821081036200034c576200034b620002ee565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620003b67fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000377565b620003c2868362000377565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b60006200040f620004096200040384620003da565b620003e4565b620003da565b9050919050565b6000819050919050565b6200042b83620003ee565b620004436200043a8262000416565b84845462000384565b825550505050565b600090565b6200045a6200044b565b6200046781848462000420565b505050565b5b818110156200048f576200048360008262000450565b6001810190506200046d565b5050565b601f821115620004de57620004a88162000352565b620004b38462000367565b81016020851015620004c3578190505b620004db620004d28562000367565b8301826200046c565b50505b505050565b600082821c905092915050565b60006200050360001984600802620004e3565b1980831691505092915050565b60006200051e8383620004f0565b9150826002028217905092915050565b6200053982620002e3565b67ffffffffffffffff8111156200055557620005546200012e565b5b6200056182546200031d565b6200056e82828562000493565b600060209050601f831160018114620005a6576000841562000591578287015190505b6200059d858262000510565b8655506200060d565b601f198416620005b68662000352565b60005b82811015620005e057848901518255600182019150602085019450602081019050620005b9565b86831015620006005784890151620005fc601f891682620004f0565b8355505b6001600288020188555050505b505050505050565b600081905092915050565b60008190508160005260206000209050919050565b6000815462000644816200031d565b62000650818662000615565b945060018216600081146200066e57600181146200068457620006bb565b60ff1983168652811515820286019350620006bb565b6200068f8562000620565b60005b83811015620006b35781548189015260018201915060208101905062000692565b838801955050505b50505092915050565b6000620006d2828462000635565b915081905092915050565b6000819050919050565b620006f281620006dd565b82525050565b6200070381620003da565b82525050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620007368262000709565b9050919050565b620007488162000729565b82525050565b600060a082019050620007656000830188620006e7565b620007746020830187620006e7565b620007836040830186620006e7565b620007926060830185620006f8565b620007a160808301846200073d565b9695505050505050565b60805160a051611f12620007d1600039600061076a015260006107360152611f126000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c806370a082311161008c578063a9059cbb11610066578063a9059cbb1461022a578063c370b0421461025a578063d505accf14610278578063dd62ed3e14610294576100cf565b806370a08231146101ac5780637ecebe00146101dc57806395d89b411461020c576100cf565b806306fdde03146100d4578063095ea7b3146100f257806318160ddd1461012257806323b872dd14610140578063313ce567146101705780633644e5151461018e575b600080fd5b6100dc6102c4565b6040516100e991906110fc565b60405180910390f35b61010c600480360381019061010791906111b7565b610356565b6040516101199190611212565b60405180910390f35b61012a610498565b604051610137919061123c565b60405180910390f35b61015a60048036038101906101559190611257565b610521565b6040516101679190611212565b60405180910390f35b610178610729565b60405161018591906112c6565b60405180910390f35b610196610732565b6040516101a391906112fa565b60405180910390f35b6101c660048036038101906101c19190611315565b61078f565b6040516101d3919061123c565b60405180910390f35b6101f660048036038101906101f19190611315565b61081c565b604051610203919061123c565b60405180910390f35b610214610834565b60405161022191906110fc565b60405180910390f35b610244600480360381019061023f91906111b7565b6108c6565b6040516102519190611212565b60405180910390f35b610262610a05565b60405161026f91906110fc565b60405180910390f35b610292600480360381019061028d919061139a565b610a93565b005b6102ae60048036038101906102a9919061143c565b610ddc565b6040516102bb919061123c565b60405180910390f35b6060600080546102d3906114ab565b80601f01602080910402602001604051908101604052809291908181526020018280546102ff906114ab565b801561034c5780601f106103215761010080835404028352916020019161034c565b820191906000526020600020905b81548152906001019060200180831161032f57829003601f168201915b5050505050905090565b6000610360610e6c565b73ffffffffffffffffffffffffffffffffffffffff16632b6b7ab5338561038686610e88565b60006040518563ffffffff1660e01b81526004016103a79493929190611688565b6020604051808303816000875af11580156103c6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103ea9190611700565b610429576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104209061179f565b60405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610486919061123c565b60405180910390a36001905092915050565b60006104a2610faa565b73ffffffffffffffffffffffffffffffffffffffff1663fe3b2b8860006040518263ffffffff1660e01b81526004016104db9190611858565b602060405180830381865afa1580156104f8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061051c919061188f565b905090565b600061052b610e6c565b73ffffffffffffffffffffffffffffffffffffffff1663fbdb0e87853360006040518463ffffffff1660e01b8152600401610568939291906118bc565b602060405180830381865afa158015610585573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105a9919061188f565b8211156105eb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105e29061196c565b60405180910390fd5b6105f3610faa565b73ffffffffffffffffffffffffffffffffffffffff166384404811858561061986610e88565b6040518463ffffffff1660e01b81526004016106379392919061198c565b6020604051808303816000875af1158015610656573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061067a9190611700565b6106b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106b090611a3c565b60405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610716919061123c565b60405180910390a3600190509392505050565b60006012905090565b60007f0000000000000000000000000000000000000000000000000000000000000000461461076857610763610fc6565b61078a565b7f00000000000000000000000000000000000000000000000000000000000000005b905090565b6000610799610faa565b73ffffffffffffffffffffffffffffffffffffffff166334d1fdaf8360006040518363ffffffff1660e01b81526004016107d4929190611a5c565b602060405180830381865afa1580156107f1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610815919061188f565b9050919050565b60016020528060005260406000206000915090505481565b606060008054610843906114ab565b80601f016020809104026020016040519081016040528092919081815260200182805461086f906114ab565b80156108bc5780601f10610891576101008083540402835291602001916108bc565b820191906000526020600020905b81548152906001019060200180831161089f57829003601f168201915b5050505050905090565b60006108d0610faa565b73ffffffffffffffffffffffffffffffffffffffff16638440481133856108f686610e88565b6040518463ffffffff1660e01b81526004016109149392919061198c565b6020604051808303816000875af1158015610933573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109579190611700565b610996576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161098d90611afe565b60405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516109f3919061123c565b60405180910390a36001905092915050565b60008054610a12906114ab565b80601f0160208091040260200160405190810160405280929190818152602001828054610a3e906114ab565b8015610a8b5780601f10610a6057610100808354040283529160200191610a8b565b820191906000526020600020905b815481529060010190602001808311610a6e57829003601f168201915b505050505081565b42841015610ad6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610acd90611b90565b60405180910390fd5b60006001610ae2610732565b7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98a8a8a600160008f73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000815480929190600101919050558b604051602001610b6a96959493929190611bb0565b60405160208183030381529060405280519060200120604051602001610b91929190611c89565b6040516020818303038152906040528051906020012085858560405160008152602001604052604051610bc79493929190611cc0565b6020604051602081039080840390855afa158015610be9573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614158015610c5d57508773ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b610c9c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c9390611d51565b60405180910390fd5b610ca4610e6c565b73ffffffffffffffffffffffffffffffffffffffff16632b6b7ab58289610cca8a610e88565b60006040518563ffffffff1660e01b8152600401610ceb9493929190611688565b6020604051808303816000875af1158015610d0a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d2e9190611700565b610d6d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d649061179f565b60405180910390fd5b508573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92587604051610dcb919061123c565b60405180910390a350505050505050565b6000610de6610e6c565b73ffffffffffffffffffffffffffffffffffffffff1663fbdb0e87848460006040518463ffffffff1660e01b8152600401610e23939291906118bc565b602060405180830381865afa158015610e40573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e64919061188f565b905092915050565b600073bdf49c3c3882102fc017ffb661108c63a836d065905090565b60606000600167ffffffffffffffff811115610ea757610ea6611d71565b5b604051908082528060200260200182016040528015610ee057816020015b610ecd611052565b815260200190600190039081610ec55790505b509050604051806040016040528084815260200160008054610f01906114ab565b80601f0160208091040260200160405190810160405280929190818152602001828054610f2d906114ab565b8015610f7a5780601f10610f4f57610100808354040283529160200191610f7a565b820191906000526020600020905b815481529060010190602001808311610f5d57829003601f168201915b505050505081525081600081518110610f9657610f95611da0565b5b602002602001018190525080915050919050565b6000734381dc2ab14285160c808659aee005d51255add7905090565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6000604051610ff89190611e72565b60405180910390207fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc64630604051602001611037959493929190611e89565b60405160208183030381529060405280519060200120905090565b604051806040016040528060008152602001606081525090565b600081519050919050565b600082825260208201905092915050565b60005b838110156110a657808201518184015260208101905061108b565b60008484015250505050565b6000601f19601f8301169050919050565b60006110ce8261106c565b6110d88185611077565b93506110e8818560208601611088565b6110f1816110b2565b840191505092915050565b6000602082019050818103600083015261111681846110c3565b905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061114e82611123565b9050919050565b61115e81611143565b811461116957600080fd5b50565b60008135905061117b81611155565b92915050565b6000819050919050565b61119481611181565b811461119f57600080fd5b50565b6000813590506111b18161118b565b92915050565b600080604083850312156111ce576111cd61111e565b5b60006111dc8582860161116c565b92505060206111ed858286016111a2565b9150509250929050565b60008115159050919050565b61120c816111f7565b82525050565b60006020820190506112276000830184611203565b92915050565b61123681611181565b82525050565b6000602082019050611251600083018461122d565b92915050565b6000806000606084860312156112705761126f61111e565b5b600061127e8682870161116c565b935050602061128f8682870161116c565b92505060406112a0868287016111a2565b9150509250925092565b600060ff82169050919050565b6112c0816112aa565b82525050565b60006020820190506112db60008301846112b7565b92915050565b6000819050919050565b6112f4816112e1565b82525050565b600060208201905061130f60008301846112eb565b92915050565b60006020828403121561132b5761132a61111e565b5b60006113398482850161116c565b91505092915050565b61134b816112aa565b811461135657600080fd5b50565b60008135905061136881611342565b92915050565b611377816112e1565b811461138257600080fd5b50565b6000813590506113948161136e565b92915050565b600080600080600080600060e0888a0312156113b9576113b861111e565b5b60006113c78a828b0161116c565b97505060206113d88a828b0161116c565b96505060406113e98a828b016111a2565b95505060606113fa8a828b016111a2565b945050608061140b8a828b01611359565b93505060a061141c8a828b01611385565b92505060c061142d8a828b01611385565b91505092959891949750929550565b600080604083850312156114535761145261111e565b5b60006114618582860161116c565b92505060206114728582860161116c565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806114c357607f821691505b6020821081036114d6576114d561147c565b5b50919050565b6114e581611143565b82525050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61152081611181565b82525050565b600082825260208201905092915050565b60006115428261106c565b61154c8185611526565b935061155c818560208601611088565b611565816110b2565b840191505092915050565b60006040830160008301516115886000860182611517565b50602083015184820360208601526115a08282611537565b9150508091505092915050565b60006115b98383611570565b905092915050565b6000602082019050919050565b60006115d9826114eb565b6115e381856114f6565b9350836020820285016115f585611507565b8060005b85811015611631578484038952815161161285826115ad565b945061161d836115c1565b925060208a019950506001810190506115f9565b50829750879550505050505092915050565b6000819050919050565b6000819050919050565b600061167261166d61166884611643565b61164d565b611181565b9050919050565b61168281611657565b82525050565b600060808201905061169d60008301876114dc565b6116aa60208301866114dc565b81810360408301526116bc81856115ce565b90506116cb6060830184611679565b95945050505050565b6116dd816111f7565b81146116e857600080fd5b50565b6000815190506116fa816116d4565b92915050565b6000602082840312156117165761171561111e565b5b6000611724848285016116eb565b91505092915050565b7f506f6c6172697345524332303a206661696c656420746f20617070726f76652060008201527f7370656e64000000000000000000000000000000000000000000000000000000602082015250565b6000611789602583611077565b91506117948261172d565b604082019050919050565b600060208201905081810360008301526117b88161177c565b9050919050565b60008190508160005260206000209050919050565b600081546117e1816114ab565b6117eb8186611077565b94506001821660008114611806576001811461181c5761184f565b60ff19831686528115156020028601935061184f565b611825856117bf565b60005b8381101561184757815481890152600182019150602081019050611828565b808801955050505b50505092915050565b6000602082019050818103600083015261187281846117d4565b905092915050565b6000815190506118898161118b565b92915050565b6000602082840312156118a5576118a461111e565b5b60006118b38482850161187a565b91505092915050565b60006060820190506118d160008301866114dc565b6118de60208301856114dc565b81810360408301526118f081846117d4565b9050949350505050565b7f506f6c6172697345524332303a20696e73756666696369656e7420617070726f60008201527f76616c0000000000000000000000000000000000000000000000000000000000602082015250565b6000611956602383611077565b9150611961826118fa565b604082019050919050565b6000602082019050818103600083015261198581611949565b9050919050565b60006060820190506119a160008301866114dc565b6119ae60208301856114dc565b81810360408301526119c081846115ce565b9050949350505050565b7f506f6c6172697345524332303a206661696c656420746f2073656e642062616e60008201527f6b20746f6b656e73000000000000000000000000000000000000000000000000602082015250565b6000611a26602883611077565b9150611a31826119ca565b604082019050919050565b60006020820190508181036000830152611a5581611a19565b9050919050565b6000604082019050611a7160008301856114dc565b8181036020830152611a8381846117d4565b90509392505050565b7f506f6c6172697345524332303a206661696c656420746f2073656e6420746f6b60008201527f656e730000000000000000000000000000000000000000000000000000000000602082015250565b6000611ae8602383611077565b9150611af382611a8c565b604082019050919050565b60006020820190508181036000830152611b1781611adb565b9050919050565b7f506f6c6172697345524332303a205045524d49545f444541444c494e455f455860008201527f5049524544000000000000000000000000000000000000000000000000000000602082015250565b6000611b7a602583611077565b9150611b8582611b1e565b604082019050919050565b60006020820190508181036000830152611ba981611b6d565b9050919050565b600060c082019050611bc560008301896112eb565b611bd260208301886114dc565b611bdf60408301876114dc565b611bec606083018661122d565b611bf9608083018561122d565b611c0660a083018461122d565b979650505050505050565b600081905092915050565b7f1901000000000000000000000000000000000000000000000000000000000000600082015250565b6000611c52600283611c11565b9150611c5d82611c1c565b600282019050919050565b6000819050919050565b611c83611c7e826112e1565b611c68565b82525050565b6000611c9482611c45565b9150611ca08285611c72565b602082019150611cb08284611c72565b6020820191508190509392505050565b6000608082019050611cd560008301876112eb565b611ce260208301866112b7565b611cef60408301856112eb565b611cfc60608301846112eb565b95945050505050565b7f506f6c6172697345524332303a20494e56414c49445f5349474e455200000000600082015250565b6000611d3b601c83611077565b9150611d4682611d05565b602082019050919050565b60006020820190508181036000830152611d6a81611d2e565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600081905092915050565b60008190508160005260206000209050919050565b60008154611dfc816114ab565b611e068186611dcf565b94506001821660008114611e215760018114611e3657611e69565b60ff1983168652811515820286019350611e69565b611e3f85611dda565b60005b83811015611e6157815481890152600182019150602081019050611e42565b838801955050505b50505092915050565b6000611e7e8284611def565b915081905092915050565b600060a082019050611e9e60008301886112eb565b611eab60208301876112eb565b611eb860408301866112eb565b611ec5606083018561122d565b611ed260808301846114dc565b969550505050505056fea264697066735822122020dfcce5b13876f5d8e9ab09537141735409d2830eed167304868cf12785b40964736f6c63430008130033",
}

// PolarisERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use PolarisERC20MetaData.ABI instead.
var PolarisERC20ABI = PolarisERC20MetaData.ABI

// PolarisERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PolarisERC20MetaData.Bin instead.
var PolarisERC20Bin = PolarisERC20MetaData.Bin

// DeployPolarisERC20 deploys a new Ethereum contract, binding an instance of PolarisERC20 to it.
func DeployPolarisERC20(auth *bind.TransactOpts, backend bind.ContractBackend, _denom string) (common.Address, *types.Transaction, *PolarisERC20, error) {
	parsed, err := PolarisERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PolarisERC20Bin), backend, _denom)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PolarisERC20{PolarisERC20Caller: PolarisERC20Caller{contract: contract}, PolarisERC20Transactor: PolarisERC20Transactor{contract: contract}, PolarisERC20Filterer: PolarisERC20Filterer{contract: contract}}, nil
}

// PolarisERC20 is an auto generated Go binding around an Ethereum contract.
type PolarisERC20 struct {
	PolarisERC20Caller     // Read-only binding to the contract
	PolarisERC20Transactor // Write-only binding to the contract
	PolarisERC20Filterer   // Log filterer for contract events
}

// PolarisERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type PolarisERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolarisERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type PolarisERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolarisERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolarisERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolarisERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolarisERC20Session struct {
	Contract     *PolarisERC20     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PolarisERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolarisERC20CallerSession struct {
	Contract *PolarisERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PolarisERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolarisERC20TransactorSession struct {
	Contract     *PolarisERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PolarisERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type PolarisERC20Raw struct {
	Contract *PolarisERC20 // Generic contract binding to access the raw methods on
}

// PolarisERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolarisERC20CallerRaw struct {
	Contract *PolarisERC20Caller // Generic read-only contract binding to access the raw methods on
}

// PolarisERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolarisERC20TransactorRaw struct {
	Contract *PolarisERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPolarisERC20 creates a new instance of PolarisERC20, bound to a specific deployed contract.
func NewPolarisERC20(address common.Address, backend bind.ContractBackend) (*PolarisERC20, error) {
	contract, err := bindPolarisERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PolarisERC20{PolarisERC20Caller: PolarisERC20Caller{contract: contract}, PolarisERC20Transactor: PolarisERC20Transactor{contract: contract}, PolarisERC20Filterer: PolarisERC20Filterer{contract: contract}}, nil
}

// NewPolarisERC20Caller creates a new read-only instance of PolarisERC20, bound to a specific deployed contract.
func NewPolarisERC20Caller(address common.Address, caller bind.ContractCaller) (*PolarisERC20Caller, error) {
	contract, err := bindPolarisERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolarisERC20Caller{contract: contract}, nil
}

// NewPolarisERC20Transactor creates a new write-only instance of PolarisERC20, bound to a specific deployed contract.
func NewPolarisERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*PolarisERC20Transactor, error) {
	contract, err := bindPolarisERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolarisERC20Transactor{contract: contract}, nil
}

// NewPolarisERC20Filterer creates a new log filterer instance of PolarisERC20, bound to a specific deployed contract.
func NewPolarisERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*PolarisERC20Filterer, error) {
	contract, err := bindPolarisERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolarisERC20Filterer{contract: contract}, nil
}

// bindPolarisERC20 binds a generic wrapper to an already deployed contract.
func bindPolarisERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PolarisERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PolarisERC20 *PolarisERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PolarisERC20.Contract.PolarisERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PolarisERC20 *PolarisERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolarisERC20.Contract.PolarisERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PolarisERC20 *PolarisERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PolarisERC20.Contract.PolarisERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PolarisERC20 *PolarisERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PolarisERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PolarisERC20 *PolarisERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolarisERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PolarisERC20 *PolarisERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PolarisERC20.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_PolarisERC20 *PolarisERC20Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PolarisERC20.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_PolarisERC20 *PolarisERC20Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _PolarisERC20.Contract.DOMAINSEPARATOR(&_PolarisERC20.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_PolarisERC20 *PolarisERC20CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _PolarisERC20.Contract.DOMAINSEPARATOR(&_PolarisERC20.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PolarisERC20 *PolarisERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PolarisERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PolarisERC20 *PolarisERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PolarisERC20.Contract.Allowance(&_PolarisERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PolarisERC20 *PolarisERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PolarisERC20.Contract.Allowance(&_PolarisERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_PolarisERC20 *PolarisERC20Caller) BalanceOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PolarisERC20.contract.Call(opts, &out, "balanceOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_PolarisERC20 *PolarisERC20Session) BalanceOf(user common.Address) (*big.Int, error) {
	return _PolarisERC20.Contract.BalanceOf(&_PolarisERC20.CallOpts, user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_PolarisERC20 *PolarisERC20CallerSession) BalanceOf(user common.Address) (*big.Int, error) {
	return _PolarisERC20.Contract.BalanceOf(&_PolarisERC20.CallOpts, user)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_PolarisERC20 *PolarisERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PolarisERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_PolarisERC20 *PolarisERC20Session) Decimals() (uint8, error) {
	return _PolarisERC20.Contract.Decimals(&_PolarisERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_PolarisERC20 *PolarisERC20CallerSession) Decimals() (uint8, error) {
	return _PolarisERC20.Contract.Decimals(&_PolarisERC20.CallOpts)
}

// Denom is a free data retrieval call binding the contract method 0xc370b042.
//
// Solidity: function denom() view returns(string)
func (_PolarisERC20 *PolarisERC20Caller) Denom(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PolarisERC20.contract.Call(opts, &out, "denom")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Denom is a free data retrieval call binding the contract method 0xc370b042.
//
// Solidity: function denom() view returns(string)
func (_PolarisERC20 *PolarisERC20Session) Denom() (string, error) {
	return _PolarisERC20.Contract.Denom(&_PolarisERC20.CallOpts)
}

// Denom is a free data retrieval call binding the contract method 0xc370b042.
//
// Solidity: function denom() view returns(string)
func (_PolarisERC20 *PolarisERC20CallerSession) Denom() (string, error) {
	return _PolarisERC20.Contract.Denom(&_PolarisERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PolarisERC20 *PolarisERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PolarisERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PolarisERC20 *PolarisERC20Session) Name() (string, error) {
	return _PolarisERC20.Contract.Name(&_PolarisERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PolarisERC20 *PolarisERC20CallerSession) Name() (string, error) {
	return _PolarisERC20.Contract.Name(&_PolarisERC20.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_PolarisERC20 *PolarisERC20Caller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PolarisERC20.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_PolarisERC20 *PolarisERC20Session) Nonces(arg0 common.Address) (*big.Int, error) {
	return _PolarisERC20.Contract.Nonces(&_PolarisERC20.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_PolarisERC20 *PolarisERC20CallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _PolarisERC20.Contract.Nonces(&_PolarisERC20.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PolarisERC20 *PolarisERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PolarisERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PolarisERC20 *PolarisERC20Session) Symbol() (string, error) {
	return _PolarisERC20.Contract.Symbol(&_PolarisERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PolarisERC20 *PolarisERC20CallerSession) Symbol() (string, error) {
	return _PolarisERC20.Contract.Symbol(&_PolarisERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PolarisERC20 *PolarisERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PolarisERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PolarisERC20 *PolarisERC20Session) TotalSupply() (*big.Int, error) {
	return _PolarisERC20.Contract.TotalSupply(&_PolarisERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PolarisERC20 *PolarisERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _PolarisERC20.Contract.TotalSupply(&_PolarisERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_PolarisERC20 *PolarisERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PolarisERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_PolarisERC20 *PolarisERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PolarisERC20.Contract.Approve(&_PolarisERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_PolarisERC20 *PolarisERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PolarisERC20.Contract.Approve(&_PolarisERC20.TransactOpts, spender, amount)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_PolarisERC20 *PolarisERC20Transactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PolarisERC20.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_PolarisERC20 *PolarisERC20Session) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PolarisERC20.Contract.Permit(&_PolarisERC20.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_PolarisERC20 *PolarisERC20TransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _PolarisERC20.Contract.Permit(&_PolarisERC20.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_PolarisERC20 *PolarisERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PolarisERC20.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_PolarisERC20 *PolarisERC20Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PolarisERC20.Contract.Transfer(&_PolarisERC20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_PolarisERC20 *PolarisERC20TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PolarisERC20.Contract.Transfer(&_PolarisERC20.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_PolarisERC20 *PolarisERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PolarisERC20.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_PolarisERC20 *PolarisERC20Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PolarisERC20.Contract.TransferFrom(&_PolarisERC20.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_PolarisERC20 *PolarisERC20TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PolarisERC20.Contract.TransferFrom(&_PolarisERC20.TransactOpts, from, to, amount)
}

// PolarisERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PolarisERC20 contract.
type PolarisERC20ApprovalIterator struct {
	Event *PolarisERC20Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PolarisERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolarisERC20Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PolarisERC20Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PolarisERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolarisERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolarisERC20Approval represents a Approval event raised by the PolarisERC20 contract.
type PolarisERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PolarisERC20 *PolarisERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PolarisERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PolarisERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PolarisERC20ApprovalIterator{contract: _PolarisERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PolarisERC20 *PolarisERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PolarisERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PolarisERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolarisERC20Approval)
				if err := _PolarisERC20.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PolarisERC20 *PolarisERC20Filterer) ParseApproval(log types.Log) (*PolarisERC20Approval, error) {
	event := new(PolarisERC20Approval)
	if err := _PolarisERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolarisERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PolarisERC20 contract.
type PolarisERC20TransferIterator struct {
	Event *PolarisERC20Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PolarisERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolarisERC20Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PolarisERC20Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PolarisERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolarisERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolarisERC20Transfer represents a Transfer event raised by the PolarisERC20 contract.
type PolarisERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PolarisERC20 *PolarisERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PolarisERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PolarisERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PolarisERC20TransferIterator{contract: _PolarisERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PolarisERC20 *PolarisERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PolarisERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PolarisERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolarisERC20Transfer)
				if err := _PolarisERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PolarisERC20 *PolarisERC20Filterer) ParseTransfer(log types.Log) (*PolarisERC20Transfer, error) {
	event := new(PolarisERC20Transfer)
	if err := _PolarisERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
