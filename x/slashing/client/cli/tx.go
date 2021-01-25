package cli

import (
	"github.com/spf13/cobra"

	"github.com/osiz-blockchainapp/bitcoiva-sdk/client"
	"github.com/osiz-blockchainapp/bitcoiva-sdk/client/context"
	"github.com/osiz-blockchainapp/bitcoiva-sdk/codec"
	sdk "github.com/osiz-blockchainapp/bitcoiva-sdk/types"
	"github.com/osiz-blockchainapp/bitcoiva-sdk/x/auth"
	"github.com/osiz-blockchainapp/bitcoiva-sdk/x/auth/client/utils"
	"github.com/osiz-blockchainapp/bitcoiva-sdk/x/slashing/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	slashingTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Slashing transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	slashingTxCmd.AddCommand(client.PostCommands(
		GetCmdUnjail(cdc),
	)...)

	return slashingTxCmd
}

// GetCmdUnjail implements the create unjail validator command.
func GetCmdUnjail(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "unjail",
		Args:  cobra.NoArgs,
		Short: "unjail validator previously jailed for downtime",
		Long: `unjail a jailed validator:

$ <appcli> tx slashing unjail --from mykey
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			valAddr := cliCtx.GetFromAddress()

			msg := types.NewMsgUnjail(sdk.ValAddress(valAddr))
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
