const solc = require("solc") // 跟solidity版本要一致
const fs = require("fs")

const sourceLotteryContract = fs.readFileSync("./contracts/LotteryContract.sol", "utf-8")


let input = {
    language: "Solidity",
    sources: {
        "LotteryContract": {
            content: sourceLotteryContract
        }
    },
    settings: {
        outputSelection: {
          '*': {
            '*': ['*']
          }
        }
      }
}

const output = JSON.parse(solc.compile(JSON.stringify(input)))

module.exports = {
    LotteryContract: output["contracts"]["LotteryContract"]["LotteryContract"]
}