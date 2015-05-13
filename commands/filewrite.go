//http://qiita.com/kazamai/items/3699c24b16ed4fb883d4
package commands

import (
	"github.com/codegangsta/cli"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func filewrite(c *cli.Context) {
	args := os.Args
	if len(args) != 3 {
		log.Fatal("Usage: ", path.Base(args[0]), " filename text")
	}

	data := []byte(args[2])
	err := ioutil.WriteFile(args[1], data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
