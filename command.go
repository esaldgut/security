package security

import (
	"log"
	"os"
	"os/exec"
)

func ExecCommand(cmd string) (command *exec.Cmd) {
	command = exec.Command(cmd)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()

	if err != nil {
		log.Fatal(err)
	}

	return
}
