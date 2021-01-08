module github.com/veeruns/rpcserver/calculator

go 1.15

require (
	github.com/hashicorp/go-plugin v1.4.0
	github.com/veeruns/rpcserver/calculator/calcs v0.0.0-00010101000000-000000000000
)

replace github.com/veeruns/rpcserver/calculator/calcs => ./calcs
