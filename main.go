package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
        "io/ioutil"
        "strings"
	"github.com/herber523/yamlenv/lib/env"
	"github.com/herber523/yamlenv/lib/file"
)

type (
	// Config information.
	Config struct {
		readpath string
		writepath string
	}
)

var config Config

func main() {
	app := cli.NewApp()
	app.Name = "YAMLENV"
	app.Action = run
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "readpath,r",
			Usage: "read dir path",
		},
		cli.StringFlag{
			Name:  "writepath,w",
			Usage: "write dir path",
		},
	}

	app.Run(os.Args)
}

func run(c *cli.Context) error {
	config = Config{
		readpath: c.String("readpath"),
		writepath: c.String("writepath"),
	}

	return exec()
}

func exec() error{
        dirpath := config.readpath
        writepath := config.writepath
        _, err := os.Stat(writepath)
        if err != nil {
            fmt.Println(err)
            os.MkdirAll(writepath, os.ModePerm)
        }
        envs := env.GetEnv()
        files := file.GetFile(dirpath)
        for _, f := range files{
            dat, _ := ioutil.ReadFile(dirpath + "/" + f)
            str := string(dat)
            for envk, envv := range envs{
                envkey := fmt.Sprintf("${%s}", envk)
                str = strings.Replace(str, envkey, envv, -1)
            }
            fmt.Println(str)
            sb := []byte(str)
            err := ioutil.WriteFile(writepath + "/" + f, sb, 0644)
            if err != nil{
                fmt.Println(err)
            }
        }
        return nil
}

