pragma solidity ^0.8.15 ;

contract WindowShade {

    bool private motorstate ;
    uint64 private openPercentage ;

    constructor()   {
        motorstate = false;
        openPercentage= 0 ;
    }

    function  getOpenPercentage() public view returns (uint){
        return openPercentage ;
    }

    function getMotorState() public view returns (bool){
        return motorstate;

    }

    function SwitchMotorState() public returns (bool){
        motorstate = !motorstate ;
        return motorstate ;
    }

    function Update(uint64 update) public returns(uint) {
        openPercentage = update;
        return openPercentage ;
    }


}