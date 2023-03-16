package main

import (
	"os"
	"os/exec"
)

func main() {
	GitUsuali()
}

func GitUsuali() {
	cmd := exec.Command("git", "status")
	cmd.Stdout = os.Stdout
	cmd.Run()
	cmd = exec.Command("git", "add", ".")
	cmd.Stdout = os.Stdout
	cmd.Run()
	cmd = exec.Command("git", "commit", "-m", "'test'")
	cmd.Stdout = os.Stdout
	cmd.Run()
	cmd = exec.Command("git", "push")
	cmd.Stdout = os.Stdout
	cmd.Run()
	cmd = exec.Command("git", "status")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
