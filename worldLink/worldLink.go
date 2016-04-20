package worldLink

import "net"

// WorldLink wraps a connection
type WorldLink struct {
	Connection net.Listener
	Port       int
}

// New creates a new WorldLink
func New() *WorldLink {
	return &WorldLink{}
}

// Connect - Bind the worldLink to a local port
func (wl *WorldLink) Connect() {
	// Start connection
	conn, err := net.Listen("tcp", ":0")

	if err != nil {
		panic(err)
	}

	// Assign worldLink Connection
	wl.Connection = conn

	// Extract Port from Connection
	port := conn.Addr().(*net.TCPAddr).Port
	wl.Port = port
}

// Disconnect - Close the WorldLink connection
func (wl *WorldLink) Disconnect() {
	if wl != nil {
		wl.Connection.Close()
	}
}
