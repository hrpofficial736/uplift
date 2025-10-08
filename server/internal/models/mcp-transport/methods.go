package mcptransport

import (
	"context"
	"encoding/json"
	"fmt"
)

func NewInMemoryTransport(serverId string) *InMemoryTransport {
	return &InMemoryTransport{
		ServerId:      serverId,
		ClientChannel: make(chan map[string]interface{}),
		ServerChannel: make(chan interface{}),
		Closed:        false,
	}
}

// Methods to be called by the MCP Client to the Server
func (t *InMemoryTransport) Send(ctx context.Context, message map[string]interface{}) error {
	t.Mu.RLock()

	if t.Closed {
		t.Mu.RUnlock()
		return fmt.Errorf("transport closed")
	}

	t.Mu.RUnlock()
	select {
	case t.ClientChannel <- message:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (t *InMemoryTransport) Receive(ctx context.Context) (interface{}, error) {
	t.Mu.RLock()

	if t.Closed {
		t.Mu.RUnlock()
		return nil, fmt.Errorf("transport closed")
	}

	t.Mu.RUnlock()

	select {
	case message := <-t.ServerChannel:
		var decoded interface{}
		if err := json.Unmarshal(message.([]byte), &decoded); err != nil {
			return nil, fmt.Errorf("failed to unmarshal message: %s", err)
		}
		return decoded, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// Methods to be called by MCP Server to the Client
func (t *InMemoryTransport) ReceiveFromClient(ctx context.Context) (map[string]interface{}, error) {
	t.Mu.RLock()

	if t.Closed {
		t.Mu.RUnlock()
		return nil, fmt.Errorf("transport closed")
	}

	t.Mu.RUnlock()

	select {
	case message := <-t.ClientChannel:
		return message, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (t *InMemoryTransport) SendToClient(ctx context.Context, message interface{}) error {
	t.Mu.RLock()

	if t.Closed {
		t.Mu.RUnlock()
		return fmt.Errorf("transport closed")
	}

	t.Mu.RUnlock()

	select {
	case t.ServerChannel <- message:
		return nil

	case <-ctx.Done():
		return ctx.Err()
	}
}

func (t *InMemoryTransport) Close() error {
	t.Mu.Lock()

	defer t.Mu.Unlock()

	if !t.Closed {
		t.Closed = true
		close(t.ClientChannel)
		close(t.ServerChannel)
	}

	return nil
}

func (t *InMemoryTransport) GetServerId() string {
	return t.ServerId
}

// Manager methods

func (tm *TransportManager) CreateTransport(agentType string) *InMemoryTransport {
	tm.Mu.Lock()

	defer tm.Mu.Unlock()

	transport := NewInMemoryTransport(agentType)
	tm.Transports[agentType] = transport

	return transport
}

func (tm *TransportManager) GetTransport(agentType string) (*InMemoryTransport, error) {
	tm.Mu.RLock()

	defer tm.Mu.RUnlock()

	transport, exists := tm.Transports[agentType]

	if !exists {
		return nil, fmt.Errorf("transport for the agent %s not found", agentType)
	}

	return transport, nil
}

func (tm *TransportManager) CloseAll() error {
	tm.Mu.Lock()

	defer tm.Mu.Unlock()

	for agentType, transport := range tm.Transports {
		if err := transport.Close(); err != nil {
			return fmt.Errorf("failed to close transport for the %s: %w", agentType, err)
		}
	}

	tm.Transports = make(map[string]*InMemoryTransport)

	return nil
}
