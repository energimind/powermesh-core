package mongo

import "github.com/energimind/powermesh-core/modules/users"

var (
	validModelUser = users.User{
		ID:       "user-id",
		Username: "username",
		Email:    "user@somewhere.com",
	}
	validStoreUser = storeUser{
		ID:       validModelUser.ID,
		Username: validModelUser.Username,
		Email:    validModelUser.Email,
	}
)
