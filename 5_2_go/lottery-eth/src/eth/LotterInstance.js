const web3 = require("../utils/web3Instance")
const { LotterContractAbi } = require("./abi")
const dappAddress = "0x9Fa751E932f9A38B063AB916E9De11964A4581b9" // 这个地址一般是数据库存起来返回来的

const LotterContractInstance = new web3.eth.Contract(LotterContractAbi, dappAddress)

// 这是返回的Promise对象
module.exports = LotterContractInstance

