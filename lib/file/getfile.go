package file

import (
	"io/ioutil"
	"log"
        "strings"
)

func Contains(s, substr string) bool {
	return strings.Index(s, substr) != -1
}


func GetFile(dirpath string) []string{
        var rfiles[]string
	files, err := ioutil.ReadDir(dirpath)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		fname := f.Name()
		if Contains(fname, "yaml") {
                        rfiles = append(rfiles, f.Name())
		}
	}
        return rfiles
}
