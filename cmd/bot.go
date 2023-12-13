package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/artyomtugaryov/vpnbot/pkg/clients/telegram"
)

func main() {
	tgClient := telegram.New("api.telegram.org", mustToken())
	fmt.Println(tgClient)
	// fetcher = fetcher.New(tgClient)
	// processor = processor.New(tgClient)
	// consumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String("token", "", "Telegram token for access")

	flag.Parse()

	if *token == "" {
		log.Fatal("Token was not specified")
	}

	return *token
}
