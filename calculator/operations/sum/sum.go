package main

import (
	"log"

	"github.com/hashicorp/go-plugin"
	"github.com/veeruns/rpcserver/calculator/calcs"
)

func main() {

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "CALCULATOR_SIMPLE",
			MagicCookieValue: "28ad59a0-f6e2-446f-b64b-c48de4ab721e",
		},
		Plugins: map[string]plugin.Plugin{
			"calculator": calcs.CalcsPlugin{
				Impl: Calcssum{},
			},
		},
	})
	log.Println("Plugin Serving")
}

type Calcssum struct{}

func (Calcssum) Operation(input []float32) float32 {
	var sum float32
	log.Printf("Operation being called")
	for _, v := range input {
		sum = sum + v
	}

	return sum
}
