const { abi, evm } = require("./01_compile")
const Web3 = require("web3")

const web3 = new Web3()





web3.setProvider(getProvider())


let contract = new web3.eth.Contract(abi)

async function init() {
    let account = await getAccount()

    // //2. 执行部署
    let instance = await contract.deploy({
        data: evm.bytecode.object, //合约的bytecode
        arguments: ["部署的时候设置的构造器里面的值"]
    })
        .send({
            from: account,
            gas: "3000000",
            gasPrice: '20000000000'
        })
    return instance
}



function getProvider() {
    //     let terms = 'scout same naive genius cannon maze differ acquire penalty habit surround ice'
    // let netIp = 'https://ropsten.infura.io/v3/02cd1e3c295c425597fa105999493baa'

    // let provider = new HDWalletProvider(terms, netIp)
    let provider = "HTTP://127.0.0.1:7545"
    return provider
}


async function getAccount() {

    let account = await web3.eth.getAccounts()
    return account[0]
}

init()

module.exports = {
    getProvider,
    getAccount
}