package main

import (
	"fmt"
	"os"

	"../github.com/codegangsta/cli"
	"./util"
)

func main() {
	app := cli.NewApp()
	app.Name = "oss-starter"
	app.Usage = "Create scaffold of OSS"
	app.Version = "0.1.0"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "dryrun, d",
			Usage: "oss --dry-run start react-native",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "start",
			Aliases: []string{"s"},
			Usage:   "Create new scaffold",
			Action:  start,
		},
	}

	app.Before = func(c *cli.Context) error {
		// fmt.Println("OSS Starting...")
		return nil
	}

	app.After = func(c *cli.Context) error {
		// fmt.Println("Created!")
		return nil
	}

	app.Run(os.Args)
}

func start(c *cli.Context) {

	// var isDry = c.GlobalBool("dryrun")
	// if isDry {
	// 	fmt.Println("this is dry-run")
	// }

	var paramFirst = ""
	if len(c.Args()) > 0 {
		paramFirst = c.Args().First()
		if util.TypeChecker(paramFirst) {
			fmt.Printf("Contained! %s\n", paramFirst)
		}
	}
}
