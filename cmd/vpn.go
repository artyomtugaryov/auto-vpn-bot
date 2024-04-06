package main

import (
	"fmt"

	"github.com/artyomtugaryov/vpnbot/pkg/vpn_assistent"
)

func main() {
	u := vpn_assistent.NewAdmin("Artyom", "123456", "artyomtugaryov")
	fmt.Println(u)
}
