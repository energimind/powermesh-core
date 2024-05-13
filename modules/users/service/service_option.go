package service

// Option defines the option for the service.
type Option func(*UserService)

// WithListener sets the listener for the service.
func WithListener(listener listener) Option {
	return func(s *UserService) {
		s.listener = listener
	}
}
