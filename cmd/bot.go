package main

import (
	"flag"
	"log"

	"github.com/artyomtugaryov/vpnbot/pkg/entities"
	"github.com/artyomtugaryov/vpnbot/pkg/storage"
	sqlite_storage "github.com/artyomtugaryov/vpnbot/pkg/storage/sqlite"
)

func main() {
	// tgClient := telegram.New("api.telegram.org", mustToken())
	storage := getStorage()
	customer := entities.Customer{
		Username: "artyomtugaryov",
	}

	if err := storage.SaveCusromer(&customer); err != nil {
		panic(err)
	}
	if err := storage.DisableCusromer(&customer); err != nil {
		panic(err)
	}

	// fetcher = fetcher.New(tgClient)
	// processor = processor.New(tgClient)
	// consumer.Start(fetcher, processor)
	storage.Close()
}

func getStorage() storage.Storage {
	return sqlite_storage.New("./data")
}

func mustToken() string {
	token := flag.String("token", "", "Telegram token for access")

	flag.Parse()

	if *token == "" {
		log.Fatal("Token was not specified")
	}

	return *token
}
