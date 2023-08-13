import { ethers } from "hardhat";

async function main() {
    const vaultFactory = await ethers.getContractFactory("Vault");
    const signers = await ethers.getSigners();
    const vaultContract = await vaultFactory.connect(signers[0]).deploy();
    console.log(await vaultContract.getAddress());
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
