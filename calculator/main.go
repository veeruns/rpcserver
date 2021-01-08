package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"

	plugin "github.com/hashicorp/go-plugin"

	"github.com/veeruns/rpcserver/calculator/calcs"
)

func main() {
	operation := "sum"
	argu := []float32{1, 2}
	if len(os.Args) > 1 {
		operation = os.Args[1]
		argu = nil
		for _, v := range os.Args[2:] {
			if s, err := strconv.ParseFloat(v, 32); err == nil {
				argu = append(argu, float32(s))
			}

		}

	}

	// Note this is BAD, but demonstrates the concept. This will load plugins any sprintf is bad but for now its ok
	pluginCmd := fmt.Sprintf("./calculator-%s", operation)

	//log.SetOutput(ioutil.Discard)

	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "CALCULATOR_SIMPLE",
			MagicCookieValue: "PI",
		},
		Plugins: map[string]plugin.Plugin{
			"calculator": new(calcs.CalcsPlugin),
		},
		Cmd: exec.Command(pluginCmd),
	})
	log.Println("Client Created")
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("calculator")
	log.Printf("Dispensing caculator")
	if err != nil {
		log.Fatal(err)
	}

	// We should have a Greeter now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	calcs := raw.(calcs.Calcs)
	test := calcs.Operation(argu)
	fmt.Printf("%v\n", test)
}
