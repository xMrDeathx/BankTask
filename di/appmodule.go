package di

import (
	"BankTask/wallet/cmd"
	"BankTask/wallet/impl/di"
	"log"
	"net/http"
	"os"
)

func InitAppModule() {
	walletConn, err := di.InitWalletModule(cmd.NewConfig())
	defer walletConn.Close()

	appPort := os.Getenv("BACKEND_PORT")
	if appPort == "" {
		appPort = "8080"
	}

	log.Println("INFO: Start wallet server")
	err = http.ListenAndServe(":"+appPort, nil)
	if err != nil {
		log.Panic("ListenAndServe: ", err)
	}
}

func Migrate() {
	err := cmd.Migrate(cmd.NewConfig())
	if err != nil {
		log.Fatal("Failed to migrate wallet module: ", err)
	}

	log.Println("INFO: Database migration completed")
}
