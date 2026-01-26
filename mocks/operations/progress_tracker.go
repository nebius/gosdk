package operations

import (
	time "time"

	"github.com/nebius/gosdk/operations"
	commonpb "github.com/nebius/gosdk/proto/nebius/common/v1"
)

func (s Stub) ProgressTracker() operations.OperationProgressTracker {
	return &ptStub{op: s.op}
}

type ptStub struct {
	op *commonpb.Operation
}

type currentStepStub struct {
	step *commonpb.ProgressTracker_Step
}

// Description implements [operations.CurrentStep].
func (c *currentStepStub) Description() string {
	return c.step.GetDescription()
}

// StartedAt implements [operations.CurrentStep].
func (c *currentStepStub) StartedAt() time.Time {
	startedAt := c.step.GetStartedAt()
	if startedAt == nil {
		return time.Time{}
	}
	return startedAt.AsTime()
}

// FinishedAt implements [operations.CurrentStep].
func (c *currentStepStub) FinishedAt() time.Time {
	finishedAt := c.step.GetFinishedAt()
	if finishedAt == nil {
		return time.Time{}
	}
	return finishedAt.AsTime()
}

// WorkDone implements [operations.CurrentStep].
func (c *currentStepStub) WorkDone() *commonpb.ProgressTracker_WorkDone {
	return c.step.GetWorkDone()
}

// WorkFraction implements [operations.CurrentStep].
func (c *currentStepStub) WorkFraction() (float64, bool) {
	if c.step.GetWorkDone() == nil {
		return 0.0, false
	}
	workDone := c.step.GetWorkDone()
	if workDone.GetTotalTickCount() == 0 {
		return 0.0, false
	}
	return float64(workDone.GetDoneTickCount()) / float64(workDone.GetTotalTickCount()), true
}

// Steps implements [operations.OperationProgressTracker].
func (pt *ptStub) Steps() []operations.CurrentStep {
	steps := pt.op.GetProgressTracker().GetSteps()
	ret := make([]operations.CurrentStep, 0, len(steps))
	for _, step := range steps {
		ret = append(ret, &currentStepStub{step: step})
	}
	return ret
}

// TimeFraction implements [operations.OperationProgressTracker].
func (pt *ptStub) TimeFraction() (float64, bool) {
	pt.op.GetProgressTracker()
	startedAt := pt.op.GetProgressTracker().GetStartedAt()
	estimatedFinishedAt := pt.op.GetProgressTracker().GetEstimatedFinishedAt()
	if startedAt == nil || estimatedFinishedAt == nil {
		return 0.0, false
	}
	startTime := startedAt.AsTime()
	estFinishTime := estimatedFinishedAt.AsTime()
	if estFinishTime.Before(startTime) {
		return 0.0, false
	}
	totalDuration := estFinishTime.Sub(startTime)
	elapsedDuration := time.Since(startTime)
	if elapsedDuration < 0 || totalDuration <= 0 {
		return 0.0, false
	}
	if elapsedDuration > totalDuration {
		return 1.0, true
	}
	return float64(elapsedDuration) / float64(totalDuration), true
}

// WorkDone implements [operations.OperationProgressTracker].
func (pt *ptStub) WorkDone() *commonpb.ProgressTracker_WorkDone {
	return pt.op.GetProgressTracker().GetWorkDone()
}

// WorkFraction implements [operations.OperationProgressTracker].
func (pt *ptStub) WorkFraction() (float64, bool) {
	workDone := pt.op.GetProgressTracker().GetWorkDone()
	if workDone == nil || workDone.GetTotalTickCount() == 0 {
		return 0.0, false
	}
	return float64(workDone.GetDoneTickCount()) / float64(workDone.GetTotalTickCount()), true
}

func (pt *ptStub) FinishedAt() time.Time {
	finishedAt := pt.op.GetProgressTracker().GetFinishedAt()
	if finishedAt == nil {
		return time.Time{}
	}
	return finishedAt.AsTime()
}

// Description returns a human-readable description of the progress tracker.
func (pt *ptStub) Description() string {
	return pt.op.GetProgressTracker().GetDescription()
}

// EstimatedFinishedAt returns the estimated finished at time of the progress tracker.
// If not set, returns zero time.
// For the finished operation, it returns the actual finished at time.
func (pt *ptStub) EstimatedFinishedAt() time.Time {
	if actual := pt.op.GetProgressTracker().GetFinishedAt(); actual != nil {
		return actual.AsTime()
	}

	if opActual := pt.op.GetFinishedAt(); opActual != nil {
		return opActual.AsTime()
	}

	if est := pt.op.GetProgressTracker().GetEstimatedFinishedAt(); est != nil {
		return est.AsTime()
	}
	return time.Time{}
}

// StartedAt returns the start timestamp of the progress tracker.
func (pt *ptStub) StartedAt() time.Time {
	startedAt := pt.op.GetProgressTracker().GetStartedAt()
	if startedAt == nil {
		return time.Time{}
	}
	return startedAt.AsTime()
}
