package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Establish RPC Connection
	client, err := ethclient.Dial("http://192.168.1.3:8545")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to Private Blockchain", client)
	}

	//Test if RPC Connection is working
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Header is ", header.Number.String())

}
