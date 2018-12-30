package main

import (
	"blockchain_project/core"
)

func main() {
	bc := core.NewBlockChain()
    defer bc.Db.Close()

	cli := core.CLI{bc}
	cli.Run()
}
