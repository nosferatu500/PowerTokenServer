package controllers

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"PowerTokenServer/model"
	"github.com/ethereum/go-ethereum/common"
	"PowerTokenServer/gocontracts/token"
	"github.com/ethereum/go-ethereum/crypto"
	"fmt"
	"strings"
	"io/ioutil"
	"encoding/json"
	"strconv"
	"github.com/ethereum/go-ethereum/ethclient"
	"net/http"
)

// TODO: Add condition for catch exceptions

const (
	contractAddress = "0x035483Bc81b0982a07966522510AA6Ff761dE848"
	connectionString = "https://rinkeby.infura.io/dUQe3kE7aGgdnkBmEAny"
)

func generateAuth() (*bind.TransactOpts, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	return bind.NewKeyedTransactor(key), err
}

func createPowerTokenSession(auth *bind.TransactOpts) (*token.PowerTokenSession, error) {

	var contract *token.PowerToken
	var err error

	if model.Simulation {
		contract, err = token.NewPowerToken(common.HexToAddress(model.TokenAddress), model.Simulator)
		if err != nil {
			return nil, err
		}
	} else {
		contract, err = token.NewPowerToken(common.HexToAddress(contractAddress), model.Connect)
		if err != nil {
			return nil, err
		}
	}

	session := token.PowerTokenSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
			From:    auth.From,
		},
		TransactOpts: bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasLimit: model.Gaslimit,
		},
	}
	return &session, err
}

func GetTotalSupply(w http.ResponseWriter, r *http.Request) {
	rpc_cli, _ := ethclient.Dial(connectionString)
	model.Connect = rpc_cli

	auth, _ := generateAuth()

	session, _ := createPowerTokenSession(auth)

	total, _ := session.TotalSupply()

	data := &model.BasicStringResp{}

	data.Result = total.String()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func GetSymbol(w http.ResponseWriter, r *http.Request) {
	rpc_cli, _ := ethclient.Dial(connectionString)
	model.Connect = rpc_cli

	auth, _ := generateAuth()

	session, _ := createPowerTokenSession(auth)

	total, _ := session.Symbol()

	data := &model.BasicStringResp{}

	data.Result = total
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func GetPausedStatus(w http.ResponseWriter, r *http.Request) {
	rpc_cli, _ := ethclient.Dial(connectionString)
	model.Connect = rpc_cli

	auth, _ := generateAuth()

	session, _ := createPowerTokenSession(auth)

	total, _ := session.Paused()

	data := &model.BasicBoolResp{}

	data.Result = total
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}










///////////////////////////////////////////////////////////////
/////////////////TEST_FUNC/////////////////////////////////////
///////////////////////////////////////////////////////////////

func SayhelloName2(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  // parse arguments, you have to call this by yourself
	fmt.Println(r.Form)  // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie2!") // send data to client side
}

func GetBlockNumber(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Get(`https://api.infura.io/v1/jsonrpc/mainnet/eth_blockNumber?token=dUQe3kE7aGgdnkBmEAny`)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)


	fmt.Println(resp.Body)
	fmt.Println(body)

	var res = new(model.BasicStringResp)
	_ = json.Unmarshal(body, &res)

	data := &model.BasicStringResp{}

	hex := strings.TrimPrefix(res.Result, "0x")
	x, _ := strconv.ParseInt(hex, 16, 64)
	s := strconv.Itoa(int(x))

	data.Result = s
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
	//fmt.Fprintf(w, res.Result)
}

func GetTotalSupplyTest(w http.ResponseWriter, r *http.Request) {
	rpc_cli, _ := ethclient.Dial(connectionString)

	model.Connect = rpc_cli
	fmt.Println("connect", model.Connect)

	auth, _ := generateAuth()
	fmt.Println("auth",auth)

	session, _ := createPowerTokenSession(auth)
	fmt.Println("session",session)

	total, _ := session.TotalSupply()
	fmt.Println("total",total)

	balance, _ := session.BalanceOf(common.HexToAddress("0xd0a6e6c54dbc68db5db3a091b171a77407ff7ccf"))
	fmt.Println("balance",balance)

	var res = new(model.BasicStringResp)
	//_ = json.Unmarshal(total, &res)

	data := &model.BasicStringResp{}

	hex := strings.TrimPrefix(res.Result, "0x")
	x, _ := strconv.ParseInt(hex, 16, 64)
	s := strconv.Itoa(int(x))

	data.Result = s
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
	//fmt.Fprintf(w, res.Result)
}
