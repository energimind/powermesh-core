package service

// ModelServiceOption defines the option for the service.
type ModelServiceOption func(*ModelService)

// WithModelListener sets the listener for the service.
func WithModelListener(listener modelListener) ModelServiceOption {
	return func(s *ModelService) {
		s.listener = listener
	}
}
