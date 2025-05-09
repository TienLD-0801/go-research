package main

import (
	configs "go-backend/internal/configs"
	env_configs "go-backend/internal/configs/env"
)

func main() {
	env_configs.LoadEnv()
	server := configs.NewServer()
	server.Start()
}
