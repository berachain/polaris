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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_stakingprecompile\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validatorAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Data\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"HELLO\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"Success\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"compareMock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveValidatorsMock\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staking\",\"outputs\":[{\"internalType\":\"contractIStakingModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalDelegated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b5060405162002d0138038062002d01833981810160405281019062000037919062000471565b8383601282600090816200004c91906200076c565b5081600190816200005e91906200076c565b508060ff1660808160ff16815250504660a0818152505062000085620001e960201b60201c565b60c08181525050505050600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603620000f6576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036200015d576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050620009b5565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60006040516200021d919062000902565b60405180910390207fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc646306040516020016200025e95949392919062000958565b60405160208183030381529060405280519060200120905090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b620002e28262000297565b810181811067ffffffffffffffff82111715620003045762000303620002a8565b5b80604052505050565b60006200031962000279565b9050620003278282620002d7565b919050565b600067ffffffffffffffff8211156200034a5762000349620002a8565b5b620003558262000297565b9050602081019050919050565b60005b838110156200038257808201518184015260208101905062000365565b60008484015250505050565b6000620003a56200039f846200032c565b6200030d565b905082815260208101848484011115620003c457620003c362000292565b5b620003d184828562000362565b509392505050565b600082601f830112620003f157620003f06200028d565b5b8151620004038482602086016200038e565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600062000439826200040c565b9050919050565b6200044b816200042c565b81146200045757600080fd5b50565b6000815190506200046b8162000440565b92915050565b600080600080608085870312156200048e576200048d62000283565b5b600085015167ffffffffffffffff811115620004af57620004ae62000288565b5b620004bd87828801620003d9565b945050602085015167ffffffffffffffff811115620004e157620004e062000288565b5b620004ef87828801620003d9565b935050604062000502878288016200045a565b925050606062000515878288016200045a565b91505092959194509250565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200057457607f821691505b6020821081036200058a57620005896200052c565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620005f47fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620005b5565b620006008683620005b5565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b60006200064d62000647620006418462000618565b62000622565b62000618565b9050919050565b6000819050919050565b62000669836200062c565b62000681620006788262000654565b848454620005c2565b825550505050565b600090565b6200069862000689565b620006a58184846200065e565b505050565b5b81811015620006cd57620006c16000826200068e565b600181019050620006ab565b5050565b601f8211156200071c57620006e68162000590565b620006f184620005a5565b8101602085101562000701578190505b620007196200071085620005a5565b830182620006aa565b50505b505050565b600082821c905092915050565b6000620007416000198460080262000721565b1980831691505092915050565b60006200075c83836200072e565b9150826002028217905092915050565b620007778262000521565b67ffffffffffffffff811115620007935762000792620002a8565b5b6200079f82546200055b565b620007ac828285620006d1565b600060209050601f831160018114620007e45760008415620007cf578287015190505b620007db85826200074e565b8655506200084b565b601f198416620007f48662000590565b60005b828110156200081e57848901518255600182019150602085019450602081019050620007f7565b868310156200083e57848901516200083a601f8916826200072e565b8355505b6001600288020188555050505b505050505050565b600081905092915050565b60008190508160005260206000209050919050565b6000815462000882816200055b565b6200088e818662000853565b94506001821660008114620008ac5760018114620008c257620008f9565b60ff1983168652811515820286019350620008f9565b620008cd856200085e565b60005b83811015620008f157815481890152600182019150602081019050620008d0565b838801955050505b50505092915050565b600062000910828462000873565b915081905092915050565b6000819050919050565b62000930816200091b565b82525050565b620009418162000618565b82525050565b62000952816200042c565b82525050565b600060a0820190506200096f600083018862000925565b6200097e602083018762000925565b6200098d604083018662000925565b6200099c606083018562000936565b620009ab608083018462000947565b9695505050505050565b60805160a05160c05161231c620009e5600039600061099e0152600061096a01526000610944015261231c6000f3fe6080604052600436106101235760003560e01c80637ecebe00116100a0578063a3d1a57111610064578063a3d1a571146103eb578063a9059cbb14610416578063d505accf14610453578063dc4124c01461047c578063dd62ed3e146104a75761012a565b80637ecebe001461031157806380d04de81461034e57806395d89b41146103795780639de70258146103a45780639fa6dd35146103cf5761012a565b8063313ce567116100e7578063313ce567146102285780633644e515146102535780633fe4676e1461027e5780634cf088d9146102a957806370a08231146102d45761012a565b806306fdde031461012f578063095ea7b31461015a57806318160ddd1461019757806323b872dd146101c25780632e1a7d4d146101ff5761012a565b3661012a57005b600080fd5b34801561013b57600080fd5b506101446104e4565b60405161015191906115f4565b60405180910390f35b34801561016657600080fd5b50610181600480360381019061017c91906116be565b610572565b60405161018e9190611719565b60405180910390f35b3480156101a357600080fd5b506101ac610664565b6040516101b99190611743565b60405180910390f35b3480156101ce57600080fd5b506101e960048036038101906101e4919061175e565b61066a565b6040516101f69190611719565b60405180910390f35b34801561020b57600080fd5b50610226600480360381019061022191906117b1565b6108b4565b005b34801561023457600080fd5b5061023d610942565b60405161024a91906117fa565b60405180910390f35b34801561025f57600080fd5b50610268610966565b604051610275919061182e565b60405180910390f35b34801561028a57600080fd5b506102936109c3565b6040516102a09190611858565b60405180910390f35b3480156102b557600080fd5b506102be6109e9565b6040516102cb91906118d2565b60405180910390f35b3480156102e057600080fd5b506102fb60048036038101906102f691906118ed565b610a0f565b6040516103089190611743565b60405180910390f35b34801561031d57600080fd5b50610338600480360381019061033391906118ed565b610a27565b6040516103459190611743565b60405180910390f35b34801561035a57600080fd5b50610363610a3f565b6040516103709190611743565b60405180910390f35b34801561038557600080fd5b5061038e610b06565b60405161039b91906115f4565b60405180910390f35b3480156103b057600080fd5b506103b9610b94565b6040516103c691906119d8565b60405180910390f35b6103e960048036038101906103e491906117b1565b610c31565b005b3480156103f757600080fd5b50610400610d45565b60405161040d91906119d8565b60405180910390f35b34801561042257600080fd5b5061043d600480360381019061043891906116be565b610dff565b60405161044a9190611719565b60405180910390f35b34801561045f57600080fd5b5061047a60048036038101906104759190611a52565b610f13565b005b34801561048857600080fd5b5061049161120c565b60405161049e9190611719565b60405180910390f35b3480156104b357600080fd5b506104ce60048036038101906104c99190611af4565b611313565b6040516104db9190611743565b60405180910390f35b600080546104f190611b63565b80601f016020809104026020016040519081016040528092919081815260200182805461051d90611b63565b801561056a5780601f1061053f5761010080835404028352916020019161056a565b820191906000526020600020905b81548152906001019060200180831161054d57829003601f168201915b505050505081565b600081600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516106529190611743565b60405180910390a36001905092915050565b60025481565b600080600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146107a057828161071f9190611bc3565b600460008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b82600360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546107ef9190611bc3565b9250508190555082600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef856040516108a09190611743565b60405180910390a360019150509392505050565b600081036108ee576040517f1f2a200500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108f83382611338565b3373ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f1935050505015801561093e573d6000803e3d6000fd5b5050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60007f0000000000000000000000000000000000000000000000000000000000000000461461099c57610997611408565b6109be565b7f00000000000000000000000000000000000000000000000000000000000000005b905090565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60036020528060005260406000206000915090505481565b60056020528060005260406000206000915090505481565b6000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166315049a5a30600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518363ffffffff1660e01b8152600401610ac0929190611bf7565b602060405180830381865afa158015610add573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b019190611c35565b905090565b60018054610b1390611b63565b80601f0160208091040260200160405190810160405280929190818152602001828054610b3f90611b63565b8015610b8c5780601f10610b6157610100808354040283529160200191610b8c565b820191906000526020600020905b815481529060010190602001808311610b6f57829003601f168201915b505050505081565b6060600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639de702586040518163ffffffff1660e01b8152600401600060405180830381865afa158015610c03573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610c2c9190611dbf565b905090565b6000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663026e402b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518363ffffffff1660e01b8152600401610cb2929190611e08565b6020604051808303816000875af1158015610cd1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cf59190611e5d565b905080610d37576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d2e90611ed6565b60405180910390fd5b610d413383611494565b5050565b60606000600167ffffffffffffffff811115610d6457610d63611c67565b5b604051908082528060200260200182016040528015610d925781602001602082028036833780820191505090505b509050734fc768b13a96e8ebfdecfb85d86d9dc498f5b06281600081518110610dbe57610dbd611ef6565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508091505090565b600081600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610e509190611bc3565b9250508190555081600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610f019190611743565b60405180910390a36001905092915050565b42841015610f56576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f4d90611f71565b60405180910390fd5b60006001610f62610966565b7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98a8a8a600560008f73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000815480929190600101919050558b604051602001610fea96959493929190611f91565b6040516020818303038152906040528051906020012060405160200161101192919061206a565b604051602081830303815290604052805190602001208585856040516000815260200160405260405161104794939291906120a1565b6020604051602081039080840390855afa158015611069573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141580156110dd57508773ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b61111c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161111390612132565b60405180910390fd5b85600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550508573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925876040516111fb9190611743565b60405180910390a350505050505050565b60008060003073ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527fa3d1a571000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516112b89190612199565b600060405180830381855afa9150503d80600081146112f3576040519150601f19603f3d011682016040523d82523d6000602084013e6112f8565b606091505b50915091506000611307610b94565b90506001935050505090565b6004602052816000526040600020602052806000526040600020600091509150505481565b80600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546113879190611bc3565b9250508190555080600260008282540392505081905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516113fc9190611743565b60405180910390a35050565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f600060405161143a9190612248565b60405180910390207fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6463060405160200161147995949392919061225f565b60405160208183030381529060405280519060200120905090565b80600260008282546114a691906122b2565b9250508190555080600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516115589190611743565b60405180910390a35050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561159e578082015181840152602081019050611583565b60008484015250505050565b6000601f19601f8301169050919050565b60006115c682611564565b6115d0818561156f565b93506115e0818560208601611580565b6115e9816115aa565b840191505092915050565b6000602082019050818103600083015261160e81846115bb565b905092915050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006116558261162a565b9050919050565b6116658161164a565b811461167057600080fd5b50565b6000813590506116828161165c565b92915050565b6000819050919050565b61169b81611688565b81146116a657600080fd5b50565b6000813590506116b881611692565b92915050565b600080604083850312156116d5576116d4611620565b5b60006116e385828601611673565b92505060206116f4858286016116a9565b9150509250929050565b60008115159050919050565b611713816116fe565b82525050565b600060208201905061172e600083018461170a565b92915050565b61173d81611688565b82525050565b60006020820190506117586000830184611734565b92915050565b60008060006060848603121561177757611776611620565b5b600061178586828701611673565b935050602061179686828701611673565b92505060406117a7868287016116a9565b9150509250925092565b6000602082840312156117c7576117c6611620565b5b60006117d5848285016116a9565b91505092915050565b600060ff82169050919050565b6117f4816117de565b82525050565b600060208201905061180f60008301846117eb565b92915050565b6000819050919050565b61182881611815565b82525050565b6000602082019050611843600083018461181f565b92915050565b6118528161164a565b82525050565b600060208201905061186d6000830184611849565b92915050565b6000819050919050565b600061189861189361188e8461162a565b611873565b61162a565b9050919050565b60006118aa8261187d565b9050919050565b60006118bc8261189f565b9050919050565b6118cc816118b1565b82525050565b60006020820190506118e760008301846118c3565b92915050565b60006020828403121561190357611902611620565b5b600061191184828501611673565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61194f8161164a565b82525050565b60006119618383611946565b60208301905092915050565b6000602082019050919050565b60006119858261191a565b61198f8185611925565b935061199a83611936565b8060005b838110156119cb5781516119b28882611955565b97506119bd8361196d565b92505060018101905061199e565b5085935050505092915050565b600060208201905081810360008301526119f2818461197a565b905092915050565b611a03816117de565b8114611a0e57600080fd5b50565b600081359050611a20816119fa565b92915050565b611a2f81611815565b8114611a3a57600080fd5b50565b600081359050611a4c81611a26565b92915050565b600080600080600080600060e0888a031215611a7157611a70611620565b5b6000611a7f8a828b01611673565b9750506020611a908a828b01611673565b9650506040611aa18a828b016116a9565b9550506060611ab28a828b016116a9565b9450506080611ac38a828b01611a11565b93505060a0611ad48a828b01611a3d565b92505060c0611ae58a828b01611a3d565b91505092959891949750929550565b60008060408385031215611b0b57611b0a611620565b5b6000611b1985828601611673565b9250506020611b2a85828601611673565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680611b7b57607f821691505b602082108103611b8e57611b8d611b34565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611bce82611688565b9150611bd983611688565b9250828203905081811115611bf157611bf0611b94565b5b92915050565b6000604082019050611c0c6000830185611849565b611c196020830184611849565b9392505050565b600081519050611c2f81611692565b92915050565b600060208284031215611c4b57611c4a611620565b5b6000611c5984828501611c20565b91505092915050565b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b611c9f826115aa565b810181811067ffffffffffffffff82111715611cbe57611cbd611c67565b5b80604052505050565b6000611cd1611616565b9050611cdd8282611c96565b919050565b600067ffffffffffffffff821115611cfd57611cfc611c67565b5b602082029050602081019050919050565b600080fd5b600081519050611d228161165c565b92915050565b6000611d3b611d3684611ce2565b611cc7565b90508083825260208201905060208402830185811115611d5e57611d5d611d0e565b5b835b81811015611d875780611d738882611d13565b845260208401935050602081019050611d60565b5050509392505050565b600082601f830112611da657611da5611c62565b5b8151611db6848260208601611d28565b91505092915050565b600060208284031215611dd557611dd4611620565b5b600082015167ffffffffffffffff811115611df357611df2611625565b5b611dff84828501611d91565b91505092915050565b6000604082019050611e1d6000830185611849565b611e2a6020830184611734565b9392505050565b611e3a816116fe565b8114611e4557600080fd5b50565b600081519050611e5781611e31565b92915050565b600060208284031215611e7357611e72611620565b5b6000611e8184828501611e48565b91505092915050565b7f4661696c656420746f2064656c65676174652031000000000000000000000000600082015250565b6000611ec060148361156f565b9150611ecb82611e8a565b602082019050919050565b60006020820190508181036000830152611eef81611eb3565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f5045524d49545f444541444c494e455f45585049524544000000000000000000600082015250565b6000611f5b60178361156f565b9150611f6682611f25565b602082019050919050565b60006020820190508181036000830152611f8a81611f4e565b9050919050565b600060c082019050611fa6600083018961181f565b611fb36020830188611849565b611fc06040830187611849565b611fcd6060830186611734565b611fda6080830185611734565b611fe760a0830184611734565b979650505050505050565b600081905092915050565b7f1901000000000000000000000000000000000000000000000000000000000000600082015250565b6000612033600283611ff2565b915061203e82611ffd565b600282019050919050565b6000819050919050565b61206461205f82611815565b612049565b82525050565b600061207582612026565b91506120818285612053565b6020820191506120918284612053565b6020820191508190509392505050565b60006080820190506120b6600083018761181f565b6120c360208301866117eb565b6120d0604083018561181f565b6120dd606083018461181f565b95945050505050565b7f494e56414c49445f5349474e4552000000000000000000000000000000000000600082015250565b600061211c600e8361156f565b9150612127826120e6565b602082019050919050565b6000602082019050818103600083015261214b8161210f565b9050919050565b600081519050919050565b600081905092915050565b600061217382612152565b61217d818561215d565b935061218d818560208601611580565b80840191505092915050565b60006121a58284612168565b915081905092915050565b60008190508160005260206000209050919050565b600081546121d281611b63565b6121dc818661215d565b945060018216600081146121f7576001811461220c5761223f565b60ff198316865281151582028601935061223f565b612215856121b0565b60005b8381101561223757815481890152600182019150602081019050612218565b838801955050505b50505092915050565b600061225482846121c5565b915081905092915050565b600060a082019050612274600083018861181f565b612281602083018761181f565b61228e604083018661181f565b61229b6060830185611734565b6122a86080830184611849565b9695505050505050565b60006122bd82611688565b91506122c883611688565b92508282019050808211156122e0576122df611b94565b5b9291505056fea26469706673582212201115320b4faa4587a680d206d3dfff29d0c786a0e9f0b3c1ba643a616a15718064736f6c63430008130033",
}

// LiquidStakingABI is the input ABI used to generate the binding from.
// Deprecated: Use LiquidStakingMetaData.ABI instead.
var LiquidStakingABI = LiquidStakingMetaData.ABI

// LiquidStakingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LiquidStakingMetaData.Bin instead.
var LiquidStakingBin = LiquidStakingMetaData.Bin

// DeployLiquidStaking deploys a new Ethereum contract, binding an instance of LiquidStaking to it.
func DeployLiquidStaking(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string, _stakingprecompile common.Address, _validatorAddress common.Address) (common.Address, *types.Transaction, *LiquidStaking, error) {
	parsed, err := LiquidStakingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LiquidStakingBin), backend, _name, _symbol, _stakingprecompile, _validatorAddress)
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

// CompareMock is a free data retrieval call binding the contract method 0xdc4124c0.
//
// Solidity: function compareMock() view returns(bool)
func (_LiquidStaking *LiquidStakingCaller) CompareMock(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "compareMock")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CompareMock is a free data retrieval call binding the contract method 0xdc4124c0.
//
// Solidity: function compareMock() view returns(bool)
func (_LiquidStaking *LiquidStakingSession) CompareMock() (bool, error) {
	return _LiquidStaking.Contract.CompareMock(&_LiquidStaking.CallOpts)
}

// CompareMock is a free data retrieval call binding the contract method 0xdc4124c0.
//
// Solidity: function compareMock() view returns(bool)
func (_LiquidStaking *LiquidStakingCallerSession) CompareMock() (bool, error) {
	return _LiquidStaking.Contract.CompareMock(&_LiquidStaking.CallOpts)
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

// GetActiveValidatorsMock is a free data retrieval call binding the contract method 0xa3d1a571.
//
// Solidity: function getActiveValidatorsMock() view returns(address[])
func (_LiquidStaking *LiquidStakingCaller) GetActiveValidatorsMock(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "getActiveValidatorsMock")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetActiveValidatorsMock is a free data retrieval call binding the contract method 0xa3d1a571.
//
// Solidity: function getActiveValidatorsMock() view returns(address[])
func (_LiquidStaking *LiquidStakingSession) GetActiveValidatorsMock() ([]common.Address, error) {
	return _LiquidStaking.Contract.GetActiveValidatorsMock(&_LiquidStaking.CallOpts)
}

// GetActiveValidatorsMock is a free data retrieval call binding the contract method 0xa3d1a571.
//
// Solidity: function getActiveValidatorsMock() view returns(address[])
func (_LiquidStaking *LiquidStakingCallerSession) GetActiveValidatorsMock() ([]common.Address, error) {
	return _LiquidStaking.Contract.GetActiveValidatorsMock(&_LiquidStaking.CallOpts)
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

// ValidatorAddress is a free data retrieval call binding the contract method 0x3fe4676e.
//
// Solidity: function validatorAddress() view returns(address)
func (_LiquidStaking *LiquidStakingCaller) ValidatorAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidStaking.contract.Call(opts, &out, "validatorAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorAddress is a free data retrieval call binding the contract method 0x3fe4676e.
//
// Solidity: function validatorAddress() view returns(address)
func (_LiquidStaking *LiquidStakingSession) ValidatorAddress() (common.Address, error) {
	return _LiquidStaking.Contract.ValidatorAddress(&_LiquidStaking.CallOpts)
}

// ValidatorAddress is a free data retrieval call binding the contract method 0x3fe4676e.
//
// Solidity: function validatorAddress() view returns(address)
func (_LiquidStaking *LiquidStakingCallerSession) ValidatorAddress() (common.Address, error) {
	return _LiquidStaking.Contract.ValidatorAddress(&_LiquidStaking.CallOpts)
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

// TotalDelegated is a paid mutator transaction binding the contract method 0x80d04de8.
//
// Solidity: function totalDelegated() returns(uint256 amount)
func (_LiquidStaking *LiquidStakingTransactor) TotalDelegated(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidStaking.contract.Transact(opts, "totalDelegated")
}

// TotalDelegated is a paid mutator transaction binding the contract method 0x80d04de8.
//
// Solidity: function totalDelegated() returns(uint256 amount)
func (_LiquidStaking *LiquidStakingSession) TotalDelegated() (*types.Transaction, error) {
	return _LiquidStaking.Contract.TotalDelegated(&_LiquidStaking.TransactOpts)
}

// TotalDelegated is a paid mutator transaction binding the contract method 0x80d04de8.
//
// Solidity: function totalDelegated() returns(uint256 amount)
func (_LiquidStaking *LiquidStakingTransactorSession) TotalDelegated() (*types.Transaction, error) {
	return _LiquidStaking.Contract.TotalDelegated(&_LiquidStaking.TransactOpts)
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
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
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

// LiquidStakingHELLOIterator is returned from FilterHELLO and is used to iterate over the raw logs and unpacked data for HELLO events raised by the LiquidStaking contract.
type LiquidStakingHELLOIterator struct {
	Event *LiquidStakingHELLO // Event containing the contract specifics and raw log

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
func (it *LiquidStakingHELLOIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidStakingHELLO)
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
		it.Event = new(LiquidStakingHELLO)
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
func (it *LiquidStakingHELLOIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidStakingHELLOIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidStakingHELLO represents a HELLO event raised by the LiquidStaking contract.
type LiquidStakingHELLO struct {
	Message string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterHELLO is a free log retrieval operation binding the contract event 0xeb2b994e506bc71756358633d451ce62958e0287899f16414c7f1a93d20bf67a.
//
// Solidity: event HELLO(string message)
func (_LiquidStaking *LiquidStakingFilterer) FilterHELLO(opts *bind.FilterOpts) (*LiquidStakingHELLOIterator, error) {

	logs, sub, err := _LiquidStaking.contract.FilterLogs(opts, "HELLO")
	if err != nil {
		return nil, err
	}
	return &LiquidStakingHELLOIterator{contract: _LiquidStaking.contract, event: "HELLO", logs: logs, sub: sub}, nil
}

// WatchHELLO is a free log subscription operation binding the contract event 0xeb2b994e506bc71756358633d451ce62958e0287899f16414c7f1a93d20bf67a.
//
// Solidity: event HELLO(string message)
func (_LiquidStaking *LiquidStakingFilterer) WatchHELLO(opts *bind.WatchOpts, sink chan<- *LiquidStakingHELLO) (event.Subscription, error) {

	logs, sub, err := _LiquidStaking.contract.WatchLogs(opts, "HELLO")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidStakingHELLO)
				if err := _LiquidStaking.contract.UnpackLog(event, "HELLO", log); err != nil {
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

// ParseHELLO is a log parse operation binding the contract event 0xeb2b994e506bc71756358633d451ce62958e0287899f16414c7f1a93d20bf67a.
//
// Solidity: event HELLO(string message)
func (_LiquidStaking *LiquidStakingFilterer) ParseHELLO(log types.Log) (*LiquidStakingHELLO, error) {
	event := new(LiquidStakingHELLO)
	if err := _LiquidStaking.contract.UnpackLog(event, "HELLO", log); err != nil {
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
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_LiquidStaking *LiquidStakingFilterer) ParseTransfer(log types.Log) (*LiquidStakingTransfer, error) {
	event := new(LiquidStakingTransfer)
	if err := _LiquidStaking.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
