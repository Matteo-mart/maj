package models

import (
	"fmt"
	"os"
	"os/exec"
)

func Refresh() {
	cmd := exec.Command("sudo", "zypper", "refresh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Erreur lors du refresh :", err)
	}
	fmt.Printf("refresh sans erreur\n")
}
