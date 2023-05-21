// SPDX-License-Identifier: MIT

pragma solidity >=0.8.0 <0.9.0;

contract SupportDreamMap {
    // 这个合约用来保存，每个在网络中的，在这个平台中，都参与了什么Dream众筹项目
    mapping(address => address[]) supportMapWithDreamAddress;

    function getValue(
        address personAddress
    ) public view returns (address[] memory) {
        return supportMapWithDreamAddress[personAddress];
    }

    function setValue(address personAddress, address dreamAddress) public {
        supportMapWithDreamAddress[personAddress].push(dreamAddress);
    }
}
