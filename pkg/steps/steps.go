package steps

import (
	"container/list"
	"fmt"
	"sync"

	"github.com/mfojtik/oinc/pkg/log"
)

type Step interface {
	Execute() error
	IsParallel() bool
}

type DefaultStep struct {
}

func (s *DefaultStep) IsParallel() bool { return false }

type ParallelStep struct {
}

func (s *ParallelStep) IsParallel() bool { return true }

// StepList represents a doubly linked list of configuration steps
type StepList struct {
	// Name is the list name (used for reporting errors)
	Name string
	list.List
}

// NewStepList returns an empty, initialized StepList
func NewStepList(name string) *StepList {
	return &StepList{Name: name}
}

// Add adds a new Step at the back of the list and returns the updated steps list
// back for chaining
func (l *StepList) Add(s Step) *StepList {
	l.PushBack(s)
	return l
}

// Execute iterates over the list and executes the steps in order they were
// added to the list. It returns an error when one of the steps failed, but it
// will continue executing until all steps run.
func (l *StepList) Execute() error {
	var (
		wg       sync.WaitGroup
		hasError bool
	)
	for s := l.Front(); s != nil; s = s.Next() {
		// Get the step from list
		step, _ := s.Value.(Step)
		log.Info("Step %q ...", step)

		if step.IsParallel() {
			wg.Add(1)
			go func() {
				defer wg.Done()
				if err := step.Execute(); err != nil {
					hasError = true
					log.Error("Step %q failed to execute: %v", step, err)
				}
			}()
			continue
		}

		if err := step.Execute(); err != nil {
			hasError = true
			log.Error("Step %q failed to execute: %v", step, err)
		}
	}
	wg.Wait()
	if hasError {
		return fmt.Errorf("Failed to execute steps during %q", l.Name)
	}
	return nil
}
