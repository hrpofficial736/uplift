package mcp

import (
	"context"
	"log"
)


func NewAgentCoordinator () *AgentCoordinator {
	return &AgentCoordinator{
		transportManager: NewTransportManager(),
		mcpClients: make(map[string]*AgentMCPClient),
		mcpServers: make(map[string]*AgentMCPServer),
	}
}


func (ac *AgentCoordinator) AddAgent (agentType string, callLLM func (string) (string, error), prompt string) {
	transport := ac.transportManager.CreateTransport(agentType)
	ctx, cancel := context.WithCancel(context.Background())

	client := &AgentMCPClient{
		agentType: agentType,
		callLLM: callLLM,
		transport: transport,
		ctx: ctx,
		cancel: cancel,
	}

	client.Initialize()

	server := &AgentMCPServer{
		serverId: agentType,
		transport: transport,
	}


	// server.RegisterTool("review_code", func(params map[string]interface{}) (interface{}, error) {
	// 	code := params["code"].(string)

	// 	return callLLM(code);
	// })


	ac.mcpClients[agentType] = client
	ac.mcpServers[agentType] = server


	text, err := ac.mcpClients[agentType].callLLM(prompt)

	if err != nil {
		log.Fatalf("error while prompting the model: %s\n", err)
	}

	log.Println(text)
} 