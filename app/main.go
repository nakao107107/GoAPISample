package main

import (
	"goAPISample/app/db"
	"goAPISample/app/router"
)

func main() {
	db.Init()
	router.Init()
}
