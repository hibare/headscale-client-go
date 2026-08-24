module github.com/hibare/headscale-client-example

go 1.26.4

// Comment out the following line to use the latest version of headscale-client-go from GitHub instead of the local copy.
replace github.com/hibare/headscale-client-go => ../

require (
	github.com/hibare/headscale-client-go v0.7.0
	github.com/stretchr/testify v1.12.1
)

require (
	github.com/stretchr/objx v0.5.3 // indirect
	go.yaml.in/yaml/v3 v3.0.5 // indirect
)
