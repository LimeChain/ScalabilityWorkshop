pragma solidity ^0.5.0;

import "@openzeppelin/contracts/cryptography/ECDSA.sol";

contract DemoToken {

  struct Item {
    bytes32 name;
    bytes image;
  }

  Item[] public items;

  mapping(uint256 => address) public ownerOf;
  mapping(uint256 => uint256) public nonceOf;

  constructor(bytes32 name, bytes memory image) public {
    Item memory i = Item({
      name: name,
      image: image
    });

    items.push(i);

    ownerOf[0] = msg.sender;
  }

  function safeTransfer(uint256 id, address receiver, uint256 price, bytes memory signature) public payable {
    require(id < items.length, "Invalid item");
    require(msg.value == price, "Not enough price sent");

    uint256 nonce = nonceOf[id];
    bytes32 sendingHash = keccak256(abi.encodePacked(id, receiver, price, nonce));
    address recoveredAddress = ECDSA.recover(ECDSA.toEthSignedMessageHash(sendingHash), signature);

    require(recoveredAddress == ownerOf[id], "Invalid message or message signature");

    ownerOf[id] = receiver;
    address payable wallet = address(uint160(address(recoveredAddress)));
    wallet.transfer(price);
    nonceOf[id]++;
  }
  
}