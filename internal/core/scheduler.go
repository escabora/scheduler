package core

import (
	"time"

	"github.com/escabora/scheduler/internal/models"
	"github.com/escabora/scheduler/internal/predictions"
	"github.com/escabora/scheduler/internal/utils"
)

type Scheduler struct {
	tasks     []models.Task
	predicted map[string]models.PredictedTask
	quit      chan bool
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		tasks:     []models.Task{},
		predicted: make(map[string]models.PredictedTask),
		quit:      make(chan bool),
	}
}

func (s *Scheduler) AddTask(schedule string, job func()) {
	s.tasks = append(s.tasks, models.Task{Schedule: schedule, Job: job})
}

func (s *Scheduler) Start() {
	go func() {
		for {
			select {
			case <-s.quit:
				return
			default:
				s.runTasks()
				predictions.RunPredictions(s.predicted)
				time.Sleep(time.Second)
			}
		}
	}()
}

func (s *Scheduler) Stop() {
	s.quit <- true
}

func (s *Scheduler) runTasks() {
	now := time.Now()

	for _, task := range s.tasks {
		if utils.MatchSchedule(task.Schedule, now) {
			task.Job()
		}
	}
}
