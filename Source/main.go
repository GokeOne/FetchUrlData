package main

import (
	"FetchUrlData/checker"
	"FetchUrlData/utils"
	"flag"
)

func main() {
	// Parseamos la flag -d para obtener el dominio
	domain := flag.String("d", "", "Specify the domain to scan")
	flag.Parse()

	// Validamos que el dominio haya sido proporcionado
	checker.ValidateDomain(*domain)

	// Obtener las URLs usando gau
	urls := utils.GetUrls(*domain)

	// Filtrar las URLs válidas (status 200)
	validUrls := utils.CheckValidUrls(urls)

	// Buscar patrones sensibles en las URLs válidas
	utils.SearchSensitivePatterns(validUrls)
}
