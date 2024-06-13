package service

import "github.com/energimind/powermesh-core/modules/users"

func userFromData(id string, data users.UserData) users.User {
	return users.User{
		ID:          id,
		ExternalID:  data.ExternalID,
		Username:    data.Username,
		DisplayName: data.DisplayName,
		Email:       data.Email,
	}
}
