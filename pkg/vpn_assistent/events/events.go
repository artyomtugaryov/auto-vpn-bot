package events

import "github.com/artyomtugaryov/vpnbot/pkg/vpn_assistent/users"

type Event struct {
	eventType    int
	author       users.User
	relatedUsers []users.User
}

func NewEvent(eventType int, author users.User, relatedUsers []users.User) Event {
	return Event{
		eventType:    eventType,
		author:       author,
		relatedUsers: relatedUsers,
	}
}

func (self *Event) EventType() int {
	return self.eventType
}

func (self *Event) Author() users.User {
	return self.author
}

func (self *Event) RelatedUsers() []users.User {
	return self.relatedUsers
}
