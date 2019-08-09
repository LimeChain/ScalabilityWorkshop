pragma solidity ^0.5.0;

contract DemoToken {

  struct Item {
    bytes32 name;
    bytes image;
  }

  Item[] public items;

  mapping(uint256 => address) public ownerOf;

  constructor(bytes32 name, bytes memory image) public {
    Item memory i = Item({
      name: name,
      image: image
    });

    items.push(i);

    ownerOf[0] = msg.sender;
  }
  
}