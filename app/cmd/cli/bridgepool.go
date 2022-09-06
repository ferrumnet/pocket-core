package cli

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pokt-network/pocket-core/app"
	sdk "github.com/pokt-network/pocket-core/types"
	bridgepoolTypes "github.com/pokt-network/pocket-core/x/bridgepool/types"
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
	bridgepoolCmd.AddCommand(withdrawSignedSignatureCmd)
}

var bridgepoolCmd = &cobra.Command{
	Use:   "bridgepool",
	Short: "bridgepool management",
	Long:  `The bridgepool namespace handles all bridgepool related interactions.`,
}

func init() {
	withdrawSignedCmd.Flags().StringVar(&pwd, "pwd", "", "passphrase used by the cmd, non empty usage bypass interactive prompt")
}

var setFeeCmd = &cobra.Command{
	Use:   "set-fee <fromAddr> <token> <fee-rate> <fee> <chainId>",
	Short: "Set fee for withdrawing specific token",
	Long: `Set fee for withdrawing specific token.
Will prompt the user for the <fromAddr> account passphrase.`,
	Args: cobra.ExactArgs(5),
	Run: func(cmd *cobra.Command, args []string) {
		app.InitConfig(datadir, tmNode, persistentPeers, seeds, remoteCLIURL)
		fee, err := strconv.Atoi(args[3])
		if err != nil {
			fmt.Println(err)
			return
		}
		feeRate, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Enter Password: ")
		res, err := SetFee(args[0], args[1], uint64(feeRate), app.Credentials(pwd), args[4], int64(fee))
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
		fmt.Println("Enter Password: ")
		res, err := AllowTarget(args[0], args[1], args[2], args[3], app.Credentials(pwd), args[5], int64(fee))
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

		fmt.Println("Enter Password: ")
		res, err := DisallowTarget(args[0], args[1], args[2], app.Credentials(pwd), args[5], int64(fee))
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
		fee, err := strconv.Atoi(args[3])
		if err != nil {
			fmt.Println(err)
			return
		}
		amount, ok := sdk.NewIntFromString(args[2])
		if !ok {
			fmt.Println(args[2])
			return
		}
		fmt.Println("Enter Password: ")
		res, err := AddLiquidity(args[0], args[1], amount, app.Credentials(pwd), args[4], int64(fee))
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
		fee, err := strconv.Atoi(args[3])
		if err != nil {
			fmt.Println(err)
			return
		}
		amount, ok := sdk.NewIntFromString(args[2])
		if !ok {
			fmt.Println(args[2])
			return
		}
		fmt.Println("Enter Password: ")
		res, err := RemoveLiquidity(args[0], args[1], amount, app.Credentials(pwd), args[4], int64(fee))
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
	Use:   "swap <fromAddr> <amount> <targetNetwork> <targetToken> <targetAddress> <fee> <chainId>",
	Short: "Swap token to another network",
	Long: `Swap token to another network.
Will prompt the user for the <fromAddr> account passphrase.`,
	Args: cobra.ExactArgs(8),
	Run: func(cmd *cobra.Command, args []string) {
		app.InitConfig(datadir, tmNode, persistentPeers, seeds, remoteCLIURL)
		fee, err := strconv.Atoi(args[5])
		if err != nil {
			fmt.Println(err)
			return
		}

		amount, err := sdk.ParseCoin(args[1])
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Enter Password: ")
		res, err := Swap(args[0], amount, args[2], args[3], args[4], app.Credentials(pwd), args[6], int64(fee))
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

var withdrawSignedSignatureCmd = &cobra.Command{
	Use:   "withdraw-signed-signature <chainId> <payee> <amount> <salt> <ethPrivateKey>",
	Short: "Get signature of withdraw signed message from signer private key",
	Long: `Get signature of withdraw signed message from signer private key.
	pocket_core bridgepool withdraw-signed-signature testnet 169869f67cd3f78a722fb4795b69949fb4bc9084 10000upokt "" fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19`,
	Args: cobra.ExactArgs(5),
	Run: func(cmd *cobra.Command, args []string) {
		amount, err := sdk.ParseCoin(args[2])
		if err != nil {
			fmt.Println(err)
			return
		}

		message := &bridgepoolTypes.WithdrawSignMessage{
			ChainId: args[0],
			Payee:   args[1],
			Amount:  amount,
			Salt:    args[3],
		}
		messageBytes, err := json.Marshal(message)
		if err != nil {
			fmt.Println(err)
			return
		}

		hash := accounts.TextHash(messageBytes)
		fmt.Println("signMessage:", string(messageBytes))
		fmt.Println("signHash:", hex.EncodeToString(hash))

		ecdsaKey, err := crypto.HexToECDSA(args[4])
		if err != nil {
			fmt.Println(err)
			return
		}
		signature, err := crypto.Sign(hash, ecdsaKey)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Transform yellow paper V from 0/1 to 27/28
		signature[crypto.RecoveryIDOffset] += 27

		signatureStr := hex.EncodeToString(signature)
		fmt.Println("signature:", signatureStr)
	},
}

var withdrawSignedCmd = &cobra.Command{
	Use:   "withdraw-signed <fromAddr> <payee> <amount> <salt> <signature> <fee> <chainId>",
	Short: "Withdraw via signed message from other network signer",
	Long: `Withdraw via signed message from other network signer.
Will prompt the user for the <fromAddr> account passphrase.`,
	Args: cobra.ExactArgs(7),
	Run: func(cmd *cobra.Command, args []string) {
		app.InitConfig(datadir, tmNode, persistentPeers, seeds, remoteCLIURL)
		fee, err := strconv.Atoi(args[5])
		if err != nil {
			fmt.Println(err)
			return
		}
		amount, err := sdk.ParseCoin(args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Enter Password: ")
		res, err := WithdrawSigned(args[0], args[1], amount, args[3], args[4], app.Credentials(pwd), args[6], int64(fee))
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
