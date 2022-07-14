// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;
import "hardhat/console.sol";

abstract contract GameItemC {
    function awardItem(address , string memory ) public virtual returns (uint256);
}

contract maticBatch {
    event MintItem(address _token,address _dsts,string []  _values,uint256 [] _ids);
    function awardItemBatch(address _token, address _dsts, string [] memory _values, uint256 [] memory _ids)
    public
    {
        GameItemC token = GameItemC(_token);
        for (uint256 i = 0; i < _values.length; i++) {
            uint256 id =  token.awardItem( _dsts, _values[i]);
            console.log("maticBatch done");
        }
        emit MintItem( _token, _dsts, _values, _ids);
    }
}
