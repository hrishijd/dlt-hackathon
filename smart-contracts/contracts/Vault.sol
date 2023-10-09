// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;


import {SafeERC20, IERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";


/**
 * @title Vault Contract
 * @dev A contract for managing deposits, supplying tokens, and monitoring balances.
 */
contract Vault is Ownable{
    struct DepositInfo {
        address token;
        uint32 splitTime;
        uint32 lastSupplyTime;
        uint256 totalAmount;
        uint256 supplyAmount;
        uint256 minAmount;
        address supplier;
    }

    mapping(address => DepositInfo) public deposit;

    event TokensSupplied(address indexed receiver, DepositInfo depositInfo);

    event TokensDeposited(address indexed receiver, DepositInfo depositInfo);

    event TokensWithdrawn(address indexed receiver, DepositInfo depositInfo);

    constructor(address owner) Ownable(){
        _transferOwnership(owner);
    }

    /**
     * @dev Adds deposit information for a address.
     * @param token Address of the token (use address(0) for Ether).
     * @param totalAmount Total amount of tokens or Ether to deposit.
     * @param supplyAmount Amount of tokens to supply for funding.
     * @param minAmount Minimum balance required to maintain.
     * @param splitTime Time interval for splitting supply.
     * @param lastSupplyTime Timestamp of the last supply.
     * @param receiver Address to receive funds when supplying tokens.
     */
    function depositTokens(
        address token,
        uint32 splitTime,
        uint32 lastSupplyTime,
        uint256 totalAmount,
        uint256 supplyAmount,
        uint256 minAmount,
        address receiver,
        address supplier
    ) external payable {
        require(deposit[receiver].totalAmount == 0, "Old balance should be removed before depositing new token");
        require(receiver != address(0), "receiver can't be address zero");
        if(token == address(0)) {
            require(msg.value == totalAmount, "Value should be same as amount mentioned");
        } else {
            SafeERC20.safeTransferFrom(IERC20(token), msg.sender, address(this), totalAmount);
        }
        deposit[receiver] = DepositInfo(token, splitTime, lastSupplyTime, totalAmount, supplyAmount, minAmount, supplier);
        emit TokensDeposited(msg.sender, deposit[msg.sender]);
    }

    /**
     * @dev Refills tokens in the contract.
     * @param amount Amount of tokens to refill.
     */
    function refillTokens(address receiver, uint256 amount) external payable {
        DepositInfo memory depositInfo = deposit[receiver];
        require(receiver != address(0), "VAMOS_VAULT: Deposit info must be registered before refilling");
        if (depositInfo.token == address(0)) {
            require(msg.value == amount, "VAMOS_VAULT: Value should be same as amount mentioned");
        } else {
            SafeERC20.safeTransferFrom(IERC20(depositInfo.token), msg.sender, address(this), amount);
        }
        depositInfo.totalAmount += amount;
        deposit[msg.sender] = depositInfo;
        emit TokensDeposited(msg.sender, depositInfo);
    }

    /**
     * @dev Internal function to supply tokens to an account.
     * @param token Address of the token (use address(0) for Ether).
     * @param to Receiver of the tokens or Ether.
     * @param amount Amount of tokens or Ether to supply.
     */
    function _supply(address token, address to, uint256 amount) internal {
        if (token == address(0)) {
            (bool sent,) = to.call{value: amount}("");
            require(sent, "Failed to send Ether");
        } else {
            SafeERC20.safeTransfer(IERC20(token), to, amount);
        }
    }

    /**
     * @dev Supplies tokens to an address.
     * @param receiver Address of the account to supply tokens to.
     */
    function supplyTokens(address receiver) external onlyOwner{
        DepositInfo memory depositInfo = deposit[receiver];
        require(depositInfo.totalAmount >= depositInfo.supplyAmount, "VAMOS_VAULT: Not enough balance for this receiver");
        depositInfo.totalAmount -= depositInfo.supplyAmount;
        depositInfo.lastSupplyTime = uint32(block.timestamp);
        deposit[receiver] = depositInfo;
        _supply(depositInfo.token, receiver, depositInfo.supplyAmount);
        emit TokensSupplied(receiver, depositInfo);
    }

    /**
     * @dev Withdraws the balance of tokens or Ether to a specified address.
     * @param receiver Address that was registered to recieve tokens.
     * @param to Receiver of the tokens or Ether.
     */
    function withrawBalance(address receiver, address to) external {
        DepositInfo memory depositInfo = deposit[receiver];
        require(msg.sender == depositInfo.supplier, "VAMOS_VAULT: Can only be withdrawn by supplier");
        require(depositInfo.totalAmount > 0, "VAMOS_VAULT: Not enough balance");
        uint256 amount = depositInfo.totalAmount;
        depositInfo.totalAmount = 0;
        delete deposit[msg.sender];
        _supply(depositInfo.token, to, amount);
        emit TokensWithdrawn(msg.sender, depositInfo);
    }
}
