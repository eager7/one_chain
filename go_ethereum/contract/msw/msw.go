// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package msw

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// MswABI is the input ABI used to generate the binding from.
const MswABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"removeOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"revokeConfirmation\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"confirmations\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"calcMaxWithdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"pending\",\"type\":\"bool\"},{\"name\":\"executed\",\"type\":\"bool\"}],\"name\":\"getTransactionCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dailyLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastDay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"addOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"isConfirmed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"getConfirmationCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transactions\",\"outputs\":[{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"executed\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"from\",\"type\":\"uint256\"},{\"name\":\"to\",\"type\":\"uint256\"},{\"name\":\"pending\",\"type\":\"bool\"},{\"name\":\"executed\",\"type\":\"bool\"}],\"name\":\"getTransactionIds\",\"outputs\":[{\"name\":\"_transactionIds\",\"type\":\"uint256[]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"getConfirmations\",\"outputs\":[{\"name\":\"_confirmations\",\"type\":\"address[]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transactionCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_required\",\"type\":\"uint256\"}],\"name\":\"changeRequirement\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"confirmTransaction\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"submitTransaction\",\"outputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_dailyLimit\",\"type\":\"uint256\"}],\"name\":\"changeDailyLimit\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_OWNER_COUNT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"required\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"replaceOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spentToday\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owners\",\"type\":\"address[]\"},{\"name\":\"_required\",\"type\":\"uint256\"},{\"name\":\"_dailyLimit\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"constructor\"},{\"payable\":true,\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"dailyLimit\",\"type\":\"uint256\"}],\"name\":\"DailyLimitChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Confirmation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Revocation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Submission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Execution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"ExecutionFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"RequirementChange\",\"type\":\"event\"}]"

// MswBin is the compiled bytecode used for deploying new contracts.
const MswBin = `606060405234156200000d57fe5b6040516200258438038062002584833981016040528080518201919060200180519060200190919080519060200190919050505b82825b600082518260328211806200005857508181115b80620000645750600081145b80620000705750600082145b156200007c5760006000fd5b600092505b8451831015620001b5576002600086858151811015156200009e57fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16806200012a5750600085848151811015156200010857fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff16145b15620001365760006000fd5b60016002600087868151811015156200014b57fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b828060010193505062000081565b8460039080519060200190620001cd929190620001ed565b50836004819055505b5b5050505050806006819055505b505050620002c2565b82805482825590600052602060002090810192821562000269579160200282015b82811115620002685782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550916020019190600101906200020e565b5b5090506200027891906200027c565b5090565b620002bf91905b80821115620002bb57600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690555060010162000283565b5090565b90565b6122b280620002d26000396000f30060606040523615610152576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063025e7c27146101b3578063173825d91461021357806320ea8d86146102495780632f54bf6e146102695780633411c81c146102b75780634bc9fdc21461030e578063547415251461033457806367eeba0c146103755780636b0c932d1461039b5780637065cb48146103c1578063784547a7146103f75780638b51d13f1461042f5780639ace38c214610463578063a0e67e2b1461055e578063a8abe69a146105d3578063b5dc40c314610675578063b77bf600146106f8578063ba51a6df1461071e578063c01a8c841461073e578063c64274741461075e578063cea08621146107f4578063d74f8edd14610814578063dc8452cd1461083a578063e20056e614610860578063ee22610b146108b5578063f059cf2b146108d5575b6101b15b60003411156101ae573373ffffffffffffffffffffffffffffffffffffffff167fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c346040518082815260200191505060405180910390a25b5b565b005b34156101bb57fe5b6101d160048080359060200190919050506108fb565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561021b57fe5b610247600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190505061093b565b005b341561025157fe5b6102676004808035906020019091905050610be3565b005b341561027157fe5b61029d600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610d94565b604051808215151515815260200191505060405180910390f35b34156102bf57fe5b6102f4600480803590602001909190803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610db4565b604051808215151515815260200191505060405180910390f35b341561031657fe5b61031e610de3565b6040518082815260200191505060405180910390f35b341561033c57fe5b61035f600480803515159060200190919080351515906020019091905050610e20565b6040518082815260200191505060405180910390f35b341561037d57fe5b610385610eb7565b6040518082815260200191505060405180910390f35b34156103a357fe5b6103ab610ebd565b6040518082815260200191505060405180910390f35b34156103c957fe5b6103f5600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610ec3565b005b34156103ff57fe5b61041560048080359060200190919050506110c6565b604051808215151515815260200191505060405180910390f35b341561043757fe5b61044d60048080359060200190919050506111b0565b6040518082815260200191505060405180910390f35b341561046b57fe5b6104816004808035906020019091905050611280565b604051808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001848152602001806020018315151515815260200182810382528481815460018160011615610100020316600290048152602001915080546001816001161561010002031660029004801561054c5780601f106105215761010080835404028352916020019161054c565b820191906000526020600020905b81548152906001019060200180831161052f57829003601f168201915b50509550505050505060405180910390f35b341561056657fe5b61056e6112dc565b60405180806020018281038252838181518152602001915080519060200190602002808383600083146105c0575b8051825260208311156105c05760208201915060208101905060208303925061059c565b5050509050019250505060405180910390f35b34156105db57fe5b610610600480803590602001909190803590602001909190803515159060200190919080351515906020019091905050611371565b6040518080602001828103825283818151815260200191508051906020019060200280838360008314610662575b8051825260208311156106625760208201915060208101905060208303925061063e565b5050509050019250505060405180910390f35b341561067d57fe5b61069360048080359060200190919050506114d5565b60405180806020018281038252838181518152602001915080519060200190602002808383600083146106e5575b8051825260208311156106e5576020820191506020810190506020830392506106c1565b5050509050019250505060405180910390f35b341561070057fe5b610708611707565b6040518082815260200191505060405180910390f35b341561072657fe5b61073c600480803590602001909190505061170d565b005b341561074657fe5b61075c60048080359060200190919050506117c4565b005b341561076657fe5b6107de600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001909190803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919050506119a9565b6040518082815260200191505060405180910390f35b34156107fc57fe5b61081260048080359060200190919050506119c9565b005b341561081c57fe5b610824611a47565b6040518082815260200191505060405180910390f35b341561084257fe5b61084a611a4c565b6040518082815260200191505060405180910390f35b341561086857fe5b6108b3600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050611a52565b005b34156108bd57fe5b6108d36004808035906020019091905050611d77565b005b34156108dd57fe5b6108e5611fb2565b6040518082815260200191505060405180910390f35b60038181548110151561090a57fe5b906000526020600020900160005b915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60003073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156109785760006000fd5b81600260008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615156109d25760006000fd5b6000600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600091505b600160038054905003821015610b5e578273ffffffffffffffffffffffffffffffffffffffff16600383815481101515610a6557fe5b906000526020600020900160005b9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610b50576003600160038054905003815481101515610ac557fe5b906000526020600020900160005b9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600383815481101515610b0157fe5b906000526020600020900160005b6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610b5e565b5b8180600101925050610a2f565b6001600381818054905003915081610b769190612161565b506003805490506004541115610b9557610b9460038054905061170d565b5b8273ffffffffffffffffffffffffffffffffffffffff167f8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b9060405180905060405180910390a25b5b505b5050565b33600260008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515610c3d5760006000fd5b81336001600083815260200190815260200160002060008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515610ca95760006000fd5b836000600082815260200190815260200160002060030160009054906101000a900460ff1615610cd95760006000fd5b60006001600087815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550843373ffffffffffffffffffffffffffffffffffffffff167ff6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e960405180905060405180910390a35b5b505b50505b5050565b60026020528060005260406000206000915054906101000a900460ff1681565b60016020528160005260406000206020528060005260406000206000915091509054906101000a900460ff1681565b60006201518060075401421115610dfe576006549050610e1d565b6008546006541015610e135760009050610e1d565b6008546006540390505b90565b60006000600090505b600554811015610eaf57838015610e6157506000600082815260200190815260200160002060030160009054906101000a900460ff16155b80610e955750828015610e9457506000600082815260200190815260200160002060030160009054906101000a900460ff165b5b15610ea1576001820191505b5b8080600101915050610e29565b5b5092915050565b60065481565b60075481565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610efe5760006000fd5b80600260008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615610f575760006000fd5b8160008173ffffffffffffffffffffffffffffffffffffffff161415610f7d5760006000fd5b6001600380549050016004546032821180610f9757508181115b80610fa25750600081145b80610fad5750600082145b15610fb85760006000fd5b6001600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555060038054806001018281611024919061218d565b916000526020600020900160005b87909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550508473ffffffffffffffffffffffffffffffffffffffff167ff39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d60405180905060405180910390a25b5b50505b505b505b50565b60006000600060009150600090505b6003805490508110156111a85760016000858152602001908152602001600020600060038381548110151561110657fe5b906000526020600020900160005b9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615611187576001820191505b60045482141561119a57600192506111a9565b5b80806001019150506110d5565b5b5050919050565b60006000600090505b600380549050811015611279576001600084815260200190815260200160002060006003838154811015156111ea57fe5b906000526020600020900160005b9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161561126b576001820191505b5b80806001019150506111b9565b5b50919050565b60006020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600101549080600201908060030160009054906101000a900460ff16905084565b6112e46121b9565b600380548060200260200160405190810160405280929190818152602001828054801561136657602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161131c575b505050505090505b90565b6113796121cd565b6113816121cd565b600060006005546040518059106113955750595b908082528060200260200182016040525b50925060009150600090505b600554811015611455578580156113ea57506000600082815260200190815260200160002060030160009054906101000a900460ff16155b8061141e575084801561141d57506000600082815260200190815260200160002060030160009054906101000a900460ff165b5b156114475780838381518110151561143257fe5b90602001906020020181815250506001820191505b5b80806001019150506113b2565b8787036040518059106114655750595b908082528060200260200182016040525b5093508790505b868110156114c957828181518110151561149357fe5b90602001906020020151848983038151811015156114ad57fe5b90602001906020020181815250505b808060010191505061147d565b5b505050949350505050565b6114dd6121b9565b6114e56121b9565b600060006003805490506040518059106114fc5750595b908082528060200260200182016040525b50925060009150600090505b60038054905081101561165f5760016000868152602001908152602001600020600060038381548110151561154a57fe5b906000526020600020900160005b9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615611651576003818154811015156115d357fe5b906000526020600020900160005b9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16838381518110151561160e57fe5b9060200190602002019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506001820191505b5b8080600101915050611519565b8160405180591061166d5750595b908082528060200260200182016040525b509350600090505b818110156116fe57828181518110151561169c57fe5b9060200190602002015184828151811015156116b457fe5b9060200190602002019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250505b8080600101915050611686565b5b505050919050565b60055481565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156117485760006000fd5b60038054905081603282118061175d57508181115b806117685750600081145b806117735750600082145b1561177e5760006000fd5b826004819055507fa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a836040518082815260200191505060405180910390a15b5b50505b50565b33600260008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151561181e5760006000fd5b8160006000600083815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561187a5760006000fd5b82336001600083815260200190815260200160002060008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156118e55760006000fd5b60016001600087815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550843373ffffffffffffffffffffffffffffffffffffffff167f4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef60405180905060405180910390a361199e85611d77565b5b5b50505b505b5050565b60006119b6848484611fb8565b90506119c1816117c4565b5b9392505050565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611a045760006000fd5b806006819055507fc71bdc6afaf9b1aa90a7078191d4fc1adf3bf680fca3183697df6b0dc226bca2816040518082815260200191505060405180910390a15b5b50565b603281565b60045481565b60003073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611a8f5760006000fd5b82600260008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515611ae95760006000fd5b82600260008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615611b425760006000fd5b600092505b600380549050831015611c30578473ffffffffffffffffffffffffffffffffffffffff16600384815481101515611b7a57fe5b906000526020600020900160005b9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415611c225783600384815481101515611bd357fe5b906000526020600020900160005b6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611c30565b5b8280600101935050611b47565b6000600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508473ffffffffffffffffffffffffffffffffffffffff167f8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b9060405180905060405180910390a28373ffffffffffffffffffffffffffffffffffffffff167ff39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d60405180905060405180910390a25b5b505b505b505050565b60006000826000600082815260200190815260200160002060030160009054906101000a900460ff1615611dab5760006000fd5b600060008581526020019081526020016000209250611dc9846110c6565b91508180611e045750600083600201805460018160011615610100020316600290049050148015611e035750611e02836001015461210f565b5b5b15611faa5760018360030160006101000a81548160ff021916908315150217905550811515611e425782600101546008600082825401925050819055505b8260000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168360010154846002016040518082805460018160011615610100020316600290048015611eeb5780601f10611ec057610100808354040283529160200191611eeb565b820191906000526020600020905b815481529060010190602001808311611ece57829003601f168201915b505091505060006040518083038185876185025a03f19250505015611f3f57837f33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed7560405180905060405180910390a2611fa9565b837f526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b7923660405180905060405180910390a260008360030160006101000a81548160ff021916908315150217905550811515611fa85782600101546008600082825403925050819055505b5b5b5b5b50505050565b60085481565b60008360008173ffffffffffffffffffffffffffffffffffffffff161415611fe05760006000fd5b60055491506080604051908101604052808673ffffffffffffffffffffffffffffffffffffffff168152602001858152602001848152602001600015158152506000600084815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020190805190602001906120a09291906121e1565b5060608201518160030160006101000a81548160ff0219169083151502179055509050506001600560008282540192505081905550817fc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e5160405180905060405180910390a25b5b509392505050565b60006201518060075401421115612130574260078190555060006008819055505b6006548260085401118061214957506008548260085401105b15612157576000905061215c565b600190505b919050565b815481835581811511612188578183600052602060002091820191016121879190612261565b5b505050565b8154818355818115116121b4578183600052602060002091820191016121b39190612261565b5b505050565b602060405190810160405280600081525090565b602060405190810160405280600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061222257805160ff1916838001178555612250565b82800160010185558215612250579182015b8281111561224f578251825591602001919060010190612234565b5b50905061225d9190612261565b5090565b61228391905b8082111561227f576000816000905550600101612267565b5090565b905600a165627a7a72305820ea4e2bd4cd3bd1317fe18a6300906f54aea41437433d2941ccd0320c8bc088bb002900000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000040000000000000000000000009f7dfab2222a473284205cddf08a677726d786a00000000000000000000000005210c4dcd7eb899a1274fd6471adec9896ae05aa0000000000000000000000001d805bc00b8fa3c96ae6c8fa97b2fd24b19a9801000000000000000000000000aca7bd07a8c207f7964261c2cf1e0fbfcff37836`

// DeployMsw deploys a new Ethereum contract, binding an instance of Msw to it.
func DeployMsw(auth *bind.TransactOpts, backend bind.ContractBackend, _owners []common.Address, _required *big.Int, _dailyLimit *big.Int) (common.Address, *types.Transaction, *Msw, error) {
	parsed, err := abi.JSON(strings.NewReader(MswABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MswBin), backend, _owners, _required, _dailyLimit)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Msw{MswCaller: MswCaller{contract: contract}, MswTransactor: MswTransactor{contract: contract}, MswFilterer: MswFilterer{contract: contract}}, nil
}

// Msw is an auto generated Go binding around an Ethereum contract.
type Msw struct {
	MswCaller     // Read-only binding to the contract
	MswTransactor // Write-only binding to the contract
	MswFilterer   // Log filterer for contract events
}

// MswCaller is an auto generated read-only Go binding around an Ethereum contract.
type MswCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MswTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MswTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MswFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MswFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MswSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MswSession struct {
	Contract     *Msw              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MswCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MswCallerSession struct {
	Contract *MswCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MswTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MswTransactorSession struct {
	Contract     *MswTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MswRaw is an auto generated low-level Go binding around an Ethereum contract.
type MswRaw struct {
	Contract *Msw // Generic contract binding to access the raw methods on
}

// MswCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MswCallerRaw struct {
	Contract *MswCaller // Generic read-only contract binding to access the raw methods on
}

// MswTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MswTransactorRaw struct {
	Contract *MswTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMsw creates a new instance of Msw, bound to a specific deployed contract.
func NewMsw(address common.Address, backend bind.ContractBackend) (*Msw, error) {
	contract, err := bindMsw(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Msw{MswCaller: MswCaller{contract: contract}, MswTransactor: MswTransactor{contract: contract}, MswFilterer: MswFilterer{contract: contract}}, nil
}

// NewMswCaller creates a new read-only instance of Msw, bound to a specific deployed contract.
func NewMswCaller(address common.Address, caller bind.ContractCaller) (*MswCaller, error) {
	contract, err := bindMsw(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MswCaller{contract: contract}, nil
}

// NewMswTransactor creates a new write-only instance of Msw, bound to a specific deployed contract.
func NewMswTransactor(address common.Address, transactor bind.ContractTransactor) (*MswTransactor, error) {
	contract, err := bindMsw(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MswTransactor{contract: contract}, nil
}

// NewMswFilterer creates a new log filterer instance of Msw, bound to a specific deployed contract.
func NewMswFilterer(address common.Address, filterer bind.ContractFilterer) (*MswFilterer, error) {
	contract, err := bindMsw(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MswFilterer{contract: contract}, nil
}

// bindMsw binds a generic wrapper to an already deployed contract.
func bindMsw(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MswABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Msw *MswRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Msw.Contract.MswCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Msw *MswRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Msw.Contract.MswTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Msw *MswRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Msw.Contract.MswTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Msw *MswCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Msw.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Msw *MswTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Msw.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Msw *MswTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Msw.Contract.contract.Transact(opts, method, params...)
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() constant returns(uint256)
func (_Msw *MswCaller) MAXOWNERCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Msw.contract.Call(opts, out, "MAX_OWNER_COUNT")
	return *ret0, err
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() constant returns(uint256)
func (_Msw *MswSession) MAXOWNERCOUNT() (*big.Int, error) {
	return _Msw.Contract.MAXOWNERCOUNT(&_Msw.CallOpts)
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() constant returns(uint256)
func (_Msw *MswCallerSession) MAXOWNERCOUNT() (*big.Int, error) {
	return _Msw.Contract.MAXOWNERCOUNT(&_Msw.CallOpts)
}

// CalcMaxWithdraw is a free data retrieval call binding the contract method 0x4bc9fdc2.
//
// Solidity: function calcMaxWithdraw() constant returns(uint256)
func (_Msw *MswCaller) CalcMaxWithdraw(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Msw.contract.Call(opts, out, "calcMaxWithdraw")
	return *ret0, err
}

// CalcMaxWithdraw is a free data retrieval call binding the contract method 0x4bc9fdc2.
//
// Solidity: function calcMaxWithdraw() constant returns(uint256)
func (_Msw *MswSession) CalcMaxWithdraw() (*big.Int, error) {
	return _Msw.Contract.CalcMaxWithdraw(&_Msw.CallOpts)
}

// CalcMaxWithdraw is a free data retrieval call binding the contract method 0x4bc9fdc2.
//
// Solidity: function calcMaxWithdraw() constant returns(uint256)
func (_Msw *MswCallerSession) CalcMaxWithdraw() (*big.Int, error) {
	return _Msw.Contract.CalcMaxWithdraw(&_Msw.CallOpts)
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations(uint256 , address ) constant returns(bool)
func (_Msw *MswCaller) Confirmations(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Msw.contract.Call(opts, out, "confirmations", arg0, arg1)
	return *ret0, err
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations(uint256 , address ) constant returns(bool)
func (_Msw *MswSession) Confirmations(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Msw.Contract.Confirmations(&_Msw.CallOpts, arg0, arg1)
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations(uint256 , address ) constant returns(bool)
func (_Msw *MswCallerSession) Confirmations(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Msw.Contract.Confirmations(&_Msw.CallOpts, arg0, arg1)
}

// DailyLimit is a free data retrieval call binding the contract method 0x67eeba0c.
//
// Solidity: function dailyLimit() constant returns(uint256)
func (_Msw *MswCaller) DailyLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Msw.contract.Call(opts, out, "dailyLimit")
	return *ret0, err
}

// DailyLimit is a free data retrieval call binding the contract method 0x67eeba0c.
//
// Solidity: function dailyLimit() constant returns(uint256)
func (_Msw *MswSession) DailyLimit() (*big.Int, error) {
	return _Msw.Contract.DailyLimit(&_Msw.CallOpts)
}

// DailyLimit is a free data retrieval call binding the contract method 0x67eeba0c.
//
// Solidity: function dailyLimit() constant returns(uint256)
func (_Msw *MswCallerSession) DailyLimit() (*big.Int, error) {
	return _Msw.Contract.DailyLimit(&_Msw.CallOpts)
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 transactionId) constant returns(uint256 count)
func (_Msw *MswCaller) GetConfirmationCount(opts *bind.CallOpts, transactionId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Msw.contract.Call(opts, out, "getConfirmationCount", transactionId)
	return *ret0, err
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 transactionId) constant returns(uint256 count)
func (_Msw *MswSession) GetConfirmationCount(transactionId *big.Int) (*big.Int, error) {
	return _Msw.Contract.GetConfirmationCount(&_Msw.CallOpts, transactionId)
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 transactionId) constant returns(uint256 count)
func (_Msw *MswCallerSession) GetConfirmationCount(transactionId *big.Int) (*big.Int, error) {
	return _Msw.Contract.GetConfirmationCount(&_Msw.CallOpts, transactionId)
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(uint256 transactionId) constant returns(address[] _confirmations)
func (_Msw *MswCaller) GetConfirmations(opts *bind.CallOpts, transactionId *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Msw.contract.Call(opts, out, "getConfirmations", transactionId)
	return *ret0, err
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(uint256 transactionId) constant returns(address[] _confirmations)
func (_Msw *MswSession) GetConfirmations(transactionId *big.Int) ([]common.Address, error) {
	return _Msw.Contract.GetConfirmations(&_Msw.CallOpts, transactionId)
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(uint256 transactionId) constant returns(address[] _confirmations)
func (_Msw *MswCallerSession) GetConfirmations(transactionId *big.Int) ([]common.Address, error) {
	return _Msw.Contract.GetConfirmations(&_Msw.CallOpts, transactionId)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_Msw *MswCaller) GetOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Msw.contract.Call(opts, out, "getOwners")
	return *ret0, err
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_Msw *MswSession) GetOwners() ([]common.Address, error) {
	return _Msw.Contract.GetOwners(&_Msw.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_Msw *MswCallerSession) GetOwners() ([]common.Address, error) {
	return _Msw.Contract.GetOwners(&_Msw.CallOpts)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(bool pending, bool executed) constant returns(uint256 count)
func (_Msw *MswCaller) GetTransactionCount(opts *bind.CallOpts, pending bool, executed bool) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Msw.contract.Call(opts, out, "getTransactionCount", pending, executed)
	return *ret0, err
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(bool pending, bool executed) constant returns(uint256 count)
func (_Msw *MswSession) GetTransactionCount(pending bool, executed bool) (*big.Int, error) {
	return _Msw.Contract.GetTransactionCount(&_Msw.CallOpts, pending, executed)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(bool pending, bool executed) constant returns(uint256 count)
func (_Msw *MswCallerSession) GetTransactionCount(pending bool, executed bool) (*big.Int, error) {
	return _Msw.Contract.GetTransactionCount(&_Msw.CallOpts, pending, executed)
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(uint256 from, uint256 to, bool pending, bool executed) constant returns(uint256[] _transactionIds)
func (_Msw *MswCaller) GetTransactionIds(opts *bind.CallOpts, from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Msw.contract.Call(opts, out, "getTransactionIds", from, to, pending, executed)
	return *ret0, err
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(uint256 from, uint256 to, bool pending, bool executed) constant returns(uint256[] _transactionIds)
func (_Msw *MswSession) GetTransactionIds(from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	return _Msw.Contract.GetTransactionIds(&_Msw.CallOpts, from, to, pending, executed)
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(uint256 from, uint256 to, bool pending, bool executed) constant returns(uint256[] _transactionIds)
func (_Msw *MswCallerSession) GetTransactionIds(from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	return _Msw.Contract.GetTransactionIds(&_Msw.CallOpts, from, to, pending, executed)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 transactionId) constant returns(bool)
func (_Msw *MswCaller) IsConfirmed(opts *bind.CallOpts, transactionId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Msw.contract.Call(opts, out, "isConfirmed", transactionId)
	return *ret0, err
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 transactionId) constant returns(bool)
func (_Msw *MswSession) IsConfirmed(transactionId *big.Int) (bool, error) {
	return _Msw.Contract.IsConfirmed(&_Msw.CallOpts, transactionId)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 transactionId) constant returns(bool)
func (_Msw *MswCallerSession) IsConfirmed(transactionId *big.Int) (bool, error) {
	return _Msw.Contract.IsConfirmed(&_Msw.CallOpts, transactionId)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) constant returns(bool)
func (_Msw *MswCaller) IsOwner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Msw.contract.Call(opts, out, "isOwner", arg0)
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) constant returns(bool)
func (_Msw *MswSession) IsOwner(arg0 common.Address) (bool, error) {
	return _Msw.Contract.IsOwner(&_Msw.CallOpts, arg0)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) constant returns(bool)
func (_Msw *MswCallerSession) IsOwner(arg0 common.Address) (bool, error) {
	return _Msw.Contract.IsOwner(&_Msw.CallOpts, arg0)
}

// LastDay is a free data retrieval call binding the contract method 0x6b0c932d.
//
// Solidity: function lastDay() constant returns(uint256)
func (_Msw *MswCaller) LastDay(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Msw.contract.Call(opts, out, "lastDay")
	return *ret0, err
}

// LastDay is a free data retrieval call binding the contract method 0x6b0c932d.
//
// Solidity: function lastDay() constant returns(uint256)
func (_Msw *MswSession) LastDay() (*big.Int, error) {
	return _Msw.Contract.LastDay(&_Msw.CallOpts)
}

// LastDay is a free data retrieval call binding the contract method 0x6b0c932d.
//
// Solidity: function lastDay() constant returns(uint256)
func (_Msw *MswCallerSession) LastDay() (*big.Int, error) {
	return _Msw.Contract.LastDay(&_Msw.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) constant returns(address)
func (_Msw *MswCaller) Owners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Msw.contract.Call(opts, out, "owners", arg0)
	return *ret0, err
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) constant returns(address)
func (_Msw *MswSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _Msw.Contract.Owners(&_Msw.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) constant returns(address)
func (_Msw *MswCallerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _Msw.Contract.Owners(&_Msw.CallOpts, arg0)
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() constant returns(uint256)
func (_Msw *MswCaller) Required(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Msw.contract.Call(opts, out, "required")
	return *ret0, err
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() constant returns(uint256)
func (_Msw *MswSession) Required() (*big.Int, error) {
	return _Msw.Contract.Required(&_Msw.CallOpts)
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() constant returns(uint256)
func (_Msw *MswCallerSession) Required() (*big.Int, error) {
	return _Msw.Contract.Required(&_Msw.CallOpts)
}

// SpentToday is a free data retrieval call binding the contract method 0xf059cf2b.
//
// Solidity: function spentToday() constant returns(uint256)
func (_Msw *MswCaller) SpentToday(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Msw.contract.Call(opts, out, "spentToday")
	return *ret0, err
}

// SpentToday is a free data retrieval call binding the contract method 0xf059cf2b.
//
// Solidity: function spentToday() constant returns(uint256)
func (_Msw *MswSession) SpentToday() (*big.Int, error) {
	return _Msw.Contract.SpentToday(&_Msw.CallOpts)
}

// SpentToday is a free data retrieval call binding the contract method 0xf059cf2b.
//
// Solidity: function spentToday() constant returns(uint256)
func (_Msw *MswCallerSession) SpentToday() (*big.Int, error) {
	return _Msw.Contract.SpentToday(&_Msw.CallOpts)
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() constant returns(uint256)
func (_Msw *MswCaller) TransactionCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Msw.contract.Call(opts, out, "transactionCount")
	return *ret0, err
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() constant returns(uint256)
func (_Msw *MswSession) TransactionCount() (*big.Int, error) {
	return _Msw.Contract.TransactionCount(&_Msw.CallOpts)
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() constant returns(uint256)
func (_Msw *MswCallerSession) TransactionCount() (*big.Int, error) {
	return _Msw.Contract.TransactionCount(&_Msw.CallOpts)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) constant returns(address destination, uint256 value, bytes data, bool executed)
func (_Msw *MswCaller) Transactions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	ret := new(struct {
		Destination common.Address
		Value       *big.Int
		Data        []byte
		Executed    bool
	})
	out := ret
	err := _Msw.contract.Call(opts, out, "transactions", arg0)
	return *ret, err
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) constant returns(address destination, uint256 value, bytes data, bool executed)
func (_Msw *MswSession) Transactions(arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	return _Msw.Contract.Transactions(&_Msw.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) constant returns(address destination, uint256 value, bytes data, bool executed)
func (_Msw *MswCallerSession) Transactions(arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	return _Msw.Contract.Transactions(&_Msw.CallOpts, arg0)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner) returns()
func (_Msw *MswTransactor) AddOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _Msw.contract.Transact(opts, "addOwner", owner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner) returns()
func (_Msw *MswSession) AddOwner(owner common.Address) (*types.Transaction, error) {
	return _Msw.Contract.AddOwner(&_Msw.TransactOpts, owner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner) returns()
func (_Msw *MswTransactorSession) AddOwner(owner common.Address) (*types.Transaction, error) {
	return _Msw.Contract.AddOwner(&_Msw.TransactOpts, owner)
}

// ChangeDailyLimit is a paid mutator transaction binding the contract method 0xcea08621.
//
// Solidity: function changeDailyLimit(uint256 _dailyLimit) returns()
func (_Msw *MswTransactor) ChangeDailyLimit(opts *bind.TransactOpts, _dailyLimit *big.Int) (*types.Transaction, error) {
	return _Msw.contract.Transact(opts, "changeDailyLimit", _dailyLimit)
}

// ChangeDailyLimit is a paid mutator transaction binding the contract method 0xcea08621.
//
// Solidity: function changeDailyLimit(uint256 _dailyLimit) returns()
func (_Msw *MswSession) ChangeDailyLimit(_dailyLimit *big.Int) (*types.Transaction, error) {
	return _Msw.Contract.ChangeDailyLimit(&_Msw.TransactOpts, _dailyLimit)
}

// ChangeDailyLimit is a paid mutator transaction binding the contract method 0xcea08621.
//
// Solidity: function changeDailyLimit(uint256 _dailyLimit) returns()
func (_Msw *MswTransactorSession) ChangeDailyLimit(_dailyLimit *big.Int) (*types.Transaction, error) {
	return _Msw.Contract.ChangeDailyLimit(&_Msw.TransactOpts, _dailyLimit)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _required) returns()
func (_Msw *MswTransactor) ChangeRequirement(opts *bind.TransactOpts, _required *big.Int) (*types.Transaction, error) {
	return _Msw.contract.Transact(opts, "changeRequirement", _required)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _required) returns()
func (_Msw *MswSession) ChangeRequirement(_required *big.Int) (*types.Transaction, error) {
	return _Msw.Contract.ChangeRequirement(&_Msw.TransactOpts, _required)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _required) returns()
func (_Msw *MswTransactorSession) ChangeRequirement(_required *big.Int) (*types.Transaction, error) {
	return _Msw.Contract.ChangeRequirement(&_Msw.TransactOpts, _required)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 transactionId) returns()
func (_Msw *MswTransactor) ConfirmTransaction(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _Msw.contract.Transact(opts, "confirmTransaction", transactionId)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 transactionId) returns()
func (_Msw *MswSession) ConfirmTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _Msw.Contract.ConfirmTransaction(&_Msw.TransactOpts, transactionId)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 transactionId) returns()
func (_Msw *MswTransactorSession) ConfirmTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _Msw.Contract.ConfirmTransaction(&_Msw.TransactOpts, transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 transactionId) returns()
func (_Msw *MswTransactor) ExecuteTransaction(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _Msw.contract.Transact(opts, "executeTransaction", transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 transactionId) returns()
func (_Msw *MswSession) ExecuteTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _Msw.Contract.ExecuteTransaction(&_Msw.TransactOpts, transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 transactionId) returns()
func (_Msw *MswTransactorSession) ExecuteTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _Msw.Contract.ExecuteTransaction(&_Msw.TransactOpts, transactionId)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address owner) returns()
func (_Msw *MswTransactor) RemoveOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _Msw.contract.Transact(opts, "removeOwner", owner)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address owner) returns()
func (_Msw *MswSession) RemoveOwner(owner common.Address) (*types.Transaction, error) {
	return _Msw.Contract.RemoveOwner(&_Msw.TransactOpts, owner)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address owner) returns()
func (_Msw *MswTransactorSession) RemoveOwner(owner common.Address) (*types.Transaction, error) {
	return _Msw.Contract.RemoveOwner(&_Msw.TransactOpts, owner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(address owner, address newOwner) returns()
func (_Msw *MswTransactor) ReplaceOwner(opts *bind.TransactOpts, owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _Msw.contract.Transact(opts, "replaceOwner", owner, newOwner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(address owner, address newOwner) returns()
func (_Msw *MswSession) ReplaceOwner(owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _Msw.Contract.ReplaceOwner(&_Msw.TransactOpts, owner, newOwner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(address owner, address newOwner) returns()
func (_Msw *MswTransactorSession) ReplaceOwner(owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _Msw.Contract.ReplaceOwner(&_Msw.TransactOpts, owner, newOwner)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 transactionId) returns()
func (_Msw *MswTransactor) RevokeConfirmation(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _Msw.contract.Transact(opts, "revokeConfirmation", transactionId)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 transactionId) returns()
func (_Msw *MswSession) RevokeConfirmation(transactionId *big.Int) (*types.Transaction, error) {
	return _Msw.Contract.RevokeConfirmation(&_Msw.TransactOpts, transactionId)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 transactionId) returns()
func (_Msw *MswTransactorSession) RevokeConfirmation(transactionId *big.Int) (*types.Transaction, error) {
	return _Msw.Contract.RevokeConfirmation(&_Msw.TransactOpts, transactionId)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address destination, uint256 value, bytes data) returns(uint256 transactionId)
func (_Msw *MswTransactor) SubmitTransaction(opts *bind.TransactOpts, destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Msw.contract.Transact(opts, "submitTransaction", destination, value, data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address destination, uint256 value, bytes data) returns(uint256 transactionId)
func (_Msw *MswSession) SubmitTransaction(destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Msw.Contract.SubmitTransaction(&_Msw.TransactOpts, destination, value, data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address destination, uint256 value, bytes data) returns(uint256 transactionId)
func (_Msw *MswTransactorSession) SubmitTransaction(destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Msw.Contract.SubmitTransaction(&_Msw.TransactOpts, destination, value, data)
}

// MswConfirmationIterator is returned from FilterConfirmation and is used to iterate over the raw logs and unpacked data for Confirmation events raised by the Msw contract.
type MswConfirmationIterator struct {
	Event *MswConfirmation // Event containing the contract specifics and raw log

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
func (it *MswConfirmationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MswConfirmation)
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
		it.Event = new(MswConfirmation)
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
func (it *MswConfirmationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MswConfirmationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MswConfirmation represents a Confirmation event raised by the Msw contract.
type MswConfirmation struct {
	Sender        common.Address
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterConfirmation is a free log retrieval operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(address indexed sender, uint256 indexed transactionId)
func (_Msw *MswFilterer) FilterConfirmation(opts *bind.FilterOpts, sender []common.Address, transactionId []*big.Int) (*MswConfirmationIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _Msw.contract.FilterLogs(opts, "Confirmation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MswConfirmationIterator{contract: _Msw.contract, event: "Confirmation", logs: logs, sub: sub}, nil
}

// WatchConfirmation is a free log subscription operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(address indexed sender, uint256 indexed transactionId)
func (_Msw *MswFilterer) WatchConfirmation(opts *bind.WatchOpts, sink chan<- *MswConfirmation, sender []common.Address, transactionId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _Msw.contract.WatchLogs(opts, "Confirmation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MswConfirmation)
				if err := _Msw.contract.UnpackLog(event, "Confirmation", log); err != nil {
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

// MswDailyLimitChangeIterator is returned from FilterDailyLimitChange and is used to iterate over the raw logs and unpacked data for DailyLimitChange events raised by the Msw contract.
type MswDailyLimitChangeIterator struct {
	Event *MswDailyLimitChange // Event containing the contract specifics and raw log

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
func (it *MswDailyLimitChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MswDailyLimitChange)
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
		it.Event = new(MswDailyLimitChange)
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
func (it *MswDailyLimitChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MswDailyLimitChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MswDailyLimitChange represents a DailyLimitChange event raised by the Msw contract.
type MswDailyLimitChange struct {
	DailyLimit *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDailyLimitChange is a free log retrieval operation binding the contract event 0xc71bdc6afaf9b1aa90a7078191d4fc1adf3bf680fca3183697df6b0dc226bca2.
//
// Solidity: event DailyLimitChange(uint256 dailyLimit)
func (_Msw *MswFilterer) FilterDailyLimitChange(opts *bind.FilterOpts) (*MswDailyLimitChangeIterator, error) {

	logs, sub, err := _Msw.contract.FilterLogs(opts, "DailyLimitChange")
	if err != nil {
		return nil, err
	}
	return &MswDailyLimitChangeIterator{contract: _Msw.contract, event: "DailyLimitChange", logs: logs, sub: sub}, nil
}

// WatchDailyLimitChange is a free log subscription operation binding the contract event 0xc71bdc6afaf9b1aa90a7078191d4fc1adf3bf680fca3183697df6b0dc226bca2.
//
// Solidity: event DailyLimitChange(uint256 dailyLimit)
func (_Msw *MswFilterer) WatchDailyLimitChange(opts *bind.WatchOpts, sink chan<- *MswDailyLimitChange) (event.Subscription, error) {

	logs, sub, err := _Msw.contract.WatchLogs(opts, "DailyLimitChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MswDailyLimitChange)
				if err := _Msw.contract.UnpackLog(event, "DailyLimitChange", log); err != nil {
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

// MswDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Msw contract.
type MswDepositIterator struct {
	Event *MswDeposit // Event containing the contract specifics and raw log

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
func (it *MswDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MswDeposit)
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
		it.Event = new(MswDeposit)
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
func (it *MswDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MswDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MswDeposit represents a Deposit event raised by the Msw contract.
type MswDeposit struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 value)
func (_Msw *MswFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address) (*MswDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Msw.contract.FilterLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &MswDepositIterator{contract: _Msw.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 value)
func (_Msw *MswFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *MswDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Msw.contract.WatchLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MswDeposit)
				if err := _Msw.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// MswExecutionIterator is returned from FilterExecution and is used to iterate over the raw logs and unpacked data for Execution events raised by the Msw contract.
type MswExecutionIterator struct {
	Event *MswExecution // Event containing the contract specifics and raw log

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
func (it *MswExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MswExecution)
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
		it.Event = new(MswExecution)
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
func (it *MswExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MswExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MswExecution represents a Execution event raised by the Msw contract.
type MswExecution struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecution is a free log retrieval operation binding the contract event 0x33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed75.
//
// Solidity: event Execution(uint256 indexed transactionId)
func (_Msw *MswFilterer) FilterExecution(opts *bind.FilterOpts, transactionId []*big.Int) (*MswExecutionIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _Msw.contract.FilterLogs(opts, "Execution", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MswExecutionIterator{contract: _Msw.contract, event: "Execution", logs: logs, sub: sub}, nil
}

// WatchExecution is a free log subscription operation binding the contract event 0x33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed75.
//
// Solidity: event Execution(uint256 indexed transactionId)
func (_Msw *MswFilterer) WatchExecution(opts *bind.WatchOpts, sink chan<- *MswExecution, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _Msw.contract.WatchLogs(opts, "Execution", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MswExecution)
				if err := _Msw.contract.UnpackLog(event, "Execution", log); err != nil {
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

// MswExecutionFailureIterator is returned from FilterExecutionFailure and is used to iterate over the raw logs and unpacked data for ExecutionFailure events raised by the Msw contract.
type MswExecutionFailureIterator struct {
	Event *MswExecutionFailure // Event containing the contract specifics and raw log

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
func (it *MswExecutionFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MswExecutionFailure)
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
		it.Event = new(MswExecutionFailure)
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
func (it *MswExecutionFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MswExecutionFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MswExecutionFailure represents a ExecutionFailure event raised by the Msw contract.
type MswExecutionFailure struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecutionFailure is a free log retrieval operation binding the contract event 0x526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b79236.
//
// Solidity: event ExecutionFailure(uint256 indexed transactionId)
func (_Msw *MswFilterer) FilterExecutionFailure(opts *bind.FilterOpts, transactionId []*big.Int) (*MswExecutionFailureIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _Msw.contract.FilterLogs(opts, "ExecutionFailure", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MswExecutionFailureIterator{contract: _Msw.contract, event: "ExecutionFailure", logs: logs, sub: sub}, nil
}

// WatchExecutionFailure is a free log subscription operation binding the contract event 0x526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b79236.
//
// Solidity: event ExecutionFailure(uint256 indexed transactionId)
func (_Msw *MswFilterer) WatchExecutionFailure(opts *bind.WatchOpts, sink chan<- *MswExecutionFailure, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _Msw.contract.WatchLogs(opts, "ExecutionFailure", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MswExecutionFailure)
				if err := _Msw.contract.UnpackLog(event, "ExecutionFailure", log); err != nil {
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

// MswOwnerAdditionIterator is returned from FilterOwnerAddition and is used to iterate over the raw logs and unpacked data for OwnerAddition events raised by the Msw contract.
type MswOwnerAdditionIterator struct {
	Event *MswOwnerAddition // Event containing the contract specifics and raw log

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
func (it *MswOwnerAdditionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MswOwnerAddition)
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
		it.Event = new(MswOwnerAddition)
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
func (it *MswOwnerAdditionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MswOwnerAdditionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MswOwnerAddition represents a OwnerAddition event raised by the Msw contract.
type MswOwnerAddition struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOwnerAddition is a free log retrieval operation binding the contract event 0xf39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d.
//
// Solidity: event OwnerAddition(address indexed owner)
func (_Msw *MswFilterer) FilterOwnerAddition(opts *bind.FilterOpts, owner []common.Address) (*MswOwnerAdditionIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Msw.contract.FilterLogs(opts, "OwnerAddition", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MswOwnerAdditionIterator{contract: _Msw.contract, event: "OwnerAddition", logs: logs, sub: sub}, nil
}

// WatchOwnerAddition is a free log subscription operation binding the contract event 0xf39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d.
//
// Solidity: event OwnerAddition(address indexed owner)
func (_Msw *MswFilterer) WatchOwnerAddition(opts *bind.WatchOpts, sink chan<- *MswOwnerAddition, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Msw.contract.WatchLogs(opts, "OwnerAddition", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MswOwnerAddition)
				if err := _Msw.contract.UnpackLog(event, "OwnerAddition", log); err != nil {
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

// MswOwnerRemovalIterator is returned from FilterOwnerRemoval and is used to iterate over the raw logs and unpacked data for OwnerRemoval events raised by the Msw contract.
type MswOwnerRemovalIterator struct {
	Event *MswOwnerRemoval // Event containing the contract specifics and raw log

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
func (it *MswOwnerRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MswOwnerRemoval)
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
		it.Event = new(MswOwnerRemoval)
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
func (it *MswOwnerRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MswOwnerRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MswOwnerRemoval represents a OwnerRemoval event raised by the Msw contract.
type MswOwnerRemoval struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOwnerRemoval is a free log retrieval operation binding the contract event 0x8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90.
//
// Solidity: event OwnerRemoval(address indexed owner)
func (_Msw *MswFilterer) FilterOwnerRemoval(opts *bind.FilterOpts, owner []common.Address) (*MswOwnerRemovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Msw.contract.FilterLogs(opts, "OwnerRemoval", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MswOwnerRemovalIterator{contract: _Msw.contract, event: "OwnerRemoval", logs: logs, sub: sub}, nil
}

// WatchOwnerRemoval is a free log subscription operation binding the contract event 0x8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90.
//
// Solidity: event OwnerRemoval(address indexed owner)
func (_Msw *MswFilterer) WatchOwnerRemoval(opts *bind.WatchOpts, sink chan<- *MswOwnerRemoval, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Msw.contract.WatchLogs(opts, "OwnerRemoval", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MswOwnerRemoval)
				if err := _Msw.contract.UnpackLog(event, "OwnerRemoval", log); err != nil {
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

// MswRequirementChangeIterator is returned from FilterRequirementChange and is used to iterate over the raw logs and unpacked data for RequirementChange events raised by the Msw contract.
type MswRequirementChangeIterator struct {
	Event *MswRequirementChange // Event containing the contract specifics and raw log

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
func (it *MswRequirementChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MswRequirementChange)
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
		it.Event = new(MswRequirementChange)
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
func (it *MswRequirementChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MswRequirementChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MswRequirementChange represents a RequirementChange event raised by the Msw contract.
type MswRequirementChange struct {
	Required *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRequirementChange is a free log retrieval operation binding the contract event 0xa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a.
//
// Solidity: event RequirementChange(uint256 required)
func (_Msw *MswFilterer) FilterRequirementChange(opts *bind.FilterOpts) (*MswRequirementChangeIterator, error) {

	logs, sub, err := _Msw.contract.FilterLogs(opts, "RequirementChange")
	if err != nil {
		return nil, err
	}
	return &MswRequirementChangeIterator{contract: _Msw.contract, event: "RequirementChange", logs: logs, sub: sub}, nil
}

// WatchRequirementChange is a free log subscription operation binding the contract event 0xa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a.
//
// Solidity: event RequirementChange(uint256 required)
func (_Msw *MswFilterer) WatchRequirementChange(opts *bind.WatchOpts, sink chan<- *MswRequirementChange) (event.Subscription, error) {

	logs, sub, err := _Msw.contract.WatchLogs(opts, "RequirementChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MswRequirementChange)
				if err := _Msw.contract.UnpackLog(event, "RequirementChange", log); err != nil {
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

// MswRevocationIterator is returned from FilterRevocation and is used to iterate over the raw logs and unpacked data for Revocation events raised by the Msw contract.
type MswRevocationIterator struct {
	Event *MswRevocation // Event containing the contract specifics and raw log

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
func (it *MswRevocationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MswRevocation)
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
		it.Event = new(MswRevocation)
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
func (it *MswRevocationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MswRevocationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MswRevocation represents a Revocation event raised by the Msw contract.
type MswRevocation struct {
	Sender        common.Address
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRevocation is a free log retrieval operation binding the contract event 0xf6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9.
//
// Solidity: event Revocation(address indexed sender, uint256 indexed transactionId)
func (_Msw *MswFilterer) FilterRevocation(opts *bind.FilterOpts, sender []common.Address, transactionId []*big.Int) (*MswRevocationIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _Msw.contract.FilterLogs(opts, "Revocation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MswRevocationIterator{contract: _Msw.contract, event: "Revocation", logs: logs, sub: sub}, nil
}

// WatchRevocation is a free log subscription operation binding the contract event 0xf6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9.
//
// Solidity: event Revocation(address indexed sender, uint256 indexed transactionId)
func (_Msw *MswFilterer) WatchRevocation(opts *bind.WatchOpts, sink chan<- *MswRevocation, sender []common.Address, transactionId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _Msw.contract.WatchLogs(opts, "Revocation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MswRevocation)
				if err := _Msw.contract.UnpackLog(event, "Revocation", log); err != nil {
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

// MswSubmissionIterator is returned from FilterSubmission and is used to iterate over the raw logs and unpacked data for Submission events raised by the Msw contract.
type MswSubmissionIterator struct {
	Event *MswSubmission // Event containing the contract specifics and raw log

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
func (it *MswSubmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MswSubmission)
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
		it.Event = new(MswSubmission)
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
func (it *MswSubmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MswSubmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MswSubmission represents a Submission event raised by the Msw contract.
type MswSubmission struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSubmission is a free log retrieval operation binding the contract event 0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51.
//
// Solidity: event Submission(uint256 indexed transactionId)
func (_Msw *MswFilterer) FilterSubmission(opts *bind.FilterOpts, transactionId []*big.Int) (*MswSubmissionIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _Msw.contract.FilterLogs(opts, "Submission", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MswSubmissionIterator{contract: _Msw.contract, event: "Submission", logs: logs, sub: sub}, nil
}

// WatchSubmission is a free log subscription operation binding the contract event 0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51.
//
// Solidity: event Submission(uint256 indexed transactionId)
func (_Msw *MswFilterer) WatchSubmission(opts *bind.WatchOpts, sink chan<- *MswSubmission, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _Msw.contract.WatchLogs(opts, "Submission", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MswSubmission)
				if err := _Msw.contract.UnpackLog(event, "Submission", log); err != nil {
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
