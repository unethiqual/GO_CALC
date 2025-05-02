package main

import (
	"github.com/unethiqual/GO_CALC/internal/agent"
	"github.com/unethiqual/GO_CALC/internal/config"
)

func main() {
	cfg := config.LoadConfig()

	agent := agent.New(cfg)
	agent.Run()
}
