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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_denom\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denom\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801562000010575f80fd5b506040516200208d3803806200208d833981810160405281019062000036919062000280565b805f908162000046919062000506565b5046608081815250506200005f6200006d60201b60201c565b60a081815250505062000773565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f5f6040516200009f919062000692565b60405180910390207fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc64630604051602001620000e095949392919062000718565b60405160208183030381529060405280519060200120905090565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6200015c8262000114565b810181811067ffffffffffffffff821117156200017e576200017d62000124565b5b80604052505050565b5f62000192620000fb565b9050620001a0828262000151565b919050565b5f67ffffffffffffffff821115620001c257620001c162000124565b5b620001cd8262000114565b9050602081019050919050565b5f5b83811015620001f9578082015181840152602081019050620001dc565b5f8484015250505050565b5f6200021a6200021484620001a5565b62000187565b90508281526020810184848401111562000239576200023862000110565b5b62000246848285620001da565b509392505050565b5f82601f8301126200026557620002646200010c565b5b81516200027784826020860162000204565b91505092915050565b5f6020828403121562000298576200029762000104565b5b5f82015167ffffffffffffffff811115620002b857620002b762000108565b5b620002c6848285016200024e565b91505092915050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806200031e57607f821691505b602082108103620003345762000333620002d9565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302620003987fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826200035b565b620003a486836200035b565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f620003ee620003e8620003e284620003bc565b620003c5565b620003bc565b9050919050565b5f819050919050565b6200040983620003ce565b620004216200041882620003f5565b84845462000367565b825550505050565b5f90565b6200043762000429565b62000444818484620003fe565b505050565b5b818110156200046b576200045f5f826200042d565b6001810190506200044a565b5050565b601f821115620004ba5762000484816200033a565b6200048f846200034c565b810160208510156200049f578190505b620004b7620004ae856200034c565b83018262000449565b50505b505050565b5f82821c905092915050565b5f620004dc5f1984600802620004bf565b1980831691505092915050565b5f620004f68383620004cb565b9150826002028217905092915050565b6200051182620002cf565b67ffffffffffffffff8111156200052d576200052c62000124565b5b62000539825462000306565b620005468282856200046f565b5f60209050601f8311600181146200057c575f841562000567578287015190505b620005738582620004e9565b865550620005e2565b601f1984166200058c866200033a565b5f5b82811015620005b5578489015182556001820191506020850194506020810190506200058e565b86831015620005d55784890151620005d1601f891682620004cb565b8355505b6001600288020188555050505b505050505050565b5f81905092915050565b5f819050815f5260205f209050919050565b5f8154620006148162000306565b620006208186620005ea565b9450600182165f81146200063d5760018114620006535762000689565b60ff198316865281151582028601935062000689565b6200065e85620005f4565b5f5b83811015620006815781548189015260018201915060208101905062000660565b838801955050505b50505092915050565b5f6200069f828462000606565b915081905092915050565b5f819050919050565b620006be81620006aa565b82525050565b620006cf81620003bc565b82525050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6200070082620006d5565b9050919050565b6200071281620006f4565b82525050565b5f60a0820190506200072d5f830188620006b3565b6200073c6020830187620006b3565b6200074b6040830186620006b3565b6200075a6060830185620006c4565b62000769608083018462000707565b9695505050505050565b60805160a0516118f8620007955f395f61076201525f61072e01526118f85ff3fe608060405234801561000f575f80fd5b50600436106100cd575f3560e01c806370a082311161008a578063a9059cbb11610064578063a9059cbb14610227578063c370b04214610257578063d505accf14610275578063dd62ed3e14610291576100cd565b806370a08231146101a95780637ecebe00146101d957806395d89b4114610209576100cd565b806306fdde03146100d1578063095ea7b3146100ef57806318160ddd1461011f57806323b872dd1461013d578063313ce5671461016d5780633644e5151461018b575b5f80fd5b6100d96102c1565b6040516100e69190610eb8565b60405180910390f35b61010960048036038101906101049190610f69565b610350565b6040516101169190610fc1565b60405180910390f35b61012761043d565b6040516101349190610fe9565b60405180910390f35b61015760048036038101906101529190611002565b6104c2565b6040516101649190610fc1565b60405180910390f35b610175610723565b604051610182919061106d565b60405180910390f35b61019361072b565b6040516101a0919061109e565b60405180910390f35b6101c360048036038101906101be91906110b7565b610787565b6040516101d09190610fe9565b60405180910390f35b6101f360048036038101906101ee91906110b7565b610810565b6040516102009190610fe9565b60405180910390f35b610211610825565b60405161021e9190610eb8565b60405180910390f35b610241600480360381019061023c9190610f69565b6108b4565b60405161024e9190610fc1565b60405180910390f35b61025f6109e7565b60405161026c9190610eb8565b60405180910390f35b61028f600480360381019061028a9190611136565b610a72565b005b6102ab60048036038101906102a691906111d3565b610d5f565b6040516102b89190610fe9565b60405180910390f35b60605f80546102cf9061123e565b80601f01602080910402602001604051908101604052809291908181526020018280546102fb9061123e565b80156103465780601f1061031d57610100808354040283529160200191610346565b820191905f5260205f20905b81548152906001019060200180831161032957829003601f168201915b5050505050905090565b5f8160015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161042b9190610fe9565b60405180910390a36001905092915050565b5f610446610d7f565b73ffffffffffffffffffffffffffffffffffffffff1663fe3b2b885f6040518263ffffffff1660e01b815260040161047e9190611301565b602060405180830381865afa158015610499573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104bd9190611335565b905090565b5f8060015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146105ef578281610572919061138d565b60015f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055505b6105f7610d9a565b73ffffffffffffffffffffffffffffffffffffffff166393a3e3e98686866040518463ffffffff1660e01b8152600401610633939291906113cf565b6020604051808303815f875af115801561064f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610673919061142e565b6106b2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106a9906114c9565b60405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8560405161070f9190610fe9565b60405180910390a360019150509392505050565b5f6012905090565b5f7f000000000000000000000000000000000000000000000000000000000000000046146107605761075b610da4565b610782565b7f00000000000000000000000000000000000000000000000000000000000000005b905090565b5f610790610d7f565b73ffffffffffffffffffffffffffffffffffffffff166334d1fdaf835f6040518363ffffffff1660e01b81526004016107ca9291906114e7565b602060405180830381865afa1580156107e5573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108099190611335565b9050919050565b6002602052805f5260405f205f915090505481565b60605f80546108339061123e565b80601f016020809104026020016040519081016040528092919081815260200182805461085f9061123e565b80156108aa5780601f10610881576101008083540402835291602001916108aa565b820191905f5260205f20905b81548152906001019060200180831161088d57829003601f168201915b5050505050905090565b5f6108bd610d9a565b73ffffffffffffffffffffffffffffffffffffffff166393a3e3e93385856040518463ffffffff1660e01b81526004016108f9939291906113cf565b6020604051808303815f875af1158015610915573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610939919061142e565b610978576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161096f90611585565b60405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516109d59190610fe9565b60405180910390a36001905092915050565b5f80546109f39061123e565b80601f0160208091040260200160405190810160405280929190818152602001828054610a1f9061123e565b8015610a6a5780601f10610a4157610100808354040283529160200191610a6a565b820191905f5260205f20905b815481529060010190602001808311610a4d57829003601f168201915b505050505081565b42841015610ab5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aac906115ed565b60405180910390fd5b5f6001610ac061072b565b7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98a8a8a60025f8f73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f815480929190600101919050558b604051602001610b459695949392919061160b565b60405160208183030381529060405280519060200120604051602001610b6c9291906116de565b604051602081830303815290604052805190602001208585856040515f8152602001604052604051610ba19493929190611714565b6020604051602081039080840390855afa158015610bc1573d5f803e3d5ffd5b5050506020604051035190505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614158015610c3457508773ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b610c73576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c6a906117a1565b60405180910390fd5b8560015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550508573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92587604051610d4e9190610fe9565b60405180910390a350505050505050565b6001602052815f5260405f20602052805f5260405f205f91509150505481565b5f734381dc2ab14285160c808659aee005d51255add7905090565b5f62696969905090565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f5f604051610dd4919061185b565b60405180910390207fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc64630604051602001610e13959493929190611871565b60405160208183030381529060405280519060200120905090565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610e65578082015181840152602081019050610e4a565b5f8484015250505050565b5f601f19601f8301169050919050565b5f610e8a82610e2e565b610e948185610e38565b9350610ea4818560208601610e48565b610ead81610e70565b840191505092915050565b5f6020820190508181035f830152610ed08184610e80565b905092915050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610f0582610edc565b9050919050565b610f1581610efb565b8114610f1f575f80fd5b50565b5f81359050610f3081610f0c565b92915050565b5f819050919050565b610f4881610f36565b8114610f52575f80fd5b50565b5f81359050610f6381610f3f565b92915050565b5f8060408385031215610f7f57610f7e610ed8565b5b5f610f8c85828601610f22565b9250506020610f9d85828601610f55565b9150509250929050565b5f8115159050919050565b610fbb81610fa7565b82525050565b5f602082019050610fd45f830184610fb2565b92915050565b610fe381610f36565b82525050565b5f602082019050610ffc5f830184610fda565b92915050565b5f805f6060848603121561101957611018610ed8565b5b5f61102686828701610f22565b935050602061103786828701610f22565b925050604061104886828701610f55565b9150509250925092565b5f60ff82169050919050565b61106781611052565b82525050565b5f6020820190506110805f83018461105e565b92915050565b5f819050919050565b61109881611086565b82525050565b5f6020820190506110b15f83018461108f565b92915050565b5f602082840312156110cc576110cb610ed8565b5b5f6110d984828501610f22565b91505092915050565b6110eb81611052565b81146110f5575f80fd5b50565b5f81359050611106816110e2565b92915050565b61111581611086565b811461111f575f80fd5b50565b5f813590506111308161110c565b92915050565b5f805f805f805f60e0888a03121561115157611150610ed8565b5b5f61115e8a828b01610f22565b975050602061116f8a828b01610f22565b96505060406111808a828b01610f55565b95505060606111918a828b01610f55565b94505060806111a28a828b016110f8565b93505060a06111b38a828b01611122565b92505060c06111c48a828b01611122565b91505092959891949750929550565b5f80604083850312156111e9576111e8610ed8565b5b5f6111f685828601610f22565b925050602061120785828601610f22565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061125557607f821691505b60208210810361126857611267611211565b5b50919050565b5f819050815f5260205f209050919050565b5f815461128c8161123e565b6112968186610e38565b9450600182165f81146112b057600181146112c6576112f8565b60ff1983168652811515602002860193506112f8565b6112cf8561126e565b5f5b838110156112f0578154818901526001820191506020810190506112d1565b808801955050505b50505092915050565b5f6020820190508181035f8301526113198184611280565b905092915050565b5f8151905061132f81610f3f565b92915050565b5f6020828403121561134a57611349610ed8565b5b5f61135784828501611321565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61139782610f36565b91506113a283610f36565b92508282039050818111156113ba576113b9611360565b5b92915050565b6113c981610efb565b82525050565b5f6060820190506113e25f8301866113c0565b6113ef60208301856113c0565b6113fc6040830184610fda565b949350505050565b61140d81610fa7565b8114611417575f80fd5b50565b5f8151905061142881611404565b92915050565b5f6020828403121561144357611442610ed8565b5b5f6114508482850161141a565b91505092915050565b7f506f6c6172697345524332303a206661696c656420746f2073656e642062616e5f8201527f6b20746f6b656e73000000000000000000000000000000000000000000000000602082015250565b5f6114b3602883610e38565b91506114be82611459565b604082019050919050565b5f6020820190508181035f8301526114e0816114a7565b9050919050565b5f6040820190506114fa5f8301856113c0565b818103602083015261150c8184611280565b90509392505050565b7f506f6c6172697345524332303a206661696c656420746f2073656e6420746f6b5f8201527f656e730000000000000000000000000000000000000000000000000000000000602082015250565b5f61156f602383610e38565b915061157a82611515565b604082019050919050565b5f6020820190508181035f83015261159c81611563565b9050919050565b7f5045524d49545f444541444c494e455f455850495245440000000000000000005f82015250565b5f6115d7601783610e38565b91506115e2826115a3565b602082019050919050565b5f6020820190508181035f830152611604816115cb565b9050919050565b5f60c08201905061161e5f83018961108f565b61162b60208301886113c0565b61163860408301876113c0565b6116456060830186610fda565b6116526080830185610fda565b61165f60a0830184610fda565b979650505050505050565b5f81905092915050565b7f19010000000000000000000000000000000000000000000000000000000000005f82015250565b5f6116a860028361166a565b91506116b382611674565b600282019050919050565b5f819050919050565b6116d86116d382611086565b6116be565b82525050565b5f6116e88261169c565b91506116f482856116c7565b60208201915061170482846116c7565b6020820191508190509392505050565b5f6080820190506117275f83018761108f565b611734602083018661105e565b611741604083018561108f565b61174e606083018461108f565b95945050505050565b7f494e56414c49445f5349474e45520000000000000000000000000000000000005f82015250565b5f61178b600e83610e38565b915061179682611757565b602082019050919050565b5f6020820190508181035f8301526117b88161177f565b9050919050565b5f81905092915050565b5f819050815f5260205f209050919050565b5f81546117e78161123e565b6117f181866117bf565b9450600182165f811461180b576001811461182057611852565b60ff1983168652811515820286019350611852565b611829856117c9565b5f5b8381101561184a5781548189015260018201915060208101905061182b565b838801955050505b50505092915050565b5f61186682846117db565b915081905092915050565b5f60a0820190506118845f83018861108f565b611891602083018761108f565b61189e604083018661108f565b6118ab6060830185610fda565b6118b860808301846113c0565b969550505050505056fea2646970667358221220c9e323a029a387286d49bd14cacaecb17d3b802337e2c28054393c23b099ed2064736f6c63430008140033",
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
