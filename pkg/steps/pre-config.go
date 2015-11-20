package steps

type PreConfigStep struct {
	DefaultStep
}

func (s *PreConfigStep) String() string { return "pre-config" }

func (s *PreConfigStep) Execute() error {
	return NewStepList(s.String()).
		Add(&ModifyDockerConfigStep{}).
		Add(&PrepareDirsStep{}).
		Add(&RestartDockerStep{}).
		Execute()
}
