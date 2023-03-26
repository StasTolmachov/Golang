package main

import (
	"os"
	"os/exec"
)

func main() {
	//qnap MTGFTP
	cmd := exec.Command("open", "smb://admin:Stasvv-1@nas522d50/mtgftp")
	cmd.Stdout = os.Stdout
	cmd.Run()

	//qnap media
	cmd = exec.Command("open", "smb://admin:Stasvv-1@nas522d50/media")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
