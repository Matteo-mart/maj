package models

import (
	"fmt"
)

/*
Permet de sélectionner l'action souhaité
*/
func Choix() {

	for {
		fmt.Println("------------------------")
		fmt.Println("1: Mise à jour")
		fmt.Println("2: Refresh")
		fmt.Println("3: Reboot")
		fmt.Println("0: Quitter")
		fmt.Println("------------------------")
		fmt.Print("Choix : ")

		var choix int
		fmt.Scan(&choix)

		ClearTerminal()

		switch choix {
		case 1:
			fmt.Println("--- Mise à jour ---")
			Maj()
		case 2:
			fmt.Println("--- Rafraîchissement ---")
			Refresh()
		case 3:
			fmt.Println("--- Reboot ---")
			Reboot()
		case 0:
			fmt.Println("Au revoir ")
			return
		default:
			fmt.Println("Choix invalide")
		}
	}
}
