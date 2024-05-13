package service

// Option defines the option for the service.
type Option func(service *PermissionService)

// WithListener sets the listener for the service.
func WithListener(listener listener) Option {
	return func(s *PermissionService) {
		s.listener = listener
	}
}
