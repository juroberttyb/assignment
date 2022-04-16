package main

import (
	"server/infra"
)

func main() {
	infra.InitRouter()
	infra.Router.Run()
}
