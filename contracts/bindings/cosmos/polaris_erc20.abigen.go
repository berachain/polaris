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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_denom\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denom\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200234e3803806200234e8339818101604052810190620000379190620001e3565b80600090816200004891906200047f565b505062000566565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b620000b9826200006e565b810181811067ffffffffffffffff82111715620000db57620000da6200007f565b5b80604052505050565b6000620000f062000050565b9050620000fe8282620000ae565b919050565b600067ffffffffffffffff8211156200012157620001206200007f565b5b6200012c826200006e565b9050602081019050919050565b60005b83811015620001595780820151818401526020810190506200013c565b60008484015250505050565b60006200017c620001768462000103565b620000e4565b9050828152602081018484840111156200019b576200019a62000069565b5b620001a884828562000139565b509392505050565b600082601f830112620001c857620001c762000064565b5b8151620001da84826020860162000165565b91505092915050565b600060208284031215620001fc57620001fb6200005a565b5b600082015167ffffffffffffffff8111156200021d576200021c6200005f565b5b6200022b84828501620001b0565b91505092915050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200028757607f821691505b6020821081036200029d576200029c6200023f565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620003077fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620002c8565b620003138683620002c8565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000620003606200035a62000354846200032b565b62000335565b6200032b565b9050919050565b6000819050919050565b6200037c836200033f565b620003946200038b8262000367565b848454620002d5565b825550505050565b600090565b620003ab6200039c565b620003b881848462000371565b505050565b5b81811015620003e057620003d4600082620003a1565b600181019050620003be565b5050565b601f8211156200042f57620003f981620002a3565b6200040484620002b8565b8101602085101562000414578190505b6200042c6200042385620002b8565b830182620003bd565b50505b505050565b600082821c905092915050565b6000620004546000198460080262000434565b1980831691505092915050565b60006200046f838362000441565b9150826002028217905092915050565b6200048a8262000234565b67ffffffffffffffff811115620004a657620004a56200007f565b5b620004b282546200026e565b620004bf828285620003e4565b600060209050601f831160018114620004f75760008415620004e2578287015190505b620004ee858262000461565b8655506200055e565b601f1984166200050786620002a3565b60005b8281101562000531578489015182556001820191506020850194506020810190506200050a565b868310156200055157848901516200054d601f89168262000441565b8355505b6001600288020188555050505b505050505050565b611dd880620005766000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c806370a082311161006657806370a082311461015d57806395d89b411461018d578063a9059cbb146101ab578063c370b042146101db578063dd62ed3e146101f95761009e565b806306fdde03146100a3578063095ea7b3146100c157806318160ddd146100f157806323b872dd1461010f578063313ce5671461013f575b600080fd5b6100ab610229565b6040516100b89190610e11565b60405180910390f35b6100db60048036038101906100d69190610edb565b6102bb565b6040516100e89190610f36565b60405180910390f35b6100f961047e565b6040516101069190610f60565b60405180910390f35b61012960048036038101906101249190610f7b565b610507565b6040516101369190610f36565b60405180910390f35b61014761070f565b6040516101549190610fea565b60405180910390f35b61017760048036038101906101729190611005565b6107c0565b6040516101849190610f60565b60405180910390f35b61019561084d565b6040516101a29190610e11565b60405180910390f35b6101c560048036038101906101c09190610edb565b6108df565b6040516101d29190610f36565b60405180910390f35b6101e3610a1e565b6040516101f09190610e11565b60405180910390f35b610213600480360381019061020e9190611032565b610aac565b6040516102209190610f60565b60405180910390f35b6060610233610ad1565b73ffffffffffffffffffffffffffffffffffffffff166352a6ea0460006040518263ffffffff1660e01b815260040161026c919061116b565b600060405180830381865afa158015610289573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906102b291906116a4565b60600151905090565b60006102c5610aed565b73ffffffffffffffffffffffffffffffffffffffff16632b6b7ab533856102eb86610b09565b60006040518563ffffffff1660e01b815260040161030c9493929190611899565b6020604051808303816000875af115801561032b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061034f9190611911565b61038e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610385906119b0565b60405180910390fd5b81600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161046c9190610f60565b60405180910390a36001905092915050565b6000610488610ad1565b73ffffffffffffffffffffffffffffffffffffffff1663fe3b2b8860006040518263ffffffff1660e01b81526004016104c1919061116b565b602060405180830381865afa1580156104de573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061050291906119e5565b905090565b6000610511610aed565b73ffffffffffffffffffffffffffffffffffffffff1663fbdb0e87853360006040518463ffffffff1660e01b815260040161054e93929190611a12565b602060405180830381865afa15801561056b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061058f91906119e5565b8211156105d1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105c890611ac2565b60405180910390fd5b6105d9610ad1565b73ffffffffffffffffffffffffffffffffffffffff16638440481185856105ff86610c2b565b6040518463ffffffff1660e01b815260040161061d93929190611be1565b6020604051808303816000875af115801561063c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106609190611911565b61069f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161069690611c91565b60405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516106fc9190610f60565b60405180910390a3600190509392505050565b6000610719610ad1565b73ffffffffffffffffffffffffffffffffffffffff166352a6ea0460006040518263ffffffff1660e01b8152600401610752919061116b565b600060405180830381865afa15801561076f573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061079891906116a4565b602001516000815181106107af576107ae611cb1565b5b602002602001015160400151905090565b60006107ca610ad1565b73ffffffffffffffffffffffffffffffffffffffff166334d1fdaf8360006040518363ffffffff1660e01b8152600401610805929190611ce0565b602060405180830381865afa158015610822573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061084691906119e5565b9050919050565b6060610857610ad1565b73ffffffffffffffffffffffffffffffffffffffff166352a6ea0460006040518263ffffffff1660e01b8152600401610890919061116b565b600060405180830381865afa1580156108ad573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906108d691906116a4565b60a00151905090565b60006108e9610ad1565b73ffffffffffffffffffffffffffffffffffffffff166384404811338561090f86610c2b565b6040518463ffffffff1660e01b815260040161092d93929190611be1565b6020604051808303816000875af115801561094c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109709190611911565b6109af576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109a690611d82565b60405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610a0c9190610f60565b60405180910390a36001905092915050565b60008054610a2b906110a1565b80601f0160208091040260200160405190810160405280929190818152602001828054610a57906110a1565b8015610aa45780601f10610a7957610100808354040283529160200191610aa4565b820191906000526020600020905b815481529060010190602001808311610a8757829003601f168201915b505050505081565b6001602052816000526040600020602052806000526040600020600091509150505481565b6000734381dc2ab14285160c808659aee005d51255add7905090565b600073bdf49c3c3882102fc017ffb661108c63a836d065905090565b60606000600167ffffffffffffffff811115610b2857610b27611192565b5b604051908082528060200260200182016040528015610b6157816020015b610b4e610d4d565b815260200190600190039081610b465790505b509050604051806040016040528084815260200160008054610b82906110a1565b80601f0160208091040260200160405190810160405280929190818152602001828054610bae906110a1565b8015610bfb5780601f10610bd057610100808354040283529160200191610bfb565b820191906000526020600020905b815481529060010190602001808311610bde57829003601f168201915b505050505081525081600081518110610c1757610c16611cb1565b5b602002602001018190525080915050919050565b60606000600167ffffffffffffffff811115610c4a57610c49611192565b5b604051908082528060200260200182016040528015610c8357816020015b610c70610d67565b815260200190600190039081610c685790505b509050604051806040016040528084815260200160008054610ca4906110a1565b80601f0160208091040260200160405190810160405280929190818152602001828054610cd0906110a1565b8015610d1d5780601f10610cf257610100808354040283529160200191610d1d565b820191906000526020600020905b815481529060010190602001808311610d0057829003601f168201915b505050505081525081600081518110610d3957610d38611cb1565b5b602002602001018190525080915050919050565b604051806040016040528060008152602001606081525090565b604051806040016040528060008152602001606081525090565b600081519050919050565b600082825260208201905092915050565b60005b83811015610dbb578082015181840152602081019050610da0565b60008484015250505050565b6000601f19601f8301169050919050565b6000610de382610d81565b610ded8185610d8c565b9350610dfd818560208601610d9d565b610e0681610dc7565b840191505092915050565b60006020820190508181036000830152610e2b8184610dd8565b905092915050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610e7282610e47565b9050919050565b610e8281610e67565b8114610e8d57600080fd5b50565b600081359050610e9f81610e79565b92915050565b6000819050919050565b610eb881610ea5565b8114610ec357600080fd5b50565b600081359050610ed581610eaf565b92915050565b60008060408385031215610ef257610ef1610e3d565b5b6000610f0085828601610e90565b9250506020610f1185828601610ec6565b9150509250929050565b60008115159050919050565b610f3081610f1b565b82525050565b6000602082019050610f4b6000830184610f27565b92915050565b610f5a81610ea5565b82525050565b6000602082019050610f756000830184610f51565b92915050565b600080600060608486031215610f9457610f93610e3d565b5b6000610fa286828701610e90565b9350506020610fb386828701610e90565b9250506040610fc486828701610ec6565b9150509250925092565b600060ff82169050919050565b610fe481610fce565b82525050565b6000602082019050610fff6000830184610fdb565b92915050565b60006020828403121561101b5761101a610e3d565b5b600061102984828501610e90565b91505092915050565b6000806040838503121561104957611048610e3d565b5b600061105785828601610e90565b925050602061106885828601610e90565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806110b957607f821691505b6020821081036110cc576110cb611072565b5b50919050565b60008190508160005260206000209050919050565b600081546110f4816110a1565b6110fe8186610d8c565b94506001821660008114611119576001811461112f57611162565b60ff198316865281151560200286019350611162565b611138856110d2565b60005b8381101561115a5781548189015260018201915060208101905061113b565b808801955050505b50505092915050565b6000602082019050818103600083015261118581846110e7565b905092915050565b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6111ca82610dc7565b810181811067ffffffffffffffff821117156111e9576111e8611192565b5b80604052505050565b60006111fc610e33565b905061120882826111c1565b919050565b600080fd5b600080fd5b600080fd5b600067ffffffffffffffff82111561123757611236611192565b5b61124082610dc7565b9050602081019050919050565b600061126061125b8461121c565b6111f2565b90508281526020810184848401111561127c5761127b611217565b5b611287848285610d9d565b509392505050565b600082601f8301126112a4576112a3611212565b5b81516112b484826020860161124d565b91505092915050565b600067ffffffffffffffff8211156112d8576112d7611192565b5b602082029050602081019050919050565b600080fd5b600067ffffffffffffffff82111561130957611308611192565b5b602082029050602081019050919050565b600061132d611328846112ee565b6111f2565b905080838252602082019050602084028301858111156113505761134f6112e9565b5b835b8181101561139757805167ffffffffffffffff81111561137557611374611212565b5b808601611382898261128f565b85526020850194505050602081019050611352565b5050509392505050565b600082601f8301126113b6576113b5611212565b5b81516113c684826020860161131a565b91505092915050565b600063ffffffff82169050919050565b6113e8816113cf565b81146113f357600080fd5b50565b600081519050611405816113df565b92915050565b6000606082840312156114215761142061118d565b5b61142b60606111f2565b9050600082015167ffffffffffffffff81111561144b5761144a61120d565b5b6114578482850161128f565b600083015250602082015167ffffffffffffffff81111561147b5761147a61120d565b5b611487848285016113a1565b602083015250604061149b848285016113f6565b60408301525092915050565b60006114ba6114b5846112bd565b6111f2565b905080838252602082019050602084028301858111156114dd576114dc6112e9565b5b835b8181101561152457805167ffffffffffffffff81111561150257611501611212565b5b80860161150f898261140b565b855260208501945050506020810190506114df565b5050509392505050565b600082601f83011261154357611542611212565b5b81516115538482602086016114a7565b91505092915050565b600060c082840312156115725761157161118d565b5b61157c60c06111f2565b9050600082015167ffffffffffffffff81111561159c5761159b61120d565b5b6115a88482850161128f565b600083015250602082015167ffffffffffffffff8111156115cc576115cb61120d565b5b6115d88482850161152e565b602083015250604082015167ffffffffffffffff8111156115fc576115fb61120d565b5b6116088482850161128f565b604083015250606082015167ffffffffffffffff81111561162c5761162b61120d565b5b6116388482850161128f565b606083015250608082015167ffffffffffffffff81111561165c5761165b61120d565b5b6116688482850161128f565b60808301525060a082015167ffffffffffffffff81111561168c5761168b61120d565b5b6116988482850161128f565b60a08301525092915050565b6000602082840312156116ba576116b9610e3d565b5b600082015167ffffffffffffffff8111156116d8576116d7610e42565b5b6116e48482850161155c565b91505092915050565b6116f681610e67565b82525050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61173181610ea5565b82525050565b600082825260208201905092915050565b600061175382610d81565b61175d8185611737565b935061176d818560208601610d9d565b61177681610dc7565b840191505092915050565b60006040830160008301516117996000860182611728565b50602083015184820360208601526117b18282611748565b9150508091505092915050565b60006117ca8383611781565b905092915050565b6000602082019050919050565b60006117ea826116fc565b6117f48185611707565b93508360208202850161180685611718565b8060005b85811015611842578484038952815161182385826117be565b945061182e836117d2565b925060208a0199505060018101905061180a565b50829750879550505050505092915050565b6000819050919050565b6000819050919050565b600061188361187e61187984611854565b61185e565b610ea5565b9050919050565b61189381611868565b82525050565b60006080820190506118ae60008301876116ed565b6118bb60208301866116ed565b81810360408301526118cd81856117df565b90506118dc606083018461188a565b95945050505050565b6118ee81610f1b565b81146118f957600080fd5b50565b60008151905061190b816118e5565b92915050565b60006020828403121561192757611926610e3d565b5b6000611935848285016118fc565b91505092915050565b7f506f6c6172697345524332303a206661696c656420746f20617070726f76652060008201527f7370656e64000000000000000000000000000000000000000000000000000000602082015250565b600061199a602583610d8c565b91506119a58261193e565b604082019050919050565b600060208201905081810360008301526119c98161198d565b9050919050565b6000815190506119df81610eaf565b92915050565b6000602082840312156119fb576119fa610e3d565b5b6000611a09848285016119d0565b91505092915050565b6000606082019050611a2760008301866116ed565b611a3460208301856116ed565b8181036040830152611a4681846110e7565b9050949350505050565b7f506f6c6172697345524332303a20696e73756666696369656e7420617070726f60008201527f76616c0000000000000000000000000000000000000000000000000000000000602082015250565b6000611aac602383610d8c565b9150611ab782611a50565b604082019050919050565b60006020820190508181036000830152611adb81611a9f565b9050919050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6000604083016000830151611b266000860182611728565b5060208301518482036020860152611b3e8282611748565b9150508091505092915050565b6000611b578383611b0e565b905092915050565b6000602082019050919050565b6000611b7782611ae2565b611b818185611aed565b935083602082028501611b9385611afe565b8060005b85811015611bcf5784840389528151611bb08582611b4b565b9450611bbb83611b5f565b925060208a01995050600181019050611b97565b50829750879550505050505092915050565b6000606082019050611bf660008301866116ed565b611c0360208301856116ed565b8181036040830152611c158184611b6c565b9050949350505050565b7f506f6c6172697345524332303a206661696c656420746f2073656e642062616e60008201527f6b20746f6b656e73000000000000000000000000000000000000000000000000602082015250565b6000611c7b602883610d8c565b9150611c8682611c1f565b604082019050919050565b60006020820190508181036000830152611caa81611c6e565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000604082019050611cf560008301856116ed565b8181036020830152611d0781846110e7565b90509392505050565b7f506f6c6172697345524332303a206661696c656420746f2073656e6420746f6b60008201527f656e730000000000000000000000000000000000000000000000000000000000602082015250565b6000611d6c602383610d8c565b9150611d7782611d10565b604082019050919050565b60006020820190508181036000830152611d9b81611d5f565b905091905056fea264697066735822122023913ce1d19a639895eb2e20415bd6e84cdcf0a6da9ce4ff251cdfc69f6d38df64736f6c63430008130033",
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
