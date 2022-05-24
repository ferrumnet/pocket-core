package keeper

// Implements withdrawing tokens into pocket network keeper and utility functions

// function withdrawSigned(
//         address token,
//         address payee,
//         uint256 amount,
//         bytes32 salt,
//         bytes memory signature)
// external returns(uint256) {
//     bytes32 message = withdrawSignedMessage(token, payee, amount, salt);
//     address _signer = signerUnique(message, signature);
//     require(signers[_signer], "BridgePool: Invalid signer");

//     uint256 fee = 0;
//     address _feeDistributor = feeDistributor;
//     if (_feeDistributor != address(0)) {
//         fee = amount.mul(fees[token]).div(10000);
//         amount = amount.sub(fee);
//         if (fee != 0) {
//             IERC20(token).safeTransfer(_feeDistributor, fee);
//             IGeneralTaxDistributor(_feeDistributor).distributeTax(token, msg.sender);
//         }
//     }
//     IERC20(token).safeTransfer(payee, amount);
//     emit TransferBySignature(_signer, payee, token, amount, fee);
//     return amount;
// }

// function withdrawSignedVerify(
//         address token,
//         address payee,
//         uint256 amount,
//         bytes32 salt,
//         bytes calldata signature)
// external view returns (bytes32, address) {
//     bytes32 message = withdrawSignedMessage(token, payee, amount, salt);
//     (bytes32 digest, address _signer) = signer(message, signature);
//     return (digest, _signer);
// }

// function withdrawSignedMessage(
//         address token,
//         address payee,
//         uint256 amount,
//         bytes32 salt)
// internal pure returns (bytes32) {
//     return keccak256(abi.encode(
//       WITHDRAW_SIGNED_METHOD,
//       token,
//       payee,
//       amount,
//       salt
//     ));
// }
