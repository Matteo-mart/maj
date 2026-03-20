package main

import (
	"log"
	"utils/models"
)

func main() {

	log.Printf("lancement")
	models.ClearTerminal()
	models.Root()
	models.Maj()
	models.Refresh()
	// models.Reboot()

}
