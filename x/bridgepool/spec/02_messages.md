# Messages

## MsgSetFee

`MsgSetFee` is the message to set token's fee rate by module admin.

```protobuf
message MsgSetFee {
  bytes fromAddress = 1 [
    (gogoproto.jsontag) = "fromAddress",
    (gogoproto.casttype) = "github.com/pokt-network/pocket-core/types.Address"
  ];
  string token = 2;
  uint64 fee10000 = 3;
}
```

## MsgAllowTarget

`MsgAllowTarget` is the message to allow token swap to target `chainId` and `targetToken` for `token`.

```protobuf
message MsgAllowTarget {
  bytes fromAddress = 1 [
    (gogoproto.jsontag) = "fromAddress",
    (gogoproto.casttype) = "github.com/pokt-network/pocket-core/types.Address"
  ];
  string token = 2;
  string chainId = 3;
  string targetToken = 4;
}
```

## MsgDisallowTarget

`MsgDisallowTarget` is the message to disallow token swap to target `chainId` for `token`.

```protobuf
message MsgDisallowTarget {
  bytes fromAddress = 1 [
    (gogoproto.jsontag) = "fromAddress",
    (gogoproto.casttype) = "github.com/pokt-network/pocket-core/types.Address"
  ];
  string token = 2;
  string chainId = 3;
}
```

## MsgAddLiquidity

`MsgAddLiquidity` is the message to add liquidity for the token by `amount`.

```protobuf
message MsgAddLiquidity {
  bytes fromAddress = 1 [
    (gogoproto.jsontag) = "fromAddress",
    (gogoproto.casttype) = "github.com/pokt-network/pocket-core/types.Address"
  ];
  string token = 2;
  uint64 amount = 3;
}
```

## MsgRemoveLiquidity

`MsgAddLiquidity` is the message to remove liquidity for the token by `amount`.

```protobuf
message MsgRemoveLiquidity {
  bytes fromAddress = 1 [
    (gogoproto.jsontag) = "fromAddress",
    (gogoproto.casttype) = "github.com/pokt-network/pocket-core/types.Address"
  ];
  string token = 2;
  uint64 amount = 3;
}
```

## MsgAddSigner

`MsgAddSigner` is the message to add signer by module admin.

```protobuf
message MsgAddSigner {
  bytes fromAddress = 1 [
    (gogoproto.jsontag) = "fromAddress",
    (gogoproto.casttype) = "github.com/pokt-network/pocket-core/types.Address"
  ];
  string signer = 2;
}
```

## MsgRemoveSigner

`MsgAddSigner` is the message to remove signer by module admin.

```protobuf
message MsgRemoveSigner {
  bytes fromAddress = 1 [
    (gogoproto.jsontag) = "fromAddress",
    (gogoproto.casttype) = "github.com/pokt-network/pocket-core/types.Address"
  ];
  string signer = 2;
}
```

## MsgSwap

`MsgSwap` is the message to swap tokens from current network to target network. Any user can execute it by putting tokens on current network.

```protobuf
message MsgSwap {
  option (gogoproto.messagename) = true;
  option (gogoproto.goproto_getters) = false;

  bytes fromAddress = 1 [
    (gogoproto.jsontag) = "fromAddress",
    (gogoproto.casttype) = "github.com/pokt-network/pocket-core/types.Address"
  ];
  string token = 2;
  uint64 amount = 3;
  string targetChainId = 4;
  string targetToken = 5;
  string targetAddress = 6;
}
```

## MsgWithdrawSigned

`WithdrawSignMessage` is the message to withdraw tokens from target network to current network. Signature of a signer for `WithdrawSignMessage` is needed to withdraw the funds. Fees are applied on withdrawal based on withdrawing token's fee rate.
Fee distribution is done on `bridgefee` module.

```protobuf
message MsgWithdrawSigned {
  option (gogoproto.messagename) = true;
  option (gogoproto.goproto_getters) = false;

  bytes fromAddress = 1 [
    (gogoproto.jsontag) = "fromAddress",
    (gogoproto.casttype) = "github.com/pokt-network/pocket-core/types.Address"
  ];
  string payee = 2;
  types.Coin amount = 3 [ (gogoproto.nullable) = false ];
  string salt = 4;
  bytes signature = 5;
}
message WithdrawSignMessage {
  string chainId = 1;
  string payee = 2;
  types.Coin amount = 3 [(gogoproto.nullable) = false];
  string salt = 4;
}
```
