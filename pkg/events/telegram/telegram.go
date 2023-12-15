package telegram

import "github.com/artyomtugaryov/vpnbot/pkg/clients/telegram"

type Process struct {
	tgClient *telegram.Client
	offset   int
}
