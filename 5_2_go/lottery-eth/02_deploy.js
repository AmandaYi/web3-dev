const { LotteryContract } = require("./01_compile")
const Web3 = require("web3")
const HDWalletProvider = require("@truffle/hdwallet-provider");
const config = require("./config/index")

const { evm, abi } = LotteryContract
const { bytecode, deployedBytecode } = evm
const web3 = new Web3()

let mnemonicPhrase = config.mnemonicPhrase
// let netIp = 'https://ropsten.infura.io/v3/02cd1e3c295c425597fa105999493baa'
let providerOrUrl = 'http://127.0.0.1:7545'

let provider = new HDWalletProvider({
    mnemonic: {
        phrase: mnemonicPhrase
    },
    providerOrUrl: providerOrUrl
})
web3.setProvider(provider);


// 获取合约

let ethLotteryContract = new web3.eth.Contract(abi)

async function init() {
    let accounts = await web3.eth.getAccounts()
    let instance = await ethLotteryContract.deploy({
        data: bytecode.object
    }).send({
        from: accounts[0],
        gas: "3000000",
        // gasPrice: '200000000'
    })
    return instance
}
init().then(v => {
    console.log("部署完成", v._address)
}).finally(() => {
    provider.engine.stop()
})