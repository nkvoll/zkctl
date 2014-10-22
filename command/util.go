package command

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/samuel/go-zookeeper/zk"
	"strings"
	"time"
)

type Client struct {
	conn      *zk.Conn
	events    <-chan zk.Event
	namespace string
}

func connect(c *cli.Context) Client {
	connectionString := c.GlobalString("zookeeper")

	namespace := ""
	if strings.Contains(connectionString, "/") {
		connectionString, namespace = splitLink(connectionString, "/")
	}

	servers := strings.Split(connectionString, ",")

	conn, events, err := zk.Connect(servers, 10*time.Second)

	if err != nil {
		panic(err)
	}

	return Client{
		conn,
		events,
		"/" + namespace,
	}
}

func splitLink(s, sep string) (string, string) {
	x := strings.Split(s, sep)
	return x[0], x[1]
}

func printStat(stat *zk.Stat) {
	fmt.Println("Czxid", stat.Czxid)
	fmt.Println("Mzxid", stat.Mzxid)
	fmt.Println("Ctime", stat.Ctime)
	fmt.Println("Mtime", stat.Mtime)
	fmt.Println("Version", stat.Version)
	fmt.Println("Cversion", stat.Cversion)
	fmt.Println("Aversion", stat.Aversion)
	fmt.Println("EphemeralOwner", stat.EphemeralOwner)
	fmt.Println("DataLength", stat.DataLength)
	fmt.Println("NumChildren", stat.NumChildren)
	fmt.Println("Pzxid", stat.Pzxid)
}

func ensureStartsWithSlash(s string) string {
	if strings.HasPrefix(s, "/") {
		return s
	}
	return "/" + s
}
