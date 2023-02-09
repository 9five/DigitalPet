package main

import (
	"goDDD/router"
)

func main() {
	r := router.NewRouter()
	r.Run(":8000")
}
