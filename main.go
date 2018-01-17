//go:generate abigen --sol ./contracts/SingleTokenCoin.sol --pkg token  --out ./gocontracts/token/token.go

package main

import (
	"net/http"
	"fmt"
	"strings"
	"io/ioutil"
	"encoding/json"
	"strconv"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"PowerTokenServer/gocontracts/token"
	"github.com/ethereum/go-ethereum/common"
	"PowerTokenServer/model"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/labstack/gommon/log"

	"github.com/ethereum/go-ethereum/ethclient"
)

// token=dUQe3kE7aGgdnkBmEAny

// "https://api.infura.io/v1/jsonrpc/mainnet/eth_blockNumber?token=dUQe3kE7aGgdnkBmEAny"

// Mainnet	test network	https://mainnet.infura.io/dUQe3kE7aGgdnkBmEAny
// Ropsten	test network	https://ropsten.infura.io/dUQe3kE7aGgdnkBmEAny
// INFURAnet	test network	https://infuranet.infura.io/dUQe3kE7aGgdnkBmEAny
// Kovan	test network	https://kovan.infura.io/dUQe3kE7aGgdnkBmEAny
// Rinkeby	test network	https://rinkeby.infura.io/dUQe3kE7aGgdnkBmEAny
// IPFS	gateway	https://mainnet.infura.io/dUQe3kE7aGgdnkBmEAny

type Routes struct {
}

type Response struct {
	Result string `json:"result"`
}

func (p *Routes)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloName2(w, r)
		return
	}

	if r.URL.Path == "/test" {
		sayhelloName2(w, r)
		return
	}

	if r.URL.Path == "/blockNumber" {
		getBlockNumber(w, r)
		return
	}

	if r.URL.Path == "/total" {
		getTotalSupply(w, r)
		return
	}

	http.NotFound(w, r)
	return
}

func sayhelloName2(w http.ResponseWriter, r *http.Request) {
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

func getBlockNumber(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Get(`https://api.infura.io/v1/jsonrpc/mainnet/eth_blockNumber?token=dUQe3kE7aGgdnkBmEAny`)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)


	fmt.Println(resp.Body)
	fmt.Println(body)

	var res = new(Response)
	_ = json.Unmarshal(body, &res)

	data := &Response{}

	hex := strings.TrimPrefix(res.Result, "0x")
	x, _ := strconv.ParseInt(hex, 16, 64)
	s := strconv.Itoa(int(x))

	data.Result = s
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
	//fmt.Fprintf(w, res.Result)
}


func getTotalSupply(w http.ResponseWriter, r *http.Request) {
	rpc_cli, _ := ethclient.Dial("https://rinkeby.infura.io/dUQe3kE7aGgdnkBmEAny")

	model.Connect = rpc_cli
	fmt.Println("connect", model.Connect)

	auth, _ := generate_auth()
	fmt.Println("auth",auth)

	session, _ := CreateDSTokenSession(auth)
	fmt.Println("session",session)

	total, _ := session.TotalSupply()
	fmt.Println("total",total)

	balance, _ := session.BalanceOf(common.HexToAddress("0xd0a6e6c54dbc68db5db3a091b171a77407ff7ccf"))
	fmt.Println("balance",balance)

	var res = new(Response)
	//_ = json.Unmarshal(total, &res)

	data := &Response{}

	hex := strings.TrimPrefix(res.Result, "0x")
	x, _ := strconv.ParseInt(hex, 16, 64)
	s := strconv.Itoa(int(x))

	data.Result = s
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
	//fmt.Fprintf(w, res.Result)
}

func CreateDSTokenSession(auth *bind.TransactOpts) (*token.PowerTokenSession, error) {

	var contract *token.PowerToken
	var err error

	if model.Simulation {
		contract, err = token.NewPowerToken(common.HexToAddress(model.TokenAddress), model.Simulator)
		if err != nil {
			return nil, err
		}
	} else {
		contract, err = token.NewPowerToken(common.HexToAddress("0x035483Bc81b0982a07966522510AA6Ff761dE848"), model.Connect)
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

func generate_auth() (*bind.TransactOpts, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	return bind.NewKeyedTransactor(key), err
}


func main() {

	router := &Routes{} // set router
	err := http.ListenAndServe(":9090", router) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
