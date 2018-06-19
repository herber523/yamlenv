package main

import (
	"fmt"
	"os"
        "io/ioutil"
        "strings"
	"github.com/herber523/yamlenv/lib/env"
	"github.com/herber523/yamlenv/lib/file"
)


func main() {
	dirpath := os.Args[1:][0]
        envs := env.GetEnv()
        files := file.GetFile(dirpath)
        for _, f := range files{
            dat, _ := ioutil.ReadFile(dirpath + "/" + f)
            str := string(dat)
            for envk, envv := range envs{
                envkey := fmt.Sprintf("${%s}", envk)
                str = strings.Replace(str, envkey, envv, -1)
            }
            sb := []byte(str)
            err := ioutil.WriteFile("./" + f, sb, 0644)
            if err != nil{
                fmt.Println(err)
            }
        }
}
