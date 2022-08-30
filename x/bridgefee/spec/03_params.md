# Params

`Params` hold the on-chain parameters to manage bridgefee module.
At this phase, it holds `owner` field that manages administrative transactions of the module.
`Params` fields are administrated by network delegators' governance.

```protobuf
message Params {
  string owner = 1;
}
```
