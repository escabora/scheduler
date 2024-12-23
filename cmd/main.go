package main

import (
	"fmt"
	"time"

	"github.com/escabora/scheduler/internal/core"
)

func main() {
	s := core.NewScheduler()
	s.AddTask("* * * * *", func() {
		fmt.Println("Running static task")
	})

	s.Start()
	time.Sleep(10 * time.Minute)
	s.Stop()
}
