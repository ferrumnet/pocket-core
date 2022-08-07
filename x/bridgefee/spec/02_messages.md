# Messages

## MsgSetTokenInfo

`MsgSetTokenInfo` is used to set token fee buffer size by module admin.

```protobuf
message MsgSetTokenInfo {
  bytes fromAddress = 1 [
    (gogoproto.jsontag) = "fromAddress",
    (gogoproto.casttype) = "github.com/pokt-network/pocket-core/types.Address"
  ];
  TokenInfo info = 2 [ (gogoproto.nullable) = false ];
}
```

## MsgSetTokenTargetInfos

`MsgSetTokenTargetInfos` is used to set fee distribution target per token by module admin.

```protobuf
message MsgSetTokenTargetInfos {
  bytes fromAddress = 1 [
    (gogoproto.jsontag) = "fromAddress",
    (gogoproto.casttype) = "github.com/pokt-network/pocket-core/types.Address"
  ];
  string token = 2;
  repeated TargetInfo targets = 3 [ (gogoproto.nullable) = false ];
}
```

## MsgSetGlobalTargetInfos

`MsgSetTokenTargetInfos` is used to set fee distribution target for any token when it does not have specific target.
Module admin has access to execute the message.

```protobuf
message MsgSetGlobalTargetInfos {
  bytes fromAddress = 1 [
    (gogoproto.jsontag) = "fromAddress",
    (gogoproto.casttype) = "github.com/pokt-network/pocket-core/types.Address"
  ];
  repeated TargetInfo targets = 2 [ (gogoproto.nullable) = false ];
}
```
