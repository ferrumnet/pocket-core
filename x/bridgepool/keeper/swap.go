package keeper

// Implements swap tokens on pocket network to other network

// function swap(address token, uint256 amount, uint256 targetNetwork, address targetToken)
// external returns(uint256) {
//     return _swap(msg.sender, token, amount, targetNetwork, targetToken, msg.sender);
// }

// function swapToAddress(address token,
//     uint256 amount,
//     uint256 targetNetwork,
//     address targetToken,
//     address targetAddress)
// external returns(uint256) {
//     require(targetAddress != address(0), "BridgePool: targetAddress is required");
//     return _swap(msg.sender, token, amount, targetNetwork, targetToken, targetAddress);
// }

// function _swap(address from, address token, uint256 amount, uint256 targetNetwork,
//     address targetToken, address targetAddress) internal returns(uint256) {
// 	require(targetNetwork != 0, "BP: targetNetwork is requried");
// 	require(allowedTargets[token][targetNetwork] == targetToken, "BP: target not allowed");
//     amount = SafeAmount.safeTransferFrom(token, from, address(this), amount);
//     emit BridgeSwap(from, token, targetNetwork, targetToken, targetAddress, amount);
//     return amount;
// }
