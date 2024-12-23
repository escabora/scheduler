# Scheduler Package

## Overview
This package provides a task scheduler in Go with support for both static and predictive scheduling of tasks.

### Features
- **Static Scheduling:** Schedule tasks with cron-like syntax.
- **Predictive Scheduling:** Automatically estimate future execution times based on runtime behavior.
- **Extensible Design:** Easy to add custom matchers or prediction models.

### Installation
```sh
go get github.com/escabora/scheduler
```

### Usage
```go
import (
	"github.com/escabora/scheduler/internal/core"
)

func main() {
	s := core.NewScheduler()

	s.AddTask("* * * * *", func() {
		fmt.Println("Running static task")
	})

	s.Start()
	defer s.Stop()
}
```

### Testing
Run tests with:
```sh
go test ./tests/...
```