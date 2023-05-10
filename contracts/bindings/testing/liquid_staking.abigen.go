// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testing

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

// LiquidStakingMetaData contains all meta data concerning the LiquidStaking contract.
var LiquidStakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Data\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"Success\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staking\",\"outputs\":[{\"internalType\":\"contractIStakingModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"name\":\"totalDelegated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x61010060405273d9a998cac66092748ffec7cfbd155aae1737c2ff73ffffffffffffffffffffffffffffffffffffffff1660e09073ffffffffffffffffffffffffffffffffffffffff168152503480156200005957600080fd5b50604051620029803803806200298083398181016040528101906200007f919062000302565b818160128260009081620000949190620005d2565b508160019081620000a69190620005d2565b508060ff1660808160ff16815250504660a08181525050620000cd620000df60201b60201c565b60c0818152505050505050506200084f565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f600060405162000113919062000768565b60405180910390207fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6463060405160200162000154959493929190620007f2565b60405160208183030381529060405280519060200120905090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b620001d8826200018d565b810181811067ffffffffffffffff82111715620001fa57620001f96200019e565b5b80604052505050565b60006200020f6200016f565b90506200021d8282620001cd565b919050565b600067ffffffffffffffff82111562000240576200023f6200019e565b5b6200024b826200018d565b9050602081019050919050565b60005b83811015620002785780820151818401526020810190506200025b565b60008484015250505050565b60006200029b620002958462000222565b62000203565b905082815260208101848484011115620002ba57620002b962000188565b5b620002c784828562000258565b509392505050565b600082601f830112620002e757620002e662000183565b5b8151620002f984826020860162000284565b91505092915050565b600080604083850312156200031c576200031b62000179565b5b600083015167ffffffffffffffff8111156200033d576200033c6200017e565b5b6200034b85828601620002cf565b925050602083015167ffffffffffffffff8111156200036f576200036e6200017e565b5b6200037d85828601620002cf565b9150509250929050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680620003da57607f821691505b602082108103620003f057620003ef62000392565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026200045a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826200041b565b6200046686836200041b565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000620004b3620004ad620004a7846200047e565b62000488565b6200047e565b9050919050565b6000819050919050565b620004cf8362000492565b620004e7620004de82620004ba565b84845462000428565b825550505050565b600090565b620004fe620004ef565b6200050b818484620004c4565b505050565b5b81811015620005335762000527600082620004f4565b60018101905062000511565b5050565b601f82111562000582576200054c81620003f6565b62000557846200040b565b8101602085101562000567578190505b6200057f62000576856200040b565b83018262000510565b50505b505050565b600082821c905092915050565b6000620005a76000198460080262000587565b1980831691505092915050565b6000620005c2838362000594565b9150826002028217905092915050565b620005dd8262000387565b67ffffffffffffffff811115620005f957620005f86200019e565b5b620006058254620003c1565b6200061282828562000537565b600060209050601f8311600181146200064a576000841562000635578287015190505b620006418582620005b4565b865550620006b1565b601f1984166200065a86620003f6565b60005b8281101562000684578489015182556001820191506020850194506020810190506200065d565b86831015620006a45784890151620006a0601f89168262000594565b8355505b6001600288020188555050505b505050505050565b600081905092915050565b60008190508160005260206000209050919050565b60008154620006e881620003c1565b620006f48186620006b9565b9450600182166000811462000712576001811462000728576200075f565b60ff19831686528115158202860193506200075f565b6200073385620006c4565b60005b83811015620007575781548189015260018201915060208101905062000736565b838801955050505b50505092915050565b6000620007768284620006d9565b915081905092915050565b6000819050919050565b620007968162000781565b82525050565b620007a7816200047e565b82525050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620007da82620007ad565b9050919050565b620007ec81620007cd565b82525050565b600060a0820190506200080960008301886200078b565b6200081860208301876200078b565b6200082760408301866200078b565b6200083660608301856200079c565b620008456080830184620007e1565b9695505050505050565b60805160a05160c05160e0516120db620008a56000396000818161093501528181610a1901528181610aee01528181610ba101526110c40152600061090e015260006108da015260006108b401526120db6000f3fe6080604052600436106101025760003560e01c806370a08231116100955780639fa6dd35116100645780639fa6dd3514610358578063a9059cbb14610374578063d505accf146103b1578063dd62ed3e146103da578063f639187e1461041757610109565b806370a08231146102885780637ecebe00146102c557806395d89b41146103025780639de702581461032d57610109565b80632e1a7d4d116100d15780632e1a7d4d146101de578063313ce567146102075780633644e515146102325780634cf088d91461025d57610109565b806306fdde031461010e578063095ea7b31461013957806318160ddd1461017657806323b872dd146101a157610109565b3661010957005b600080fd5b34801561011a57600080fd5b50610123610454565b6040516101309190611421565b60405180910390f35b34801561014557600080fd5b50610160600480360381019061015b91906114eb565b6104e2565b60405161016d9190611546565b60405180910390f35b34801561018257600080fd5b5061018b6105d4565b6040516101989190611570565b60405180910390f35b3480156101ad57600080fd5b506101c860048036038101906101c3919061158b565b6105da565b6040516101d59190611546565b60405180910390f35b3480156101ea57600080fd5b50610205600480360381019061020091906115de565b610824565b005b34801561021357600080fd5b5061021c6108b2565b6040516102299190611627565b60405180910390f35b34801561023e57600080fd5b506102476108d6565b604051610254919061165b565b60405180910390f35b34801561026957600080fd5b50610272610933565b60405161027f91906116d5565b60405180910390f35b34801561029457600080fd5b506102af60048036038101906102aa91906116f0565b610957565b6040516102bc9190611570565b60405180910390f35b3480156102d157600080fd5b506102ec60048036038101906102e791906116f0565b61096f565b6040516102f99190611570565b60405180910390f35b34801561030e57600080fd5b50610317610987565b6040516103249190611421565b60405180910390f35b34801561033957600080fd5b50610342610a15565b60405161034f91906117db565b60405180910390f35b610372600480360381019061036d91906115de565b610ab0565b005b34801561038057600080fd5b5061039b600480360381019061039691906114eb565b610c8e565b6040516103a89190611546565b60405180910390f35b3480156103bd57600080fd5b506103d860048036038101906103d39190611855565b610da2565b005b3480156103e657600080fd5b5061040160048036038101906103fc91906118f7565b61109b565b60405161040e9190611570565b60405180910390f35b34801561042357600080fd5b5061043e600480360381019061043991906116f0565b6110c0565b60405161044b9190611570565b60405180910390f35b6000805461046190611966565b80601f016020809104026020016040519081016040528092919081815260200182805461048d90611966565b80156104da5780601f106104af576101008083540402835291602001916104da565b820191906000526020600020905b8154815290600101906020018083116104bd57829003601f168201915b505050505081565b600081600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516105c29190611570565b60405180910390a36001905092915050565b60025481565b600080600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461071057828161068f91906119c6565b600460008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b82600360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461075f91906119c6565b9250508190555082600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef856040516108109190611570565b60405180910390a360019150509392505050565b6000810361085e576040517f1f2a200500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108683382611165565b3373ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156108ae573d6000803e3d6000fd5b5050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60007f0000000000000000000000000000000000000000000000000000000000000000461461090c57610907611235565b61092e565b7f00000000000000000000000000000000000000000000000000000000000000005b905090565b7f000000000000000000000000000000000000000000000000000000000000000081565b60036020528060005260406000206000915090505481565b60056020528060005260406000206000915090505481565b6001805461099490611966565b80601f01602080910402602001604051908101604052809291908181526020018280546109c090611966565b8015610a0d5780601f106109e257610100808354040283529160200191610a0d565b820191906000526020600020905b8154815290600101906020018083116109f057829003601f168201915b505050505081565b60607f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16639de702586040518163ffffffff1660e01b8152600401600060405180830381865afa158015610a82573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610aab9190611b57565b905090565b60008103610aea576040517f1f2a200500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16639de702586040518163ffffffff1660e01b8152600401600060405180830381865afa158015610b57573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610b809190611b57565b600081518110610b9357610b92611ba0565b5b6020026020010151905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663026e402b83856040518363ffffffff1660e01b8152600401610bfa929190611bde565b6020604051808303816000875af1158015610c19573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c3d9190611c33565b905080610c7f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c7690611cac565b60405180910390fd5b610c8933846112c1565b505050565b600081600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610cdf91906119c6565b9250508190555081600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610d909190611570565b60405180910390a36001905092915050565b42841015610de5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ddc90611d18565b60405180910390fd5b60006001610df16108d6565b7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98a8a8a600560008f73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000815480929190600101919050558b604051602001610e7996959493929190611d38565b60405160208183030381529060405280519060200120604051602001610ea0929190611e11565b6040516020818303038152906040528051906020012085858560405160008152602001604052604051610ed69493929190611e48565b6020604051602081039080840390855afa158015610ef8573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614158015610f6c57508773ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b610fab576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fa290611ed9565b60405180910390fd5b85600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550508573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258760405161108a9190611570565b60405180910390a350505050505050565b6004602052816000526040600020602052806000526040600020600091509150505481565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166315049a5a30846040518363ffffffff1660e01b815260040161111d929190611ef9565b602060405180830381865afa15801561113a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061115e9190611f37565b9050919050565b80600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546111b491906119c6565b9250508190555080600260008282540392505081905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516112299190611570565b60405180910390a35050565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60006040516112679190612007565b60405180910390207fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc646306040516020016112a695949392919061201e565b60405160208183030381529060405280519060200120905090565b80600260008282546112d39190612071565b9250508190555080600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516113859190611570565b60405180910390a35050565b600081519050919050565b600082825260208201905092915050565b60005b838110156113cb5780820151818401526020810190506113b0565b60008484015250505050565b6000601f19601f8301169050919050565b60006113f382611391565b6113fd818561139c565b935061140d8185602086016113ad565b611416816113d7565b840191505092915050565b6000602082019050818103600083015261143b81846113e8565b905092915050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061148282611457565b9050919050565b61149281611477565b811461149d57600080fd5b50565b6000813590506114af81611489565b92915050565b6000819050919050565b6114c8816114b5565b81146114d357600080fd5b50565b6000813590506114e5816114bf565b92915050565b600080604083850312156115025761150161144d565b5b6000611510858286016114a0565b9250506020611521858286016114d6565b9150509250929050565b60008115159050919050565b6115408161152b565b82525050565b600060208201905061155b6000830184611537565b92915050565b61156a816114b5565b82525050565b60006020820190506115856000830184611561565b92915050565b6000806000606084860312156115a4576115a361144d565b5b60006115b2868287016114a0565b93505060206115c3868287016114a0565b92505060406115d4868287016114d6565b9150509250925092565b6000602082840312156115f4576115f361144d565b5b6000611602848285016114d6565b91505092915050565b600060ff82169050919050565b6116218161160b565b82525050565b600060208201905061163c6000830184611618565b92915050565b6000819050919050565b61165581611642565b82525050565b6000602082019050611670600083018461164c565b92915050565b6000819050919050565b600061169b61169661169184611457565b611676565b611457565b9050919050565b60006116ad82611680565b9050919050565b60006116bf826116a2565b9050919050565b6116cf816116b4565b82525050565b60006020820190506116ea60008301846116c6565b92915050565b6000602082840312156117065761170561144d565b5b6000611714848285016114a0565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61175281611477565b82525050565b60006117648383611749565b60208301905092915050565b6000602082019050919050565b60006117888261171d565b6117928185611728565b935061179d83611739565b8060005b838110156117ce5781516117b58882611758565b97506117c083611770565b9250506001810190506117a1565b5085935050505092915050565b600060208201905081810360008301526117f5818461177d565b905092915050565b6118068161160b565b811461181157600080fd5b50565b600081359050611823816117fd565b92915050565b61183281611642565b811461183d57600080fd5b50565b60008135905061184f81611829565b92915050565b600080600080600080600060e0888a0312156118745761187361144d565b5b60006118828a828b016114a0565b97505060206118938a828b016114a0565b96505060406118a48a828b016114d6565b95505060606118b58a828b016114d6565b94505060806118c68a828b01611814565b93505060a06118d78a828b01611840565b92505060c06118e88a828b01611840565b91505092959891949750929550565b6000806040838503121561190e5761190d61144d565b5b600061191c858286016114a0565b925050602061192d858286016114a0565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061197e57607f821691505b60208210810361199157611990611937565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006119d1826114b5565b91506119dc836114b5565b92508282039050818111156119f4576119f3611997565b5b92915050565b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b611a37826113d7565b810181811067ffffffffffffffff82111715611a5657611a556119ff565b5b80604052505050565b6000611a69611443565b9050611a758282611a2e565b919050565b600067ffffffffffffffff821115611a9557611a946119ff565b5b602082029050602081019050919050565b600080fd5b600081519050611aba81611489565b92915050565b6000611ad3611ace84611a7a565b611a5f565b90508083825260208201905060208402830185811115611af657611af5611aa6565b5b835b81811015611b1f5780611b0b8882611aab565b845260208401935050602081019050611af8565b5050509392505050565b600082601f830112611b3e57611b3d6119fa565b5b8151611b4e848260208601611ac0565b91505092915050565b600060208284031215611b6d57611b6c61144d565b5b600082015167ffffffffffffffff811115611b8b57611b8a611452565b5b611b9784828501611b29565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b611bd881611477565b82525050565b6000604082019050611bf36000830185611bcf565b611c006020830184611561565b9392505050565b611c108161152b565b8114611c1b57600080fd5b50565b600081519050611c2d81611c07565b92915050565b600060208284031215611c4957611c4861144d565b5b6000611c5784828501611c1e565b91505092915050565b7f4661696c656420746f2064656c65676174650000000000000000000000000000600082015250565b6000611c9660128361139c565b9150611ca182611c60565b602082019050919050565b60006020820190508181036000830152611cc581611c89565b9050919050565b7f5045524d49545f444541444c494e455f45585049524544000000000000000000600082015250565b6000611d0260178361139c565b9150611d0d82611ccc565b602082019050919050565b60006020820190508181036000830152611d3181611cf5565b9050919050565b600060c082019050611d4d600083018961164c565b611d5a6020830188611bcf565b611d676040830187611bcf565b611d746060830186611561565b611d816080830185611561565b611d8e60a0830184611561565b979650505050505050565b600081905092915050565b7f1901000000000000000000000000000000000000000000000000000000000000600082015250565b6000611dda600283611d99565b9150611de582611da4565b600282019050919050565b6000819050919050565b611e0b611e0682611642565b611df0565b82525050565b6000611e1c82611dcd565b9150611e288285611dfa565b602082019150611e388284611dfa565b6020820191508190509392505050565b6000608082019050611e5d600083018761164c565b611e6a6020830186611618565b611e77604083018561164c565b611e84606083018461164c565b95945050505050565b7f494e56414c49445f5349474e4552000000000000000000000000000000000000600082015250565b6000611ec3600e8361139c565b9150611ece82611e8d565b602082019050919050565b60006020820190508181036000830152611ef281611eb6565b9050919050565b6000604082019050611f0e6000830185611bcf565b611f1b6020830184611bcf565b9392505050565b600081519050611f31816114bf565b92915050565b600060208284031215611f4d57611f4c61144d565b5b6000611f5b84828501611f22565b91505092915050565b600081905092915050565b60008190508160005260206000209050919050565b60008154611f9181611966565b611f9b8186611f64565b94506001821660008114611fb65760018114611fcb57611ffe565b60ff1983168652811515820286019350611ffe565b611fd485611f6f565b60005b83811015611ff657815481890152600182019150602081019050611fd7565b838801955050505b50505092915050565b60006120138284611f84565b915081905092915050565b600060a082019050612033600083018861164c565b612040602083018761164c565b61204d604083018661164c565b61205a6060830185611561565b6120676080830184611bcf565b9695505050505050565b600061207c826114b5565b9150612087836114b5565b925082820190508082111561209f5761209e611997565b5b9291505056fea264697066735822122075d6134187e3cf2809c948798ceda54b514a12b923fe6cb1ebe5e7915bda3c6a64736f6c63430008130033",
}

// LiquidStakingABI is the input ABI used to generate the binding from.
// Deprecated: Use LiquidStakingMetaData.ABI instead.
var LiquidStakingABI = LiquidStakingMetaData.ABI

// LiquidStakingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LiquidStakingMetaData.Bin instead.
var LiquidStakingBin = LiquidStakingMetaData.Bin

// DeployLiquidStaking deploys a new Ethereum contract, binding an instance of LiquidStaking to it.
func DeployLiquidStaking(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string) (common.Address, *types.Transaction, *LiquidStaking, error) {
	parsed, err := LiquidStakingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LiquidStakingBin), backend, _name, _symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LiquidStaking{LiquidStakingCaller: LiquidStakingCaller{contract: contract}, LiquidStakingTransactor: LiquidStakingTransactor{contract: contract}, LiquidStakingFilterer: LiquidStakingFilterer{contract: contract}}, nil
}

// LiquidStaking is an auto generated Go binding around an Ethereum contract.
type LiquidStaking struct {
	LiquidStakingCaller     // Read-only binding to the contract
	LiquidStakingTransactor // Write-only binding to the contract
	LiquidStakingFilterer   // Log filterer for contract events
}

// LiquidStakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type LiquidStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidStakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LiquidStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidStakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LiquidStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidStakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LiquidStakingSession struct {
	Contract     *LiquidStaking    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LiquidStakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LiquidStakingCallerSession struct {
	Contract *LiquidStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// LiquidStakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LiquidStakingTransactorSession struct {
	Contract     *LiquidStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// LiquidStakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type LiquidStakingRaw struct {
	Contract *LiquidStaking // Generic contract binding to access the raw methods on
}

// LiquidStakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LiquidStakingCallerRaw struct {
	Contract *LiquidStakingCaller // Generic read-only contract binding to access the raw methods on
}

// LiquidStakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LiquidStakingTransactorRaw struct {
	Contract *LiquidStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLiquidStaking creates a new instance of LiquidStaking, bound to a specific deployed contract.
func NewLiquidStaking(address common.Address, backend bind.ContractBackend) (*LiquidStaking, error) {
	contract, err := bindLiquidStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LiquidStaking{LiquidStakingCaller: LiquidStakingCaller{contract: contract}, LiquidStakingTransactor: LiquidStakingTransactor{contract: contract}, LiquidStakingFilterer: LiquidStakingFilterer{contract: contract}}, nil
}

// NewLiquidStakingCaller creates a new read-only instance of LiquidStaking, bound to a specific deployed contract.
func NewLiquidStakingCaller(address common.Address, caller bind.ContractCaller) (*LiquidStakingCaller, error) {
	contract, err := bindLiquidStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidStakingCaller{contract: contract}, nil
}

// NewLiquidStakingTransactor creates a new write-only instance of LiquidStaking, bound to a specific deployed contract.
func NewLiquidStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*LiquidStakingTransactor, error) {
	contract, err := bindLiquidStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidStakingTransactor{contract: contract}, nil
}

// NewLiquidStakingFilterer creates a new log filterer instance of LiquidStaking, bound to a specific deployed contract.
func NewLiquidStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*LiquidStakingFilterer, error) {
	contract, err := bindLiquidStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LiquidStakingFilterer{contract: contract}, nil
}

// bindLiquidStaking binds a generic wrapper to an already deployed contract.
func bindLiquidStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LiquidStakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquidStaking *LiquidStakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquidStaking.Contract.LiquidStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquidStaking *LiquidStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidStaking.Contract.LiquidStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquidStaking *LiquidStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquidStaking.Contract.LiquidStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquidStaking *LiquidStakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquidStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquidStaking *LiquidStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquidStaking *LiquidStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquidStaking.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_LiquidStaking *LiquidStakingCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_LiquidStaking *LiquidStakingSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _LiquidStaking.Contract.DOMAINSEPARATOR(&_LiquidStaking.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_LiquidStaking *LiquidStakingCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _LiquidStaking.Contract.DOMAINSEPARATOR(&_LiquidStaking.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_LiquidStaking *LiquidStakingCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_LiquidStaking *LiquidStakingSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _LiquidStaking.Contract.Allowance(&_LiquidStaking.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_LiquidStaking *LiquidStakingCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _LiquidStaking.Contract.Allowance(&_LiquidStaking.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_LiquidStaking *LiquidStakingCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_LiquidStaking *LiquidStakingSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _LiquidStaking.Contract.BalanceOf(&_LiquidStaking.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_LiquidStaking *LiquidStakingCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _LiquidStaking.Contract.BalanceOf(&_LiquidStaking.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LiquidStaking *LiquidStakingCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LiquidStaking *LiquidStakingSession) Decimals() (uint8, error) {
	return _LiquidStaking.Contract.Decimals(&_LiquidStaking.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LiquidStaking *LiquidStakingCallerSession) Decimals() (uint8, error) {
	return _LiquidStaking.Contract.Decimals(&_LiquidStaking.CallOpts)
}

// GetActiveValidators is a free data retrieval call binding the contract method 0x9de70258.
//
// Solidity: function getActiveValidators() view returns(address[])
func (_LiquidStaking *LiquidStakingCaller) GetActiveValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "getActiveValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetActiveValidators is a free data retrieval call binding the contract method 0x9de70258.
//
// Solidity: function getActiveValidators() view returns(address[])
func (_LiquidStaking *LiquidStakingSession) GetActiveValidators() ([]common.Address, error) {
	return _LiquidStaking.Contract.GetActiveValidators(&_LiquidStaking.CallOpts)
}

// GetActiveValidators is a free data retrieval call binding the contract method 0x9de70258.
//
// Solidity: function getActiveValidators() view returns(address[])
func (_LiquidStaking *LiquidStakingCallerSession) GetActiveValidators() ([]common.Address, error) {
	return _LiquidStaking.Contract.GetActiveValidators(&_LiquidStaking.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LiquidStaking *LiquidStakingCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LiquidStaking *LiquidStakingSession) Name() (string, error) {
	return _LiquidStaking.Contract.Name(&_LiquidStaking.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LiquidStaking *LiquidStakingCallerSession) Name() (string, error) {
	return _LiquidStaking.Contract.Name(&_LiquidStaking.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_LiquidStaking *LiquidStakingCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_LiquidStaking *LiquidStakingSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _LiquidStaking.Contract.Nonces(&_LiquidStaking.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_LiquidStaking *LiquidStakingCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _LiquidStaking.Contract.Nonces(&_LiquidStaking.CallOpts, arg0)
}

// Staking is a free data retrieval call binding the contract method 0x4cf088d9.
//
// Solidity: function staking() view returns(address)
func (_LiquidStaking *LiquidStakingCaller) Staking(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "staking")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Staking is a free data retrieval call binding the contract method 0x4cf088d9.
//
// Solidity: function staking() view returns(address)
func (_LiquidStaking *LiquidStakingSession) Staking() (common.Address, error) {
	return _LiquidStaking.Contract.Staking(&_LiquidStaking.CallOpts)
}

// Staking is a free data retrieval call binding the contract method 0x4cf088d9.
//
// Solidity: function staking() view returns(address)
func (_LiquidStaking *LiquidStakingCallerSession) Staking() (common.Address, error) {
	return _LiquidStaking.Contract.Staking(&_LiquidStaking.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LiquidStaking *LiquidStakingCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LiquidStaking *LiquidStakingSession) Symbol() (string, error) {
	return _LiquidStaking.Contract.Symbol(&_LiquidStaking.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LiquidStaking *LiquidStakingCallerSession) Symbol() (string, error) {
	return _LiquidStaking.Contract.Symbol(&_LiquidStaking.CallOpts)
}

// TotalDelegated is a free data retrieval call binding the contract method 0xf639187e.
//
// Solidity: function totalDelegated(address validatorAddress) view returns(uint256 amount)
func (_LiquidStaking *LiquidStakingCaller) TotalDelegated(opts *bind.CallOpts, validatorAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "totalDelegated", validatorAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDelegated is a free data retrieval call binding the contract method 0xf639187e.
//
// Solidity: function totalDelegated(address validatorAddress) view returns(uint256 amount)
func (_LiquidStaking *LiquidStakingSession) TotalDelegated(validatorAddress common.Address) (*big.Int, error) {
	return _LiquidStaking.Contract.TotalDelegated(&_LiquidStaking.CallOpts, validatorAddress)
}

// TotalDelegated is a free data retrieval call binding the contract method 0xf639187e.
//
// Solidity: function totalDelegated(address validatorAddress) view returns(uint256 amount)
func (_LiquidStaking *LiquidStakingCallerSession) TotalDelegated(validatorAddress common.Address) (*big.Int, error) {
	return _LiquidStaking.Contract.TotalDelegated(&_LiquidStaking.CallOpts, validatorAddress)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LiquidStaking *LiquidStakingCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LiquidStaking *LiquidStakingSession) TotalSupply() (*big.Int, error) {
	return _LiquidStaking.Contract.TotalSupply(&_LiquidStaking.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LiquidStaking *LiquidStakingCallerSession) TotalSupply() (*big.Int, error) {
	return _LiquidStaking.Contract.TotalSupply(&_LiquidStaking.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LiquidStaking *LiquidStakingTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LiquidStaking.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LiquidStaking *LiquidStakingSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LiquidStaking.Contract.Approve(&_LiquidStaking.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LiquidStaking *LiquidStakingTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LiquidStaking.Contract.Approve(&_LiquidStaking.TransactOpts, spender, amount)
}

// Delegate is a paid mutator transaction binding the contract method 0x9fa6dd35.
//
// Solidity: function delegate(uint256 amount) payable returns()
func (_LiquidStaking *LiquidStakingTransactor) Delegate(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LiquidStaking.contract.Transact(opts, "delegate", amount)
}

// Delegate is a paid mutator transaction binding the contract method 0x9fa6dd35.
//
// Solidity: function delegate(uint256 amount) payable returns()
func (_LiquidStaking *LiquidStakingSession) Delegate(amount *big.Int) (*types.Transaction, error) {
	return _LiquidStaking.Contract.Delegate(&_LiquidStaking.TransactOpts, amount)
}

// Delegate is a paid mutator transaction binding the contract method 0x9fa6dd35.
//
// Solidity: function delegate(uint256 amount) payable returns()
func (_LiquidStaking *LiquidStakingTransactorSession) Delegate(amount *big.Int) (*types.Transaction, error) {
	return _LiquidStaking.Contract.Delegate(&_LiquidStaking.TransactOpts, amount)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_LiquidStaking *LiquidStakingTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _LiquidStaking.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_LiquidStaking *LiquidStakingSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _LiquidStaking.Contract.Permit(&_LiquidStaking.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_LiquidStaking *LiquidStakingTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _LiquidStaking.Contract.Permit(&_LiquidStaking.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_LiquidStaking *LiquidStakingTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LiquidStaking.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_LiquidStaking *LiquidStakingSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LiquidStaking.Contract.Transfer(&_LiquidStaking.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_LiquidStaking *LiquidStakingTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LiquidStaking.Contract.Transfer(&_LiquidStaking.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_LiquidStaking *LiquidStakingTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LiquidStaking.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_LiquidStaking *LiquidStakingSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LiquidStaking.Contract.TransferFrom(&_LiquidStaking.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_LiquidStaking *LiquidStakingTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LiquidStaking.Contract.TransferFrom(&_LiquidStaking.TransactOpts, from, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_LiquidStaking *LiquidStakingTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LiquidStaking.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_LiquidStaking *LiquidStakingSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _LiquidStaking.Contract.Withdraw(&_LiquidStaking.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_LiquidStaking *LiquidStakingTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _LiquidStaking.Contract.Withdraw(&_LiquidStaking.TransactOpts, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LiquidStaking *LiquidStakingTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidStaking.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LiquidStaking *LiquidStakingSession) Receive() (*types.Transaction, error) {
	return _LiquidStaking.Contract.Receive(&_LiquidStaking.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LiquidStaking *LiquidStakingTransactorSession) Receive() (*types.Transaction, error) {
	return _LiquidStaking.Contract.Receive(&_LiquidStaking.TransactOpts)
}

// LiquidStakingApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the LiquidStaking contract.
type LiquidStakingApprovalIterator struct {
	Event *LiquidStakingApproval // Event containing the contract specifics and raw log

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
func (it *LiquidStakingApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidStakingApproval)
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
		it.Event = new(LiquidStakingApproval)
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
func (it *LiquidStakingApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidStakingApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidStakingApproval represents a Approval event raised by the LiquidStaking contract.
type LiquidStakingApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LiquidStaking *LiquidStakingFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*LiquidStakingApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LiquidStaking.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &LiquidStakingApprovalIterator{contract: _LiquidStaking.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LiquidStaking *LiquidStakingFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LiquidStakingApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LiquidStaking.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidStakingApproval)
				if err := _LiquidStaking.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_LiquidStaking *LiquidStakingFilterer) ParseApproval(log types.Log) (*LiquidStakingApproval, error) {
	event := new(LiquidStakingApproval)
	if err := _LiquidStaking.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidStakingDataIterator is returned from FilterData and is used to iterate over the raw logs and unpacked data for Data events raised by the LiquidStaking contract.
type LiquidStakingDataIterator struct {
	Event *LiquidStakingData // Event containing the contract specifics and raw log

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
func (it *LiquidStakingDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidStakingData)
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
		it.Event = new(LiquidStakingData)
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
func (it *LiquidStakingDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidStakingDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidStakingData represents a Data event raised by the LiquidStaking contract.
type LiquidStakingData struct {
	Data []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterData is a free log retrieval operation binding the contract event 0x0b76c48be4e2908f4c9d4eabaf7538e91577fd9ae26db46693fa8d861c6a42fb.
//
// Solidity: event Data(bytes data)
func (_LiquidStaking *LiquidStakingFilterer) FilterData(opts *bind.FilterOpts) (*LiquidStakingDataIterator, error) {

	logs, sub, err := _LiquidStaking.contract.FilterLogs(opts, "Data")
	if err != nil {
		return nil, err
	}
	return &LiquidStakingDataIterator{contract: _LiquidStaking.contract, event: "Data", logs: logs, sub: sub}, nil
}

// WatchData is a free log subscription operation binding the contract event 0x0b76c48be4e2908f4c9d4eabaf7538e91577fd9ae26db46693fa8d861c6a42fb.
//
// Solidity: event Data(bytes data)
func (_LiquidStaking *LiquidStakingFilterer) WatchData(opts *bind.WatchOpts, sink chan<- *LiquidStakingData) (event.Subscription, error) {

	logs, sub, err := _LiquidStaking.contract.WatchLogs(opts, "Data")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidStakingData)
				if err := _LiquidStaking.contract.UnpackLog(event, "Data", log); err != nil {
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

// ParseData is a log parse operation binding the contract event 0x0b76c48be4e2908f4c9d4eabaf7538e91577fd9ae26db46693fa8d861c6a42fb.
//
// Solidity: event Data(bytes data)
func (_LiquidStaking *LiquidStakingFilterer) ParseData(log types.Log) (*LiquidStakingData, error) {
	event := new(LiquidStakingData)
	if err := _LiquidStaking.contract.UnpackLog(event, "Data", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidStakingSuccessIterator is returned from FilterSuccess and is used to iterate over the raw logs and unpacked data for Success events raised by the LiquidStaking contract.
type LiquidStakingSuccessIterator struct {
	Event *LiquidStakingSuccess // Event containing the contract specifics and raw log

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
func (it *LiquidStakingSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidStakingSuccess)
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
		it.Event = new(LiquidStakingSuccess)
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
func (it *LiquidStakingSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidStakingSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidStakingSuccess represents a Success event raised by the LiquidStaking contract.
type LiquidStakingSuccess struct {
	Success bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSuccess is a free log retrieval operation binding the contract event 0x3b0a8ddef325df2bfdfa6b430ae4c8421841cd135bfa8fb5e432f200787520bb.
//
// Solidity: event Success(bool indexed success)
func (_LiquidStaking *LiquidStakingFilterer) FilterSuccess(opts *bind.FilterOpts, success []bool) (*LiquidStakingSuccessIterator, error) {

	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}

	logs, sub, err := _LiquidStaking.contract.FilterLogs(opts, "Success", successRule)
	if err != nil {
		return nil, err
	}
	return &LiquidStakingSuccessIterator{contract: _LiquidStaking.contract, event: "Success", logs: logs, sub: sub}, nil
}

// WatchSuccess is a free log subscription operation binding the contract event 0x3b0a8ddef325df2bfdfa6b430ae4c8421841cd135bfa8fb5e432f200787520bb.
//
// Solidity: event Success(bool indexed success)
func (_LiquidStaking *LiquidStakingFilterer) WatchSuccess(opts *bind.WatchOpts, sink chan<- *LiquidStakingSuccess, success []bool) (event.Subscription, error) {

	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}

	logs, sub, err := _LiquidStaking.contract.WatchLogs(opts, "Success", successRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidStakingSuccess)
				if err := _LiquidStaking.contract.UnpackLog(event, "Success", log); err != nil {
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

// ParseSuccess is a log parse operation binding the contract event 0x3b0a8ddef325df2bfdfa6b430ae4c8421841cd135bfa8fb5e432f200787520bb.
//
// Solidity: event Success(bool indexed success)
func (_LiquidStaking *LiquidStakingFilterer) ParseSuccess(log types.Log) (*LiquidStakingSuccess, error) {
	event := new(LiquidStakingSuccess)
	if err := _LiquidStaking.contract.UnpackLog(event, "Success", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidStakingTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the LiquidStaking contract.
type LiquidStakingTransferIterator struct {
	Event *LiquidStakingTransfer // Event containing the contract specifics and raw log

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
func (it *LiquidStakingTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidStakingTransfer)
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
		it.Event = new(LiquidStakingTransfer)
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
func (it *LiquidStakingTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidStakingTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidStakingTransfer represents a Transfer event raised by the LiquidStaking contract.
type LiquidStakingTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LiquidStaking *LiquidStakingFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LiquidStakingTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LiquidStaking.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LiquidStakingTransferIterator{contract: _LiquidStaking.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LiquidStaking *LiquidStakingFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LiquidStakingTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LiquidStaking.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidStakingTransfer)
				if err := _LiquidStaking.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_LiquidStaking *LiquidStakingFilterer) ParseTransfer(log types.Log) (*LiquidStakingTransfer, error) {
	event := new(LiquidStakingTransfer)
	if err := _LiquidStaking.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
