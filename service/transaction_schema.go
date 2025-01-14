package service

import (
	"github.com/cawabunga/terra.go/types"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"
)

type QueryTxRequest struct {
	Page  *int64  `json:"page"`
	Limit *int64  `json:"limit"`
	Query types.Q `json:"query"`
}

type QueryTxResponse struct {
	TotalCount cosmostypes.Int          `json:"total_count"`
	Count      cosmostypes.Int          `json:"count"`
	PageNumber cosmostypes.Int          `json:"page_number"`
	PageTotal  cosmostypes.Int          `json:"page_total"`
	Limit      cosmostypes.Int          `json:"limit"`
	Txs        []cosmostypes.TxResponse `json:"txs"`
}
