package e2e

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
	s.NotEmpty(resp.Policy, "Expected policy to be updated")
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
	s.NotEmpty(resp.Policy, "Expected policy with groups to be updated")
}
