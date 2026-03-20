package models

import (
	"fmt"
	"os"
	"os/exec"
)

func Maj() {
	cmd := exec.Command("sudo", "zypper", "dup")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Erreur lors de la mise à jour :", err)
	}
	fmt.Printf("maj sans erreur\n")
}
