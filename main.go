package main

import (
	"server/infra"
)

// curl localhost:8080
// curl localhost:8080 --request "POST"
// curl localhost:8080/...?vip=1 --request "PATCH"

func main() {
	infra.InitRouter()
	infra.Router.Run()
}
