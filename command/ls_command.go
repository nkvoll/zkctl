package command

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/samuel/go-zookeeper/zk"
	"strings"
)

func NewLsCommand() cli.Command {
	return cli.Command{
		Name:  "ls",
		Usage: "list a directory",

		Flags: []cli.Flag{
			cli.BoolFlag{"recursive, r", "false", "list directories recursively"},
			cli.BoolFlag{"stat, s", "false", "show ZooKeeper stat"},
		},

		Action: func(c *cli.Context) {
			handleLs(c)
		},
	}
}

type LsJson struct {
	Path     string
	Stat     *zk.Stat `json:",omitempty"`
	Children []string
}

func doLs(client Client, path string, recursive bool) ([]string, *zk.Stat, error) {
	children, stat, err := client.conn.Children(path)
	if err != nil {
		return children, stat, err
	}

	if recursive {
		for _, child := range children {
			childPath := path + "/" + child
			nestedChildren, _, err := doLs(client, childPath, recursive)
			if err != nil {
				return children, stat, err
			}
			for _, nestedChild := range nestedChildren {
				children = append(children, child+"/"+nestedChild)
			}
		}
	}

	return children, stat, nil
}

func handleLs(c *cli.Context) {
	client := connect(c)

	path := ensureStartsWithSlash(c.Args().First())

	children, rootStat, err := doLs(client, client.namespacedPath(path), c.Bool("recursive"))
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	if c.GlobalString("format") == "json" {
		result := LsJson{path, rootStat, children}
		if !c.Bool("stat") {
			result.Stat = nil
		}

		js, err := json.Marshal(result)
		if err != nil {
			fmt.Println("err:", err)
		} else {
			fmt.Println(string(js))
		}
	} else {
		fmt.Println(strings.Join(children, "\n"))
	}
}
