package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
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
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	var plugindir string
	envplugindir, ok := os.LookupEnv("PLUGIN_DIR")
	if !ok {
		plugindir = fmt.Sprintf("%s/.config/plugins/", usr.HomeDir)
	} else {
		plugindir = envplugindir
	}

	enumerateplugins, err := listplugins(plugindir)
	if err != nil {
		log.Fatalf("Unable to list plugins %s", err)
	}
	var pluginCmd string
	for _, a := range enumerateplugins {
		fmt.Printf("Plugin names are %s\n", a)
		plugincheck := fmt.Sprintf("%s/%s", plugindir, operation)
		if plugincheck == a {

			fmt.Printf("Plugin found")
			pluginCmd = a
			break
		} else {
			fmt.Printf("%s: %s", plugincheck, a)
		}
	}
	fmt.Printf("Plugin is %s\n", pluginCmd)
	// Note this is BAD, but demonstrates the concept. This will load plugins any sprintf is bad but for now its ok

	//log.SetOutput(ioutil.Discard)

	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "CALCULATOR_SIMPLE",
			MagicCookieValue: "28ad59a0-f6e2-446f-b64b-c48de4ab721e",
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
	fmt.Printf("Output value is %f\n", test)
}

func listplugins(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	var plugins []string
	if err != nil {
		log.Fatal("Unable to read the plugin directory: %s\n", err)
	}
	for _, v := range files {
		fqdn := fmt.Sprintf("%s/%s", path, v.Name())
		plugins = append(plugins, fqdn)
	}

	return plugins, nil
}
