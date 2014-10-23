package main

import (
	"github.com/codegangsta/cli"
	"github.com/nkvoll/zkctl/command"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Name = "zkctl"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "zookeeper, z", Value: "", Usage: "a ZooKeeper connection string", EnvVar: "ZOOKEEPER_SERVERS"},
		cli.StringFlag{Name: "format, f", Value: "", Usage: "output format", EnvVar: "ZKCTL_OUTPUT_FORMAT"},
	}

	app.Commands = []cli.Command{
		command.NewLsCommand(),
		command.NewGetCommand(),
		command.NewSetCommand(),
	}

	app.Run(os.Args)
}
