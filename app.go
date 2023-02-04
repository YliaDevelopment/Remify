package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/YliaDevelopment/Remify/injectors"
	"github.com/YliaDevelopment/Remify/proxy"
	"github.com/YliaDevelopment/Remify/utils"
	"github.com/sandertv/gophertunnel/minecraft/auth"
	"github.com/sirupsen/logrus"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: remify <SERVER ADDR> <LISTEN ADDR>")
		return
	}

	serverAddress := os.Args[1]
	listenAddr := os.Args[2]

	if !strings.Contains(serverAddress, ":") {
		serverAddress = fmt.Sprintf("%s:19132", serverAddress)
	}

	if !strings.Contains(listenAddr, ":") {
		listenAddr = fmt.Sprintf("%s:9999", listenAddr)
	}

	token, err := utils.FetchToken()
	if err != nil {
		logrus.Fatalf("failed to get auth token: %v", err)
	}

	context := proxy.Context{
		Token:            auth.RefreshTokenSource(token),
		ServerAddress:    serverAddress,
		ListenAddress:    listenAddr,
		EnabledInjectors: []injectors.Injector{&injectors.LatencyShow{}},
	}

	err = context.Start()
	fmt.Println(err)
}
