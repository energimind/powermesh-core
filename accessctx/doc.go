// Package accessctx provides functionality for managing access control information within a context.
// It allows for the injection and retrieval of an actor's access control details within a context.
// This can be used to pass access control information between services.
//
// The main types are:
// - Actor: Represents an actor with access control details.
//
// The main functions are:
// - WithActor: Injects an actor into a context.
// - Actor: Retrieves an actor from a context.
//
// This package interacts with the 'access' package for defining the access control details.
package accessctx
