package main

import (
	env "license-service/pkg/env"
	logs "license-service/pkg/log/logger"
)

func main() {

	env.Load()
	_ = logs.NewLogger()

}
