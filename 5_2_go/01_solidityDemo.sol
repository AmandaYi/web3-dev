pragma solidity ^0.4.24;

contract Inbox {
    string public message;

    function Inbox()payable {

    }
    function setMessage(string newMessage) public {

    }
    function getMessage() public constant returns(string) {
        return (message);
    }
}