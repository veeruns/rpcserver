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
			MagicCookieValue: "PI",
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
