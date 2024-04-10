package service

import "github.com/energimind/powermesh-core/access"

// canAct returns true if the actor can act.
// For the user service, only admins can perform actions.
func canAct(actor access.Actor) bool {
	return actor.Role == access.RoleAdmin
}
