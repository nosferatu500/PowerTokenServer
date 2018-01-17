package model

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
)

var (
	Secret           = "hdPQ23!-_!!&"
	CrowdsaleAddress string
	WrapperAddress   string
	TokenAddress     string
	KeystoreReal     = "keystore-real"
	Connect          *ethclient.Client
	Simulator        *backends.SimulatedBackend
	Simulation       bool
	Gaslimit         uint64
	Port string
)

const (
	MAIN = iota + 1
	ROPSTEN
	RINKEBY
	KOVAN
	INFURANET
)

type BasicStringResp struct {
	Result string `json:"result"`
}

type BasicBoolResp struct {
	Result bool `json:"result"`
}

type BasicAddressResp struct {
	Address common.Address `json:"address"`
}







type User struct {
	TransactOpts *bind.TransactOpts
	PrivateKey   string
}

type LoginJsonReq struct {
	KeyJson string `json:"key_json"`
	Pass    string `json:"pass"`
}

type LoginPrivateKeyReq struct {
	PrivateKey string `json:"private_key"`
}

type CrowdSaleDeployReq struct {
	TokenAddress string `json:"token_address"`
}

type SetTotalSupplyReq_Token struct {
	Address string `json:"address"`
}

type TokenTransferOwnershipReq_Token struct {
	Address string `json:"address"`
}

type MintReq_Token struct {
	Address string `json:"address"`
	Tokens  string `json:"tokens"`
}

type SetFreezeForReq_Token struct {
	Time    string `json:"time"`
	Address string `json:"address"`
	Tokens  string `json:"tokens"`
}

type WithdrowTokensReq_Token struct {
	Address string `json:"address"`
	Tokens  string `json:"tokens"`
}

type SetFreezeReq_Token struct {
	Address string `json:"address"`
}

type RemoveFreezeReq_Token struct {
	Address string `json:"address"`
}

type MoveUnsoldReq_Token struct {
	Address string `json:"address"`
}

type NewTransferManualTokensnewTransferReq_Token struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Value string `json:"value"`
}

type IsFreezeReq_Token struct {
	Address string `json:"address"`
	Value   string `json:"value"`
}

type TransferReq_Token struct {
	To    string `json:"to"`
	Value string `json:"value"`
}

type NewTransferReq_Token struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Value string `json:"value"`
}

type BalanceOfReq_Token struct {
	Address string `json:"address"`
}

type TransferFromReq_Token struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Value string `json:"value"`
}

type ApproveReq_Token struct {
	Spender string `json:"spender"`
	Value   string `json:"value"`
}

type AllowanceReq_Token struct {
	Owner   string `json:"owner"`
	Spender string `json:"spender"`
}

type TotalSupplyResp struct {
	Result string `json:"result"`
}

type AllowanceResp_Token struct {
	Remaining string `json:"remaining"`
}

type BalanceOfResp_Token struct {
	Result string `json:"result"`
}

type GetOwnerTokenResp_Token struct {
	Address string `json:"address"`
}

type TransactionResp struct {
	Transaction string `json:"transaction"`
	GasUsed     string `json:"gas_used"`
	Status      uint   `json:"status"`
}

type ErrorResp struct {
	Error string `json:"error"`
}

type LoginResp struct {
	Access_token string `json:"access_token"`
	Exp          int64  `json:"exp"`
}

type DeployContractResp struct {
	Address     string `json:"address"`
	Transaction string `json:"transaction"`
	GasUsed     string `json:"gas_used"`
	Status      uint   `json:"status"`

}

//Crowdsale
type SendToAddressReq_Crowdsale struct {
	Address string `json:"address"`
	Tokens  string `json:"tokens"`
	Type    string `json:"type"`
}

type NewTransferManualTokensnewTransferReq_Crowdsale struct {
	FromAddress string `json:"from_address"`
	ToAddress   string `json:"to_address"`
	Value       string `json:"value"`
}

type AddrReq_Crowdsale struct {
	Address string `json:"address"`
}

type TimeReq_Crowdsale struct {
	Time string `json:"time"`
}

type ValueReq_Crowdsale struct {
	Value string `json:"value"`
}

type GetInvestorsTokensReq_Crowdsale struct {
	Address string `json:"address"`
	Type    string `json:"type"`
}

type GetInvestorsCountReq_Crowdsale struct {
	Type string `json:"type"`
}

type GetInvestorByIndexReq_Crowdsale struct {
	Index string `json:"index"`
	Type  string `json:"type"`
}

type GetInvestorByIndexResp_Crowdsale struct {
	Address string `json:"address"`
}

type GetWrapperDataResp_Crowdsale struct {
	Data []byte `json:"data"`
}

type Uint256Resp_Crowdsale struct {
	Uint256 string `json:"uint_256"`
}

type GetInvestorsTokensResp_Crowdsale struct {
	Uint256 string `json:"uint_256"`
}

type CalculateRateResp_Crowdsale struct {
	Rate string `json:"rate"`
}

type AddrResp_Crowdsale struct {
	Address string `json:"address"`
}

type TokenResp_Crowdsale struct {
	Address string `json:"address"`
}

type ETHUSDResp_Crowdsale struct {
	ETHUSD string `json:"ethusd"`
}

type GetFinishStatusResp_Crowdsale struct {
	Status bool `json:"status"`
}

type GetBalanceContractResp_Crowdsale struct {
	Balance string `json:"balance"`
}

type GetStartDatesResp_Crowdsale struct {
	Angel_sale_start  string `json:"angel_sale_start"`
	Pre_sale_start    string `json:"pre_sale_start"`
	Public_sale_start string `json:"public_sale_start"`
}

type GetFinishDatesResp_Crowdsale struct {
	Angel_sale_finish  string `json:"angel_sale_finish"`
	Pre_sale_finish    string `json:"pre_sale_finish"`
	Public_sale_finish string `json:"public_sale_finish"`
}

type GetSoldTokenResp_Crowdsale struct {
	SoldTokens       string `json:"sold_tokens"`
	Angel_sale_sold  string `json:"angel_sale_sold"`
	Pre_sale_sold    string `json:"pre_sale_sold"`
	Public_sale_sold string `json:"public_sale_sold"`
	Founding_sold    string `json:"founding_sold"`
	PeInvestors_sold string `json:"pe_investors_sold"`
}

type GetLeftTokenResp_Crowdsale struct {
	All_left              string `json:"all_left"`
	Founding_left         string `json:"founding_left"`
	Angel_left            string `json:"angel_left"`
	PreSaleAmount_left    string `json:"pre_sale_amount_left"`
	PEInvestorAmount_left string `json:"pe_investor_amount_left"`
	PublicSaleAmount_left string `json:"public_sale_amount_left"`
}

type GetTotalTokenResp_Crowdsale struct {
	TotalToken       string `json:"total_token"`
	FoundingAmount   string `json:"founding_amount"`
	AngelAmount      string `json:"angel_amount"`
	PreSaleAmount    string `json:"pre_sale_amount"`
	PEInvestorAmount string `json:"pe_investor_amount"`
	PublicSaleAmount string `json:"public_sale_amount"`
}

type GetTotalETHResp_Crowdsale struct {
	TotalETH             string `json:"total_eth"`
	Angel_sale_totalETH  string `json:"angel_sale_total_eth"`
	Pre_sale_totalETH    string `json:"pre_sale_total_eth"`
	Public_sale_totalETH string `json:"public_sale_total_eth"`
}

type SendTransactionReq struct {
	ToAddress string `json:"to_address"`
	Value     string `json:"value"`
}
