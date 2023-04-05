package main

import (
	"os"
	"os/exec"
)

func main() {
	GitRegularly()
}

func GitRegularly() {
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
