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
	Bin: "0x61010060405273d9a998cac66092748ffec7cfbd155aae1737c2ff73ffffffffffffffffffffffffffffffffffffffff1660e09073ffffffffffffffffffffffffffffffffffffffff1681525034801562000058575f80fd5b50604051620028583803806200285883398181016040528101906200007e9190620002f0565b81816012825f9081620000929190620005aa565b508160019081620000a49190620005aa565b508060ff1660808160ff16815250504660a08181525050620000cb620000dd60201b60201c565b60c08181525050505050505062000817565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f5f6040516200010f919062000736565b60405180910390207fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6463060405160200162000150959493929190620007bc565b60405160208183030381529060405280519060200120905090565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b620001cc8262000184565b810181811067ffffffffffffffff82111715620001ee57620001ed62000194565b5b80604052505050565b5f620002026200016b565b9050620002108282620001c1565b919050565b5f67ffffffffffffffff82111562000232576200023162000194565b5b6200023d8262000184565b9050602081019050919050565b5f5b83811015620002695780820151818401526020810190506200024c565b5f8484015250505050565b5f6200028a620002848462000215565b620001f7565b905082815260208101848484011115620002a957620002a862000180565b5b620002b68482856200024a565b509392505050565b5f82601f830112620002d557620002d46200017c565b5b8151620002e784826020860162000274565b91505092915050565b5f806040838503121562000309576200030862000174565b5b5f83015167ffffffffffffffff81111562000329576200032862000178565b5b6200033785828601620002be565b925050602083015167ffffffffffffffff8111156200035b576200035a62000178565b5b6200036985828601620002be565b9150509250929050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680620003c257607f821691505b602082108103620003d857620003d76200037d565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026200043c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620003ff565b620004488683620003ff565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f620004926200048c620004868462000460565b62000469565b62000460565b9050919050565b5f819050919050565b620004ad8362000472565b620004c5620004bc8262000499565b8484546200040b565b825550505050565b5f90565b620004db620004cd565b620004e8818484620004a2565b505050565b5b818110156200050f57620005035f82620004d1565b600181019050620004ee565b5050565b601f8211156200055e576200052881620003de565b6200053384620003f0565b8101602085101562000543578190505b6200055b6200055285620003f0565b830182620004ed565b50505b505050565b5f82821c905092915050565b5f620005805f198460080262000563565b1980831691505092915050565b5f6200059a83836200056f565b9150826002028217905092915050565b620005b58262000373565b67ffffffffffffffff811115620005d157620005d062000194565b5b620005dd8254620003aa565b620005ea82828562000513565b5f60209050601f83116001811462000620575f84156200060b578287015190505b6200061785826200058d565b86555062000686565b601f1984166200063086620003de565b5f5b82811015620006595784890151825560018201915060208501945060208101905062000632565b8683101562000679578489015162000675601f8916826200056f565b8355505b6001600288020188555050505b505050505050565b5f81905092915050565b5f819050815f5260205f209050919050565b5f8154620006b881620003aa565b620006c481866200068e565b9450600182165f8114620006e15760018114620006f7576200072d565b60ff19831686528115158202860193506200072d565b620007028562000698565b5f5b83811015620007255781548189015260018201915060208101905062000704565b838801955050505b50505092915050565b5f620007438284620006aa565b915081905092915050565b5f819050919050565b62000762816200074e565b82525050565b620007738162000460565b82525050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f620007a48262000779565b9050919050565b620007b68162000798565b82525050565b5f60a082019050620007d15f83018862000757565b620007e0602083018762000757565b620007ef604083018662000757565b620007fe606083018562000768565b6200080d6080830184620007ab565b9695505050505050565b60805160a05160c05160e051611ff1620008675f395f8181610907015281816109e301528181610ab201528181610b5f015261106601525f6108e001525f6108ac01525f6108870152611ff15ff3fe608060405260043610610101575f3560e01c806370a08231116100945780639fa6dd35116100635780639fa6dd351461034a578063a9059cbb14610366578063d505accf146103a2578063dd62ed3e146103ca578063f639187e1461040657610108565b806370a082311461027e5780637ecebe00146102ba57806395d89b41146102f65780639de702581461032057610108565b80632e1a7d4d116100d05780632e1a7d4d146101d8578063313ce567146102005780633644e5151461022a5780634cf088d91461025457610108565b806306fdde031461010c578063095ea7b31461013657806318160ddd1461017257806323b872dd1461019c57610108565b3661010857005b5f80fd5b348015610117575f80fd5b50610120610442565b60405161012d91906113af565b60405180910390f35b348015610141575f80fd5b5061015c6004803603810190610157919061146d565b6104cd565b60405161016991906114c5565b60405180910390f35b34801561017d575f80fd5b506101866105ba565b60405161019391906114ed565b60405180910390f35b3480156101a7575f80fd5b506101c260048036038101906101bd9190611506565b6105c0565b6040516101cf91906114c5565b60405180910390f35b3480156101e3575f80fd5b506101fe60048036038101906101f99190611556565b6107fb565b005b34801561020b575f80fd5b50610214610885565b604051610221919061159c565b60405180910390f35b348015610235575f80fd5b5061023e6108a9565b60405161024b91906115cd565b60405180910390f35b34801561025f575f80fd5b50610268610905565b6040516102759190611641565b60405180910390f35b348015610289575f80fd5b506102a4600480360381019061029f919061165a565b610929565b6040516102b191906114ed565b60405180910390f35b3480156102c5575f80fd5b506102e060048036038101906102db919061165a565b61093e565b6040516102ed91906114ed565b60405180910390f35b348015610301575f80fd5b5061030a610953565b60405161031791906113af565b60405180910390f35b34801561032b575f80fd5b506103346109df565b604051610341919061173c565b60405180910390f35b610364600480360381019061035f9190611556565b610a76565b005b348015610371575f80fd5b5061038c6004803603810190610387919061146d565b610c49565b60405161039991906114c5565b60405180910390f35b3480156103ad575f80fd5b506103c860048036038101906103c391906117b0565b610d56565b005b3480156103d5575f80fd5b506103f060048036038101906103eb919061184d565b611043565b6040516103fd91906114ed565b60405180910390f35b348015610411575f80fd5b5061042c6004803603810190610427919061165a565b611063565b60405161043991906114ed565b60405180910390f35b5f805461044e906118b8565b80601f016020809104026020016040519081016040528092919081815260200182805461047a906118b8565b80156104c55780601f1061049c576101008083540402835291602001916104c5565b820191905f5260205f20905b8154815290600101906020018083116104a857829003601f168201915b505050505081565b5f8160045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516105a891906114ed565b60405180910390a36001905092915050565b60025481565b5f8060045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146106ed5782816106709190611915565b60045f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055505b8260035f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546107399190611915565b925050819055508260035f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef856040516107e791906114ed565b60405180910390a360019150509392505050565b5f8103610834576040517f1f2a200500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61083e3382611105565b3373ffffffffffffffffffffffffffffffffffffffff166108fc8290811502906040515f60405180830381858888f19350505050158015610881573d5f803e3d5ffd5b5050565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f7f000000000000000000000000000000000000000000000000000000000000000046146108de576108d96111d0565b610900565b7f00000000000000000000000000000000000000000000000000000000000000005b905090565b7f000000000000000000000000000000000000000000000000000000000000000081565b6003602052805f5260405f205f915090505481565b6005602052805f5260405f205f915090505481565b60018054610960906118b8565b80601f016020809104026020016040519081016040528092919081815260200182805461098c906118b8565b80156109d75780601f106109ae576101008083540402835291602001916109d7565b820191905f5260205f20905b8154815290600101906020018083116109ba57829003601f168201915b505050505081565b60607f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16639de702586040518163ffffffff1660e01b81526004015f60405180830381865afa158015610a49573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190610a719190611a9c565b905090565b5f8103610aaf576040517f1f2a200500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16639de702586040518163ffffffff1660e01b81526004015f60405180830381865afa158015610b18573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190610b409190611a9c565b5f81518110610b5257610b51611ae3565b5b602002602001015190505f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663026e402b83856040518363ffffffff1660e01b8152600401610bb8929190611b1f565b6020604051808303815f875af1158015610bd4573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610bf89190611b70565b905080610c3a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c3190611be5565b60405180910390fd5b610c44338461125a565b505050565b5f8160035f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610c969190611915565b925050819055508160035f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610d4491906114ed565b60405180910390a36001905092915050565b42841015610d99576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d9090611c4d565b60405180910390fd5b5f6001610da46108a9565b7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98a8a8a60055f8f73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f815480929190600101919050558b604051602001610e2996959493929190611c6b565b60405160208183030381529060405280519060200120604051602001610e50929190611d3e565b604051602081830303815290604052805190602001208585856040515f8152602001604052604051610e859493929190611d74565b6020604051602081039080840390855afa158015610ea5573d5f803e3d5ffd5b5050506020604051035190505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614158015610f1857508773ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b610f57576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f4e90611e01565b60405180910390fd5b8560045f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550508573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258760405161103291906114ed565b60405180910390a350505050505050565b6004602052815f5260405f20602052805f5260405f205f91509150505481565b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166315049a5a30846040518363ffffffff1660e01b81526004016110bf929190611e1f565b602060405180830381865afa1580156110da573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110fe9190611e5a565b9050919050565b8060035f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546111519190611915565b925050819055508060025f82825403925050819055505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516111c491906114ed565b60405180910390a35050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f5f6040516112009190611f21565b60405180910390207fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6463060405160200161123f959493929190611f37565b60405160208183030381529060405280519060200120905090565b8060025f82825461126b9190611f88565b925050819055508060035f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055508173ffffffffffffffffffffffffffffffffffffffff165f73ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161131991906114ed565b60405180910390a35050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561135c578082015181840152602081019050611341565b5f8484015250505050565b5f601f19601f8301169050919050565b5f61138182611325565b61138b818561132f565b935061139b81856020860161133f565b6113a481611367565b840191505092915050565b5f6020820190508181035f8301526113c78184611377565b905092915050565b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f611409826113e0565b9050919050565b611419816113ff565b8114611423575f80fd5b50565b5f8135905061143481611410565b92915050565b5f819050919050565b61144c8161143a565b8114611456575f80fd5b50565b5f8135905061146781611443565b92915050565b5f8060408385031215611483576114826113d8565b5b5f61149085828601611426565b92505060206114a185828601611459565b9150509250929050565b5f8115159050919050565b6114bf816114ab565b82525050565b5f6020820190506114d85f8301846114b6565b92915050565b6114e78161143a565b82525050565b5f6020820190506115005f8301846114de565b92915050565b5f805f6060848603121561151d5761151c6113d8565b5b5f61152a86828701611426565b935050602061153b86828701611426565b925050604061154c86828701611459565b9150509250925092565b5f6020828403121561156b5761156a6113d8565b5b5f61157884828501611459565b91505092915050565b5f60ff82169050919050565b61159681611581565b82525050565b5f6020820190506115af5f83018461158d565b92915050565b5f819050919050565b6115c7816115b5565b82525050565b5f6020820190506115e05f8301846115be565b92915050565b5f819050919050565b5f6116096116046115ff846113e0565b6115e6565b6113e0565b9050919050565b5f61161a826115ef565b9050919050565b5f61162b82611610565b9050919050565b61163b81611621565b82525050565b5f6020820190506116545f830184611632565b92915050565b5f6020828403121561166f5761166e6113d8565b5b5f61167c84828501611426565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6116b7816113ff565b82525050565b5f6116c883836116ae565b60208301905092915050565b5f602082019050919050565b5f6116ea82611685565b6116f4818561168f565b93506116ff8361169f565b805f5b8381101561172f57815161171688826116bd565b9750611721836116d4565b925050600181019050611702565b5085935050505092915050565b5f6020820190508181035f83015261175481846116e0565b905092915050565b61176581611581565b811461176f575f80fd5b50565b5f813590506117808161175c565b92915050565b61178f816115b5565b8114611799575f80fd5b50565b5f813590506117aa81611786565b92915050565b5f805f805f805f60e0888a0312156117cb576117ca6113d8565b5b5f6117d88a828b01611426565b97505060206117e98a828b01611426565b96505060406117fa8a828b01611459565b955050606061180b8a828b01611459565b945050608061181c8a828b01611772565b93505060a061182d8a828b0161179c565b92505060c061183e8a828b0161179c565b91505092959891949750929550565b5f8060408385031215611863576118626113d8565b5b5f61187085828601611426565b925050602061188185828601611426565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806118cf57607f821691505b6020821081036118e2576118e161188b565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61191f8261143a565b915061192a8361143a565b9250828203905081811115611942576119416118e8565b5b92915050565b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61198282611367565b810181811067ffffffffffffffff821117156119a1576119a061194c565b5b80604052505050565b5f6119b36113cf565b90506119bf8282611979565b919050565b5f67ffffffffffffffff8211156119de576119dd61194c565b5b602082029050602081019050919050565b5f80fd5b5f81519050611a0181611410565b92915050565b5f611a19611a14846119c4565b6119aa565b90508083825260208201905060208402830185811115611a3c57611a3b6119ef565b5b835b81811015611a655780611a5188826119f3565b845260208401935050602081019050611a3e565b5050509392505050565b5f82601f830112611a8357611a82611948565b5b8151611a93848260208601611a07565b91505092915050565b5f60208284031215611ab157611ab06113d8565b5b5f82015167ffffffffffffffff811115611ace57611acd6113dc565b5b611ada84828501611a6f565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b611b19816113ff565b82525050565b5f604082019050611b325f830185611b10565b611b3f60208301846114de565b9392505050565b611b4f816114ab565b8114611b59575f80fd5b50565b5f81519050611b6a81611b46565b92915050565b5f60208284031215611b8557611b846113d8565b5b5f611b9284828501611b5c565b91505092915050565b7f4661696c656420746f2064656c656761746500000000000000000000000000005f82015250565b5f611bcf60128361132f565b9150611bda82611b9b565b602082019050919050565b5f6020820190508181035f830152611bfc81611bc3565b9050919050565b7f5045524d49545f444541444c494e455f455850495245440000000000000000005f82015250565b5f611c3760178361132f565b9150611c4282611c03565b602082019050919050565b5f6020820190508181035f830152611c6481611c2b565b9050919050565b5f60c082019050611c7e5f8301896115be565b611c8b6020830188611b10565b611c986040830187611b10565b611ca560608301866114de565b611cb260808301856114de565b611cbf60a08301846114de565b979650505050505050565b5f81905092915050565b7f19010000000000000000000000000000000000000000000000000000000000005f82015250565b5f611d08600283611cca565b9150611d1382611cd4565b600282019050919050565b5f819050919050565b611d38611d33826115b5565b611d1e565b82525050565b5f611d4882611cfc565b9150611d548285611d27565b602082019150611d648284611d27565b6020820191508190509392505050565b5f608082019050611d875f8301876115be565b611d94602083018661158d565b611da160408301856115be565b611dae60608301846115be565b95945050505050565b7f494e56414c49445f5349474e45520000000000000000000000000000000000005f82015250565b5f611deb600e8361132f565b9150611df682611db7565b602082019050919050565b5f6020820190508181035f830152611e1881611ddf565b9050919050565b5f604082019050611e325f830185611b10565b611e3f6020830184611b10565b9392505050565b5f81519050611e5481611443565b92915050565b5f60208284031215611e6f57611e6e6113d8565b5b5f611e7c84828501611e46565b91505092915050565b5f81905092915050565b5f819050815f5260205f209050919050565b5f8154611ead816118b8565b611eb78186611e85565b9450600182165f8114611ed15760018114611ee657611f18565b60ff1983168652811515820286019350611f18565b611eef85611e8f565b5f5b83811015611f1057815481890152600182019150602081019050611ef1565b838801955050505b50505092915050565b5f611f2c8284611ea1565b915081905092915050565b5f60a082019050611f4a5f8301886115be565b611f5760208301876115be565b611f6460408301866115be565b611f7160608301856114de565b611f7e6080830184611b10565b9695505050505050565b5f611f928261143a565b9150611f9d8361143a565b9250828201905080821115611fb557611fb46118e8565b5b9291505056fea2646970667358221220246d75a3dad02560e8760551bc6f091f11e7905bb3cd7f7bd5c4e388aaeb334764736f6c63430008140033",
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
