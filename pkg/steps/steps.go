package steps

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
