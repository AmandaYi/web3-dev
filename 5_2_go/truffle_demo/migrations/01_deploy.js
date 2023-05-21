const myContract= truffle.require("Simple");
module.exports = function(deployer){
    deployer.deploy(myContract);
}