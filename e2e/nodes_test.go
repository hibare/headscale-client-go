package e2e

import (
	"github.com/hibare/headscale-client-go/v1/nodes"
)

func (s *E2ESuite) TestNodes_List() {
	nodeList, err := s.client.Nodes().List(s.T().Context(), nodes.NodeListFilter{})
	s.Require().NoError(err)
	s.GreaterOrEqual(len(nodeList.Nodes), 2, "Expected at least 2 nodes from Tailscale containers")
}

func (s *E2ESuite) TestNodes_Get() {
	ctx := s.T().Context()

	nodeList, err := s.client.Nodes().List(ctx, nodes.NodeListFilter{})
	s.Require().NoError(err)
	s.Require().GreaterOrEqual(len(nodeList.Nodes), 1, "Expected at least one node")

	node, err := s.client.Nodes().Get(ctx, nodeList.Nodes[0].ID)
	s.Require().NoError(err)
	s.Equal(nodeList.Nodes[0].ID, node.Node.ID, "Expected node ID to match")
}

func (s *E2ESuite) TestNodes_ListWithFilter() {
	ctx := s.T().Context()

	nodeList, err := s.client.Nodes().List(ctx, nodes.NodeListFilter{
		User: s.testUser.Name,
	})
	s.Require().NoError(err)
	s.GreaterOrEqual(len(nodeList.Nodes), 2, "Expected at least 2 nodes for test user")
}

func (s *E2ESuite) TestNodes_Rename() {
	ctx := s.T().Context()

	nodeList, err := s.client.Nodes().List(ctx, nodes.NodeListFilter{})
	s.Require().NoError(err)
	s.Require().GreaterOrEqual(len(nodeList.Nodes), 1, "Expected at least one node")

	originalName := nodeList.Nodes[0].GivenName
	newName := "renamed-test-node"

	renamedNode, err := s.client.Nodes().Rename(ctx, nodeList.Nodes[0].ID, newName)
	s.Require().NoError(err)
	s.Equal(newName, renamedNode.Node.GivenName, "Expected node to be renamed")

	_, err = s.client.Nodes().Rename(ctx, nodeList.Nodes[0].ID, originalName)
	s.Require().NoError(err)
}

func (s *E2ESuite) TestNodes_AddTags() {
	ctx := s.T().Context()

	nodeList, err := s.client.Nodes().List(ctx, nodes.NodeListFilter{})
	s.Require().NoError(err)
	s.Require().GreaterOrEqual(len(nodeList.Nodes), 1, "Expected at least one node")

	tags := []string{"tag:e2e-test"}
	taggedNode, err := s.client.Nodes().AddTags(ctx, nodeList.Nodes[0].ID, tags)
	s.Require().NoError(err)
	s.NotEmpty(taggedNode.Node.Tags, "Expected node to have tags")
}

func (s *E2ESuite) TestNodes_BackFillIP() {
	ctx := s.T().Context()

	resp, err := s.client.Nodes().BackFillIP(ctx, true)
	s.Require().NoError(err)
	s.NotNil(resp, "Expected backfill response")
}

func (s *E2ESuite) TestNodes_NodeProperties() {
	ctx := s.T().Context()

	nodeList, err := s.client.Nodes().List(ctx, nodes.NodeListFilter{})
	s.Require().NoError(err)
	s.Require().GreaterOrEqual(len(nodeList.Nodes), 1, "Expected at least one node")

	node := nodeList.Nodes[0]
	s.NotEmpty(node.ID, "Expected node ID")
	s.NotEmpty(node.Name, "Expected node name")
	s.NotEmpty(node.User.Name, "Expected node user name")
	s.NotEmpty(node.IPAddresses, "Expected node to have IP addresses")
}
