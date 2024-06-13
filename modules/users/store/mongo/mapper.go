package mongo

import "github.com/energimind/powermesh-core/modules/users"

func toStoreUser(u users.User) storeUser {
	return storeUser{
		ID:          u.ID,
		ExternalID:  u.ExternalID,
		Username:    u.Username,
		DisplayName: u.DisplayName,
		Email:       u.Email,
	}
}

func fromStoreUser(u storeUser) users.User {
	return users.User{
		ID:          u.ID,
		ExternalID:  u.ExternalID,
		Username:    u.Username,
		DisplayName: u.DisplayName,
		Email:       u.Email,
	}
}
