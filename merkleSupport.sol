//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Counters.sol";
import "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";

contract MerkleSupport is Ownable {
    using Counters for Counters.Counter;
    bytes32 public addressbyte;
    int Done ; 

    Counters.Counter public addressCnt ; 
    Counters.Counter public rootDepth ; 

    mapping(uint256 => bytes32) rootHierarcy;
    event NewRoot(address indexed NewAddress,bytes32 indexed NewRoot);
    constructor() {
    }

    function addToList(address newAddress)onlyOwner public{

        bytes32 addressHash = keccak256(toBytes(newAddress));
        bytes32 lastRoot = addressHash ;

        uint256 numberDivision = addressCnt.current();
        
        for(uint256 i=0 ; i <= rootDepth.current() ;i++){
            if(numberDivision%2 == 0){
                rootHierarcy[i] = lastRoot;
                lastRoot = hashPair(lastRoot,lastRoot);
                numberDivision = numberDivision/2 ;

            }else {

                lastRoot = hashPair(rootHierarcy[i],lastRoot);
                numberDivision = (numberDivision)/2 ; 
                if(numberDivision%2 ==0 ){
                    rootHierarcy[i+1] = lastRoot;
                }
            }
        }
        addressCnt.increment();
        if(addressCnt.current() > 2**rootDepth.current() ){
            rootDepth.increment();

        }
        emit NewRoot(newAddress,rootHierarcy[rootDepth.current()]);
        
    }
    function getRoot() public view returns(bytes32){
        return rootHierarcy[rootDepth.current()];
    }

    function getReward(bytes32[] memory proof) public {
        require(MerkleProof.verify(proof, rootHierarcy[rootDepth.current()], keccak256(toBytes(msg.sender))),"Don't Have Access");
        Done++;//some action, for simplicity I just increamented the Done.
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
    function hashPair(bytes32 a, bytes32 b) private pure returns (bytes32) {
        return a < b ? efficientHash(a, b) : efficientHash(b, a);
    }
    function efficientHash(bytes32 a, bytes32 b) private pure returns (bytes32 value) {
        /// @solidity memory-safe-assembly
        assembly {
            mstore(0x00, a)
            mstore(0x20, b)
            value := keccak256(0x00, 0x40)
        }
    }
}


