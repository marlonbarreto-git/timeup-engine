package main

import (
	"fmt"
	"timeup-engine/cmd/api/internal/infra"
	"timeup-engine/cmd/api/internal/routes"
)

var Ctx AppContext

type AppContext struct {
	firebase infra.FirebaseConn
}

func main() {
	Ctx.firebase = infra.NewFirebaseConn()

	err := routes.StartGin()
	if err != nil {
		fmt.Println("Fail")
	}
}
