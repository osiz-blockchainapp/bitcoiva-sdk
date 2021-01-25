package client

import (
	"github.com/osiz-blockchainapp/bitcoiva-sdk/x/distribution/client/cli"
	"github.com/osiz-blockchainapp/bitcoiva-sdk/x/distribution/client/rest"
	govclient "github.com/osiz-blockchainapp/bitcoiva-sdk/x/gov/client"
)

// param change proposal handler
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
)
