package main

import (
	"sample/api"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	api.Run()
}
