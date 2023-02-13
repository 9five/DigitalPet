package main

import (
	"digitalPet/router"
)

func main() {
	r := router.NewRouter()
	r.Run(":8000")
}
