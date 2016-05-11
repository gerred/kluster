package main

import (
	"github.com/gerred/kluster/cmd"
	"os"
	"fmt"
)

func main() {
    if err := cmd.RootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }
}