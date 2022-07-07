package cli

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/pokt-network/pocket-core/app"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(bridgefeeCmd)
	bridgefeeCmd.AddCommand(setTokenInfo)
	bridgefeeCmd.AddCommand(setTokenTargetInfos)
	bridgefeeCmd.AddCommand(setGlobalTargetInfos)
}

var bridgefeeCmd = &cobra.Command{
	Use:   "bridgefee",
	Short: "bridgefee management",
	Long:  `The bridgefee namespace handles all bridgefee related interactions.`,
}

func init() {
	setTokenInfo.Flags().StringVar(&pwd, "pwd", "", "passphrase used by the cmd, non empty usage bypass interactive prompt")
}

var setTokenInfo = &cobra.Command{
	Use:   "set-token-info <fromAddr> <token> <bufferSize> <tokenSpecificConfig> <fee> <chainId>",
	Short: "Set token info for bridgefee",
	Long: `Set token info for bridgefee.
Will prompt the user for the <fromAddr> account passphrase.`,
	Args: cobra.ExactArgs(6),
	Run: func(cmd *cobra.Command, args []string) {
		app.InitConfig(datadir, tmNode, persistentPeers, seeds, remoteCLIURL)
		fee, err := strconv.Atoi(args[4])
		if err != nil {
			fmt.Println(err)
			return
		}
		bufferSize, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		tokenSpecificConfig, err := strconv.Atoi(args[3])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Enter Password: ")

		res, err := SetTokenInfo(args[0], args[1], uint64(bufferSize), uint32(tokenSpecificConfig), app.Credentials(pwd), args[5], int64(fee))
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

var setTokenTargetInfos = &cobra.Command{
	Use:   "set-token-target-infos <fromAddr> <token> <targets> <weights> <targetTypes> <fee> <chainId>",
	Short: "Allow target network for withdrawal",
	Long: `Allow target network for withdrawal.
Will prompt the user for the <fromAddr> account passphrase.`,
	Args: cobra.ExactArgs(7),
	Run: func(cmd *cobra.Command, args []string) {
		app.InitConfig(datadir, tmNode, persistentPeers, seeds, remoteCLIURL)
		fee, err := strconv.Atoi(args[5])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Enter Password: ")
		res, err := SetTokenTargetInfos(args[0], args[1], args[2], args[3], args[4], app.Credentials(pwd), args[6], int64(fee))
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

var setGlobalTargetInfos = &cobra.Command{
	Use:   "set-token-target-infos <fromAddr> <targets> <weights> <targetTypes> <fee> <chainId>",
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
		res, err := SetGlobalTargetInfos(args[0], args[1], args[2], args[3], app.Credentials(pwd), args[5], int64(fee))
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
