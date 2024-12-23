package main

import (
	"fmt"
	"time"

	"github.com/escabora/scheduler/internal/schedule"
)

func main() {
	s := schedule.NewScheduler()
	s.AddTask("* * * * *", func() {
		fmt.Println("Running static task")
	})

	s.Start()
	time.Sleep(10 * time.Minute)
	s.Stop()
}
