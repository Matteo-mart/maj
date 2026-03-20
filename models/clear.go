package models

import (
	"os"
	"os/exec"
	"runtime"
)

/*
Commande clear pour terminal
*/
func ClearTerminal() {
	var cmd *exec.Cmd

	// Pour Windows
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		// Pour Linux et macOS
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
