const solc = require("solc") // 这个版本跟solidity一一对应

const fs = require("fs")

// 读取合约
let sourceCode = fs.readFileSync("./constracts/SimpleStorage.sol",'utf-8')

let input = {
    language: "Solidity",
    sources: {
        "SimpleStorage.sol": {
            content: sourceCode
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

const output = JSON.parse(solc.compile(JSON.stringify(input)));
// console.log(output["contracts"]["SimpleStorage.sol"]["SimpleStorage"]
// )
module.exports = output["contracts"]["SimpleStorage.sol"]["SimpleStorage"]

