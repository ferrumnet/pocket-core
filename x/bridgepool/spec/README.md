# Bridge Pool

Bridge pool is where users can add/remove liquidity, send/withdraw tokens from other networks.

## Administration

- manage fee rate per token
- manage allow/disallow target token sending

## Liquidity

### Add liquidity

Transfer tokens from end users' wallet into the pool module to be used as liquidity to send/receive.

### Remove liquidity

Transfer tokens from pool module to end user's wallet

## Send

- send tokens from users' wallet to pool account
- emit `BridgeSwap` event for sending tokens on the target network

## Withdraw

- Ensure that the user has offchain signature that the signer on the other chain received the funds
- Send withdrawal fees to fee pool
- Distribute fees
- Transfer tokens to end user's wallet after minting

## Fees

- Fees are collected when users withdraw tokens from other network to pocket network
- Fees distribution mechanism could be discussed
  - Who are eligible for fees distribution
  - How much fees to be spent to each address
  - How often fees will be distributed (per tx, per block, per day, per week, per month)

## Questions

- What tokens will be supported? (Pocket network token + other BSC network tokens?)
- Who will be minting the tokens on pocket network for the amount that is locked on BSC network?
- Who will be the minters of wPOKT for the amount that is locked on bridge pool?
- Why liquidity is needed? (Is it just the amount where admin can move between chains?)
