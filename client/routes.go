package client

import (
	"github.com/gorilla/mux"

	"github.com/osiz-blockchainapp/bitcoiva-sdk/client/context"
)

// Register routes
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	RegisterRPCRoutes(cliCtx, r)
}
