package command

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/samuel/go-zookeeper/zk"
)

func NewSetCommand() cli.Command {
	return cli.Command{
		Name:  "set",
		Usage: "set the contents of a node",

		Flags: []cli.Flag{
			cli.IntFlag{Name: "version, v", Value: -1, Usage: "version to overwrite"},
		},

		Action: func(c *cli.Context) {
			handleSet(c)
		},
	}
}

type SetJson struct {
	Path string
	Stat *zk.Stat `json:",omitempty"`
}

func handleSet(c *cli.Context) {
	client := connect(c)

	path := ensureStartsWithSlash(c.Args().First())
	value := c.Args()[1]
	version := int32(c.Int("version"))

	stat, err := client.conn.Set(client.namespacedPath(path), []byte(value), version)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	if c.GlobalString("format") == "json" {
		result := SetJson{path, stat}

		js, err := json.Marshal(result)
		if err != nil {
			fmt.Println("err:", err)
		} else {
			fmt.Println(string(js))
		}
	} else {
		printStat(stat)
	}
}
