package mcptransport

import "sync"

type InMemoryTransport struct {
	ServerId      string
	ClientChannel chan map[string]interface{}
	ServerChannel chan interface{}
	Closed        bool
	Mu            sync.RWMutex
}

type TransportManager struct {
	Transports map[string]*InMemoryTransport
	Mu         sync.RWMutex
}
