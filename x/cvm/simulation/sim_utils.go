package simulation

import (
	"encoding/hex"
	"fmt"
	"math/rand"

	"github.com/gogo/protobuf/proto"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sim "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/shentufoundation/shentu/v2/x/cvm/keeper"
	"github.com/shentufoundation/shentu/v2/x/cvm/types"
)

const (
	Hello55Code  = "6080604052348015600f57600080fd5b5060888061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80630c49c36c14602d575b600080fd5b60336049565b6040518082815260200191505060405180910390f35b6000603790509056fea2646970667358221220be801fb2205f223dcf6751ff8f3d1996fa2aa8bd72fec7015b75e7c4826e09a264736f6c634300060a0033"
	Hello55Abi   = "[{\"inputs\":[],\"name\":\"sayHi\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"
	Hello55SayHi = "0c49c36c"

	SimpleCode      = "608060405234801561001057600080fd5b5060c78061001f6000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c806360fe47b11460375780636d4ce63c146062575b600080fd5b606060048036036020811015604b57600080fd5b8101908080359060200190929190505050607e565b005b60686088565b6040518082815260200191505060405180910390f35b8060008190555050565b6000805490509056fea264697066735822122009c206964bec9aab615f7e1679af3050ffceeb7925dc31b5aae347317d0a74c164736f6c634300060a0033"
	SimpleAbi       = "[{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	SimpleSetPrefix = "60fe47b1"
	SimpleSetSample = "60fe47b10000000000000000000000000000000000000000000000000000000000000429" // hex for 1065
	SimpleGet       = "6d4ce63c"

	SimpleeventCode      = "608060405234801561001057600080fd5b50610275806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806360fe47b11461003b5780636d4ce63c14610069575b600080fd5b6100676004803603602081101561005157600080fd5b8101908080359060200190929190505050610087565b005b610071610214565b6040518082815260200191505060405180910390f35b806000819055507f6c2b4666ba8da5a95717621d879a77de725f3d816709b9cbe9f059b8f875e2846001600054016040518082815260200191505060405180910390a17f6c96f7d2523b7c8c2eaeba423e4234f7443e80284664c7aa5e285aac4d2613bd60005460020260405180828152602001806020018281038252600e8152602001807f7365636f6e644576656e742121210000000000000000000000000000000000008152506020019250505060405180910390a161014761021d565b6040518060600160405280600160ff168152602001600260ff168152602001600360ff1681525090507f2e71aa6814b90353bea2dc6b23acde4924508f6646936e13a137601134a73d9560036000540182306040518084815260200183600360200280838360005b838110156101ca5780820151818401526020810190506101af565b505050509050018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001935050505060405180910390a15050565b60008054905090565b604051806060016040528060039060208202803683378082019150509050509056fea2646970667358221220b22a748b3370f1eee88fe8672cc53af0fed90a6552bd0e861665ae873bd967fe64736f6c634300060a0033"
	SimpleeventAbi       = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"myNumber\",\"type\":\"uint256\"}],\"name\":\"MyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mySecondNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"s\",\"type\":\"string\"}],\"name\":\"MySecondEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"myThirdNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8[3]\",\"name\":\"arr\",\"type\":\"uint8[3]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"MyThirdEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	SimpleeventSetPrefix = "60fe47b1"
	SimpleeventSetSample = "60fe47b100000000000000000000000000000000000000000000000000000000000003e8" // hex for 1000
	SimpleeventGet       = "6d4ce63c"

	StorageCode        = "608060405234801561001057600080fd5b50610121806100206000396000f3fe6080604052348015600f57600080fd5b5060043610603c5760003560e01c80632e64cec11460415780636057361d14605d5780638f3eff7b146088575b600080fd5b604760d0565b6040518082815260200191505060405180910390f35b608660048036036020811015607157600080fd5b810190808035906020019092919050505060d9565b005b608e60e3565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b60008054905090565b8060008190555050565b60003390509056fea2646970667358221220870d6416d9143f210b232ec3a5a9df94095b8ee3803922bab202cf85f2e6a46264736f6c634300060a0033"
	StorageAbi         = "[{\"inputs\":[],\"name\":\"retrieve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sayMyAddres\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	StorageStorePrefix = "6057361d"
	StorageStoreSample = "6057361d0000000000000000000000000000000000000000000000000000000000000225" // hex for 549
	StorageRetrieve    = "2e64cec1"
	StorageSayMyAddres = "8f3eff7b"

	StringtestCode               = "608060405234801561001057600080fd5b50610c1e806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c806316a872f11461005c5780633d7405f41461007a5780635dd8e1c914610098578063dac97065146101cc578063ebea52e614610300575b600080fd5b610064610383565b6040518082815260200191505060405180910390f35b6100826108a3565b6040518082815260200191505060405180910390f35b610151600480360360208110156100ae57600080fd5b81019080803590602001906401000000008111156100cb57600080fd5b8201836020820111156100dd57600080fd5b803590602001918460018302840111640100000000831117156100ff57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506108c0565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610191578082015181840152602081019050610176565b50505050905090810190601f1680156101be5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610285600480360360208110156101e257600080fd5b81019080803590602001906401000000008111156101ff57600080fd5b82018360208201111561021157600080fd5b8035906020019184600183028401116401000000008311171561023357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929050505061097b565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156102c55780820151818401526020810190506102aa565b50505050905090810190601f1680156102f25780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610308610aa1565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561034857808201518184015260208101905061032d565b50505050905090810190601f1680156103755780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b600060606040518060400160405280600981526020017f6173646664643030300000000000000000000000000000000000000000000000815250905060606040518060400160405280600981526020017f6173646664643030300000000000000000000000000000000000000000000000815250905060606040518060400160405280600981526020017f61736466646430303100000000000000000000000000000000000000000000008152509050816040516020018082805190602001908083835b6020831061046a5780518252602082019150602081019050602083039250610447565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405160208183030381529060405280519060200120836040516020018082805190602001908083835b602083106104dc57805182526020820191506020810190506020830392506104b9565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051602081830303815290604052805190602001201461052157600080fd5b806040516020018082805190602001908083835b602083106105585780518252602082019150602081019050602083039250610535565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405160208183030381529060405280519060200120826040516020018082805190602001908083835b602083106105ca57805182526020820191506020810190506020830392506105a7565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405160208183030381529060405280519060200120141561061057600080fd5b60606040518060400160405280600a81526020017fe282ac62c3816ec3a72100000000000000000000000000000000000000000000815250905060606040518060400160405280600a81526020017f41626362c3816ec3a721000000000000000000000000000000000000000000008152509050606081905060608390507f4100000000000000000000000000000000000000000000000000000000000000816000815181106106bc57fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f62000000000000000000000000000000000000000000000000000000000000008160018151811061071957fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f63000000000000000000000000000000000000000000000000000000000000008160028151811061077657fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350806040516020018082805190602001908083835b602083106107dc57805182526020820191506020810190506020830392506107b9565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405160208183030381529060405280519060200120826040516020018082805190602001908083835b6020831061084e578051825260208201915060208101905060208303925061082b565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051602081830303815290604052805190602001201461089357600080fd5b6201e0f397505050505050505090565b600080805460018160011615610100020316600290049050905090565b606081600090805190602001906108d8929190610b43565b5060008054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561096f5780601f106109445761010080835404028352916020019161096f565b820191906000526020600020905b81548152906001019060200180831161095257829003601f168201915b50505050509050919050565b6060808290507f4100000000000000000000000000000000000000000000000000000000000000816000815181106109af57fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f620000000000000000000000000000000000000000000000000000000000000081600181518110610a0c57fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f630000000000000000000000000000000000000000000000000000000000000081600281518110610a6957fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080915050919050565b606060008054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610b395780601f10610b0e57610100808354040283529160200191610b39565b820191906000526020600020905b815481529060010190602001808311610b1c57829003601f168201915b5050505050905090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610b8457805160ff1916838001178555610bb2565b82800160010185558215610bb2579182015b82811115610bb1578251825591602001919060010190610b96565b5b509050610bbf9190610bc3565b5090565b610be591905b80821115610be1576000816000905550600101610bc9565b5090565b9056fea264697066735822122068af65928b83299f2cab64aa384e0516020be54f4dffc36dc6d965a1e47de35d64736f6c634300060a0033"
	StringtestAbi                = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_s\",\"type\":\"string\"}],\"name\":\"changeGiven\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_c\",\"type\":\"string\"}],\"name\":\"changeString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getl\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gets\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"testStuff\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]\nMeta: [CodeHash:\"\\255\\311\\037E\\306[\\361G\\256s\\357z\\264\\333\\257\\276\\001t\\305\\2341\\211\\353\\377Y\\006\\220\\337\\024\\347\\235\\332\" Meta:\"{\\\"ContractName\\\":\\\"StringTest\\\",\\\"SourceFile\\\":\\\"stringtest.sol\\\",\\\"CompilerVersion\\\":\\\"0.6.10+commit.00c0fcaf\\\",\\\"Abi\\\":[{\\\"inputs\\\":[{\\\"internalType\\\":\\\"string\\\",\\\"name\\\":\\\"_s\\\",\\\"type\\\":\\\"string\\\"}],\\\"name\\\":\\\"changeGiven\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"string\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"string\\\"}],\\\"stateMutability\\\":\\\"nonpayable\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[{\\\"internalType\\\":\\\"string\\\",\\\"name\\\":\\\"_c\\\",\\\"type\\\":\\\"string\\\"}],\\\"name\\\":\\\"changeString\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"string\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"string\\\"}],\\\"stateMutability\\\":\\\"nonpayable\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[],\\\"name\\\":\\\"getl\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"uint256\\\"}],\\\"stateMutability\\\":\\\"nonpayable\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[],\\\"name\\\":\\\"gets\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"string\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"string\\\"}],\\\"stateMutability\\\":\\\"nonpayable\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[],\\\"name\\\":\\\"testStuff\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"uint256\\\"}],\\\"stateMutability\\\":\\\"nonpayable\\\",\\\"type\\\":\\\"function\\\"}]}\" ]"
	StringtestChangeStringPrefix = "5dd8e1c90000000000000000000000000000000000000000000000000000000000000020"
	StringtestChangeStringSample = "5dd8e1c9 0000000000000000000000000000000000000000000000000000000000000020 0000000000000000000000000000000000000000000000000000000000000003 6162630000000000000000000000000000000000000000000000000000000000" // abc
	StringtestChangeGivenPrefix  = "dac970650000000000000000000000000000000000000000000000000000000000000020"
	StringtestChangeGivenSample  = "dac97065 0000000000000000000000000000000000000000000000000000000000000020 0000000000000000000000000000000000000000000000000000000000000006 6465666768690000000000000000000000000000000000000000000000000000" // defghi
	StringtestGets               = "ebea52e6"
	StringtestGetl               = "3d7405f4"
	StringtestTestStuff          = "16a872f1"
)

// DeployContract delivers a deploy tx and returns msg, contract address and error.
func DeployContract(caller sim.Account, contractCode string, contractAbi string, k keeper.Keeper, bk types.BankKeeper, r *rand.Rand,
	ctx sdk.Context, chainID string, app *baseapp.BaseApp) (msg types.MsgDeploy, contractAddr sdk.AccAddress, err error) {
	code, err := hex.DecodeString(contractCode)
	if err != nil {
		return msg, nil, err
	}

	msg = types.NewMsgDeploy(caller.Address.String(), uint64(0), code, contractAbi, nil, false, false)

	account := k.GetAccount(ctx, caller.Address)
	fees, err := RandomReasonableFees(r, ctx, bk.SpendableCoins(ctx, caller.Address))
	if err != nil {
		return msg, nil, err
	}

	txGen := simappparams.MakeTestEncodingConfig().TxConfig
	tx, err := helpers.GenTx(
		txGen,
		[]sdk.Msg{&msg},
		fees,
		helpers.DefaultGenTxGas,
		chainID,
		[]uint64{account.GetAccountNumber()},
		[]uint64{account.GetSequence()},
		caller.PrivKey,
	)
	if err != nil {
		return msg, nil, err
	}

	_, res, err := app.Deliver(txGen.TxEncoder(), tx)
	if err != nil {
		return msg, nil, err
	}

	// Unmarshal response
	var msgData sdk.TxMsgData
	err = proto.Unmarshal(res.Data, &msgData)
	if err != nil {
		panic(err)
	}

	var deployResponse types.MsgDeployResponse
	err = proto.Unmarshal(msgData.Data[0].Data, &deployResponse)
	if err != nil {
		panic(err)
	}

	return msg, deployResponse.Result, nil
}

func RandomReasonableFees(r *rand.Rand, ctx sdk.Context, spendableCoins sdk.Coins) (sdk.Coins, error) {
	if spendableCoins.Empty() {
		return nil, nil
	}

	perm := r.Perm(len(spendableCoins))
	var randCoin sdk.Coin
	for _, index := range perm {
		randCoin = spendableCoins[index]
		if !randCoin.Amount.IsZero() {
			break
		}
	}

	if randCoin.Amount.IsZero() {
		return nil, fmt.Errorf("no coins found for random fees")
	}
	//To limit the fee not exceeding ninth of remaining balances.
	//Without this limitation, it's easy to run out of accounts balance
	//  given so many simulate operations introduced by shentu modules.
	amt, err := sim.RandPositiveInt(r, randCoin.Amount.Add(sdk.NewInt(8)).Quo(sdk.NewInt(9)))
	if err != nil {
		return nil, err
	}

	// Create a random fee and verify the fees are within the account's spendable
	// balance.
	fees := sdk.NewCoins(sdk.NewCoin(randCoin.Denom, amt))

	return fees, nil
}

// CallFunction delivers a call tx and returns msg, contract address and error.
func CallFunction(caller sim.Account, prefix string, input string, contractAddr sdk.AccAddress, k keeper.Keeper, bk types.BankKeeper,
	ctx sdk.Context, r *rand.Rand, chainID string, app *baseapp.BaseApp) (msg types.MsgCall, ret []byte, err error) {
	data, err := hex.DecodeString(prefix + input)
	if err != nil {
		return msg, nil, err
	}

	msg = types.NewMsgCall(caller.Address.String(), contractAddr.String(), 0, data)

	account := k.GetAccount(ctx, caller.Address)
	fees, err := RandomReasonableFees(r, ctx, bk.SpendableCoins(ctx, caller.Address))
	if err != nil {
		return msg, nil, err
	}

	txGen := simappparams.MakeTestEncodingConfig().TxConfig
	tx, err := helpers.GenTx(
		txGen,
		[]sdk.Msg{&msg},
		fees,
		helpers.DefaultGenTxGas,
		chainID,
		[]uint64{account.GetAccountNumber()},
		[]uint64{account.GetSequence()},
		caller.PrivKey,
	)
	if err != nil {
		return msg, nil, err
	}

	_, res, err := app.Deliver(txGen.TxEncoder(), tx)
	if err != nil {
		return msg, nil, err
	}

	// Unmarshal response
	var msgData sdk.TxMsgData
	err = proto.Unmarshal(res.Data, &msgData)
	if err != nil {
		panic(err)
	}

	var callResponse types.MsgCallResponse
	err = proto.Unmarshal(msgData.Data[0].Data, &callResponse)
	if err != nil {
		panic(err)
	}

	return msg, callResponse.Result, nil
}
