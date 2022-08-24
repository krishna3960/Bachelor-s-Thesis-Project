package main

import (
	"context"
	"contract/api"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ev3go/ev3dev"
)

type MotorUpdate func(bool)

func main() {
	// Establish RPC Connection
	client, err := ethclient.Dial("http://192.168.1.3:7545")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to Private Blockchain", client)
	}

	//Test if RPC Connection is working
	/* _, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		panic(err)
	} */
	//fmt.Println("Header is ", header.Number.String())

	// Contract Address
	addresse := common.HexToAddress("0xbA7c076b53C0dd40d478153616DB2c6C75089407")

	//New Api with our Smart Contract Functions
	instance, error := api.NewApi(addresse, client)
	if error != nil {
		panic(error)
	}

	// Connect to an account using private key
	privateKey, err := crypto.HexToECDSA("9713cd8bc107fb0458a728e45e8b4bf29ca2fdd15bfd21fa93bf8e71099e1769")
	if err != nil {
		panic(err)
	}

	publicAddress := LinkToAccount(privateKey)
	fmt.Println("Address of account is ", publicAddress)
	/* publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("Error casting public key to ecdsa")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	fmt.Println("######")

	fmt.Println("Address of account is ", fromAddress) */

	// Beginning of the loop
	for {

		//Getting the ev3 touch sensor
		in1, err := ev3dev.SensorFor("ev3-ports:in1", "lego-ev3-touch")
		if err != nil {
			panic(err)
		}

		// Check the value of the button, 1 = pressed, 0 = not pressed
		strVal, _ := in1.Value(0)
		intVal, _ := strconv.Atoi(strVal)

		// If pressed

		if intVal == 1 {

			auth := Auth(privateKey, publicAddress, *client)

			/* auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(5777))
			auth.Nonce = big.NewInt(int64(nonce))
			auth.Value = big.NewInt(0)
			auth.GasLimit = uint64(3000000)
			auth.GasPrice = gasPrice */

			fmt.Println("Auth completed")
			fmt.Println("Button Pressed")
			tx, err := instance.FlipSwitch(auth)
			if err != nil {
				panic(err)
			}
			fmt.Println("transaction sent :", tx.Hash().Hex())
			time.Sleep(3 * time.Second)
			result, err := instance.GetSwitch(nil)
			if err != nil {
				panic(err)
			}
			fmt.Println("Result is ", result)

		}
	}

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
	//fmt.Println("Pending nonce : ", nonce)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}
	//fmt.Println("Suggested Gas Price : ", gasPrice)

	auth, _ := bind.NewKeyedTransactorWithChainID(privatekey_, big.NewInt(5777))
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice

	return auth
}
