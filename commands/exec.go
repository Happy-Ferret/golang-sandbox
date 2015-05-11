//http://qiita.com/kazamai/items/3699c24b16ed4fb883d4
package commands

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
)

func execscript() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("Usage: ", path.Base(args[0]), " cmd")
	}
	out, err := exec.Command(args[1], args[2:]...).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Out: ", string(out))
}
