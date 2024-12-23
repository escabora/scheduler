package predictions

import (
	"math/rand"
	"time"

	"github.com/escabora/scheduler/models"
)

func RunPredictions(predicted map[string]models.PredictedTask) {
	now := time.Now()
	for name, task := range predicted {
		if now.After(task.NextRun) {
			task.Job()
			task.LastRun = now
			predicted[name] = predictedNextRun(task)
		}
	}
}

func predictedNextRun(task models.PredictedTask) models.PredictedTask {
	task.Frequency = time.Duration(rand.Intn(5)+1) * time.Minute
	task.NextRun = time.Now().Add(task.Frequency)
	return task
}
