import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import {config as dotConfig} from "dotenv";
dotConfig();
const config: HardhatUserConfig = {
  solidity: "0.8.19",
  networks: {
    hederaTestnet:{
        url: `https://pool.arkhia.io/hedera/testnet/json-rpc/v1/${process.env.ARKHIA_KEY}`,
        accounts: [`${process.env.PRIVATE_KEY}`],
    },
    hederaMainnet:{
        url: `https://pool.arkhia.io/hedera/mainnet/json-rpc/v1/${process.env.ARKHIA_KEY}`,
        accounts: [`${process.env.PRIVATE_KEY}`]
    },
  },
};

export default config;
