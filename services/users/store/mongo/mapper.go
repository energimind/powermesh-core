package mongo

import "github.com/energimind/powermesh-core/services/users"

func toStoreUser(u users.User) storeUser {
	return storeUser{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
	}
}

func fromStoreUser(u storeUser) users.User {
	return users.User{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
	}
}
