package main

import (
	"os"
	"fmt"
	"os/exec"
)

func main() {
	args := os.Args
	numArgs := len(args)

	if (numArgs <= 1) {
		os.Exit(1)
	}
	cmdName := args[1]
	cmdArgs := args[2:numArgs]

	var (
		cmdOut []byte
		err error
	)
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprint(os.Stderr, "Execution failed: ", err)
		os.Exit(1)
	}

	result := string(cmdOut)
	fmt.Print(result)
}
