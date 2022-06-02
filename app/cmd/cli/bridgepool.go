package cli

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/pokt-network/pocket-core/app"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(bridgepoolCmd)
	bridgepoolCmd.AddCommand(withdrawSignedCmd)
}

var bridgepoolCmd = &cobra.Command{
	Use:   "bridgepool",
	Short: "node management",
	Long: `The node namespace handles all node related interactions,
from staking and unstaking; to unjailing.`,
}

func init() {
	withdrawSignedCmd.Flags().StringVar(&pwd, "pwd", "", "passphrase used by the cmd, non empty usage bypass interactive prompt")
}

// TODO: MsgAddSigner
// TODO: MsgRemoveSigner
// TODO: MsgSetFee
// TODO: MsgAllowTarget
// TODO: MsgDisallowTarget
// TODO: MsgAddLiquidity
// TODO: MsgRemoveLiquidity
// TODO: MsgSwap
// TODO: MsgWithdrawSigned

var withdrawSignedCmd = &cobra.Command{
	Use:   "unstake <operatorAddr> <fromAddr> <networkID> <fee> <isBefore8.0>",
	Short: "Unstake a node in the network",
	Long: `Unstake a node from the network, changing it's status to Unstaking.
Will prompt the user for the <fromAddr> account passphrase.`,
	Args: cobra.ExactArgs(5),
	Run: func(cmd *cobra.Command, args []string) {
		app.InitConfig(datadir, tmNode, persistentPeers, seeds, remoteCLIURL)
		fee, err := strconv.Atoi(args[3])
		if err != nil {
			fmt.Println(err)
			return
		}
		isBefore8, err := strconv.ParseBool(args[4])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Enter Password: ")
		res, err := UnstakeNode(args[0], args[1], app.Credentials(pwd), args[2], int64(fee), isBefore8)
		if err != nil {
			fmt.Println(err)
			return
		}
		j, err := json.Marshal(res)
		if err != nil {
			fmt.Println(err)
			return
		}
		resp, err := QueryRPC(SendRawTxPath, j)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(resp)
	},
}
