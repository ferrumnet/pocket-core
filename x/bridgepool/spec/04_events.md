# Events

## MsgSetFee

| Type         | Attribute Key | Attribute Value |
| ------------ | ------------- | --------------- |
| set_fee_rate | token         | {token}         |
| set_fee_rate | fee           | {fee_rate}      |

## MsgAllowTarget

| Type         | Attribute Key | Attribute Value |
| ------------ | ------------- | --------------- |
| allow_target | token         | {token}         |
| allow_target | chain_id      | {chain_id}      |
| allow_target | target_token  | {target_token}  |

## MsgDisallowTarget

| Type            | Attribute Key | Attribute Value |
| --------------- | ------------- | --------------- |
| disallow_target | token         | {token}         |
| disallow_target | chain_id      | {chain_id}      |

## MsgAddLiquidity

| Type                   | Attribute Key | Attribute Value |
| ---------------------- | ------------- | --------------- |
| bridge_liquidity_added | actor         | {actor}         |
| bridge_liquidity_added | token         | {token}         |
| bridge_liquidity_added | amount        | {amount}        |

## MsgRemoveLiquidity

| Type                     | Attribute Key | Attribute Value |
| ------------------------ | ------------- | --------------- |
| bridge_liquidity_removed | actor         | {actor}         |
| bridge_liquidity_removed | token         | {token}         |
| bridge_liquidity_removed | amount        | {amount}        |

## MsgAddSigner

| Type       | Attribute Key | Attribute Value |
| ---------- | ------------- | --------------- |
| set_signer | signer        | {signer}        |

## MsgRemoveSigner

| Type          | Attribute Key | Attribute Value |
| ------------- | ------------- | --------------- |
| remove_signer | signer        | {signer}        |

## MsgSwap

| Type        | Attribute Key   | Attribute Value   |
| ----------- | --------------- | ----------------- |
| bridge_swap | from            | {from}            |
| bridge_swap | token           | {token}           |
| bridge_swap | target_chain_id | {target_chain_id} |
| bridge_swap | target_token    | {target_token}    |
| bridge_swap | target_address  | {target_address}  |
| bridge_swap | amount          | {amount}          |

## MsgWithdrawSigned

| Type                  | Attribute Key | Attribute Value |
| --------------------- | ------------- | --------------- |
| transfer_by_signature | signer        | {signer}        |
| transfer_by_signature | receiver      | {receiver}      |
| transfer_by_signature | token         | {token}         |
| transfer_by_signature | amount        | {amount}        |
| transfer_by_signature | fee           | {fee}           |
