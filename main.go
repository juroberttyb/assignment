package main

import (
	"assignment/infra"
)

func main() {
	infra.InitRouter()
	infra.Router.Run()
}
