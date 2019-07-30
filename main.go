package main

import (
	"fmt"
	"github.com/abiosoft/ishell"
	"os"
)

func main() {

	args := os.Args

	switch args[1] {
	case "run":
		if len(args) < 3 {
			return
		}
		script := args[2]
		res, err := Exec(script)
		if err != nil {
			panic(err)
		}
		fmt.Println(res)
	case "shell":
		// create new shell.
		// by default, new shell includes 'exit', 'help' and 'clear' commands.
		shell := ishell.New()

		// display welcome info.
		shell.Println("Sample Interactive Shell")

		// register a function for "multi" command.
		shell.AddCmd(&ishell.Cmd{
			Name: "multi",
			Help: "input in multiple lines",
			Func: func(c *ishell.Context) {
				c.Println("Input multiple lines and end with semicolon ';'.")
				lines := c.ReadMultiLines(";")
				c.Println("execute result:")
				res, err := Exec(lines)
				if err != nil {
					panic(err)
				}
				c.Println(res)
			},
		})

		// run shell
		shell.Run()
	}
}
