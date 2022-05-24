package keeper

// function addLiquidity(address token, uint256 amount) external {
//     require(amount != 0, "Amount must be positive");
//     require(token != address(0), "Bad token");
//     amount = SafeAmount.safeTransferFrom(token, msg.sender, address(this), amount);
//     liquidities[token][msg.sender] = liquidities[token][msg.sender].add(amount);
//     emit BridgeLiquidityAdded(msg.sender, token, amount);
// }

// function removeLiquidityIfPossible(address token, uint256 amount) external returns (uint256) {
//     require(amount != 0, "Amount must be positive");
//     require(token != address(0), "Bad token");
//     uint256 liq = liquidities[token][msg.sender];
//     require(liq >= amount, "Not enough liquidity");
//     uint256 balance = IERC20(token).balanceOf(address(this));
//     uint256 actualLiq = balance > amount ? amount : balance;
//     liquidities[token][msg.sender] = liquidities[token][msg.sender].sub(actualLiq);
//     if (actualLiq != 0) {
//         IERC20(token).safeTransfer(msg.sender, actualLiq);
//         emit BridgeLiquidityRemoved(msg.sender, token, amount);
//     }
//     return actualLiq;
// }

// function liquidity(address token, address liquidityAdder) public view returns (uint256) {
//     return liquidities[token][liquidityAdder];
// }
