package calcs

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

//CalcsPlugin for plugins
type CalcsPlugin struct {
	Impl Calcs
}

//Server for server
func (plugin CalcsPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &CalcsrpcServer{
		Impl: plugin.Impl,
	}, nil
}

//Client for client
func (CalcsPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &Calcsrpc{client: c}, nil
}
