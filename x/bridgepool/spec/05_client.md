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
pocket_core bridgepool withdraw-signed <fromAddr> <payee> <amount> <salt> <signature> <fee> <chainId>
pocket_core bridgepool withdraw-signed 169869f67cd3f78a722fb4795b69949fb4bc9084 169869f67cd3f78a722fb4795b69949fb4bc9084 30000000upokt "" cd2deaa3531dfcfbece6e088f28dfc650b6a57caf253b89fbfad4d3d65fa4bb5064f4272c78dc47352123d100681b43efb8f9533411ac70cbd2ebf9d77a0adfa01 10000 testnet
```

### WithdrawSigned signature

```sh
pocket_core bridgepool withdraw-signed-signature <chainId> <payee> <amount> <salt> <ethPrivateKey>

pocket_core bridgepool withdraw-signed-signature testnet 169869f67cd3f78a722fb4795b69949fb4bc9084 10000upokt "" fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19

# signMessage: {"chainId":"testnet","payee":"169869f67cd3f78a722fb4795b69949fb4bc9084","amount":{"denom":"upokt","amount":"10000"}}
# signHash: 0083cf7097c01510cc46cf032a0a6071f7862b74729f68187d7bde4c3a83e1e7
# signature: 9291b53c1ea7454e91667c21788965045c6581be28d69fb27a6926e68b0d2fb103664071ea136b50e3e8dd87ee59df32ce6819ea0ca7bb645ef51db6e2328e9a1c
```
