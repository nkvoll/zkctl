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
		cli.StringFlag{"zookeeper, z", "", "a ZooKeeper connection string", "ZOOKEEPER_SERVERS"},
	}

	app.Commands = []cli.Command{
		command.NewLsCommand(),
		command.NewGetCommand(),
		command.NewSetCommand(),
	}

	app.Run(os.Args)
}
