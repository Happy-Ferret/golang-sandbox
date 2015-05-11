//http://qiita.com/kazamai/items/3699c24b16ed4fb883d4
package commands

import (
	"fmt"
	"log"
	"os"
	"path"
)

func args() {
	args := os.Args
	if len(args) <= 1 {
		log.Fatal("Usage: ", path.Base(args[0]), " anything")
	}

	fmt.Println("len: ", len(args), " values: ", args)
}
