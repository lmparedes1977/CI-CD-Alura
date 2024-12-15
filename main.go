package main

import (
	"github.com/lmparedes1977/CI-CD-Alura/database"
	"github.com/lmparedes1977/CI-CD-Alura/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
