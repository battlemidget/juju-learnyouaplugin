// Copyright 2015 Adam Stokes <adam.stokes@ubuntu.com>

package main

import (
	"fmt"
	"github.com/juju/juju/juju"
	"github.com/juju/juju/state"
	"github.com/juju/juju/juju/osenv"
	"github.com/juju/juju/cmd/envcmd"

	_ "github.com/juju/juju/provider/all"
)

type LYAPluginCommand struct {
	envcmd.EnvCommandBase
	MachineMap map[string]*state.Machine
}

func (c *LYAPluginCommand) Info() error {
	client, err := juju.NewAPIClientFromName("local")
	if err != nil {
		panic(err)
	}
	info, err := client.EnvironmentInfo()
	fmt.Printf("Connection (%s): (%s)\n", info.UUID, info.Name)
	return nil
}

func main() {
	fmt.Println("Starting learnyouaplugin!")
	err := juju.InitJujuHome()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Using JUJU_HOME: (%s)\n", osenv.JujuHome())
	fmt.Println("Grabbing State information.")
	c := &LYAPluginCommand{}
	c.Info()
}