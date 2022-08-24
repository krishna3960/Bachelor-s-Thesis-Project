package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strconv"
	"time"
	"window/api"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ev3go/ev3dev"
)

func main() {
	client, err := ethclient.Dial("http://192.168.1.3:7545")
	if err != nil {
		fmt.Println(err, "Error for dial")
	} else {
		fmt.Println("Connected to Private Blockchain", client)
	}

	addresse := common.HexToAddress("0x4C2bAbb04080190e15419e2bBAee83c4a69134Ca")
	instance, error := api.NewApi(addresse, client)
	if error != nil {
		panic(error)
	}
	privatekey, err := crypto.HexToECDSA("541a2de9c1e483ed4c3529df954d98c0a0c127b9aa1631b879fb779413f41146")
	if err != nil {
		panic(err)
	}

	publicAddress := LinkToAccount(privatekey)
	fmt.Println("Account Address is : ", publicAddress)
	auth := Auth(privatekey, publicAddress, *client)
	fmt.Println("#########")
	fmt.Println("Ev3 is Ready")
	state := false

	for {
		// Import our sensors and motor
		in1, err := ev3dev.SensorFor("ev3-ports:in1", "lego-ev3-touch")
		if err != nil {
			panic(err)
		}
		in2, err := ev3dev.SensorFor("ev3-ports:in2", "lego-ev3-touch")
		if err != nil {
			panic(err)
		}
		outA, err := ev3dev.TachoMotorFor("ev3-ports:outA", "lego-ev3-l-motor")
		if err != nil {
			fmt.Println("failed to find motor on outA:", err)
		}

		intVal1, intVal2 := ButtonValues(in1, in2)

		//Listen for input on touch sensor 1
		if intVal1 == 1 && !state {
			fmt.Println("Window shade going up")
			//Set motorstate to true on api
			auth = Auth(privatekey, publicAddress, *client)
			_, err := instance.SwitchMotorState(auth)
			time.Sleep(200 * time.Millisecond)
			state, _ = instance.GetMotorState(nil)
			if err != nil {
				panic(err)
			}
			counter, err := instance.GetOpenPercentage(nil)
			time.Sleep(100 * time.Millisecond)
			ret := MotorUp(outA, *counter, in1, in2, *instance, privatekey, publicAddress, *client)
			time.Sleep(200 * time.Millisecond)
			state = ret
			time.Sleep(200 * time.Millisecond)
			fmt.Println("Listening for new inputs")

		}

		if intVal2 == 1 && !state {
			fmt.Println("Window shade going down")
			//Set motorstate to true on api
			auth = Auth(privatekey, publicAddress, *client)
			_, err := instance.SwitchMotorState(auth)
			time.Sleep(200 * time.Millisecond)
			state, _ = instance.GetMotorState(nil)
			if err != nil {
				panic(err)
			}
			counter, err := instance.GetOpenPercentage(nil)
			time.Sleep(100 * time.Millisecond)
			ret := MotorDown(outA, *counter, in1, in2, *instance, privatekey, publicAddress, *client)
			time.Sleep(200 * time.Millisecond)
			state = ret
			time.Sleep(200 * time.Millisecond)
			fmt.Println("Listening for new inputs")

		}

	}
}

func ButtonValues(button1 *ev3dev.Sensor, button2 *ev3dev.Sensor) (int, int) {
	//Read the inputs of the touch sensors
	strVal, _ := button1.Value(0)
	intVal, _ := strconv.Atoi(strVal)
	strVal2, _ := button2.Value(0)
	intVal2, _ := strconv.Atoi(strVal2)
	return intVal, intVal2
}

func LinkToAccount(privatekey *ecdsa.PrivateKey) common.Address {
	publicKey := privatekey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic(ok)
	}

	publicAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return publicAddress

}

func Auth(privatekey_ *ecdsa.PrivateKey, publicAddress_ common.Address, client ethclient.Client) *bind.TransactOpts {
	nonce, err := client.PendingNonceAt(context.Background(), publicAddress_)
	if err != nil {
		panic(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}

	auth, _ := bind.NewKeyedTransactorWithChainID(privatekey_, big.NewInt(5777))
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice

	return auth
}

func MotorUp(motor *ev3dev.TachoMotor, counter big.Int, button1 *ev3dev.Sensor, button2 *ev3dev.Sensor, instance api.Api, privatekey_ *ecdsa.PrivateKey, publicAddress_ common.Address, client ethclient.Client) bool {
	Speed := motor.MaxSpeed()
	motor.SetSpeedSetpoint(Speed)
	motor.Command("run-forever")
	percentage := counter.Int64() * 10
	for i := percentage; i < 1000; i++ {
		intVal1, intVal2 := ButtonValues(button1, button2)
		if intVal1 == 1 || intVal2 == 1 {
			break

		}
		percentage++
	}
	motor.Command("stop")
	auth := Auth(privatekey_, publicAddress_, client)
	_, err := instance.Update(auth, uint64(percentage/10))
	if err != nil {
		panic(err)
	}
	time.Sleep(200 * time.Millisecond)
	tx, err := instance.GetOpenPercentage(nil)
	time.Sleep(200 * time.Millisecond)
	auth = Auth(privatekey_, publicAddress_, client)
	_, err = instance.SwitchMotorState(auth)
	time.Sleep(200 * time.Millisecond)
	state, _ := instance.GetMotorState(nil)
	time.Sleep(200 * time.Millisecond)

	fmt.Println("Window updated with value ", tx, " expected value is : ", percentage/10)

	return state
}

func MotorDown(motor *ev3dev.TachoMotor, counter big.Int, button1 *ev3dev.Sensor, button2 *ev3dev.Sensor, instance api.Api, privatekey_ *ecdsa.PrivateKey, publicAddress_ common.Address, client ethclient.Client) bool {
	Speed := motor.MaxSpeed()
	motor.SetSpeedSetpoint(Speed)
	motor.Command("run-forever")
	percentage := counter.Int64() * 10
	for i := percentage; i > 0; i-- {
		intVal1, intVal2 := ButtonValues(button1, button2)
		if intVal1 == 1 || intVal2 == 1 {
			break

		}
		percentage--
	}
	motor.Command("stop")
	auth := Auth(privatekey_, publicAddress_, client)
	_, err := instance.Update(auth, uint64(percentage/10))
	if err != nil {
		panic(err)
	}
	time.Sleep(200 * time.Millisecond)
	tx, err := instance.GetOpenPercentage(nil)
	time.Sleep(200 * time.Millisecond)
	auth = Auth(privatekey_, publicAddress_, client)
	_, err = instance.SwitchMotorState(auth)
	time.Sleep(200 * time.Millisecond)
	state, _ := instance.GetMotorState(nil)
	time.Sleep(200 * time.Millisecond)

	fmt.Println("Window updated with value ", tx, " expected value is : ", percentage/10)

	return state
}
