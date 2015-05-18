package main

import (
	"fmt"
	"github.com/Viperoo/golog"
	//	"github.com/jinzhu/gorm"
	"io"
	"os"
	"os/user"
	//	"strconv"
	"github.com/codegangsta/cli"
)

var logger log.Logger
var workdir string

const (
	TIME_FORMAT = "Mon Jan 02 15:04:05 MST 2006"
)

func main() {
	setWorkDir()
	setLogger()

	db := conn()
	db.LogMode(false)
	//	if *migrate != false {
	//		makeMigrate(db)
	//	}

	app := cli.NewApp()
	app.Name = "go-mailer - Send mailing with GO"
	app.Usage = "I send mailing from SMTP gateway and more ;)"
	app.Version = "0.1.1"

	app.Commands = []cli.Command{
		{
			Name:  "mailing",
			Usage: "options for mailing",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "add a new mailing",
					Action: func(c *cli.Context) {
						addMailing(db)
					},
				},
				{
					Name:  "list",
					Usage: "list all mailing with progress",
					Action: func(c *cli.Context) {
						println("removed task template: ", c.Args().First())
					},
				},
				{
					Name:  "import",
					Usage: "import mails to mailing",
					Action: func(c *cli.Context) {
						println("removed task template: ", c.Args().First())
					},
				},
				{
					Name:  "export",
					Usage: "export mails with status, send date",
					Action: func(c *cli.Context) {
						println("removed task template: ", c.Args().First())
					},
				},
				{
					Name:  "start",
					Usage: "start sending mails",
					Action: func(c *cli.Context) {
						println("removed task template: ", c.Args().First())
					},
				},
				{
					Name:  "stop",
					Usage: "stop sending mails",
					Action: func(c *cli.Context) {
						println("removed task template: ", c.Args().First())
					},
				},
				{
					Name:  "remove",
					Usage: "remove mailings with all mails",
					Action: func(c *cli.Context) {
						println("removed task template: ", c.Args().First())
					},
				},
			},
		},
		{
			Name:    "migrate",
			Aliases: []string{"m"},
			Usage:   "migrate all databases WARNING: database will be emptied",
			Action: func(c *cli.Context) {
				makeMigrate(db)
			},
		},
	}

	app.Run(os.Args)
}

func setWorkDir() {
	usr, err := user.Current()
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	workdir = usr.HomeDir + "/.go-mailer/"
	if _, err := os.Stat(workdir); os.IsNotExist(err) {
		os.Mkdir(workdir, 0777)
		os.Mkdir(workdir+"db", 0777)

		if _, err := os.Stat(workdir + "configuration.conf"); err != nil {
			file, err := os.OpenFile(workdir+"configuration.conf", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				fmt.Print(err.Error())
				os.Exit(1)
			}
			defer file.Close()

			var d string = `[Default]
From = admin@localhost
[SMTP]
Host = localhost
Port = 25
User = smtp@localhost
Password = 
From = rss@localhost
`

			if _, err = file.WriteString(d); err != nil {
				fmt.Print(err.Error())
				os.Exit(1)
			}
			logger.Info("Default configuration loaded")
		}

	}
	loadConfig(workdir + "configuration.conf")

}

func setLogger() {
	file, err := os.OpenFile(workdir+"go-mailer.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	var multi io.Writer
	multi = io.MultiWriter(file, os.Stdout)
	logger, _ = log.NewLogger(multi,
		log.TIME_FORMAT_SEC,   // Set time writting format.
		log.LOG_FORMAT_SIMPLE, // Set log writting format.
		log.LogLevel_Debug)    // Set log level.
}
