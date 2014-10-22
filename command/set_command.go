package command

import (
	"fmt"
	"github.com/codegangsta/cli"
)

func NewSetCommand() cli.Command {
	return cli.Command{
		Name:  "set",
		Usage: "set the contents of a node",

		Flags: []cli.Flag{
			cli.IntFlag{"version, v", -1, "version to overwrite", ""},
		},

		Action: func(c *cli.Context) {
			handleSet(c)
		},
	}
}

func handleSet(c *cli.Context) {
	client := connect(c)

	path := ensureStartsWithSlash(c.Args().First())
	value := c.Args()[1]
	version := int32(c.Int("version"))

	stat, err := client.conn.Set(client.namespace+path, []byte(value), version)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	printStat(stat)
}
