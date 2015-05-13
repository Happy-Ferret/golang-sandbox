package commands

import (
	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	{
		Name:        "args",
		Usage:       "Get and check command-line arguments.",
		Description: `Get and check command-line arguments.`,
		Action:      args,
	},
	{
		Name:        "exec",
		Usage:       "Execute other command.",
		Description: `Execute other command.`,
		Action:      shellexec,
	},
	{
		Name:        "fileexists",
		Usage:       "Check existing file.",
		Description: `Check existing file.`,
		Action:      fileexists,
	},
	{
		Name:        "fileread",
		Usage:       "Read a file.",
		Description: `Read a file.`,
		Action:      fileread,
	},
	{
		Name:        "filewrite",
		Usage:       "Write a file.",
		Description: `Write a file.`,
		Action:      filewrite,
	},
	{
		Name:        "listfile",
		Usage:       "List files in a directory.",
		Description: `List files in a directory.`,
		Action:      listfile,
	},
	{
		Name:        "listfileindir",
		Usage:       "List files in a directory.",
		Description: `List files in a directory.`,
		Action:      listfileindir,
	},
	{
		Name:        "now",
		Usage:       "Get current time.",
		Description: `Get current time.`,
		Action:      now,
	},
	{
		Name:        "stringtoi",
		Usage:       "Convert string to int.",
		Description: `Convert string to int.`,
		Action:      stringtoi,
	},
}
