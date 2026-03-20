package models

import (
	"fmt"
	"os"
	"os/exec"
)

func Reboot() {
	cmd := exec.Command("sudo", "reboot")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Erreur lors de la mise à jour :", err)
	}
}
