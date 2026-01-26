package operations

import (
	"time"

	common "github.com/nebius/gosdk/proto/nebius/common/v1"
)

type CurrentStep interface {
	// Description returns a human-readable description the step
	Description() string

	// StartedAt returns the step start timestamp
	StartedAt() time.Time

	// FinishedAt returns the step finished timestamp or zero time if the step is not finished yet
	FinishedAt() time.Time

	// WorkDone returns the number of done ticks and total ticks of work for the step
	// supplied with this current step, if known.
	WorkDone() *common.ProgressTracker_WorkDone

	// WorkFraction returns the fraction of work completed for the operation
	// supplied with this progress tracker.
	// If the operation has no work, it returns 0.0.
	// If the operation is finished, it returns 1.0.
	// The second return value indicates whether the work fraction is available.
	// It is false if there are no work information.
	WorkFraction() (float64, bool)
}

type OperationProgressTracker interface {
	CurrentStep

	// EstimatedFinishedAt returns the estimated finished at time of the operation
	// supplied with this progress tracker. If not set, returns zero time.
	// For the finished operation, it returns the actual finished at time.
	EstimatedFinishedAt() time.Time

	// TimeFraction returns the fraction of time elapsed for the operation
	// supplied with this progress tracker.
	// If the operation has not started yet, it returns 0.0.
	// If the operation is finished, it returns 1.0.
	// The second return value indicates whether the time fraction is available.
	// It is false if the start time is unknown.
	TimeFraction() (float64, bool)

	// Steps returns the steps of the operation supplied with this progress tracker.
	// May contain only currently running steps or both finished and currently running steps.
	// May contain only a subset of finished steps.
	// May be empty if there are no steps or the steps are not provided.
	Steps() []CurrentStep
}

func WrapProgressTracker(operation *opWrapper) OperationProgressTracker {
	if operation == nil || operation.proto.GetProgressTracker() == nil {
		return nil
	}
	return &progressTrackerWrapper{
		op: operation,
	}
}

type currentStep struct {
	step *common.ProgressTracker_Step
}

type progressTrackerWrapper struct {
	op *opWrapper
}

// Description returns a human-readable description of the progress tracker.
func (pt *progressTrackerWrapper) Description() string {
	pt.op.mu.RLock()
	defer pt.op.mu.RUnlock()
	return pt.op.proto.GetProgressTracker().GetDescription()
}

// EstimatedFinishedAt returns the estimated finished at time of the progress tracker.
// If not set, returns zero time.
// For the finished operation, it returns the actual finished at time.
func (pt *progressTrackerWrapper) EstimatedFinishedAt() time.Time {
	pt.op.mu.RLock()
	defer pt.op.mu.RUnlock()
	if actual := pt.op.proto.GetProgressTracker().GetFinishedAt(); actual != nil {
		return actual.AsTime()
	}

	if opActual := pt.op.proto.GetFinishedAt(); opActual != nil {
		return opActual.AsTime()
	}

	if est := pt.op.proto.GetProgressTracker().GetEstimatedFinishedAt(); est != nil {
		return est.AsTime()
	}
	return time.Time{}
}

// StartedAt returns the start timestamp of the progress tracker.
func (pt *progressTrackerWrapper) StartedAt() time.Time {
	pt.op.mu.RLock()
	defer pt.op.mu.RUnlock()
	startedAt := pt.op.proto.GetProgressTracker().GetStartedAt()
	if startedAt == nil {
		return time.Time{}
	}
	return startedAt.AsTime()
}

func (pt *progressTrackerWrapper) FinishedAt() time.Time {
	pt.op.mu.RLock()
	defer pt.op.mu.RUnlock()
	finishedAt := pt.op.proto.GetProgressTracker().GetFinishedAt()
	if finishedAt == nil {
		return time.Time{}
	}
	return finishedAt.AsTime()
}

// TimeFraction returns the fraction of time elapsed for the progress tracker.
// If the operation has not started yet, it returns 0.0.
// If the operation is finished, it returns 1.0.
// The second return value indicates whether the time fraction is available.
// It is false if the start time is unknown.
func (pt *progressTrackerWrapper) TimeFraction() (float64, bool) {
	pt.op.mu.RLock()
	defer pt.op.mu.RUnlock()
	if pt.op.Done() {
		return 1.0, true
	}

	startedAt := pt.op.proto.GetProgressTracker().GetStartedAt()
	if startedAt == nil {
		return 0.0, false
	}
	estFinishedAt := pt.op.proto.GetProgressTracker().GetEstimatedFinishedAt()
	if estFinishedAt == nil {
		return 0.0, false
	}
	startTime := startedAt.AsTime()
	estFinishTime := estFinishedAt.AsTime()
	now := time.Now()
	if now.Before(startTime) {
		return 0.0, true
	}
	if now.After(estFinishTime) {
		return 1.0, true
	}
	totalDuration := estFinishTime.Sub(startTime).Seconds()
	elapsedDuration := now.Sub(startTime).Seconds()
	if totalDuration <= 0 || elapsedDuration < 0 {
		return 0.0, false
	}
	return elapsedDuration / totalDuration, true
}

// WorkFraction returns the fraction of work completed for the progress tracker.
// If the operation has no work, it returns 0.0.
// If the operation is finished, it returns 1.0.
// The second return value indicates whether the work fraction is available.
// It is false if there are no work information.
func (pt *progressTrackerWrapper) WorkFraction() (float64, bool) {
	pt.op.mu.RLock()
	defer pt.op.mu.RUnlock()
	if pt.op.Done() {
		return 1.0, true
	}
	workDone := pt.op.proto.GetProgressTracker().GetWorkDone()
	if workDone == nil {
		return 0.0, false
	}
	if workDone.GetTotalTickCount() == 0 {
		return 0.0, false
	}
	return float64(workDone.GetDoneTickCount()) / float64(workDone.GetTotalTickCount()), true
}

// WorkDone returns the number of done ticks and total ticks of work for the progress tracker.
func (pt *progressTrackerWrapper) WorkDone() *common.ProgressTracker_WorkDone {
	pt.op.mu.RLock()
	defer pt.op.mu.RUnlock()
	return pt.op.proto.GetProgressTracker().GetWorkDone()
}

// Steps returns the steps of the progress tracker.
func (pt *progressTrackerWrapper) Steps() []CurrentStep {
	pt.op.mu.RLock()
	defer pt.op.mu.RUnlock()
	ret := make([]CurrentStep, 0, len(pt.op.proto.GetProgressTracker().GetSteps()))
	for _, step := range pt.op.proto.GetProgressTracker().GetSteps() {
		ret = append(ret, &currentStep{step: step})
	}
	return ret
}

func (cs *currentStep) FinishedAt() time.Time {
	finishedAt := cs.step.GetFinishedAt()
	if finishedAt == nil {
		return time.Time{}
	}
	return finishedAt.AsTime()
}

func (cs *currentStep) Description() string {
	return cs.step.GetDescription()
}

func (cs *currentStep) StartedAt() time.Time {
	startedAt := cs.step.GetStartedAt()
	if startedAt == nil {
		return time.Time{}
	}
	return startedAt.AsTime()
}

func (cs *currentStep) WorkDone() *common.ProgressTracker_WorkDone {
	return cs.step.GetWorkDone()
}

func (cs *currentStep) WorkFraction() (float64, bool) {
	if cs.step.GetWorkDone() == nil {
		return 0.0, false
	}
	workDone := cs.step.GetWorkDone()
	if workDone.GetTotalTickCount() == 0 {
		return 0.0, false
	}
	return float64(workDone.GetDoneTickCount()) / float64(workDone.GetTotalTickCount()), true
}
