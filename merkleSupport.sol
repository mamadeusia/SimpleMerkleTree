//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Counters.sol";
import "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";

contract MerkleSupport is Ownable {
    using Counters for Counters.Counter;
    bytes32 public root;
    Counters.Counter public cnt ; 
    bytes32 public addressbyte;
    constructor(bytes32 _root) {
        root = _root; 
    }

    function getReward(bytes32[] memory proof) public {

        require(MerkleProof.verify(proof, root, keccak256(toBytes(msg.sender))),"Don't Have Access");
        cnt.increment();
    }
    function getkeccak() view public returns(bytes32){
        return keccak256(toBytes(msg.sender));
    }

    function toBytes(address a) public pure returns (bytes memory b){
    assembly {
        let m := mload(0x40)
        a := and(a, 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF)
        mstore(add(m, 20), xor(0x140000000000000000000000000000000000000000, a))
        mstore(0x40, add(m, 52))
        b := m
   }
}
}
