// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polaris

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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b506040516200231e3803806200231e833981810160405281019062000037919062000357565b8181601233806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3508260019081620000e9919062000627565b508160029081620000fb919062000627565b508060ff1660808160ff16815250504660a08181525050620001226200013460201b60201c565b60c081815250505050505050620008a4565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6001604051620001689190620007bd565b60405180910390207fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc64630604051602001620001a995949392919062000847565b60405160208183030381529060405280519060200120905090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6200022d82620001e2565b810181811067ffffffffffffffff821117156200024f576200024e620001f3565b5b80604052505050565b600062000264620001c4565b905062000272828262000222565b919050565b600067ffffffffffffffff821115620002955762000294620001f3565b5b620002a082620001e2565b9050602081019050919050565b60005b83811015620002cd578082015181840152602081019050620002b0565b60008484015250505050565b6000620002f0620002ea8462000277565b62000258565b9050828152602081018484840111156200030f576200030e620001dd565b5b6200031c848285620002ad565b509392505050565b600082601f8301126200033c576200033b620001d8565b5b81516200034e848260208601620002d9565b91505092915050565b60008060408385031215620003715762000370620001ce565b5b600083015167ffffffffffffffff811115620003925762000391620001d3565b5b620003a08582860162000324565b925050602083015167ffffffffffffffff811115620003c457620003c3620001d3565b5b620003d28582860162000324565b9150509250929050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200042f57607f821691505b602082108103620004455762000444620003e7565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620004af7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000470565b620004bb868362000470565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b60006200050862000502620004fc84620004d3565b620004dd565b620004d3565b9050919050565b6000819050919050565b6200052483620004e7565b6200053c62000533826200050f565b8484546200047d565b825550505050565b600090565b6200055362000544565b6200056081848462000519565b505050565b5b8181101562000588576200057c60008262000549565b60018101905062000566565b5050565b601f821115620005d757620005a1816200044b565b620005ac8462000460565b81016020851015620005bc578190505b620005d4620005cb8562000460565b83018262000565565b50505b505050565b600082821c905092915050565b6000620005fc60001984600802620005dc565b1980831691505092915050565b6000620006178383620005e9565b9150826002028217905092915050565b6200063282620003dc565b67ffffffffffffffff8111156200064e576200064d620001f3565b5b6200065a825462000416565b620006678282856200058c565b600060209050601f8311600181146200069f57600084156200068a578287015190505b62000696858262000609565b86555062000706565b601f198416620006af866200044b565b60005b82811015620006d957848901518255600182019150602085019450602081019050620006b2565b86831015620006f95784890151620006f5601f891682620005e9565b8355505b6001600288020188555050505b505050505050565b600081905092915050565b60008190508160005260206000209050919050565b600081546200073d8162000416565b6200074981866200070e565b945060018216600081146200076757600181146200077d57620007b4565b60ff1983168652811515820286019350620007b4565b620007888562000719565b60005b83811015620007ac578154818901526001820191506020810190506200078b565b838801955050505b50505092915050565b6000620007cb82846200072e565b915081905092915050565b6000819050919050565b620007eb81620007d6565b82525050565b620007fc81620004d3565b82525050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006200082f8262000802565b9050919050565b620008418162000822565b82525050565b600060a0820190506200085e6000830188620007e0565b6200086d6020830187620007e0565b6200087c6040830186620007e0565b6200088b6060830185620007f1565b6200089a608083018462000836565b9695505050505050565b60805160a05160c051611a4a620008d46000396000610775015260006107410152600061071b0152611a4a6000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80637ecebe0011610097578063a9059cbb11610066578063a9059cbb146102b1578063d505accf146102e1578063dd62ed3e146102fd578063f2fde38b1461032d57610100565b80637ecebe00146102295780638da5cb5b1461025957806395d89b41146102775780639dc29fac1461029557610100565b8063313ce567116100d3578063313ce567146101a15780633644e515146101bf57806340c10f19146101dd57806370a08231146101f957610100565b806306fdde0314610105578063095ea7b31461012357806318160ddd1461015357806323b872dd14610171575b600080fd5b61010d610349565b60405161011a91906111cd565b60405180910390f35b61013d60048036038101906101389190611288565b6103d7565b60405161014a91906112e3565b60405180910390f35b61015b6104c9565b604051610168919061130d565b60405180910390f35b61018b60048036038101906101869190611328565b6104cf565b60405161019891906112e3565b60405180910390f35b6101a9610719565b6040516101b69190611397565b60405180910390f35b6101c761073d565b6040516101d491906113cb565b60405180910390f35b6101f760048036038101906101f29190611288565b61079a565b005b610213600480360381019061020e91906113e6565b610836565b604051610220919061130d565b60405180910390f35b610243600480360381019061023e91906113e6565b61084e565b604051610250919061130d565b60405180910390f35b610261610866565b60405161026e9190611422565b60405180910390f35b61027f61088a565b60405161028c91906111cd565b60405180910390f35b6102af60048036038101906102aa9190611288565b610918565b005b6102cb60048036038101906102c69190611288565b6109b4565b6040516102d891906112e3565b60405180910390f35b6102fb60048036038101906102f69190611495565b610ac8565b005b61031760048036038101906103129190611537565b610dc1565b604051610324919061130d565b60405180910390f35b610347600480360381019061034291906113e6565b610de6565b005b60018054610356906115a6565b80601f0160208091040260200160405190810160405280929190818152602001828054610382906115a6565b80156103cf5780601f106103a4576101008083540402835291602001916103cf565b820191906000526020600020905b8154815290600101906020018083116103b257829003601f168201915b505050505081565b600081600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516104b7919061130d565b60405180910390a36001905092915050565b60035481565b600080600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146106055782816105849190611606565b600560008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b82600460008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546106549190611606565b9250508190555082600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef85604051610705919061130d565b60405180910390a360019150509392505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60007f000000000000000000000000000000000000000000000000000000000000000046146107735761076e610f11565b610795565b7f00000000000000000000000000000000000000000000000000000000000000005b905090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610828576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161081f90611686565b60405180910390fd5b6108328282610f9d565b5050565b60046020528060005260406000206000915090505481565b60066020528060005260406000206000915090505481565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60028054610897906115a6565b80601f01602080910402602001604051908101604052809291908181526020018280546108c3906115a6565b80156109105780601f106108e557610100808354040283529160200191610910565b820191906000526020600020905b8154815290600101906020018083116108f357829003601f168201915b505050505081565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161099d90611686565b60405180910390fd5b6109b0828261106d565b5050565b600081600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610a059190611606565b9250508190555081600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610ab6919061130d565b60405180910390a36001905092915050565b42841015610b0b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b02906116f2565b60405180910390fd5b60006001610b1761073d565b7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98a8a8a600660008f73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000815480929190600101919050558b604051602001610b9f96959493929190611712565b60405160208183030381529060405280519060200120604051602001610bc69291906117eb565b6040516020818303038152906040528051906020012085858560405160008152602001604052604051610bfc9493929190611822565b6020604051602081039080840390855afa158015610c1e573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614158015610c9257508773ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b610cd1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cc8906118b3565b60405180910390fd5b85600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550508573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92587604051610db0919061130d565b60405180910390a350505050505050565b6005602052816000526040600020602052806000526040600020600091509150505481565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610e74576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e6b90611686565b60405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a350565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6001604051610f439190611976565b60405180910390207fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc64630604051602001610f8295949392919061198d565b60405160208183030381529060405280519060200120905090565b8060036000828254610faf91906119e0565b9250508190555080600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051611061919061130d565b60405180910390a35050565b80600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546110bc9190611606565b9250508190555080600360008282540392505081905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051611131919061130d565b60405180910390a35050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561117757808201518184015260208101905061115c565b60008484015250505050565b6000601f19601f8301169050919050565b600061119f8261113d565b6111a98185611148565b93506111b9818560208601611159565b6111c281611183565b840191505092915050565b600060208201905081810360008301526111e78184611194565b905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061121f826111f4565b9050919050565b61122f81611214565b811461123a57600080fd5b50565b60008135905061124c81611226565b92915050565b6000819050919050565b61126581611252565b811461127057600080fd5b50565b6000813590506112828161125c565b92915050565b6000806040838503121561129f5761129e6111ef565b5b60006112ad8582860161123d565b92505060206112be85828601611273565b9150509250929050565b60008115159050919050565b6112dd816112c8565b82525050565b60006020820190506112f860008301846112d4565b92915050565b61130781611252565b82525050565b600060208201905061132260008301846112fe565b92915050565b600080600060608486031215611341576113406111ef565b5b600061134f8682870161123d565b93505060206113608682870161123d565b925050604061137186828701611273565b9150509250925092565b600060ff82169050919050565b6113918161137b565b82525050565b60006020820190506113ac6000830184611388565b92915050565b6000819050919050565b6113c5816113b2565b82525050565b60006020820190506113e060008301846113bc565b92915050565b6000602082840312156113fc576113fb6111ef565b5b600061140a8482850161123d565b91505092915050565b61141c81611214565b82525050565b60006020820190506114376000830184611413565b92915050565b6114468161137b565b811461145157600080fd5b50565b6000813590506114638161143d565b92915050565b611472816113b2565b811461147d57600080fd5b50565b60008135905061148f81611469565b92915050565b600080600080600080600060e0888a0312156114b4576114b36111ef565b5b60006114c28a828b0161123d565b97505060206114d38a828b0161123d565b96505060406114e48a828b01611273565b95505060606114f58a828b01611273565b94505060806115068a828b01611454565b93505060a06115178a828b01611480565b92505060c06115288a828b01611480565b91505092959891949750929550565b6000806040838503121561154e5761154d6111ef565b5b600061155c8582860161123d565b925050602061156d8582860161123d565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806115be57607f821691505b6020821081036115d1576115d0611577565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061161182611252565b915061161c83611252565b9250828203905081811115611634576116336115d7565b5b92915050565b7f554e415554484f52495a45440000000000000000000000000000000000000000600082015250565b6000611670600c83611148565b915061167b8261163a565b602082019050919050565b6000602082019050818103600083015261169f81611663565b9050919050565b7f5045524d49545f444541444c494e455f45585049524544000000000000000000600082015250565b60006116dc601783611148565b91506116e7826116a6565b602082019050919050565b6000602082019050818103600083015261170b816116cf565b9050919050565b600060c08201905061172760008301896113bc565b6117346020830188611413565b6117416040830187611413565b61174e60608301866112fe565b61175b60808301856112fe565b61176860a08301846112fe565b979650505050505050565b600081905092915050565b7f1901000000000000000000000000000000000000000000000000000000000000600082015250565b60006117b4600283611773565b91506117bf8261177e565b600282019050919050565b6000819050919050565b6117e56117e0826113b2565b6117ca565b82525050565b60006117f6826117a7565b915061180282856117d4565b60208201915061181282846117d4565b6020820191508190509392505050565b600060808201905061183760008301876113bc565b6118446020830186611388565b61185160408301856113bc565b61185e60608301846113bc565b95945050505050565b7f494e56414c49445f5349474e4552000000000000000000000000000000000000600082015250565b600061189d600e83611148565b91506118a882611867565b602082019050919050565b600060208201905081810360008301526118cc81611890565b9050919050565b600081905092915050565b60008190508160005260206000209050919050565b60008154611900816115a6565b61190a81866118d3565b94506001821660008114611925576001811461193a5761196d565b60ff198316865281151582028601935061196d565b611943856118de565b60005b8381101561196557815481890152600182019150602081019050611946565b838801955050505b50505092915050565b600061198282846118f3565b915081905092915050565b600060a0820190506119a260008301886113bc565b6119af60208301876113bc565b6119bc60408301866113bc565b6119c960608301856112fe565b6119d66080830184611413565b9695505050505050565b60006119eb82611252565b91506119f683611252565b9250828201905080821115611a0e57611a0d6115d7565b5b9291505056fea2646970667358221220f69b33c87c2ab3b8a88e7cd70e63bd0a31808882b0356f468f3005253353c95764736f6c63430008130033",
}

// PolarisERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use PolarisERC20MetaData.ABI instead.
var PolarisERC20ABI = PolarisERC20MetaData.ABI

// PolarisERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PolarisERC20MetaData.Bin instead.
var PolarisERC20Bin = PolarisERC20MetaData.Bin

// DeployPolarisERC20 deploys a new Ethereum contract, binding an instance of PolarisERC20 to it.
func DeployPolarisERC20(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string) (common.Address, *types.Transaction, *PolarisERC20, error) {
	parsed, err := PolarisERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PolarisERC20Bin), backend, name, symbol)
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
// Solidity: function allowance(address , address ) view returns(uint256)
func (_PolarisERC20 *PolarisERC20Caller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PolarisERC20.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_PolarisERC20 *PolarisERC20Session) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _PolarisERC20.Contract.Allowance(&_PolarisERC20.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_PolarisERC20 *PolarisERC20CallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _PolarisERC20.Contract.Allowance(&_PolarisERC20.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_PolarisERC20 *PolarisERC20Caller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PolarisERC20.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_PolarisERC20 *PolarisERC20Session) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _PolarisERC20.Contract.BalanceOf(&_PolarisERC20.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_PolarisERC20 *PolarisERC20CallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _PolarisERC20.Contract.BalanceOf(&_PolarisERC20.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
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
// Solidity: function decimals() view returns(uint8)
func (_PolarisERC20 *PolarisERC20Session) Decimals() (uint8, error) {
	return _PolarisERC20.Contract.Decimals(&_PolarisERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PolarisERC20 *PolarisERC20CallerSession) Decimals() (uint8, error) {
	return _PolarisERC20.Contract.Decimals(&_PolarisERC20.CallOpts)
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

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PolarisERC20 *PolarisERC20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PolarisERC20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PolarisERC20 *PolarisERC20Session) Owner() (common.Address, error) {
	return _PolarisERC20.Contract.Owner(&_PolarisERC20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PolarisERC20 *PolarisERC20CallerSession) Owner() (common.Address, error) {
	return _PolarisERC20.Contract.Owner(&_PolarisERC20.CallOpts)
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

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_PolarisERC20 *PolarisERC20Transactor) Burn(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PolarisERC20.contract.Transact(opts, "burn", from, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_PolarisERC20 *PolarisERC20Session) Burn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PolarisERC20.Contract.Burn(&_PolarisERC20.TransactOpts, from, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_PolarisERC20 *PolarisERC20TransactorSession) Burn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PolarisERC20.Contract.Burn(&_PolarisERC20.TransactOpts, from, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_PolarisERC20 *PolarisERC20Transactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PolarisERC20.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_PolarisERC20 *PolarisERC20Session) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PolarisERC20.Contract.Mint(&_PolarisERC20.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_PolarisERC20 *PolarisERC20TransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PolarisERC20.Contract.Mint(&_PolarisERC20.TransactOpts, to, amount)
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

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PolarisERC20 *PolarisERC20Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PolarisERC20.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PolarisERC20 *PolarisERC20Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PolarisERC20.Contract.TransferOwnership(&_PolarisERC20.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PolarisERC20 *PolarisERC20TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PolarisERC20.Contract.TransferOwnership(&_PolarisERC20.TransactOpts, newOwner)
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
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_PolarisERC20 *PolarisERC20Filterer) ParseApproval(log types.Log) (*PolarisERC20Approval, error) {
	event := new(PolarisERC20Approval)
	if err := _PolarisERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolarisERC20OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PolarisERC20 contract.
type PolarisERC20OwnershipTransferredIterator struct {
	Event *PolarisERC20OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PolarisERC20OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolarisERC20OwnershipTransferred)
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
		it.Event = new(PolarisERC20OwnershipTransferred)
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
func (it *PolarisERC20OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolarisERC20OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolarisERC20OwnershipTransferred represents a OwnershipTransferred event raised by the PolarisERC20 contract.
type PolarisERC20OwnershipTransferred struct {
	User     common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_PolarisERC20 *PolarisERC20Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, user []common.Address, newOwner []common.Address) (*PolarisERC20OwnershipTransferredIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PolarisERC20.contract.FilterLogs(opts, "OwnershipTransferred", userRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PolarisERC20OwnershipTransferredIterator{contract: _PolarisERC20.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_PolarisERC20 *PolarisERC20Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PolarisERC20OwnershipTransferred, user []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PolarisERC20.contract.WatchLogs(opts, "OwnershipTransferred", userRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolarisERC20OwnershipTransferred)
				if err := _PolarisERC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_PolarisERC20 *PolarisERC20Filterer) ParseOwnershipTransferred(log types.Log) (*PolarisERC20OwnershipTransferred, error) {
	event := new(PolarisERC20OwnershipTransferred)
	if err := _PolarisERC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_PolarisERC20 *PolarisERC20Filterer) ParseTransfer(log types.Log) (*PolarisERC20Transfer, error) {
	event := new(PolarisERC20Transfer)
	if err := _PolarisERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
