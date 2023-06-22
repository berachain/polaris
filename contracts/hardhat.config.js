/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.18",
  defaultNetwork: "localhost",
  networks: {
    hardhat: {},
    localhost: {
      url: "http://127.0.0.1:8545",
      accounts: ["0xfffdbb37105441e14b0ee6330d855d8504ff39e705c3afa8f859ac9865f99306"],
    },
  },
  paths: {
    sources: "./src",
    tests: "./src/testing",
    artifacts: "./out",
  },
};
