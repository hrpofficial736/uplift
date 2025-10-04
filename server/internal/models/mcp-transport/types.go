package mcptransport

import "sync"


type InMemoryTransport struct {
	ServerId string
	ClientChannel chan []byte
	ServerChannel chan []byte
	Closed bool
	Mu sync.RWMutex
}


type TransportManager struct {
	Transports map[string]*InMemoryTransport
	Mu sync.RWMutex
}