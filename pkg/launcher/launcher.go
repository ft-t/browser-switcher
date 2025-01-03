package launcher

type Launcher struct {
	targetURL string
}

func New(targetURL string) *Launcher {
	return &Launcher{
		targetURL: targetURL,
	}
}
