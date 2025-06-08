package main

import (
	"github.com/joho/godotenv"
	"hospify/src/config"
)

func main() {
  godotenv.Load()
  config.InitDB()
}
