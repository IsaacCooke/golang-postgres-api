package main

import (
	"api/controllers"
	"api/data"
)

func main() {
	data.Connect()
	controllers.RunServer()
}
