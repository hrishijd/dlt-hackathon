# Vault and Monitoring Service (Vamos)

## Objective

The **Vault and Monitoring Service (Vamos)** project aims to provide accounts (EOAs or smart contracts) with a constant supply of tokens or native cryptocurrency. This service helps maintain a minimum balance in accounts to ensure smooth and secure transaction processing and protect against various attacks.

## Problems Solved

- **Security:** Safeguard accounts from unidentified threats that may lead to excessive consumption of tokens or cryptocurrency. Maintaining a minimum balance helps ensure the stable operation of software systems.
- **Efficiency:** Reduce the time spent on development and gain a competitive advantage. Eliminate the need to install monitoring systems and set up alerts for checking the system's crypto wallet.
- **Flexibility:** If a bug is identified later, developers can update the software to decide whether to provide funds to contracts or not.

## Applications

1. **Automatic Funding:** Automatically provide funds to smart contracts to ensure their continuous operation without manual intervention.
2. **Bot Monitoring:** Utilize a monitoring bot to track and manage the funding process, enhancing the overall security and stability of accounts.

## Usage

Vamos is designed to assist developers in maintaining minimum balances for accounts, ensuring secure and smooth transactions. It provides flexibility, security, and efficiency by automating the funding process and offering monitoring features.

## Getting Started

To get started with Vamos:

1. Install the required dependencies.
2. Configure the accounts and contracts that require continuous funding.
3. Deploy the Vault contract to the desired network.
4. Implement the monitoring service to ensure consistent balance maintenance.

## Current Status

Vamos is ready as a prototype, and might require some enchancements before it si started to use for monetary purposes.

## Possible Improvements

- Monitoring of events using block number so that when bot is down, It could resume from the point it stopped functioning.
- Integrate prometheus for better visualization of state of DB.
- Utilize min value to process transaction only if contracts are behind min value.

## Vault Contract

### Getting Started

These instructions will help you set up and run the Vault smart contract project on your local machine.

### Prerequisites

- [Node.js](https://nodejs.org/) (>=12.0.0)

### Installation

1. Clone this repository:
    ```
    git clone https://github.com/hrishijd/dlt-hackathon.git
    ```
2. Change to the project directory:
    ```
    cd dlt-hackathon
    cd smart-contracts
    ```
3. Install the dependencies:
    ```
    npm install
    ```
4. Create .env for providing environment variables
    sample .env
    ```
    ARKHIA_KEY=<your arkhia key here>
    Private_KEY=<your account's private key here>
    ```

## Supply Service

### Watcher Service
Watcher service maintains a DB with supply info, along with their status, by watching events emmited by supply contract, it watchers over the following events:

`TokensSupplied(address indexed from, DepositInfo depositInfo)`: Event emmited when a tokens are supplied.

`TokensDeposited(address indexed from, DepositInfo depositInfo)`: Event emitted when a tokens are deposited in contract.

### Supply Bot
Bot that keeps a watch over all Deposits, maintains a priority queue for them, and works as a supplier for contracts in need of supply.

### Development Setup : Go
To set up the project for development, follow these steps:

Clone the repository to your local machine.

Create a config file from template in this folder, with required rpc and credentials
```
cd go-service
go mod tidy
go run .
```
Install the necessary dependencies using for install.

Dependencies
- Go 1.20
- Postgresql 14

