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
	bridgepoolCmd.AddCommand(setFeeCmd)
	bridgepoolCmd.AddCommand(allowTargetCmd)
	bridgepoolCmd.AddCommand(disallowTargetCmd)
	bridgepoolCmd.AddCommand(addLiquidityCmd)
	bridgepoolCmd.AddCommand(removeLiquidityCmd)
	bridgepoolCmd.AddCommand(swapCmd)
	bridgepoolCmd.AddCommand(withdrawSignedCmd)
	bridgepoolCmd.AddCommand(addSignerCmd)
	bridgepoolCmd.AddCommand(removeSignerCmd)
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

var setFeeCmd = &cobra.Command{
	Use:   "set-fee <fromAddr> <token> <fee> <fee> <chainId>",
	Short: "Set fee for withdrawing specific token",
	Long: `Set fee for withdrawing specific token.
Will prompt the user for the <fromAddr> account passphrase.`,
	Args: cobra.ExactArgs(5),
	Run: func(cmd *cobra.Command, args []string) {
		app.InitConfig(datadir, tmNode, persistentPeers, seeds, remoteCLIURL)
		fee, err := strconv.Atoi(args[4])
		if err != nil {
			fmt.Println(err)
			return
		}
		targetChainId, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Enter Password: ")
		res, err := SetFee(args[0], args[1], uint64(targetChainId), app.Credentials(pwd), args[5], int64(fee))
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

var allowTargetCmd = &cobra.Command{
	Use:   "allow-target <fromAddr> <token> <targetChainId> <targetToken> <fee> <chainId>",
	Short: "Allow target network for withdrawal",
	Long: `Allow target network for withdrawal.
Will prompt the user for the <fromAddr> account passphrase.`,
	Args: cobra.ExactArgs(6),
	Run: func(cmd *cobra.Command, args []string) {
		app.InitConfig(datadir, tmNode, persistentPeers, seeds, remoteCLIURL)
		fee, err := strconv.Atoi(args[4])
		if err != nil {
			fmt.Println(err)
			return
		}
		targetChainId, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Enter Password: ")
		res, err := AllowTarget(args[0], args[1], uint64(targetChainId), args[3], app.Credentials(pwd), args[5], int64(fee))
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

var disallowTargetCmd = &cobra.Command{
	Use:   "disallow-target <fromAddr> <token> <targetChainId> <targetToken> <fee> <chainId>",
	Short: "Disallow target network from withdrawal",
	Long: `Disallow target network from withdrawal.
Will prompt the user for the <fromAddr> account passphrase.`,
	Args: cobra.ExactArgs(6),
	Run: func(cmd *cobra.Command, args []string) {
		app.InitConfig(datadir, tmNode, persistentPeers, seeds, remoteCLIURL)
		fee, err := strconv.Atoi(args[4])
		if err != nil {
			fmt.Println(err)
			return
		}
		targetChainId, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Enter Password: ")
		res, err := DisallowTarget(args[0], args[1], uint64(targetChainId), app.Credentials(pwd), args[5], int64(fee))
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

var addLiquidityCmd = &cobra.Command{
	Use:   "add-liquidity <fromAddr> <token> <amount> <fee> <chainId>",
	Short: "Add liquidity to another network",
	Long: `Add liquidity to another network.
Will prompt the user for the <fromAddr> account passphrase.`,
	Args: cobra.ExactArgs(5),
	Run: func(cmd *cobra.Command, args []string) {
		app.InitConfig(datadir, tmNode, persistentPeers, seeds, remoteCLIURL)
		fee, err := strconv.Atoi(args[4])
		if err != nil {
			fmt.Println(err)
			return
		}
		amount, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Enter Password: ")
		res, err := AddLiquidity(args[0], args[1], uint64(amount), app.Credentials(pwd), args[5], int64(fee))
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

var removeLiquidityCmd = &cobra.Command{
	Use:   "remove-liquidity <fromAddr> <token> <amount> <fee> <chainId>",
	Short: "Remove liquidity to another network",
	Long: `Remove liquidity to another network.
Will prompt the user for the <fromAddr> account passphrase.`,
	Args: cobra.ExactArgs(5),
	Run: func(cmd *cobra.Command, args []string) {
		app.InitConfig(datadir, tmNode, persistentPeers, seeds, remoteCLIURL)
		fee, err := strconv.Atoi(args[4])
		if err != nil {
			fmt.Println(err)
			return
		}
		amount, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Enter Password: ")
		res, err := RemoveLiquidity(args[0], args[1], uint64(amount), app.Credentials(pwd), args[5], int64(fee))
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

var swapCmd = &cobra.Command{
	Use:   "swap <fromAddr> <token> <amount> <targetNetwork> <targetToken> <targetAddress> <fee> <chainId>",
	Short: "Swap token to another network",
	Long: `Swap token to another network.
Will prompt the user for the <fromAddr> account passphrase.`,
	Args: cobra.ExactArgs(8),
	Run: func(cmd *cobra.Command, args []string) {
		app.InitConfig(datadir, tmNode, persistentPeers, seeds, remoteCLIURL)
		fee, err := strconv.Atoi(args[4])
		if err != nil {
			fmt.Println(err)
			return
		}
		amount, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Enter Password: ")
		res, err := Swap(args[0], args[1], uint64(amount), args[3], args[4], args[5], app.Credentials(pwd), args[5], int64(fee))
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

var withdrawSignedCmd = &cobra.Command{
	Use:   "withdraw-signed <fromAddr> <token> <payee> <amount> <fee> <chainId>",
	Short: "Withdraw via signed message from other network signer",
	Long: `Withdraw via signed message from other network signer.
Will prompt the user for the <fromAddr> account passphrase.`,
	Args: cobra.ExactArgs(6),
	Run: func(cmd *cobra.Command, args []string) {
		app.InitConfig(datadir, tmNode, persistentPeers, seeds, remoteCLIURL)
		fee, err := strconv.Atoi(args[4])
		if err != nil {
			fmt.Println(err)
			return
		}
		amount, err := strconv.Atoi(args[3])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Enter Password: ")
		res, err := WithdrawSigned(args[0], args[1], args[2], uint64(amount), app.Credentials(pwd), args[5], int64(fee))
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

var addSignerCmd = &cobra.Command{
	Use:   "add-signer <fromAddr> <signer> <fee> <chainId>",
	Short: "Add signer for the bridge",
	Long: `Add signer for the bridge.
Will prompt the user for the <fromAddr> account passphrase.`,
	Args: cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		app.InitConfig(datadir, tmNode, persistentPeers, seeds, remoteCLIURL)
		fee, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Enter Password: ")
		res, err := AddSigner(args[0], args[1], app.Credentials(pwd), args[3], int64(fee))
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

var removeSignerCmd = &cobra.Command{
	Use:   "remove-signer <fromAddr> <signer> <fee> <chainId>",
	Short: "Remove signer for the bridge",
	Long: `Remove signer for the bridge.
Will prompt the user for the <fromAddr> account passphrase.`,
	Args: cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		app.InitConfig(datadir, tmNode, persistentPeers, seeds, remoteCLIURL)
		fee, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Enter Password: ")
		res, err := RemoveSigner(args[0], args[1], app.Credentials(pwd), args[3], int64(fee))
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
