// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity >=0.8.0 <0.9.0;

contract Simple {
  uint data;
  function get() public view returns (uint) {
    return data;
  }
  function set(uint _data) public  {
    data = _data;
  }
}