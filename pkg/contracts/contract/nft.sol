// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/utils/Counters.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract GameItem is ERC721URIStorage, Ownable {
    using Counters for Counters.Counter;
    Counters.Counter private _tokenIds;
    using SafeMath for uint256;
    constructor(

    ) ERC721("GameItem", "ITM") {
    }
    event BatchMintItem(address owner,string []  urls,uint256 [] _ids,uint256 [] tokenIds);

    function awardItem(address player, string memory tokenURI)
    public onlyOwner
    returns (uint256)
    {
        _tokenIds.increment();

        uint256 newItemId = _tokenIds.current().add(50000);
        _mint(player, newItemId);
        _setTokenURI(newItemId, tokenURI);

        return newItemId;
    }

    function awardItemBatch( address owner, string [] memory urls, uint256 [] memory _ids)
    public onlyOwner
    {
        uint[] memory tokenIds = new uint[](urls.length);

        for (uint256 i = 0; i < urls.length; i++) {
            uint256 id = awardItem( owner, urls[i]);
            tokenIds[i] = (id);
        }
        emit BatchMintItem(  owner, urls, _ids,tokenIds);
    }

}