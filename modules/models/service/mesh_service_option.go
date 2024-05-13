package service

// MeshServiceOption defines the option for the service.
type MeshServiceOption func(*MeshService)

// WithMeshListener sets the listener for the service.
func WithMeshListener(listener meshListener) MeshServiceOption {
	return func(s *MeshService) {
		s.listener = listener
	}
}
