package tests

import (
	"testing"
	"time"

	"github.com/escabora/scheduler/internal/core"
	"github.com/escabora/scheduler/internal/predictions"
	"github.com/escabora/scheduler/internal/utils"
	"github.com/escabora/scheduler/models"
)

func TestSchedulerAddTask(t *testing.T) {
	s := core.NewScheduler()

	called := false
	s.AddTask("* * * * *", func() {
		called = true
	})

	s.RunTasks()

	if !called {
		t.Errorf("Expected task to be called, but it was not.")
	}
}

func TestMatchSchedule(t *testing.T) {
	now := time.Now()

	if !utils.MatchSchedule("* * * * *", now) {
		t.Errorf("Expected schedule to match, but it did not.")
	}

	specificTime := now.Format("15:04")

	if !utils.MatchSchedule(specificTime, now) {
		t.Errorf("Expected specific schedule to match, but it did not.")
	}
}

func TestPredictor(t *testing.T) {
	predicted := make(map[string]models.PredictedTask)
	called := false

	predicted["test"] = models.PredictedTask{
		Name:    "test",
		Job:     func() { called = true },
		NextRun: time.Now().Add(-time.Minute),
	}

	predictions.RunPredictions(predicted)

	if !called {
		t.Errorf("Expected predicted task to run, but it did not.")
	}
}
