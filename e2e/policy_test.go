package e2e

import (
	"encoding/json"
)

func (s *E2ESuite) TestPolicy_Get() {
	ctx := s.T().Context()

	pol, err := s.client.Policy().Get(ctx)
	s.Require().NoError(err)
	s.NotEmpty(pol.Policy, "Expected policy content")
}

func (s *E2ESuite) TestPolicy_Update() {
	ctx := s.T().Context()

	acl := `{
		"tagOwners": {
			"tag:test": ["testuser@headscale.test"]
		},
		"acls": [
			{
				"action": "accept",
				"src": ["*"],
				"dst": ["*:*"]
			}
		]
	}`

	resp, err := s.client.Policy().Update(ctx, acl)
	s.Require().NoError(err)
	s.assertPolicyEqual(acl, resp.Policy)
}

func (s *E2ESuite) TestPolicy_UpdateWithGroups() {
	ctx := s.T().Context()

	acl := `{
		"groups": {
			"group:admins": ["testuser@headscale.test"]
		},
		"tagOwners": {
			"tag:test": ["group:admins"]
		},
		"acls": [
			{
				"action": "accept",
				"src": ["group:admins"],
				"dst": ["*:*"]
			}
		]
	}`

	resp, err := s.client.Policy().Update(ctx, acl)
	s.Require().NoError(err)
	s.assertPolicyEqual(acl, resp.Policy)
}

// assertPolicyEqual unmarshals both expected and actual policy strings into maps and compares them.
func (s *E2ESuite) assertPolicyEqual(expected, actual string) {
	var expectedMap, actualMap map[string]interface{}

	err := json.Unmarshal([]byte(expected), &expectedMap)
	s.Require().NoError(err, "Failed to unmarshal expected policy")

	err = json.Unmarshal([]byte(actual), &actualMap)
	s.Require().NoError(err, "Failed to unmarshal actual policy")

	s.Equal(expectedMap, actualMap, "Expected policy content to match")
}
