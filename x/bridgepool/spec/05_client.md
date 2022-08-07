# Client

## Query

### Params

```sh
pocket_core query bridgepool-params
```

### Signers

```sh
pocket_core query signers
```

### Liquidities

```sh
pocket_core query liquidities
```

### Fee rates

```sh
pocket_core query fee-rates
```

### Allowed targets

```sh
pocket_core query allowed-targets
```

## Messages

### MsgSetFee

```sh
pocket_core bridgepool set-fee <fromAddr> <token> <fee-rate> <fee> <chainId>
pocket_core bridgepool set-fee 169869f67cd3f78a722fb4795b69949fb4bc9084 upokt 100 10000 testnet
```

### MsgAllowTarget

```sh
pocket_core bridgepool allow-target <fromAddr> <token> <targetChainId> <targetToken> <fee> <chainId>
pocket_core bridgepool allow-target 169869f67cd3f78a722fb4795b69949fb4bc9084 upokt 137 0xa5e0829caced8ffdd4de3c43696c57f7d7a678ff 10000 testnet
```

### MsgDisallowTarget

```sh
pocket_core bridgepool disallow-target <fromAddr> <token> <targetChainId> <targetToken> <fee> <chainId>
pocket_core bridgepool disallow-target 169869f67cd3f78a722fb4795b69949fb4bc9084 upokt 137 0xa5e0829caced8ffdd4de3c43696c57f7d7a678ff 10000 testnet
```

### MsgAddLiquidity

```sh
pocket_core bridgepool add-liquidity <fromAddr> <token> <amount> <fee> <chainId>
pocket_core bridgepool add-liquidity 169869f67cd3f78a722fb4795b69949fb4bc9084 upokt 11111111 10000 testnet
```

### MsgRemoveLiquidity

```sh
pocket_core bridgepool remove-liquidity <fromAddr> <token> <amount> <fee> <chainId>
pocket_core bridgepool remove-liquidity 169869f67cd3f78a722fb4795b69949fb4bc9084 upokt 11011 10000 testnet
```

### MsgAddSigner

```sh
pocket_core bridgepool add-signer <fromAddr> <signer> <fee> <chainId>
pocket_core bridgepool add-signer 169869f67cd3f78a722fb4795b69949fb4bc9084 169869f67cd3f78a722fb4795b69949fb4bc9084 10000 testnet
```

### MsgRemoveSigner

```sh
pocket_core bridgepool remove-signer <fromAddr> <signer> <fee> <chainId>
pocket_core bridgepool remove-signer 169869f67cd3f78a722fb4795b69949fb4bc9084 169869f67cd3f78a722fb4795b69949fb4bc9084 10000 testnet
```

### MsgSwap

```sh
pocket_core bridgepool swap <fromAddr> <token> <amount> <targetNetwork> <targetToken> <targetAddress> <fee> <chainId>
pocket_core bridgepool swap 169869f67cd3f78a722fb4795b69949fb4bc9084 upokt 10000000 polygon 0xa5e0829caced8ffdd4de3c43696c57f7d7a678ff 169869f67cd3f78a722fb4795b69949fb4bc9084 10000 testnet
```

### MsgWithdrawSigned

```sh
pocket_core bridgepool withdraw-signed <fromAddr> <token> <payee> <amount> <fee> <chainId>
pocket_core bridgepool withdraw-signed 169869f67cd3f78a722fb4795b69949fb4bc9084 upokt 169869f67cd3f78a722fb4795b69949fb4bc9084 30000000upokt 10000 testnet
```
