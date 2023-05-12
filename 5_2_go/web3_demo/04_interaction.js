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
const instance = require("./03_instance")

instance.then(async v => {
    let result = await v.methods.getV().call()
    console.log(result)
    let account = await getAccount()
    result = await v.methods.setV("我是新的值").send({
        from: account
    })
    result = await v.methods.getV().call()




    const evResult = v.methods.setV("我是新的值22").send({
        from: account
    })
    evResult.on('transactionHash', function (hash) {
        console.log("hash", hash)
    });
    evResult.on('receipt', function (receipt) {
        console.log("receipt", receipt)
    });
    evResult.on('confirmation', function (confirmationNumber, receipt) {
        console.log("confirmationNumber", confirmationNumber)
    });
    evResult.on('error', function (error, receipt) {
        console.log("error", error)
    });
})