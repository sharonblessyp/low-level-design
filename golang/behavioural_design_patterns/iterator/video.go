package iterator

type Video struct {
	title string
}

func NewVideo(title string) *Video {
	return &Video{title: title}
}

func (v *Video) GetTitle() string {
	return v.title
}
