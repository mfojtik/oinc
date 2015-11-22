package steps

type ImagesStep struct {
	DefaultStep
}

func (s *ImagesStep) String() string { return "images" }

func (s *ImagesStep) Execute() error {
	return NewStepList(s.String()).
		Add(&DownloadReleaseStep{}).
		Add(&PullDockerImages{}).
		Add(&RunOpenShiftStep{}).
		Execute()
}
