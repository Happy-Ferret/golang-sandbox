//http://qiita.com/kazamai/items/3699c24b16ed4fb883d4
package commands

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"github.com/codegangsta/cli"
)

func listfileindir(c *cli.Context) {
	args := os.Args
	if len(args) != 2 {
		log.Fatal("Usage: ", path.Base(args[0]), " dir")
	}
	files, _ := ioutil.ReadDir(args[1])
	for _, f := range files {
		fmt.Println(f.Name())
	}
}
