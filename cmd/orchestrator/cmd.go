package main

import (
	"github.com/unethiqual/GO_CALC/internal/orchestrator"
)

func main() {
	app := orchestrator.New()

	app.Run()
}
