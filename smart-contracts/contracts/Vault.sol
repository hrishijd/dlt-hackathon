// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import {SafeERC20, IERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/**
 * @title Vault Contract
 * @dev A contract for managing deposits, supplying tokens, and monitoring balances.
 */
contract Vault {
    struct DepositInfo {
        address token;
        uint256 totalAmount;
        uint256 supplyAmount;
        uint256 minAmount;
        uint256 splitTime;
        uint256 lastSupplyTime;
        address receiver;
    }

    address internal constant owner = 0x50370f3BFE032Af808ac5e605cA6dEd64cbfB6E5;

    mapping(address => DepositInfo) public deposit;

    event TokensSupplied(address indexed from, DepositInfo depositInfo);

    event TokensDeposited(address indexed from, DepositInfo depositInfo);

    modifier onlyOwner() {
        require(msg.sender == owner, "Only the owner can call this function");
        _;
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
    function addDepositInfo(
        address token,
        uint256 totalAmount,
        uint256 supplyAmount,
        uint256 minAmount,
        uint256 splitTime,
        uint256 lastSupplyTime,
        address receiver
    ) external payable {
        require(deposit[msg.sender].totalAmount == 0, "Old balance should be removed before depositing new token");
        if (token == address(0)) {
            require(msg.value == totalAmount, "Value should be same as amount mentioned");
        } else {
            SafeERC20.safeTransferFrom(IERC20(token), msg.sender, address(this), totalAmount);
        }
        deposit[msg.sender] = DepositInfo(token, totalAmount, supplyAmount, minAmount, splitTime, lastSupplyTime, receiver);
        emit TokensDeposited(msg.sender, deposit[msg.sender]);
    }

    /**
     * @dev Refills tokens in the contract.
     * @param amount Amount of tokens to refill.
     */
    function refillTokens(uint256 amount) external payable {
        DepositInfo memory depositInfo = deposit[msg.sender];
        require(depositInfo.receiver != address(0), "Deposit info must be registered before refilling");
        require(depositInfo.totalAmount == 0, "Old balance should be removed before depositing new token");
        if (depositInfo.token == address(0)) {
            require(msg.value == amount, "Value should be same as amount mentioned");
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
     * @param supplier Address of the account to supply tokens to.
     * @param timestamp Timestamp of the supply.
     */
    function supplyTokens(address supplier, uint256 timestamp) external {
        DepositInfo memory depositInfo = deposit[supplier];
        require(msg.sender == owner, "Sender should be owner");
        require(depositInfo.totalAmount >= depositInfo.supplyAmount, "Not enough balance");
        depositInfo.totalAmount -= depositInfo.supplyAmount;
        depositInfo.lastSupplyTime = timestamp;
        deposit[supplier] = depositInfo;
        _supply(depositInfo.token, depositInfo.receiver, depositInfo.supplyAmount);
        emit TokensSupplied(supplier, depositInfo);
    }

    /**
     * @dev Withdraws the balance of tokens or Ether to a specified address.
     * @param to Receiver of the tokens or Ether.
     */
    function withrawBalance(address to) external {
        DepositInfo memory depositInfo = deposit[msg.sender];
        require(depositInfo.totalAmount > 0, "Not enough balance");
        uint256 amount = depositInfo.totalAmount;
        depositInfo.totalAmount = 0;
        delete deposit[msg.sender];
        _supply(depositInfo.token, to, amount);
        emit TokensSupplied(msg.sender, depositInfo);
    }
}
