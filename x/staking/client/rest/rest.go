package rest

import (
	"github.com/gorilla/mux"

	"github.com/osiz-blockchainapp/bitcoiva-sdk/client/context"
)

// RegisterRoutes registers staking-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	registerQueryRoutes(cliCtx, r)
	registerTxRoutes(cliCtx, r)
}
