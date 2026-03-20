package models

import (
	"fmt"
	"os"
	"os/exec"
)

func Root() {
	cmd := exec.Command("sudo", "-s")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Erreur root:", err)
	}
	fmt.Printf("root sans erreur\n")
}
