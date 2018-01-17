//go:generate abigen --sol ./contracts/SingleTokenCoin.sol --pkg token  --out ./gocontracts/token/token.go

package main

import (
	"net/http"
	"PowerTokenServer/controllers"
	"github.com/labstack/gommon/log"
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

func (p *Routes)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" { controllers.SayhelloName2(w, r); return }

	if r.URL.Path == "/test" { controllers.SayhelloName2(w, r);	return }

	if r.URL.Path == "/token/totalSupply" {	controllers.GetTotalSupply(w, r); return }

	if r.URL.Path == "/token/symbol" { controllers.GetSymbol(w, r);	return }

	if r.URL.Path == "/token/paused" { controllers.GetPausedStatus(w, r);	return }
/*
	if r.URL.Path == "/token/owner" { getOwner(w, r);	return }

	if r.URL.Path == "/token/name" { getName(w, r);	return }

	if r.URL.Path == "/token/mintingFinished" { getMintingFinished(w, r);	return }

	//  address _wallet
	if r.URL.Path == "/token/isFrozen" { getIsFrozen(w, r);	return }

	if r.URL.Path == "/token/decimals" { getDecimals(w, r);	return }

	// address _wallet
	if r.URL.Path == "/token/balanceOf" { getBalanceOf(w, r);	return }

	// address  _owner; address _spender
	if r.URL.Path == "/token/allowance" { getAllowance(w, r);	return }
*/

	http.NotFound(w, r)
	return
}


func main() {

	router := &Routes{} // set router
	err := http.ListenAndServe(":9090", router) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
