package users

type User struct {
	Name     string
	UniqueID string
	Username string
	isAdmin  bool
	isFrozen bool
}

func NewUser(name string, uniqueID string, username string) User {
	return User{
		Name:     name,
		UniqueID: uniqueID,
		Username: username,
		isAdmin:  false,
		isFrozen: false,
	}
}

func NewAdmin(name string, uniqueID string, username string) User {
	return User{
		Name:     name,
		UniqueID: uniqueID,
		Username: username,
		isAdmin:  true,
		isFrozen: false,
	}
}

func (self User) Defrost() {
	self.isFrozen = false
}

func (self User) Freeze() {
	self.isFrozen = true
}
