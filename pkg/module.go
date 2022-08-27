package aether

// Module is an interface that represents an Aether Module. A module is a self-contained, reuseable piece of
// functionality that servers can install. It is the fundamental building block of an Aether server.
type Module interface {
	// Configure is a function that will be called as Aether sets up the server. Configure is passed
	// a pointer to a ServerConfig, which can be used to add interceptors, register gRPC services, and more.
	Configure(c *ServerConfig) error
}
