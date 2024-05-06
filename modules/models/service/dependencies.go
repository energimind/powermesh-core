package service

// idGenerator defines the external ID generator.
type idGenerator interface {
	GenerateID() string
}
