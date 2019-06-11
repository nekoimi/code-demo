package main

import (
	"fmt"
	"os/exec"
)

func main()  {

	var (
		cmd *exec.Cmd
	)

	cmd = exec.Command("/usr/bin/bash", "-c", "echo hello world")

	output ,err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(output))

}
