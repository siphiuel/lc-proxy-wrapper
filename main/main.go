package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	proxy "github.com/siphiuel/lc-proxy-wrapper"
)

func main() {
	var testConfig = proxy.Config{
		Eth2Network:      "mainnet",
		TrustedBlockRoot: "0x6fd2f9fdc616a0755f3f88ac58e4a7871788c3128261a18bf9645eee7042eb53",
		//Web3Url:          Web3UrlType{"HttpUrl", "https://mainnet.infura.io/v3/800c641949d64d768a5070a1b0511938"},
		Web3Url:    "https://mainnet.infura.io/v3/800c641949d64d768a5070a1b0511938",
		RpcAddress: "127.0.0.1",
		RpcPort:    8545,
		LogLevel:   "INFO",
		//Eth2Network:      "prater",
		//TrustedBlockRoot: "0x017e4563ebf7fed67cff819c63d8da397b4ed0452a3bbd7cae13476abc5020e4",
	}
	ctx, cancel := context.WithCancel(context.Background())
	proxy.StartLightClient(ctx, &testConfig)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Before range signals")
	// time.Sleep(8 * time.Second)
	// cancel()
	for range signals {
		fmt.Println("Signal caught, exiting")
		cancel()
	}
	fmt.Println("Exiting")

}
