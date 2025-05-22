package main

import (
	"testing"

	"github.com/hibare/headscale-client-go/v1/client"
	"github.com/hibare/headscale-client-go/v1/nodes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListNodes(t *testing.T) {
	// Use the existing mock
	nodeMock := new(nodes.MockNodeResource)
	fakeNodes := nodes.NodesResponse{
		Nodes: []nodes.Node{{ID: "1", Name: "testnode"}},
	}
	nodeMock.On("List", mock.Anything, nodes.NodeListFilter{}).Return(fakeNodes, nil)

	// Mock client
	clientMock := new(client.MockClientV1)
	clientMock.On("Nodes").Return(nodeMock)

	output, err := listNodes(clientMock)
	assert.NoError(t, err)
	assert.Contains(t, output, "ID:1")
	assert.Contains(t, output, "Name:testnode")
	clientMock.AssertExpectations(t)
	nodeMock.AssertExpectations(t)
}
