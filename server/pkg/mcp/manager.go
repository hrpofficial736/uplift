package mcp

import "fmt"


func NewTransportManager () *TransportManager {
	return &TransportManager{
		transports: make(map[string]*InMemoryTransport),
	}
}


func (tm *TransportManager) CreateTransport (agentType string) *InMemoryTransport {
	tm.mu.Lock();

	defer tm.mu.Unlock();

	transport := NewInMemoryTransport(agentType);
	tm.transports[agentType] = transport;

	return transport;
}

func (tm *TransportManager) GetTransport (agentType string) (*InMemoryTransport, error) {
	tm.mu.RLock();

	defer tm.mu.RUnlock();

	transport, exists := tm.transports[agentType];

	if !exists {
		return nil, fmt.Errorf("transport for the agent %s not found", agentType)
	}

	return transport, nil;
}

func (tm *TransportManager) CloseAll () error {
	tm.mu.Lock();

	defer tm.mu.Unlock();

	for agentType, transport := range tm.transports {
		if err := transport.Close(); err != nil {
			return fmt.Errorf("failed to close transport for the %s: %w", agentType, err)
		}
	}

	tm.transports = make(map[string]*InMemoryTransport)

	return nil;
}