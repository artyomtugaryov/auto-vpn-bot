package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	token := mustToken()
	fmt.Println(token)
	// tgClient = telegram.New(token)
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
