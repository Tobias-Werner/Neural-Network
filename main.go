package main

import (
	"Landsat-Extractor/logger"
)

func main() {

	logger.Create()
	defer logger.Destroy()

	logger.Info.Println("Test")

}