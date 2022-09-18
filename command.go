package security

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func ExecCommand(cmd string) (command *exec.Cmd) {
	cmdAttr := strings.Split(cmd, " ")
	if len(cmdAttr) > 1 {
		command = exec.Command(cmdAttr[0], cmdAttr[1])
	} else {
		command = exec.Command(cmd)
	}
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()

	if err != nil {
		log.Fatal(err)
	}

	return
}
