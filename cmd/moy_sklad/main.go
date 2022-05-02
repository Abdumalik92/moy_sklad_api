package main

import (
	"github.com/Abdumalik92/moy_sklad_api/internal/pkg/utils"
	"github.com/Abdumalik92/moy_sklad_api/internal/routes"
)

func main() {
	utils.ReadSettings()

	routes.RunAllRoutes()
}
