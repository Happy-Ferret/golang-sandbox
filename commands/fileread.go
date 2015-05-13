//http://qiita.com/kazamai/items/3699c24b16ed4fb883d4
package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func fileread(c *cli.Context) {
	args := os.Args
	if len(args) != 2 {
		log.Fatal("Usage: ", path.Base(args[0]), " file")
	}
	data, err := ioutil.ReadFile(args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Read:[", string(data), "]")
}
