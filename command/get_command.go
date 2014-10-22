package command

import (
	"fmt"
	"github.com/codegangsta/cli"
)

func NewGetCommand() cli.Command {
	return cli.Command{
		Name:  "get",
		Usage: "get the contents of a node",

		Action: func(c *cli.Context) {
			handleGet(c)
		},
	}
}

func handleGet(c *cli.Context) {
	client := connect(c)

	path := ensureStartsWithSlash(c.Args().First())

	bytes, stat, err := client.conn.Get(client.namespace + path)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	printStat(stat)
	fmt.Println(string(bytes))
}
