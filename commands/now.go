//http://qiita.com/kazamai/items/3699c24b16ed4fb883d4
package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"time"
)

func now(c *cli.Context) {
	now := time.Now()
	fmt.Println("now: ", now)
	fmt.Println("now: ", now.Format("20060102"))       // YYYYMMDD
	fmt.Println("now: ", now.Format("20060102150405")) //YYMMDDhhmmss
}
