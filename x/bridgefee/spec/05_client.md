# Client

## Query

### Params

```sh
pocket_core query bridgefee-params
```

### All token infos

```sh
pocket_core query all-token-infos
```

### All token target infos

```sh
pocket_core query all-token-target-infos
```

## Messages

### MsgSetTokenInfo

```sh
pocket_core bridgefee set-token-info <fromAddr> <token> <bufferSize> <fee> <chainId>
pocket_core bridgefee set-token-info 169869f67cd3f78a722fb4795b69949fb4bc9084 upokt 1000 10000 testnet
```

### MsgSetTokenTargetInfos

```sh
set-token-target-infos <fromAddr> <token> <targets> <weights> <targetTypes> <fee> <chainId>
pocket_core bridgefee set-token-target-infos 169869f67cd3f78a722fb4795b69949fb4bc9084 upokt 924364bbf0347842e1e0fa9cd2167dc630fc3c0c,a9cb5f27e7a4b2be3045439f151d3e68108cd65c 1,1 Address,Address 10000 testnet
```

### MsgSetGlobalTargetInfos

```sh
pocket_core bridgefee set-global-target-infos <fromAddr> <targets> <weights> <targetTypes> <fee> <chainId>
pocket_core bridgefee set-global-target-infos 169869f67cd3f78a722fb4795b69949fb4bc9084 924364bbf0347842e1e0fa9cd2167dc630fc3c0c 1 Burn 10000 testnet
```
