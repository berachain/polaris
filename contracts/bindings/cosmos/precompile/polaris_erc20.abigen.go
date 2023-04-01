// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package precompile

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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b50604051620021f2380380620021f28339818101604052810190620000389190620002ef565b8181601282600090816200004d9190620005bf565b5081600190816200005f9190620005bf565b508060ff1660808160ff16815250504660a0818152505062000086620000cc60201b60201c565b60c081815250505050503373ffffffffffffffffffffffffffffffffffffffff1660e08173ffffffffffffffffffffffffffffffffffffffff168152505050506200083c565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f600060405162000100919062000755565b60405180910390207fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6463060405160200162000141959493929190620007df565b60405160208183030381529060405280519060200120905090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b620001c5826200017a565b810181811067ffffffffffffffff82111715620001e757620001e66200018b565b5b80604052505050565b6000620001fc6200015c565b90506200020a8282620001ba565b919050565b600067ffffffffffffffff8211156200022d576200022c6200018b565b5b62000238826200017a565b9050602081019050919050565b60005b838110156200026557808201518184015260208101905062000248565b60008484015250505050565b60006200028862000282846200020f565b620001f0565b905082815260208101848484011115620002a757620002a662000175565b5b620002b484828562000245565b509392505050565b600082601f830112620002d457620002d362000170565b5b8151620002e684826020860162000271565b91505092915050565b6000806040838503121562000309576200030862000166565b5b600083015167ffffffffffffffff8111156200032a57620003296200016b565b5b6200033885828601620002bc565b925050602083015167ffffffffffffffff8111156200035c576200035b6200016b565b5b6200036a85828601620002bc565b9150509250929050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680620003c757607f821691505b602082108103620003dd57620003dc6200037f565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620004477fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000408565b62000453868362000408565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000620004a06200049a62000494846200046b565b62000475565b6200046b565b9050919050565b6000819050919050565b620004bc836200047f565b620004d4620004cb82620004a7565b84845462000415565b825550505050565b600090565b620004eb620004dc565b620004f8818484620004b1565b505050565b5b81811015620005205762000514600082620004e1565b600181019050620004fe565b5050565b601f8211156200056f576200053981620003e3565b6200054484620003f8565b8101602085101562000554578190505b6200056c6200056385620003f8565b830182620004fd565b50505b505050565b600082821c905092915050565b6000620005946000198460080262000574565b1980831691505092915050565b6000620005af838362000581565b9150826002028217905092915050565b620005ca8262000374565b67ffffffffffffffff811115620005e657620005e56200018b565b5b620005f28254620003ae565b620005ff82828562000524565b600060209050601f83116001811462000637576000841562000622578287015190505b6200062e8582620005a1565b8655506200069e565b601f1984166200064786620003e3565b60005b8281101562000671578489015182556001820191506020850194506020810190506200064a565b868310156200069157848901516200068d601f89168262000581565b8355505b6001600288020188555050505b505050505050565b600081905092915050565b60008190508160005260206000209050919050565b60008154620006d581620003ae565b620006e18186620006a6565b94506001821660008114620006ff576001811462000715576200074c565b60ff19831686528115158202860193506200074c565b6200072085620006b1565b60005b83811015620007445781548189015260018201915060208101905062000723565b838801955050505b50505092915050565b6000620007638284620006c6565b915081905092915050565b6000819050919050565b62000783816200076e565b82525050565b62000794816200046b565b82525050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620007c7826200079a565b9050919050565b620007d981620007ba565b82525050565b600060a082019050620007f6600083018862000778565b62000805602083018762000778565b62000814604083018662000778565b62000823606083018562000789565b620008326080830184620007ce565b9695505050505050565b60805160a05160c05160e0516119676200088b6000396000818161074c015281816107dd015281816108c5015261095601526000610725015260006106f1015260006106cb01526119676000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c806370a082311161008c578063a0712d6811610066578063a0712d6814610261578063a9059cbb1461027d578063d505accf146102ad578063dd62ed3e146102c9576100ea565b806370a08231146101e35780637ecebe001461021357806395d89b4114610243576100ea565b806323b872dd116100c857806323b872dd1461015b578063313ce5671461018b5780633644e515146101a957806342966c68146101c7576100ea565b806306fdde03146100ef578063095ea7b31461010d57806318160ddd1461013d575b600080fd5b6100f76102f9565b604051610104919061106c565b60405180910390f35b61012760048036038101906101229190611127565b610387565b6040516101349190611182565b60405180910390f35b610145610479565b60405161015291906111ac565b60405180910390f35b610175600480360381019061017091906111c7565b61047f565b6040516101829190611182565b60405180910390f35b6101936106c9565b6040516101a09190611236565b60405180910390f35b6101b16106ed565b6040516101be919061126a565b60405180910390f35b6101e160048036038101906101dc9190611285565b61074a565b005b6101fd60048036038101906101f891906112b2565b610805565b60405161020a91906111ac565b60405180910390f35b61022d600480360381019061022891906112b2565b61081d565b60405161023a91906111ac565b60405180910390f35b61024b610835565b604051610258919061106c565b60405180910390f35b61027b60048036038101906102769190611285565b6108c3565b005b61029760048036038101906102929190611127565b61097e565b6040516102a49190611182565b60405180910390f35b6102c760048036038101906102c29190611337565b610a92565b005b6102e360048036038101906102de91906113d9565b610d8b565b6040516102f091906111ac565b60405180910390f35b6000805461030690611448565b80601f016020809104026020016040519081016040528092919081815260200182805461033290611448565b801561037f5780601f106103545761010080835404028352916020019161037f565b820191906000526020600020905b81548152906001019060200180831161036257829003601f168201915b505050505081565b600081600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161046791906111ac565b60405180910390a36001905092915050565b60025481565b600080600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146105b557828161053491906114a8565b600460008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b82600360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461060491906114a8565b9250508190555082600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef856040516106b591906111ac565b60405180910390a360019150509392505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60007f000000000000000000000000000000000000000000000000000000000000000046146107235761071e610db0565b610745565b7f00000000000000000000000000000000000000000000000000000000000000005b905090565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107d8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107cf90611528565b60405180910390fd5b6108027f000000000000000000000000000000000000000000000000000000000000000082610e3c565b50565b60036020528060005260406000206000915090505481565b60056020528060005260406000206000915090505481565b6001805461084290611448565b80601f016020809104026020016040519081016040528092919081815260200182805461086e90611448565b80156108bb5780601f10610890576101008083540402835291602001916108bb565b820191906000526020600020905b81548152906001019060200180831161089e57829003601f168201915b505050505081565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610951576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161094890611594565b60405180910390fd5b61097b7f000000000000000000000000000000000000000000000000000000000000000082610f0c565b50565b600081600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546109cf91906114a8565b9250508190555081600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610a8091906111ac565b60405180910390a36001905092915050565b42841015610ad5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610acc90611600565b60405180910390fd5b60006001610ae16106ed565b7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98a8a8a600560008f73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000815480929190600101919050558b604051602001610b699695949392919061162f565b60405160208183030381529060405280519060200120604051602001610b90929190611708565b6040516020818303038152906040528051906020012085858560405160008152602001604052604051610bc6949392919061173f565b6020604051602081039080840390855afa158015610be8573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614158015610c5c57508773ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b610c9b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c92906117d0565b60405180910390fd5b85600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550508573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92587604051610d7a91906111ac565b60405180910390a350505050505050565b6004602052816000526040600020602052806000526040600020600091509150505481565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6000604051610de29190611893565b60405180910390207fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc64630604051602001610e219594939291906118aa565b60405160208183030381529060405280519060200120905090565b80600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610e8b91906114a8565b9250508190555080600260008282540392505081905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610f0091906111ac565b60405180910390a35050565b8060026000828254610f1e91906118fd565b9250508190555080600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610fd091906111ac565b60405180910390a35050565b600081519050919050565b600082825260208201905092915050565b60005b83811015611016578082015181840152602081019050610ffb565b60008484015250505050565b6000601f19601f8301169050919050565b600061103e82610fdc565b6110488185610fe7565b9350611058818560208601610ff8565b61106181611022565b840191505092915050565b600060208201905081810360008301526110868184611033565b905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006110be82611093565b9050919050565b6110ce816110b3565b81146110d957600080fd5b50565b6000813590506110eb816110c5565b92915050565b6000819050919050565b611104816110f1565b811461110f57600080fd5b50565b600081359050611121816110fb565b92915050565b6000806040838503121561113e5761113d61108e565b5b600061114c858286016110dc565b925050602061115d85828601611112565b9150509250929050565b60008115159050919050565b61117c81611167565b82525050565b60006020820190506111976000830184611173565b92915050565b6111a6816110f1565b82525050565b60006020820190506111c1600083018461119d565b92915050565b6000806000606084860312156111e0576111df61108e565b5b60006111ee868287016110dc565b93505060206111ff868287016110dc565b925050604061121086828701611112565b9150509250925092565b600060ff82169050919050565b6112308161121a565b82525050565b600060208201905061124b6000830184611227565b92915050565b6000819050919050565b61126481611251565b82525050565b600060208201905061127f600083018461125b565b92915050565b60006020828403121561129b5761129a61108e565b5b60006112a984828501611112565b91505092915050565b6000602082840312156112c8576112c761108e565b5b60006112d6848285016110dc565b91505092915050565b6112e88161121a565b81146112f357600080fd5b50565b600081359050611305816112df565b92915050565b61131481611251565b811461131f57600080fd5b50565b6000813590506113318161130b565b92915050565b600080600080600080600060e0888a0312156113565761135561108e565b5b60006113648a828b016110dc565b97505060206113758a828b016110dc565b96505060406113868a828b01611112565b95505060606113978a828b01611112565b94505060806113a88a828b016112f6565b93505060a06113b98a828b01611322565b92505060c06113ca8a828b01611322565b91505092959891949750929550565b600080604083850312156113f0576113ef61108e565b5b60006113fe858286016110dc565b925050602061140f858286016110dc565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061146057607f821691505b60208210810361147357611472611419565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006114b3826110f1565b91506114be836110f1565b92508282039050818111156114d6576114d5611479565b5b92915050565b7f6f6e6c79206465706c6f7965722063616e206275726e00000000000000000000600082015250565b6000611512601683610fe7565b915061151d826114dc565b602082019050919050565b6000602082019050818103600083015261154181611505565b9050919050565b7f6f6e6c79206465706c6f7965722063616e206d696e7400000000000000000000600082015250565b600061157e601683610fe7565b915061158982611548565b602082019050919050565b600060208201905081810360008301526115ad81611571565b9050919050565b7f5045524d49545f444541444c494e455f45585049524544000000000000000000600082015250565b60006115ea601783610fe7565b91506115f5826115b4565b602082019050919050565b60006020820190508181036000830152611619816115dd565b9050919050565b611629816110b3565b82525050565b600060c082019050611644600083018961125b565b6116516020830188611620565b61165e6040830187611620565b61166b606083018661119d565b611678608083018561119d565b61168560a083018461119d565b979650505050505050565b600081905092915050565b7f1901000000000000000000000000000000000000000000000000000000000000600082015250565b60006116d1600283611690565b91506116dc8261169b565b600282019050919050565b6000819050919050565b6117026116fd82611251565b6116e7565b82525050565b6000611713826116c4565b915061171f82856116f1565b60208201915061172f82846116f1565b6020820191508190509392505050565b6000608082019050611754600083018761125b565b6117616020830186611227565b61176e604083018561125b565b61177b606083018461125b565b95945050505050565b7f494e56414c49445f5349474e4552000000000000000000000000000000000000600082015250565b60006117ba600e83610fe7565b91506117c582611784565b602082019050919050565b600060208201905081810360008301526117e9816117ad565b9050919050565b600081905092915050565b60008190508160005260206000209050919050565b6000815461181d81611448565b61182781866117f0565b9450600182166000811461184257600181146118575761188a565b60ff198316865281151582028601935061188a565b611860856117fb565b60005b8381101561188257815481890152600182019150602081019050611863565b838801955050505b50505092915050565b600061189f8284611810565b915081905092915050565b600060a0820190506118bf600083018861125b565b6118cc602083018761125b565b6118d9604083018661125b565b6118e6606083018561119d565b6118f36080830184611620565b9695505050505050565b6000611908826110f1565b9150611913836110f1565b925082820190508082111561192b5761192a611479565b5b9291505056fea2646970667358221220676eefb47e629eebdf8c7cfc8406fc975438f6f3e5adb7e46859a4fb0c18bf2464736f6c63430008130033",
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

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_PolarisERC20 *PolarisERC20Transactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _PolarisERC20.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_PolarisERC20 *PolarisERC20Session) Burn(amount *big.Int) (*types.Transaction, error) {
	return _PolarisERC20.Contract.Burn(&_PolarisERC20.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_PolarisERC20 *PolarisERC20TransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _PolarisERC20.Contract.Burn(&_PolarisERC20.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_PolarisERC20 *PolarisERC20Transactor) Mint(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _PolarisERC20.contract.Transact(opts, "mint", amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_PolarisERC20 *PolarisERC20Session) Mint(amount *big.Int) (*types.Transaction, error) {
	return _PolarisERC20.Contract.Mint(&_PolarisERC20.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_PolarisERC20 *PolarisERC20TransactorSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _PolarisERC20.Contract.Mint(&_PolarisERC20.TransactOpts, amount)
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
