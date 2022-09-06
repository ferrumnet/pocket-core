# State

## TokenInfo

`TokenInfo` has `bufferSize` which is the threshold value to distribute token to targets.

```protobuf
message TokenInfo {
  string token = 1;
  uint64 bufferSize = 2;
}
```

## TargetInfo

`TargetInfo` is the minimum unit of distribution target.
It has `tType`, which describes how to handle the target `NotSet`, `Burn`, `Address`, `RewardPool`.
And weight allocation to the target.

`Address` is used to just receive the distribution and `Burn` is used to burn the token.
`NotSet` and `RewardPool` are not implemented.

```protobuf
message TargetInfo {
  string target = 1;
  TargetType tType = 2;
  uint64 weight = 3;
}
```

## TokenTargetInfo

`TokenTargetInfo` includes all the targets for a specific token.

```protobuf
message TokenTargetInfo {
  string token = 1;
  repeated TargetInfo targets = 2 [ (gogoproto.nullable) = false ];
}
```
