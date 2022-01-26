// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package curve

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
)

// CurveMetaData contains all meta data concerning the Curve contract.
var CurveMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"PoolAdded\",\"inputs\":[{\"type\":\"address\",\"name\":\"pool\",\"indexed\":true},{\"type\":\"bytes\",\"name\":\"rate_method_id\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"PoolRemoved\",\"inputs\":[{\"type\":\"address\",\"name\":\"pool\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_address_provider\"},{\"type\":\"address\",\"name\":\"_gauge_controller\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"name\":\"find_pool_for_coins\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_from\"},{\"type\":\"address\",\"name\":\"_to\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"find_pool_for_coins\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_from\"},{\"type\":\"address\",\"name\":\"_to\"},{\"type\":\"uint256\",\"name\":\"i\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"get_n_coins\",\"outputs\":[{\"type\":\"uint256[2]\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1704},{\"name\":\"get_coins\",\"outputs\":[{\"type\":\"address[8]\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":12285},{\"name\":\"get_underlying_coins\",\"outputs\":[{\"type\":\"address[8]\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":12347},{\"name\":\"get_decimals\",\"outputs\":[{\"type\":\"uint256[8]\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":8199},{\"name\":\"get_underlying_decimals\",\"outputs\":[{\"type\":\"uint256[8]\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":8261},{\"name\":\"get_rates\",\"outputs\":[{\"type\":\"uint256[8]\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":34780},{\"name\":\"get_gauges\",\"outputs\":[{\"type\":\"address[10]\",\"name\":\"\"},{\"type\":\"int128[10]\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":20310},{\"name\":\"get_balances\",\"outputs\":[{\"type\":\"uint256[8]\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":16818},{\"name\":\"get_underlying_balances\",\"outputs\":[{\"type\":\"uint256[8]\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":158953},{\"name\":\"get_virtual_price_from_lp_token\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_token\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2080},{\"name\":\"get_A\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1198},{\"name\":\"get_parameters\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"A\"},{\"type\":\"uint256\",\"name\":\"future_A\"},{\"type\":\"uint256\",\"name\":\"fee\"},{\"type\":\"uint256\",\"name\":\"admin_fee\"},{\"type\":\"uint256\",\"name\":\"future_fee\"},{\"type\":\"uint256\",\"name\":\"future_admin_fee\"},{\"type\":\"address\",\"name\":\"future_owner\"},{\"type\":\"uint256\",\"name\":\"initial_A\"},{\"type\":\"uint256\",\"name\":\"initial_A_time\"},{\"type\":\"uint256\",\"name\":\"future_A_time\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":6458},{\"name\":\"get_fees\",\"outputs\":[{\"type\":\"uint256[2]\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1603},{\"name\":\"get_admin_balances\",\"outputs\":[{\"type\":\"uint256[8]\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":36719},{\"name\":\"get_coin_indices\",\"outputs\":[{\"type\":\"int128\",\"name\":\"\"},{\"type\":\"int128\",\"name\":\"\"},{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"},{\"type\":\"address\",\"name\":\"_from\"},{\"type\":\"address\",\"name\":\"_to\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":27456},{\"name\":\"estimate_gas_used\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"},{\"type\":\"address\",\"name\":\"_from\"},{\"type\":\"address\",\"name\":\"_to\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":32329},{\"name\":\"add_pool\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"},{\"type\":\"uint256\",\"name\":\"_n_coins\"},{\"type\":\"address\",\"name\":\"_lp_token\"},{\"type\":\"bytes32\",\"name\":\"_rate_method_id\"},{\"type\":\"uint256\",\"name\":\"_decimals\"},{\"type\":\"uint256\",\"name\":\"_underlying_decimals\"},{\"type\":\"bool\",\"name\":\"_has_initial_A\"},{\"type\":\"bool\",\"name\":\"_is_v1\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":10196577},{\"name\":\"add_pool_without_underlying\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"},{\"type\":\"uint256\",\"name\":\"_n_coins\"},{\"type\":\"address\",\"name\":\"_lp_token\"},{\"type\":\"bytes32\",\"name\":\"_rate_method_id\"},{\"type\":\"uint256\",\"name\":\"_decimals\"},{\"type\":\"uint256\",\"name\":\"_use_rates\"},{\"type\":\"bool\",\"name\":\"_has_initial_A\"},{\"type\":\"bool\",\"name\":\"_is_v1\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":5590664},{\"name\":\"add_metapool\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"},{\"type\":\"uint256\",\"name\":\"_n_coins\"},{\"type\":\"address\",\"name\":\"_lp_token\"},{\"type\":\"uint256\",\"name\":\"_decimals\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":10226976},{\"name\":\"remove_pool\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":779646579509},{\"name\":\"set_pool_gas_estimates\",\"outputs\":[],\"inputs\":[{\"type\":\"address[5]\",\"name\":\"_addr\"},{\"type\":\"uint256[2][5]\",\"name\":\"_amount\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":355578},{\"name\":\"set_coin_gas_estimates\",\"outputs\":[],\"inputs\":[{\"type\":\"address[10]\",\"name\":\"_addr\"},{\"type\":\"uint256[10]\",\"name\":\"_amount\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":357165},{\"name\":\"set_gas_estimate_contract\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"},{\"type\":\"address\",\"name\":\"_estimator\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37747},{\"name\":\"set_liquidity_gauges\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_pool\"},{\"type\":\"address[10]\",\"name\":\"_liquidity_gauges\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":365793},{\"name\":\"address_provider\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2111},{\"name\":\"gauge_controller\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2141},{\"name\":\"pool_list\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2280},{\"name\":\"pool_count\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2201},{\"name\":\"get_pool_from_lp_token\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2446},{\"name\":\"get_lp_token\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2476}]",
}

// CurveABI is the input ABI used to generate the binding from.
// Deprecated: Use CurveMetaData.ABI instead.
var CurveABI = CurveMetaData.ABI

// Curve is an auto generated Go binding around an Ethereum contract.
type Curve struct {
	CurveCaller     // Read-only binding to the contract
	CurveTransactor // Write-only binding to the contract
	CurveFilterer   // Log filterer for contract events
}

// CurveCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurveCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurveTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurveFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurveSession struct {
	Contract     *Curve            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurveCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurveCallerSession struct {
	Contract *CurveCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CurveTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurveTransactorSession struct {
	Contract     *CurveTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurveRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurveRaw struct {
	Contract *Curve // Generic contract binding to access the raw methods on
}

// CurveCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurveCallerRaw struct {
	Contract *CurveCaller // Generic read-only contract binding to access the raw methods on
}

// CurveTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurveTransactorRaw struct {
	Contract *CurveTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurve creates a new instance of Curve, bound to a specific deployed contract.
func NewCurve(address common.Address, backend bind.ContractBackend) (*Curve, error) {
	contract, err := bindCurve(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Curve{CurveCaller: CurveCaller{contract: contract}, CurveTransactor: CurveTransactor{contract: contract}, CurveFilterer: CurveFilterer{contract: contract}}, nil
}

// NewCurveCaller creates a new read-only instance of Curve, bound to a specific deployed contract.
func NewCurveCaller(address common.Address, caller bind.ContractCaller) (*CurveCaller, error) {
	contract, err := bindCurve(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurveCaller{contract: contract}, nil
}

// NewCurveTransactor creates a new write-only instance of Curve, bound to a specific deployed contract.
func NewCurveTransactor(address common.Address, transactor bind.ContractTransactor) (*CurveTransactor, error) {
	contract, err := bindCurve(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurveTransactor{contract: contract}, nil
}

// NewCurveFilterer creates a new log filterer instance of Curve, bound to a specific deployed contract.
func NewCurveFilterer(address common.Address, filterer bind.ContractFilterer) (*CurveFilterer, error) {
	contract, err := bindCurve(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurveFilterer{contract: contract}, nil
}

// bindCurve binds a generic wrapper to an already deployed contract.
func bindCurve(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CurveABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Curve *CurveRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Curve.Contract.CurveCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Curve *CurveRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Curve.Contract.CurveTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Curve *CurveRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Curve.Contract.CurveTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Curve *CurveCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Curve.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Curve *CurveTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Curve.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Curve *CurveTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Curve.Contract.contract.Transact(opts, method, params...)
}

// AddressProvider is a free data retrieval call binding the contract method 0xce50c2e7.
//
// Solidity: function address_provider() view returns(address)
func (_Curve *CurveCaller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "address_provider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProvider is a free data retrieval call binding the contract method 0xce50c2e7.
//
// Solidity: function address_provider() view returns(address)
func (_Curve *CurveSession) AddressProvider() (common.Address, error) {
	return _Curve.Contract.AddressProvider(&_Curve.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0xce50c2e7.
//
// Solidity: function address_provider() view returns(address)
func (_Curve *CurveCallerSession) AddressProvider() (common.Address, error) {
	return _Curve.Contract.AddressProvider(&_Curve.CallOpts)
}

// EstimateGasUsed is a free data retrieval call binding the contract method 0xb0bb365b.
//
// Solidity: function estimate_gas_used(address _pool, address _from, address _to) view returns(uint256)
func (_Curve *CurveCaller) EstimateGasUsed(opts *bind.CallOpts, _pool common.Address, _from common.Address, _to common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "estimate_gas_used", _pool, _from, _to)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateGasUsed is a free data retrieval call binding the contract method 0xb0bb365b.
//
// Solidity: function estimate_gas_used(address _pool, address _from, address _to) view returns(uint256)
func (_Curve *CurveSession) EstimateGasUsed(_pool common.Address, _from common.Address, _to common.Address) (*big.Int, error) {
	return _Curve.Contract.EstimateGasUsed(&_Curve.CallOpts, _pool, _from, _to)
}

// EstimateGasUsed is a free data retrieval call binding the contract method 0xb0bb365b.
//
// Solidity: function estimate_gas_used(address _pool, address _from, address _to) view returns(uint256)
func (_Curve *CurveCallerSession) EstimateGasUsed(_pool common.Address, _from common.Address, _to common.Address) (*big.Int, error) {
	return _Curve.Contract.EstimateGasUsed(&_Curve.CallOpts, _pool, _from, _to)
}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) view returns(address)
func (_Curve *CurveCaller) FindPoolForCoins(opts *bind.CallOpts, _from common.Address, _to common.Address) (common.Address, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "find_pool_for_coins", _from, _to)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) view returns(address)
func (_Curve *CurveSession) FindPoolForCoins(_from common.Address, _to common.Address) (common.Address, error) {
	return _Curve.Contract.FindPoolForCoins(&_Curve.CallOpts, _from, _to)
}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) view returns(address)
func (_Curve *CurveCallerSession) FindPoolForCoins(_from common.Address, _to common.Address) (common.Address, error) {
	return _Curve.Contract.FindPoolForCoins(&_Curve.CallOpts, _from, _to)
}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) view returns(address)
func (_Curve *CurveCaller) FindPoolForCoins0(opts *bind.CallOpts, _from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "find_pool_for_coins0", _from, _to, i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) view returns(address)
func (_Curve *CurveSession) FindPoolForCoins0(_from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	return _Curve.Contract.FindPoolForCoins0(&_Curve.CallOpts, _from, _to, i)
}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) view returns(address)
func (_Curve *CurveCallerSession) FindPoolForCoins0(_from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	return _Curve.Contract.FindPoolForCoins0(&_Curve.CallOpts, _from, _to, i)
}

// GaugeController is a free data retrieval call binding the contract method 0xd8b9a018.
//
// Solidity: function gauge_controller() view returns(address)
func (_Curve *CurveCaller) GaugeController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "gauge_controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GaugeController is a free data retrieval call binding the contract method 0xd8b9a018.
//
// Solidity: function gauge_controller() view returns(address)
func (_Curve *CurveSession) GaugeController() (common.Address, error) {
	return _Curve.Contract.GaugeController(&_Curve.CallOpts)
}

// GaugeController is a free data retrieval call binding the contract method 0xd8b9a018.
//
// Solidity: function gauge_controller() view returns(address)
func (_Curve *CurveCallerSession) GaugeController() (common.Address, error) {
	return _Curve.Contract.GaugeController(&_Curve.CallOpts)
}

// GetA is a free data retrieval call binding the contract method 0x55b30b19.
//
// Solidity: function get_A(address _pool) view returns(uint256)
func (_Curve *CurveCaller) GetA(opts *bind.CallOpts, _pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "get_A", _pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetA is a free data retrieval call binding the contract method 0x55b30b19.
//
// Solidity: function get_A(address _pool) view returns(uint256)
func (_Curve *CurveSession) GetA(_pool common.Address) (*big.Int, error) {
	return _Curve.Contract.GetA(&_Curve.CallOpts, _pool)
}

// GetA is a free data retrieval call binding the contract method 0x55b30b19.
//
// Solidity: function get_A(address _pool) view returns(uint256)
func (_Curve *CurveCallerSession) GetA(_pool common.Address) (*big.Int, error) {
	return _Curve.Contract.GetA(&_Curve.CallOpts, _pool)
}

// GetAdminBalances is a free data retrieval call binding the contract method 0xc11e45b8.
//
// Solidity: function get_admin_balances(address _pool) view returns(uint256[8])
func (_Curve *CurveCaller) GetAdminBalances(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "get_admin_balances", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetAdminBalances is a free data retrieval call binding the contract method 0xc11e45b8.
//
// Solidity: function get_admin_balances(address _pool) view returns(uint256[8])
func (_Curve *CurveSession) GetAdminBalances(_pool common.Address) ([8]*big.Int, error) {
	return _Curve.Contract.GetAdminBalances(&_Curve.CallOpts, _pool)
}

// GetAdminBalances is a free data retrieval call binding the contract method 0xc11e45b8.
//
// Solidity: function get_admin_balances(address _pool) view returns(uint256[8])
func (_Curve *CurveCallerSession) GetAdminBalances(_pool common.Address) ([8]*big.Int, error) {
	return _Curve.Contract.GetAdminBalances(&_Curve.CallOpts, _pool)
}

// GetBalances is a free data retrieval call binding the contract method 0x92e3cc2d.
//
// Solidity: function get_balances(address _pool) view returns(uint256[8])
func (_Curve *CurveCaller) GetBalances(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "get_balances", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetBalances is a free data retrieval call binding the contract method 0x92e3cc2d.
//
// Solidity: function get_balances(address _pool) view returns(uint256[8])
func (_Curve *CurveSession) GetBalances(_pool common.Address) ([8]*big.Int, error) {
	return _Curve.Contract.GetBalances(&_Curve.CallOpts, _pool)
}

// GetBalances is a free data retrieval call binding the contract method 0x92e3cc2d.
//
// Solidity: function get_balances(address _pool) view returns(uint256[8])
func (_Curve *CurveCallerSession) GetBalances(_pool common.Address) ([8]*big.Int, error) {
	return _Curve.Contract.GetBalances(&_Curve.CallOpts, _pool)
}

// GetCoinIndices is a free data retrieval call binding the contract method 0xeb85226d.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to) view returns(int128, int128, bool)
func (_Curve *CurveCaller) GetCoinIndices(opts *bind.CallOpts, _pool common.Address, _from common.Address, _to common.Address) (*big.Int, *big.Int, bool, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "get_coin_indices", _pool, _from, _to)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)

	return out0, out1, out2, err

}

// GetCoinIndices is a free data retrieval call binding the contract method 0xeb85226d.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to) view returns(int128, int128, bool)
func (_Curve *CurveSession) GetCoinIndices(_pool common.Address, _from common.Address, _to common.Address) (*big.Int, *big.Int, bool, error) {
	return _Curve.Contract.GetCoinIndices(&_Curve.CallOpts, _pool, _from, _to)
}

// GetCoinIndices is a free data retrieval call binding the contract method 0xeb85226d.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to) view returns(int128, int128, bool)
func (_Curve *CurveCallerSession) GetCoinIndices(_pool common.Address, _from common.Address, _to common.Address) (*big.Int, *big.Int, bool, error) {
	return _Curve.Contract.GetCoinIndices(&_Curve.CallOpts, _pool, _from, _to)
}

// GetCoins is a free data retrieval call binding the contract method 0x9ac90d3d.
//
// Solidity: function get_coins(address _pool) view returns(address[8])
func (_Curve *CurveCaller) GetCoins(opts *bind.CallOpts, _pool common.Address) ([8]common.Address, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "get_coins", _pool)

	if err != nil {
		return *new([8]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([8]common.Address)).(*[8]common.Address)

	return out0, err

}

// GetCoins is a free data retrieval call binding the contract method 0x9ac90d3d.
//
// Solidity: function get_coins(address _pool) view returns(address[8])
func (_Curve *CurveSession) GetCoins(_pool common.Address) ([8]common.Address, error) {
	return _Curve.Contract.GetCoins(&_Curve.CallOpts, _pool)
}

// GetCoins is a free data retrieval call binding the contract method 0x9ac90d3d.
//
// Solidity: function get_coins(address _pool) view returns(address[8])
func (_Curve *CurveCallerSession) GetCoins(_pool common.Address) ([8]common.Address, error) {
	return _Curve.Contract.GetCoins(&_Curve.CallOpts, _pool)
}

// GetDecimals is a free data retrieval call binding the contract method 0x52b51555.
//
// Solidity: function get_decimals(address _pool) view returns(uint256[8])
func (_Curve *CurveCaller) GetDecimals(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "get_decimals", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetDecimals is a free data retrieval call binding the contract method 0x52b51555.
//
// Solidity: function get_decimals(address _pool) view returns(uint256[8])
func (_Curve *CurveSession) GetDecimals(_pool common.Address) ([8]*big.Int, error) {
	return _Curve.Contract.GetDecimals(&_Curve.CallOpts, _pool)
}

// GetDecimals is a free data retrieval call binding the contract method 0x52b51555.
//
// Solidity: function get_decimals(address _pool) view returns(uint256[8])
func (_Curve *CurveCallerSession) GetDecimals(_pool common.Address) ([8]*big.Int, error) {
	return _Curve.Contract.GetDecimals(&_Curve.CallOpts, _pool)
}

// GetFees is a free data retrieval call binding the contract method 0x7cdb72b0.
//
// Solidity: function get_fees(address _pool) view returns(uint256[2])
func (_Curve *CurveCaller) GetFees(opts *bind.CallOpts, _pool common.Address) ([2]*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "get_fees", _pool)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetFees is a free data retrieval call binding the contract method 0x7cdb72b0.
//
// Solidity: function get_fees(address _pool) view returns(uint256[2])
func (_Curve *CurveSession) GetFees(_pool common.Address) ([2]*big.Int, error) {
	return _Curve.Contract.GetFees(&_Curve.CallOpts, _pool)
}

// GetFees is a free data retrieval call binding the contract method 0x7cdb72b0.
//
// Solidity: function get_fees(address _pool) view returns(uint256[2])
func (_Curve *CurveCallerSession) GetFees(_pool common.Address) ([2]*big.Int, error) {
	return _Curve.Contract.GetFees(&_Curve.CallOpts, _pool)
}

// GetGauges is a free data retrieval call binding the contract method 0x56059ffb.
//
// Solidity: function get_gauges(address _pool) view returns(address[10], int128[10])
func (_Curve *CurveCaller) GetGauges(opts *bind.CallOpts, _pool common.Address) ([10]common.Address, [10]*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "get_gauges", _pool)

	if err != nil {
		return *new([10]common.Address), *new([10]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([10]common.Address)).(*[10]common.Address)
	out1 := *abi.ConvertType(out[1], new([10]*big.Int)).(*[10]*big.Int)

	return out0, out1, err

}

// GetGauges is a free data retrieval call binding the contract method 0x56059ffb.
//
// Solidity: function get_gauges(address _pool) view returns(address[10], int128[10])
func (_Curve *CurveSession) GetGauges(_pool common.Address) ([10]common.Address, [10]*big.Int, error) {
	return _Curve.Contract.GetGauges(&_Curve.CallOpts, _pool)
}

// GetGauges is a free data retrieval call binding the contract method 0x56059ffb.
//
// Solidity: function get_gauges(address _pool) view returns(address[10], int128[10])
func (_Curve *CurveCallerSession) GetGauges(_pool common.Address) ([10]common.Address, [10]*big.Int, error) {
	return _Curve.Contract.GetGauges(&_Curve.CallOpts, _pool)
}

// GetLpToken is a free data retrieval call binding the contract method 0x37951049.
//
// Solidity: function get_lp_token(address arg0) view returns(address)
func (_Curve *CurveCaller) GetLpToken(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "get_lp_token", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLpToken is a free data retrieval call binding the contract method 0x37951049.
//
// Solidity: function get_lp_token(address arg0) view returns(address)
func (_Curve *CurveSession) GetLpToken(arg0 common.Address) (common.Address, error) {
	return _Curve.Contract.GetLpToken(&_Curve.CallOpts, arg0)
}

// GetLpToken is a free data retrieval call binding the contract method 0x37951049.
//
// Solidity: function get_lp_token(address arg0) view returns(address)
func (_Curve *CurveCallerSession) GetLpToken(arg0 common.Address) (common.Address, error) {
	return _Curve.Contract.GetLpToken(&_Curve.CallOpts, arg0)
}

// GetNCoins is a free data retrieval call binding the contract method 0x940494f1.
//
// Solidity: function get_n_coins(address _pool) view returns(uint256[2])
func (_Curve *CurveCaller) GetNCoins(opts *bind.CallOpts, _pool common.Address) ([2]*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "get_n_coins", _pool)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetNCoins is a free data retrieval call binding the contract method 0x940494f1.
//
// Solidity: function get_n_coins(address _pool) view returns(uint256[2])
func (_Curve *CurveSession) GetNCoins(_pool common.Address) ([2]*big.Int, error) {
	return _Curve.Contract.GetNCoins(&_Curve.CallOpts, _pool)
}

// GetNCoins is a free data retrieval call binding the contract method 0x940494f1.
//
// Solidity: function get_n_coins(address _pool) view returns(uint256[2])
func (_Curve *CurveCallerSession) GetNCoins(_pool common.Address) ([2]*big.Int, error) {
	return _Curve.Contract.GetNCoins(&_Curve.CallOpts, _pool)
}

// GetParameters is a free data retrieval call binding the contract method 0x1f80a957.
//
// Solidity: function get_parameters(address _pool) view returns(uint256 A, uint256 future_A, uint256 fee, uint256 admin_fee, uint256 future_fee, uint256 future_admin_fee, address future_owner, uint256 initial_A, uint256 initial_A_time, uint256 future_A_time)
func (_Curve *CurveCaller) GetParameters(opts *bind.CallOpts, _pool common.Address) (struct {
	A              *big.Int
	FutureA        *big.Int
	Fee            *big.Int
	AdminFee       *big.Int
	FutureFee      *big.Int
	FutureAdminFee *big.Int
	FutureOwner    common.Address
	InitialA       *big.Int
	InitialATime   *big.Int
	FutureATime    *big.Int
}, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "get_parameters", _pool)

	outstruct := new(struct {
		A              *big.Int
		FutureA        *big.Int
		Fee            *big.Int
		AdminFee       *big.Int
		FutureFee      *big.Int
		FutureAdminFee *big.Int
		FutureOwner    common.Address
		InitialA       *big.Int
		InitialATime   *big.Int
		FutureATime    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.A = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FutureA = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Fee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AdminFee = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.FutureFee = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.FutureAdminFee = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.FutureOwner = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.InitialA = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.InitialATime = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.FutureATime = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetParameters is a free data retrieval call binding the contract method 0x1f80a957.
//
// Solidity: function get_parameters(address _pool) view returns(uint256 A, uint256 future_A, uint256 fee, uint256 admin_fee, uint256 future_fee, uint256 future_admin_fee, address future_owner, uint256 initial_A, uint256 initial_A_time, uint256 future_A_time)
func (_Curve *CurveSession) GetParameters(_pool common.Address) (struct {
	A              *big.Int
	FutureA        *big.Int
	Fee            *big.Int
	AdminFee       *big.Int
	FutureFee      *big.Int
	FutureAdminFee *big.Int
	FutureOwner    common.Address
	InitialA       *big.Int
	InitialATime   *big.Int
	FutureATime    *big.Int
}, error) {
	return _Curve.Contract.GetParameters(&_Curve.CallOpts, _pool)
}

// GetParameters is a free data retrieval call binding the contract method 0x1f80a957.
//
// Solidity: function get_parameters(address _pool) view returns(uint256 A, uint256 future_A, uint256 fee, uint256 admin_fee, uint256 future_fee, uint256 future_admin_fee, address future_owner, uint256 initial_A, uint256 initial_A_time, uint256 future_A_time)
func (_Curve *CurveCallerSession) GetParameters(_pool common.Address) (struct {
	A              *big.Int
	FutureA        *big.Int
	Fee            *big.Int
	AdminFee       *big.Int
	FutureFee      *big.Int
	FutureAdminFee *big.Int
	FutureOwner    common.Address
	InitialA       *big.Int
	InitialATime   *big.Int
	FutureATime    *big.Int
}, error) {
	return _Curve.Contract.GetParameters(&_Curve.CallOpts, _pool)
}

// GetPoolFromLpToken is a free data retrieval call binding the contract method 0xbdf475c3.
//
// Solidity: function get_pool_from_lp_token(address arg0) view returns(address)
func (_Curve *CurveCaller) GetPoolFromLpToken(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "get_pool_from_lp_token", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolFromLpToken is a free data retrieval call binding the contract method 0xbdf475c3.
//
// Solidity: function get_pool_from_lp_token(address arg0) view returns(address)
func (_Curve *CurveSession) GetPoolFromLpToken(arg0 common.Address) (common.Address, error) {
	return _Curve.Contract.GetPoolFromLpToken(&_Curve.CallOpts, arg0)
}

// GetPoolFromLpToken is a free data retrieval call binding the contract method 0xbdf475c3.
//
// Solidity: function get_pool_from_lp_token(address arg0) view returns(address)
func (_Curve *CurveCallerSession) GetPoolFromLpToken(arg0 common.Address) (common.Address, error) {
	return _Curve.Contract.GetPoolFromLpToken(&_Curve.CallOpts, arg0)
}

// GetRates is a free data retrieval call binding the contract method 0xce99e45a.
//
// Solidity: function get_rates(address _pool) view returns(uint256[8])
func (_Curve *CurveCaller) GetRates(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "get_rates", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetRates is a free data retrieval call binding the contract method 0xce99e45a.
//
// Solidity: function get_rates(address _pool) view returns(uint256[8])
func (_Curve *CurveSession) GetRates(_pool common.Address) ([8]*big.Int, error) {
	return _Curve.Contract.GetRates(&_Curve.CallOpts, _pool)
}

// GetRates is a free data retrieval call binding the contract method 0xce99e45a.
//
// Solidity: function get_rates(address _pool) view returns(uint256[8])
func (_Curve *CurveCallerSession) GetRates(_pool common.Address) ([8]*big.Int, error) {
	return _Curve.Contract.GetRates(&_Curve.CallOpts, _pool)
}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x59f4f351.
//
// Solidity: function get_underlying_balances(address _pool) view returns(uint256[8])
func (_Curve *CurveCaller) GetUnderlyingBalances(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "get_underlying_balances", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x59f4f351.
//
// Solidity: function get_underlying_balances(address _pool) view returns(uint256[8])
func (_Curve *CurveSession) GetUnderlyingBalances(_pool common.Address) ([8]*big.Int, error) {
	return _Curve.Contract.GetUnderlyingBalances(&_Curve.CallOpts, _pool)
}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x59f4f351.
//
// Solidity: function get_underlying_balances(address _pool) view returns(uint256[8])
func (_Curve *CurveCallerSession) GetUnderlyingBalances(_pool common.Address) ([8]*big.Int, error) {
	return _Curve.Contract.GetUnderlyingBalances(&_Curve.CallOpts, _pool)
}

// GetUnderlyingCoins is a free data retrieval call binding the contract method 0xa77576ef.
//
// Solidity: function get_underlying_coins(address _pool) view returns(address[8])
func (_Curve *CurveCaller) GetUnderlyingCoins(opts *bind.CallOpts, _pool common.Address) ([8]common.Address, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "get_underlying_coins", _pool)

	if err != nil {
		return *new([8]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([8]common.Address)).(*[8]common.Address)

	return out0, err

}

// GetUnderlyingCoins is a free data retrieval call binding the contract method 0xa77576ef.
//
// Solidity: function get_underlying_coins(address _pool) view returns(address[8])
func (_Curve *CurveSession) GetUnderlyingCoins(_pool common.Address) ([8]common.Address, error) {
	return _Curve.Contract.GetUnderlyingCoins(&_Curve.CallOpts, _pool)
}

// GetUnderlyingCoins is a free data retrieval call binding the contract method 0xa77576ef.
//
// Solidity: function get_underlying_coins(address _pool) view returns(address[8])
func (_Curve *CurveCallerSession) GetUnderlyingCoins(_pool common.Address) ([8]common.Address, error) {
	return _Curve.Contract.GetUnderlyingCoins(&_Curve.CallOpts, _pool)
}

// GetUnderlyingDecimals is a free data retrieval call binding the contract method 0x4cb088f1.
//
// Solidity: function get_underlying_decimals(address _pool) view returns(uint256[8])
func (_Curve *CurveCaller) GetUnderlyingDecimals(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "get_underlying_decimals", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetUnderlyingDecimals is a free data retrieval call binding the contract method 0x4cb088f1.
//
// Solidity: function get_underlying_decimals(address _pool) view returns(uint256[8])
func (_Curve *CurveSession) GetUnderlyingDecimals(_pool common.Address) ([8]*big.Int, error) {
	return _Curve.Contract.GetUnderlyingDecimals(&_Curve.CallOpts, _pool)
}

// GetUnderlyingDecimals is a free data retrieval call binding the contract method 0x4cb088f1.
//
// Solidity: function get_underlying_decimals(address _pool) view returns(uint256[8])
func (_Curve *CurveCallerSession) GetUnderlyingDecimals(_pool common.Address) ([8]*big.Int, error) {
	return _Curve.Contract.GetUnderlyingDecimals(&_Curve.CallOpts, _pool)
}

// GetVirtualPriceFromLpToken is a free data retrieval call binding the contract method 0xc5b7074a.
//
// Solidity: function get_virtual_price_from_lp_token(address _token) view returns(uint256)
func (_Curve *CurveCaller) GetVirtualPriceFromLpToken(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "get_virtual_price_from_lp_token", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPriceFromLpToken is a free data retrieval call binding the contract method 0xc5b7074a.
//
// Solidity: function get_virtual_price_from_lp_token(address _token) view returns(uint256)
func (_Curve *CurveSession) GetVirtualPriceFromLpToken(_token common.Address) (*big.Int, error) {
	return _Curve.Contract.GetVirtualPriceFromLpToken(&_Curve.CallOpts, _token)
}

// GetVirtualPriceFromLpToken is a free data retrieval call binding the contract method 0xc5b7074a.
//
// Solidity: function get_virtual_price_from_lp_token(address _token) view returns(uint256)
func (_Curve *CurveCallerSession) GetVirtualPriceFromLpToken(_token common.Address) (*big.Int, error) {
	return _Curve.Contract.GetVirtualPriceFromLpToken(&_Curve.CallOpts, _token)
}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() view returns(uint256)
func (_Curve *CurveCaller) PoolCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "pool_count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() view returns(uint256)
func (_Curve *CurveSession) PoolCount() (*big.Int, error) {
	return _Curve.Contract.PoolCount(&_Curve.CallOpts)
}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() view returns(uint256)
func (_Curve *CurveCallerSession) PoolCount() (*big.Int, error) {
	return _Curve.Contract.PoolCount(&_Curve.CallOpts)
}

// PoolList is a free data retrieval call binding the contract method 0x3a1d5d8e.
//
// Solidity: function pool_list(uint256 arg0) view returns(address)
func (_Curve *CurveCaller) PoolList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Curve.contract.Call(opts, &out, "pool_list", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolList is a free data retrieval call binding the contract method 0x3a1d5d8e.
//
// Solidity: function pool_list(uint256 arg0) view returns(address)
func (_Curve *CurveSession) PoolList(arg0 *big.Int) (common.Address, error) {
	return _Curve.Contract.PoolList(&_Curve.CallOpts, arg0)
}

// PoolList is a free data retrieval call binding the contract method 0x3a1d5d8e.
//
// Solidity: function pool_list(uint256 arg0) view returns(address)
func (_Curve *CurveCallerSession) PoolList(arg0 *big.Int) (common.Address, error) {
	return _Curve.Contract.PoolList(&_Curve.CallOpts, arg0)
}

// AddMetapool is a paid mutator transaction binding the contract method 0xcfedad3c.
//
// Solidity: function add_metapool(address _pool, uint256 _n_coins, address _lp_token, uint256 _decimals) returns()
func (_Curve *CurveTransactor) AddMetapool(opts *bind.TransactOpts, _pool common.Address, _n_coins *big.Int, _lp_token common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _Curve.contract.Transact(opts, "add_metapool", _pool, _n_coins, _lp_token, _decimals)
}

// AddMetapool is a paid mutator transaction binding the contract method 0xcfedad3c.
//
// Solidity: function add_metapool(address _pool, uint256 _n_coins, address _lp_token, uint256 _decimals) returns()
func (_Curve *CurveSession) AddMetapool(_pool common.Address, _n_coins *big.Int, _lp_token common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _Curve.Contract.AddMetapool(&_Curve.TransactOpts, _pool, _n_coins, _lp_token, _decimals)
}

// AddMetapool is a paid mutator transaction binding the contract method 0xcfedad3c.
//
// Solidity: function add_metapool(address _pool, uint256 _n_coins, address _lp_token, uint256 _decimals) returns()
func (_Curve *CurveTransactorSession) AddMetapool(_pool common.Address, _n_coins *big.Int, _lp_token common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _Curve.Contract.AddMetapool(&_Curve.TransactOpts, _pool, _n_coins, _lp_token, _decimals)
}

// AddPool is a paid mutator transaction binding the contract method 0x290b1e48.
//
// Solidity: function add_pool(address _pool, uint256 _n_coins, address _lp_token, bytes32 _rate_method_id, uint256 _decimals, uint256 _underlying_decimals, bool _has_initial_A, bool _is_v1) returns()
func (_Curve *CurveTransactor) AddPool(opts *bind.TransactOpts, _pool common.Address, _n_coins *big.Int, _lp_token common.Address, _rate_method_id [32]byte, _decimals *big.Int, _underlying_decimals *big.Int, _has_initial_A bool, _is_v1 bool) (*types.Transaction, error) {
	return _Curve.contract.Transact(opts, "add_pool", _pool, _n_coins, _lp_token, _rate_method_id, _decimals, _underlying_decimals, _has_initial_A, _is_v1)
}

// AddPool is a paid mutator transaction binding the contract method 0x290b1e48.
//
// Solidity: function add_pool(address _pool, uint256 _n_coins, address _lp_token, bytes32 _rate_method_id, uint256 _decimals, uint256 _underlying_decimals, bool _has_initial_A, bool _is_v1) returns()
func (_Curve *CurveSession) AddPool(_pool common.Address, _n_coins *big.Int, _lp_token common.Address, _rate_method_id [32]byte, _decimals *big.Int, _underlying_decimals *big.Int, _has_initial_A bool, _is_v1 bool) (*types.Transaction, error) {
	return _Curve.Contract.AddPool(&_Curve.TransactOpts, _pool, _n_coins, _lp_token, _rate_method_id, _decimals, _underlying_decimals, _has_initial_A, _is_v1)
}

// AddPool is a paid mutator transaction binding the contract method 0x290b1e48.
//
// Solidity: function add_pool(address _pool, uint256 _n_coins, address _lp_token, bytes32 _rate_method_id, uint256 _decimals, uint256 _underlying_decimals, bool _has_initial_A, bool _is_v1) returns()
func (_Curve *CurveTransactorSession) AddPool(_pool common.Address, _n_coins *big.Int, _lp_token common.Address, _rate_method_id [32]byte, _decimals *big.Int, _underlying_decimals *big.Int, _has_initial_A bool, _is_v1 bool) (*types.Transaction, error) {
	return _Curve.Contract.AddPool(&_Curve.TransactOpts, _pool, _n_coins, _lp_token, _rate_method_id, _decimals, _underlying_decimals, _has_initial_A, _is_v1)
}

// AddPoolWithoutUnderlying is a paid mutator transaction binding the contract method 0x60c06f89.
//
// Solidity: function add_pool_without_underlying(address _pool, uint256 _n_coins, address _lp_token, bytes32 _rate_method_id, uint256 _decimals, uint256 _use_rates, bool _has_initial_A, bool _is_v1) returns()
func (_Curve *CurveTransactor) AddPoolWithoutUnderlying(opts *bind.TransactOpts, _pool common.Address, _n_coins *big.Int, _lp_token common.Address, _rate_method_id [32]byte, _decimals *big.Int, _use_rates *big.Int, _has_initial_A bool, _is_v1 bool) (*types.Transaction, error) {
	return _Curve.contract.Transact(opts, "add_pool_without_underlying", _pool, _n_coins, _lp_token, _rate_method_id, _decimals, _use_rates, _has_initial_A, _is_v1)
}

// AddPoolWithoutUnderlying is a paid mutator transaction binding the contract method 0x60c06f89.
//
// Solidity: function add_pool_without_underlying(address _pool, uint256 _n_coins, address _lp_token, bytes32 _rate_method_id, uint256 _decimals, uint256 _use_rates, bool _has_initial_A, bool _is_v1) returns()
func (_Curve *CurveSession) AddPoolWithoutUnderlying(_pool common.Address, _n_coins *big.Int, _lp_token common.Address, _rate_method_id [32]byte, _decimals *big.Int, _use_rates *big.Int, _has_initial_A bool, _is_v1 bool) (*types.Transaction, error) {
	return _Curve.Contract.AddPoolWithoutUnderlying(&_Curve.TransactOpts, _pool, _n_coins, _lp_token, _rate_method_id, _decimals, _use_rates, _has_initial_A, _is_v1)
}

// AddPoolWithoutUnderlying is a paid mutator transaction binding the contract method 0x60c06f89.
//
// Solidity: function add_pool_without_underlying(address _pool, uint256 _n_coins, address _lp_token, bytes32 _rate_method_id, uint256 _decimals, uint256 _use_rates, bool _has_initial_A, bool _is_v1) returns()
func (_Curve *CurveTransactorSession) AddPoolWithoutUnderlying(_pool common.Address, _n_coins *big.Int, _lp_token common.Address, _rate_method_id [32]byte, _decimals *big.Int, _use_rates *big.Int, _has_initial_A bool, _is_v1 bool) (*types.Transaction, error) {
	return _Curve.Contract.AddPoolWithoutUnderlying(&_Curve.TransactOpts, _pool, _n_coins, _lp_token, _rate_method_id, _decimals, _use_rates, _has_initial_A, _is_v1)
}

// RemovePool is a paid mutator transaction binding the contract method 0x474932b0.
//
// Solidity: function remove_pool(address _pool) returns()
func (_Curve *CurveTransactor) RemovePool(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _Curve.contract.Transact(opts, "remove_pool", _pool)
}

// RemovePool is a paid mutator transaction binding the contract method 0x474932b0.
//
// Solidity: function remove_pool(address _pool) returns()
func (_Curve *CurveSession) RemovePool(_pool common.Address) (*types.Transaction, error) {
	return _Curve.Contract.RemovePool(&_Curve.TransactOpts, _pool)
}

// RemovePool is a paid mutator transaction binding the contract method 0x474932b0.
//
// Solidity: function remove_pool(address _pool) returns()
func (_Curve *CurveTransactorSession) RemovePool(_pool common.Address) (*types.Transaction, error) {
	return _Curve.Contract.RemovePool(&_Curve.TransactOpts, _pool)
}

// SetCoinGasEstimates is a paid mutator transaction binding the contract method 0x237f89f2.
//
// Solidity: function set_coin_gas_estimates(address[10] _addr, uint256[10] _amount) returns()
func (_Curve *CurveTransactor) SetCoinGasEstimates(opts *bind.TransactOpts, _addr [10]common.Address, _amount [10]*big.Int) (*types.Transaction, error) {
	return _Curve.contract.Transact(opts, "set_coin_gas_estimates", _addr, _amount)
}

// SetCoinGasEstimates is a paid mutator transaction binding the contract method 0x237f89f2.
//
// Solidity: function set_coin_gas_estimates(address[10] _addr, uint256[10] _amount) returns()
func (_Curve *CurveSession) SetCoinGasEstimates(_addr [10]common.Address, _amount [10]*big.Int) (*types.Transaction, error) {
	return _Curve.Contract.SetCoinGasEstimates(&_Curve.TransactOpts, _addr, _amount)
}

// SetCoinGasEstimates is a paid mutator transaction binding the contract method 0x237f89f2.
//
// Solidity: function set_coin_gas_estimates(address[10] _addr, uint256[10] _amount) returns()
func (_Curve *CurveTransactorSession) SetCoinGasEstimates(_addr [10]common.Address, _amount [10]*big.Int) (*types.Transaction, error) {
	return _Curve.Contract.SetCoinGasEstimates(&_Curve.TransactOpts, _addr, _amount)
}

// SetGasEstimateContract is a paid mutator transaction binding the contract method 0xca991b14.
//
// Solidity: function set_gas_estimate_contract(address _pool, address _estimator) returns()
func (_Curve *CurveTransactor) SetGasEstimateContract(opts *bind.TransactOpts, _pool common.Address, _estimator common.Address) (*types.Transaction, error) {
	return _Curve.contract.Transact(opts, "set_gas_estimate_contract", _pool, _estimator)
}

// SetGasEstimateContract is a paid mutator transaction binding the contract method 0xca991b14.
//
// Solidity: function set_gas_estimate_contract(address _pool, address _estimator) returns()
func (_Curve *CurveSession) SetGasEstimateContract(_pool common.Address, _estimator common.Address) (*types.Transaction, error) {
	return _Curve.Contract.SetGasEstimateContract(&_Curve.TransactOpts, _pool, _estimator)
}

// SetGasEstimateContract is a paid mutator transaction binding the contract method 0xca991b14.
//
// Solidity: function set_gas_estimate_contract(address _pool, address _estimator) returns()
func (_Curve *CurveTransactorSession) SetGasEstimateContract(_pool common.Address, _estimator common.Address) (*types.Transaction, error) {
	return _Curve.Contract.SetGasEstimateContract(&_Curve.TransactOpts, _pool, _estimator)
}

// SetLiquidityGauges is a paid mutator transaction binding the contract method 0xef6b9788.
//
// Solidity: function set_liquidity_gauges(address _pool, address[10] _liquidity_gauges) returns()
func (_Curve *CurveTransactor) SetLiquidityGauges(opts *bind.TransactOpts, _pool common.Address, _liquidity_gauges [10]common.Address) (*types.Transaction, error) {
	return _Curve.contract.Transact(opts, "set_liquidity_gauges", _pool, _liquidity_gauges)
}

// SetLiquidityGauges is a paid mutator transaction binding the contract method 0xef6b9788.
//
// Solidity: function set_liquidity_gauges(address _pool, address[10] _liquidity_gauges) returns()
func (_Curve *CurveSession) SetLiquidityGauges(_pool common.Address, _liquidity_gauges [10]common.Address) (*types.Transaction, error) {
	return _Curve.Contract.SetLiquidityGauges(&_Curve.TransactOpts, _pool, _liquidity_gauges)
}

// SetLiquidityGauges is a paid mutator transaction binding the contract method 0xef6b9788.
//
// Solidity: function set_liquidity_gauges(address _pool, address[10] _liquidity_gauges) returns()
func (_Curve *CurveTransactorSession) SetLiquidityGauges(_pool common.Address, _liquidity_gauges [10]common.Address) (*types.Transaction, error) {
	return _Curve.Contract.SetLiquidityGauges(&_Curve.TransactOpts, _pool, _liquidity_gauges)
}

// SetPoolGasEstimates is a paid mutator transaction binding the contract method 0x0733b67a.
//
// Solidity: function set_pool_gas_estimates(address[5] _addr, uint256[2][5] _amount) returns()
func (_Curve *CurveTransactor) SetPoolGasEstimates(opts *bind.TransactOpts, _addr [5]common.Address, _amount [5][2]*big.Int) (*types.Transaction, error) {
	return _Curve.contract.Transact(opts, "set_pool_gas_estimates", _addr, _amount)
}

// SetPoolGasEstimates is a paid mutator transaction binding the contract method 0x0733b67a.
//
// Solidity: function set_pool_gas_estimates(address[5] _addr, uint256[2][5] _amount) returns()
func (_Curve *CurveSession) SetPoolGasEstimates(_addr [5]common.Address, _amount [5][2]*big.Int) (*types.Transaction, error) {
	return _Curve.Contract.SetPoolGasEstimates(&_Curve.TransactOpts, _addr, _amount)
}

// SetPoolGasEstimates is a paid mutator transaction binding the contract method 0x0733b67a.
//
// Solidity: function set_pool_gas_estimates(address[5] _addr, uint256[2][5] _amount) returns()
func (_Curve *CurveTransactorSession) SetPoolGasEstimates(_addr [5]common.Address, _amount [5][2]*big.Int) (*types.Transaction, error) {
	return _Curve.Contract.SetPoolGasEstimates(&_Curve.TransactOpts, _addr, _amount)
}

// CurvePoolAddedIterator is returned from FilterPoolAdded and is used to iterate over the raw logs and unpacked data for PoolAdded events raised by the Curve contract.
type CurvePoolAddedIterator struct {
	Event *CurvePoolAdded // Event containing the contract specifics and raw log

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
func (it *CurvePoolAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolAdded)
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
		it.Event = new(CurvePoolAdded)
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
func (it *CurvePoolAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolAdded represents a PoolAdded event raised by the Curve contract.
type CurvePoolAdded struct {
	Pool         common.Address
	RateMethodId []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPoolAdded is a free log retrieval operation binding the contract event 0xe485c16479ab7092c0b3fc4649843c06be7f072194675261590c84473ab0aea9.
//
// Solidity: event PoolAdded(address indexed pool, bytes rate_method_id)
func (_Curve *CurveFilterer) FilterPoolAdded(opts *bind.FilterOpts, pool []common.Address) (*CurvePoolAddedIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Curve.contract.FilterLogs(opts, "PoolAdded", poolRule)
	if err != nil {
		return nil, err
	}
	return &CurvePoolAddedIterator{contract: _Curve.contract, event: "PoolAdded", logs: logs, sub: sub}, nil
}

// WatchPoolAdded is a free log subscription operation binding the contract event 0xe485c16479ab7092c0b3fc4649843c06be7f072194675261590c84473ab0aea9.
//
// Solidity: event PoolAdded(address indexed pool, bytes rate_method_id)
func (_Curve *CurveFilterer) WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *CurvePoolAdded, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Curve.contract.WatchLogs(opts, "PoolAdded", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolAdded)
				if err := _Curve.contract.UnpackLog(event, "PoolAdded", log); err != nil {
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

// ParsePoolAdded is a log parse operation binding the contract event 0xe485c16479ab7092c0b3fc4649843c06be7f072194675261590c84473ab0aea9.
//
// Solidity: event PoolAdded(address indexed pool, bytes rate_method_id)
func (_Curve *CurveFilterer) ParsePoolAdded(log types.Log) (*CurvePoolAdded, error) {
	event := new(CurvePoolAdded)
	if err := _Curve.contract.UnpackLog(event, "PoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolRemovedIterator is returned from FilterPoolRemoved and is used to iterate over the raw logs and unpacked data for PoolRemoved events raised by the Curve contract.
type CurvePoolRemovedIterator struct {
	Event *CurvePoolRemoved // Event containing the contract specifics and raw log

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
func (it *CurvePoolRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolRemoved)
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
		it.Event = new(CurvePoolRemoved)
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
func (it *CurvePoolRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolRemoved represents a PoolRemoved event raised by the Curve contract.
type CurvePoolRemoved struct {
	Pool common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPoolRemoved is a free log retrieval operation binding the contract event 0x4106dfdaa577573db51c0ca93f766dbedfa0758faa2e7f5bcdb7c142be803c3f.
//
// Solidity: event PoolRemoved(address indexed pool)
func (_Curve *CurveFilterer) FilterPoolRemoved(opts *bind.FilterOpts, pool []common.Address) (*CurvePoolRemovedIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Curve.contract.FilterLogs(opts, "PoolRemoved", poolRule)
	if err != nil {
		return nil, err
	}
	return &CurvePoolRemovedIterator{contract: _Curve.contract, event: "PoolRemoved", logs: logs, sub: sub}, nil
}

// WatchPoolRemoved is a free log subscription operation binding the contract event 0x4106dfdaa577573db51c0ca93f766dbedfa0758faa2e7f5bcdb7c142be803c3f.
//
// Solidity: event PoolRemoved(address indexed pool)
func (_Curve *CurveFilterer) WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *CurvePoolRemoved, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Curve.contract.WatchLogs(opts, "PoolRemoved", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolRemoved)
				if err := _Curve.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
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

// ParsePoolRemoved is a log parse operation binding the contract event 0x4106dfdaa577573db51c0ca93f766dbedfa0758faa2e7f5bcdb7c142be803c3f.
//
// Solidity: event PoolRemoved(address indexed pool)
func (_Curve *CurveFilterer) ParsePoolRemoved(log types.Log) (*CurvePoolRemoved, error) {
	event := new(CurvePoolRemoved)
	if err := _Curve.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
