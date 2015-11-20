package steps

type PreConfigStep struct {
}

func (s *PreConfigStep) String() string        { return "pre-config" }
func (s *PreConfigStep) IsParallel() bool      { return false }
func (s *PreConfigStep) Done() <-chan struct{} { return nil }

func (s *PreConfigStep) Execute() error {
	return NewStepList(s.String()).
		Add(&PullDockerImages{}).
		Add(&ModifyDockerConfigStep{}).
		Execute()
}
