pragma solidity ^0.8.15 ;

contract LightSwitch {

    bool private lightSwitch = false;

    

    function getSwitch() public view returns (bool) {
        return lightSwitch ;
    }

    function flipSwitch() public returns(bool){
        lightSwitch = !lightSwitch ;
        return lightSwitch ;
    }

}