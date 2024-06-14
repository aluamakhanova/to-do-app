package main

import (
	"github.com/aluamakhanova/to-do-app/config"
	"github.com/aluamakhanova/to-do-app/internal"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()
)

func main() {
	defer config.DisconnectDB(db)

	internal.Routes()
}
