package mcp

import (
	"context"
	"fmt"
)

func NewInMemoryTransport (serverId string) *InMemoryTransport {
	return &InMemoryTransport{
		serverId: serverId,
		clientChannel: make(chan []byte, 100),
		serverChannel: make(chan []byte, 100),
		closed: false,
	}
}


// Methods to be called by the MCP Client to the Server
func (t *InMemoryTransport) Send (ctx context.Context, message []byte) error {
	t.mu.RLock();

	if t.closed {
		t.mu.RUnlock();
		return fmt.Errorf("transport closed")
	}


	t.mu.RUnlock();
	select {
	case t.clientChannel <- message:
		return nil
	case <-ctx.Done():
		return ctx.Err();
	}
}


func (t *InMemoryTransport) Receive (ctx context.Context) ([]byte, error) {
	t.mu.RLock();

	if t.closed {
		t.mu.RUnlock();
		return nil, fmt.Errorf("transport closed");
	}


	t.mu.RUnlock();

	select {
	case message := <-t.serverChannel:
		return message, nil
	case <-ctx.Done():
		return nil, ctx.Err();
	}
}


// Methods to be called by MCP Server to the Client
func (t *InMemoryTransport) ReceiveFromClient (ctx context.Context) ([]byte, error) {
	t.mu.RLock();

	if t.closed {
		t.mu.RUnlock();
		return nil, fmt.Errorf("transport closed")
	}

	t.mu.RUnlock();

	select {
	case message := <-t.clientChannel:
		return message, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}


func (t *InMemoryTransport) SendToClient (ctx context.Context, message []byte) error {
	t.mu.RLock();

	if t.closed {
		t.mu.RUnlock();
		return fmt.Errorf("transport closed")
	}

	t.mu.RUnlock();

	select {
	case t.serverChannel <- message:
		return nil

	case <-ctx.Done():
		return ctx.Err();
	}
}


func (t *InMemoryTransport) Close () error {
	t.mu.Lock()

	defer t.mu.Unlock();

	if !t.closed {
		t.closed = true;
		close(t.clientChannel)
		close(t.serverChannel)
	}

	return nil
}


func (t *InMemoryTransport) GetServerId () string {
	return t.serverId;
}