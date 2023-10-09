import { expect } from "chai";
import { ethers } from "hardhat";
import { IWETH, Vault } from "../typechain-types";
import {
  time,
  loadFixture,
} from "@nomicfoundation/hardhat-network-helpers";
import { Signer } from "ethers";
  
  describe("Vault", function () {
    let owner, account0, account1: Signer;
    let WETH : IWETH;
    let amount: BigInt;
    let vault: Vault;
  
    beforeEach(async function (){
      [owner, account0, account1] = await ethers.getSigners();

      WETH = await ethers.getContractAt("IWETH", "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2");
      
      const VaultContractFactory = await ethers.getContractFactory("Vault");
      vault = await VaultContractFactory.deploy(owner.address);

      amount = ethers.parseUnits("5", "ether");
      const duration = 3600;

      await WETH.connect(account0).deposit({value: amount});
      await WETH.connect(account0).approve(await vault.getAddress(), amount);

      await vault.connect(account0).depositTokens(await WETH.getAddress(), 100, await time.latest(), amount, amount/10n, ethers.parseUnits("5", "ether"), account1.address, account0.address);
    });
    it("Should supply right amount of tokens", async function () {
      const balanceBefore = await WETH.balanceOf(account1.address);
      await vault.supplyTokens(account1.address);
      expect(await WETH.balanceOf(account1.address)).to.equal(amount/10n + balanceBefore);
    });
    it("Should not supply if called not called by owner", async function () {
      await expect(vault.connect(account0).supplyTokens(account1.address)).to.be.revertedWith("Ownable: caller is not the owner");
    });
    it("Supplier should be able to withdraw tokens", async function () {
      const balanceBefore = await WETH.balanceOf(account0.address);
      await vault.connect(account0).withrawBalance(account1.address, account0.address);
      expect(await WETH.balanceOf(account0.address)).to.equal(balanceBefore + amount);
    });
    it("Supplier not be able to withdraw tokens if not called by supplier", async function () {
      await expect(vault.connect(account1).withrawBalance(account1.address, account0.address)).to.be.revertedWith("VAMOS_VAULT: Can only be withdrawn by supplier");
    });
    it("Should be able to refill tokens", async function () {
      await WETH.connect(account0).deposit({value: amount});
      await WETH.connect(account0).approve(await vault.getAddress(), amount);
      await vault.connect(account0).refillTokens(account1.address, amount);
    });
  });
