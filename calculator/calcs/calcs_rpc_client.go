package calcs

import (
	"fmt"
	"net/rpc"
)

//Calcsrpc is the client structure which will be called from main
type Calcsrpc struct{ client *rpc.Client }

//Operation that main uses to communicate with plugin
func (g *Calcsrpc) Operation(input []float32) float32 {
	var resp float32
	err := g.client.Call("Plugin.Operation", input, &resp)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		fmt.Printf("Before panic %v\n", input)
		panic(err)
	}

	return resp
}
