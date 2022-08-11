# Client

## Query

### Params

```sh
pocket_core query bridgepool-params
```

### Signers

#### To check the signers of the Bridge

```sh
pocket_core query signers
```

### Liquidities

#### To check the liquidity amount added by the liquidity providers

```sh
pocket_core query liquidities
```

### Fee rates

#### To check the current bridge-fee for each token

```sh
pocket_core query fee-rates
```

### Allowed targets

#### To check all the allowed targets such as Target Blockchain, Target Wallet, Target Token

```sh
pocket_core query allowed-targets
```

## Messages

### MsgSetFee

#### To set the bridge fee for a particular token 

```sh
pocket_core bridgepool set-fee <fromAddr> <token> <fee-rate> <fee> <chainId>
pocket_core bridgepool set-fee 169869f67cd3f78a722fb4795b69949fb4bc9084 upokt 100 10000 testnet
```

### MsgAllowTarget

#### To set the allowed targets pairs on source/target blockchains
```sh
pocket_core bridgepool allow-target <fromAddr> <token> <targetChainId> <targetToken> <fee> <chainId>
pocket_core bridgepool allow-target 169869f67cd3f78a722fb4795b69949fb4bc9084 upokt 137 0xa5e0829caced8ffdd4de3c43696c57f7d7a678ff 10000 testnet
```

### MsgDisallowTarget

#### To uset the allowed targets pairs on source/target blockchains

```sh
pocket_core bridgepool disallow-target <fromAddr> <token> <targetChainId> <targetToken> <fee> <chainId>
pocket_core bridgepool disallow-target 169869f67cd3f78a722fb4795b69949fb4bc9084 upokt 137 0xa5e0829caced8ffdd4de3c43696c57f7d7a678ff 10000 testnet
```

### MsgAddLiquidity

#### To set the liquidity of a particular token (by a liquidity provider) on the bridge

```sh
pocket_core bridgepool add-liquidity <fromAddr> <token> <amount> <fee> <chainId>
pocket_core bridgepool add-liquidity 169869f67cd3f78a722fb4795b69949fb4bc9084 upokt 11111111 10000 testnet
```

### MsgRemoveLiquidity

#### To remove the liquidity of a particular token (by a liquidity provider) on the bridge

```sh
pocket_core bridgepool remove-liquidity <fromAddr> <token> <amount> <fee> <chainId>
pocket_core bridgepool remove-liquidity 169869f67cd3f78a722fb4795b69949fb4bc9084 upokt 11011 10000 testnet
```

### MsgAddSigner

#### To add the signer of the bridge for signing transactions

```sh
pocket_core bridgepool add-signer <fromAddr> <signer> <fee> <chainId>
pocket_core bridgepool add-signer 169869f67cd3f78a722fb4795b69949fb4bc9084 169869f67cd3f78a722fb4795b69949fb4bc9084 10000 testnet
```

### MsgRemoveSigner

#### To remove the signer of the bridge

```sh
pocket_core bridgepool remove-signer <fromAddr> <signer> <fee> <chainId>
pocket_core bridgepool remove-signer 169869f67cd3f78a722fb4795b69949fb4bc9084 169869f67cd3f78a722fb4795b69949fb4bc9084 10000 testnet
```

### MsgSwap

#### To perform a swap from POKT to Target Blockchain 
```sh
pocket_core bridgepool swap <fromAddr> <token> <amount> <targetNetwork> <targetToken> <targetAddress> <fee> <chainId>
pocket_core bridgepool swap 169869f67cd3f78a722fb4795b69949fb4bc9084 upokt 10000000 polygon 0xa5e0829caced8ffdd4de3c43696c57f7d7a678ff 169869f67cd3f78a722fb4795b69949fb4bc9084 10000 testnet
```

### MsgWithdrawSigned

#### To perform withdrawal of tokens received on POKT Network after a signature verification process to ensure secure withdrawal

```sh
pocket_core bridgepool withdraw-signed <fromAddr> <token> <payee> <amount> <fee> <chainId>
pocket_core bridgepool withdraw-signed 169869f67cd3f78a722fb4795b69949fb4bc9084 upokt 169869f67cd3f78a722fb4795b69949fb4bc9084 30000000upokt 10000 testnet
```
