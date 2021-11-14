package main

import (
	"fmt"
	"os/exec"
)

func main() {
	//Shell content hardcoded
	content := "echo 'This is my script';touch bad.py; echo 'End of script'"

	//execute binary
	cmd := exec.Command("/bin/sh", "-c", content)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	fmt.Println(string(stdout))
	//fmt.Println(stdout, content)
}
