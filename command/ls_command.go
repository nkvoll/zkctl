package command

import (
	"fmt"
	"github.com/codegangsta/cli"
	"strings"
)

func NewLsCommand() cli.Command {
	return cli.Command{
		Name:  "ls",
		Usage: "list a directory",

		Action: func(c *cli.Context) {
			handleLs(c)
		},
	}
}

func handleLs(c *cli.Context) {
	client := connect(c)

	path := ensureStartsWithSlash(c.Args().First())

	children, _, err := client.conn.Children(client.namespace + path)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Println(strings.Join(children, "\n"))
}
