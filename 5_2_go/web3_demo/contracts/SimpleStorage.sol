// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.19;

contract SimpleStorage {
    string  public  v;
    constructor(string memory initV) {
        v = initV;
    }
    function getV() public view  returns(string memory) {
        return v;
    }
    function setV(string memory _str) public {
        string memory tmpStr = _str;
        v = tmpStr;
    }
}