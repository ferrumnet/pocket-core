# State

## FeeRate

`FeeRate` is used to manage fee rate per token. `rate` is put in basis points and the value of 100 is equivalent to 1%.

```protobuf
message FeeRate {
  string token = 1;
  uint64 rate = 2;
}
```

## AllowedTarget

`AllowedTarget` is used to manage target network token addresses to swap current network token to.
The object exists for 1 token and 1 chain id pair.

```protobuf
message AllowedTarget {
  string token = 1;
  string chainId = 2;
  string targetToken = 3;
}
```

## Liquidity

`Liquidity` stores the record of token liquidity put by specific address. The `amount` field is increased when user add liquidity and decreased when user remove liquidity.
The object exists for 1 address and 1 token pair.

```protobuf
message Liquidity {
  string token = 1;
  string address = 2;
  uint64 amount = 3;
}
```
