const Web3 = require("web3")
const web3 = new Web3()


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


web3.setProvider(getProvider())


// abi 取自 构建后的json
const SimpleStorageJSON =  require("./constracts/artifacts/SimpleStorage.json")
const abi = SimpleStorageJSON.abi
// 用用部署的账号逻辑
async  function init(){
    let account =  await getAccount()
    const contractInstance = new web3.eth.Contract(abi, "0xa43F41dDA1b9075F0a51B6b9dD015a391FB00aA7")
    return contractInstance
}

module.exports = init()

