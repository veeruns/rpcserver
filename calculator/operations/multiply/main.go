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
				Impl: Calcmultiply{},
			},
		},
	})
	log.Println("Plugin Serving")
}

type Calcmultiply struct{}

func (Calcmultiply) Operation(input []float32) float32 {
	var product float32
	product = 1
	log.Printf("Operation multiply being called being called")
	for _, v := range input {
		product = product * v
	}

	return product
}
