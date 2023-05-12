// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract DreamContract {
    // 项目发起人
    address creator;
    // 众筹的项目名称
    string dreamName;
    // 众筹项目目标筹集金额
    uint256 targetAmount;
    // 每次众筹时，限制的金额数值
    uint256 limitSupportAmount;
    // 众筹截止日期，到此时间如果需要的金额不足，那么众筹失败，单位是秒
    uint256 endTime;

    // 初始化
    constructor(
        string memory _dreamName,
        uint256 _targetAmount,
        uint256 _limitSupportAmount,
        uint256 _sumTime
    ) {
        dreamName = _dreamName;
        targetAmount = _targetAmount;
        limitSupportAmount = _limitSupportAmount;
        endTime = block.timestamp + _sumTime;
    }
}
