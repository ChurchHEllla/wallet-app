package main

import (
	"log"
	"wallet-app/config"
	updateWallet "wallet-app/internal/api/handler/api_1_wallet"
	createWallet "wallet-app/internal/api/handler/create_api_1_wallet"
	getWallet "wallet-app/internal/api/handler/get_api_1_wallets"
	"wallet-app/internal/db"
	"wallet-app/internal/domain/wallet/repository"
	"wallet-app/internal/domain/wallet/service"

	"github.com/gin-gonic/gin"
)

const patternCreateWallet = "/api/v1/createWallet"
const (
	patternWalletId  = ":walletId"
	patternGetWallet = "/api/v1/wallets/"
)
const patternBalance = "/api/v1/wallet"

// TODO написать тесты
func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	database, err := db.ConnectDB(cfg.PostgresDSN)
	if err != nil {
		log.Fatal(err)
	}

	walletRepo := repository.New(database, nil)

	walletService := service.New(walletRepo)

	createWalletHandler := createWallet.New(walletService)
	getWalletHandler := getWallet.New(walletService)
	updateBalanceHandler := updateWallet.New(walletService)

	server := gin.Default()

	server.POST(patternBalance, updateBalanceHandler.Handler)
	server.GET(patternGetWallet+patternWalletId, getWalletHandler.Handler)
	server.POST(patternCreateWallet, createWalletHandler.Handler)

	log.Println("server started on :8080")
	log.Fatal(server.Run(":8080"))

}
