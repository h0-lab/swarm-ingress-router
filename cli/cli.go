package main

import (
	"github.com/tpbowden/swarm-ingress-router/server"
	"github.com/tpbowden/swarm-ingress-router/version"
	"github.com/urfave/cli"
	"os"
)

type args struct {
	bind     string
	interval int
}

func main() {
	start(os.Args, server.NewServer)
}

func start(args []string, serverInit func(string, int) server.Startable) {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "bind, b",
			Value: "127.0.0.1",
			Usage: "Bind to `address`",
		},
		cli.IntFlag{
			Name:  "interval, i",
			Value: 10,
			Usage: "Poll interval in `seconds`",
		},
	}
	app.Name = "Swarm Ingress Router"
	app.Usage = "Route DNS names to Swarm services based on labels"
	app.Version = version.Version.String()

	app.Action = func(c *cli.Context) error {
		server := serverInit(c.String("bind"), c.Int("interval"))
		server.Start()
		return nil
	}

	app.Run(args)
}
