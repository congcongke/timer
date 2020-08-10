package main

import (
	"github.com/congcongke/timer/cmd/timer"
)

func main() {
	cmd := timer.NewTimerCommand()

	cmd.Execute()
}
