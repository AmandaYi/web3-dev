// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract LotteryContract {
    address public manager;
    address[] public players;
    uint256 public round;
    address public winner;

    constructor() {
        manager = msg.sender;
    }

    // 每个人可以投多次，但是每次只能投1ether
    function play() public payable {
        require(msg.value == 1 ether);
        players.push(msg.sender);
    }

    // 开奖函数：
    // 目标：从彩民池（数组）中找到一个随机彩民（找一个随机数）
    // 找到一个特别大的数（随机）， 对我们的彩民数组长度求余数。
    // 用哈数值来实现大的随机数。 v3
    // 哈希内容的随机：当前时间，区块的挖矿难度，彩民数量，作为输入
    function open() isRootAccount public  {
        bytes memory v1 = abi.encodePacked(
            block.timestamp,
            block.prevrandao,
            players.length
        );
        bytes32 v2 = keccak256(v1);
        uint256 v3 = uint256(v2);
        uint256 winnerIndex = v3 % players.length;
        winner = players[winnerIndex];

        // 给人家转钱
        uint256 giveWinnerMoney = address(this).balance * 90 / 100 ;
        uint256 giveManagerMoney = address(this).balance - giveWinnerMoney;

        payable(winner).transfer(giveWinnerMoney);
        payable(manager).transfer(giveManagerMoney);
        round++;

        // 清空本次的参与人数
        delete players; // players = address[];
    }

    // 取消本次开奖，退款函数
    function cancel() isRootAccount public  {
        for (uint256 i = 0; i < players.length; i++) {
               address player = players[i];
               payable(player).transfer(1 ether);
        }
        // 期数加一
        round++;
        delete players;
    }
    modifier isRootAccount() {
        require(msg.sender == manager);
        _;
    }

    // 获取彩民人数
    function getPlayersCount () public view returns (uint256){
        return uint256(players.length);
    }

    // 获取余额
    function getBalance() public view returns (uint256) {
        return address(this).balance;
    }
    // 获取所有彩民的数据
    function getAllPlayers() public view returns (address[] memory) {
        return players;
    }
}
