package steps

type ImagesStep struct {
	PullImages bool
	DefaultStep
}

func (s *ImagesStep) String() string { return "images" }

func (s *ImagesStep) Execute() error {
	return NewStepList(s.String()).
		Add(&DownloadReleaseStep{}).
		Add(&PullDockerImages{PullImages: s.PullImages}).
		Add(&CleanupExistingStep{}).
		Execute()
}
