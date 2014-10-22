package command

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/samuel/go-zookeeper/zk"
)

func NewGetCommand() cli.Command {
	return cli.Command{
		Name:  "get",
		Usage: "get the contents of a node",

		Flags: []cli.Flag{
			cli.BoolFlag{"stat, s", "false", "show ZooKeeper stat"},
			cli.BoolFlag{"data-json, dj", "false", "decode data as json"},
		},

		Action: func(c *cli.Context) {
			handleGet(c)
		},
	}
}

type GetJson struct {
	Path string
	Stat *zk.Stat `json:",omitempty"`
	Data interface{}
}

func handleGet(c *cli.Context) {
	client := connect(c)

	path := ensureStartsWithSlash(c.Args().First())

	bytes, stat, err := client.conn.Get(client.namespacedPath(path))
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	if c.GlobalString("format") == "json" {
		result := GetJson{path, stat, string(bytes)}
		if !c.Bool("stat") {
			result.Stat = nil
		}

		if c.Bool("data-json") {
			mapConfig := make(map[string]interface{})
			err := json.Unmarshal(bytes, &mapConfig)
			if err != nil {
				fmt.Println("err:", err)
			} else {
				result.Data = mapConfig
			}
		}

		js, err := json.Marshal(result)
		if err != nil {
			fmt.Println("err:", err)
		} else {
			fmt.Println(string(js))
		}
	} else {
		if c.Bool("stat") {
			printStat(stat)
		}
		fmt.Print(string(bytes))
	}
}
