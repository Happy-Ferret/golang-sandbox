//http://qiita.com/kazamai/items/3699c24b16ed4fb883d4
package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"os"
	"path"
	"strconv"
)

func stringtoi(c *cli.Context) {
	args := os.Args
	if len(args) != 2 {
		log.Fatal("Usage: ", path.Base(args[0]), " number")
	}
	if val, err := strconv.Atoi(args[1]); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("val: ", val)
	}
}
