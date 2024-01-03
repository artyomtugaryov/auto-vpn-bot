package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/artyomtugaryov/vpnbot/pkg/storage"

	sqlite_storage "github.com/artyomtugaryov/vpnbot/pkg/storage/sqlite"
)

func main() {
	// tgClient := telegram.New("api.telegram.org", mustToken())
	dataStorage := sqlite_storage.New("./data")
	fmt.Println(dataStorage)
	customer := storage.Customer{
		Username: "artyomtugaryov",
	}

	if err := dataStorage.SaveCusromer(&customer); err != nil {
		panic(err)
	}
	if err := dataStorage.DisableCusromer(&customer); err != nil {
		panic(err)
	}

	// fetcher = fetcher.New(tgClient)
	// processor = processor.New(tgClient)
	// consumer.Start(fetcher, processor)
	dataStorage.Close()
}

func mustToken() string {
	token := flag.String("token", "", "Telegram token for access")

	flag.Parse()

	if *token == "" {
		log.Fatal("Token was not specified")
	}

	return *token
}
