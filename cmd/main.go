package main

import (
	"fmt"
	"os"

	"github.com/saurabh2013/microservice-basic/cmd/core.go"
	"github.com/saurabh2013/microservice-basic/internal/constant"
)

func main() {

	// initialize service
	env := os.Getenv(constant.Env)
	_, err := core.InitService(env)
	if err != nil {
		fmt.Printf("env: %v, service init failed %v \n", env, err)
	}

	// configure auth middleware if any

	// configure routes

	// start http service
	// port, err := config.Get(constant.Port)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// graceful shutdown and clean up

}
