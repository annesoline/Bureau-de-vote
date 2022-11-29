package main

import (
	"fmt"

	"gitlab.utc.fr/aguilber/ia04/agt/restserveuragent"
)

func main() {
	const url1 = ":8080"

	ServeurAgt := restserveuragent.NewRestServeurAgent(url1)
	go ServeurAgt.Start()
	fmt.Println("Démarrage du serveur")
	fmt.Scanln()
}
