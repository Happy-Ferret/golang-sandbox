//http://qiita.com/kazamai/items/3699c24b16ed4fb883d4
package commands

import (
	"fmt"
	"log"
	"os"
	"path"
	"github.com/codegangsta/cli"
)

func fileexists(c *cli.Context) {
	args := os.Args
	if len(args) != 2 {
		log.Fatal("Usage: ", path.Base(args[0]), " filename")
	}
	if _, err := os.Stat(args[1]); err == nil {
		fmt.Println("File exists")
	} else {
		fmt.Println("File not found.")
	}
}
