const Web3 = require("web3")
const web3 = new Web3();

web3.setProvider(window.web3.currentProvider)

module.exports = web3
