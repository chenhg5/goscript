package main

import (
	"fmt"
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
	}
}
