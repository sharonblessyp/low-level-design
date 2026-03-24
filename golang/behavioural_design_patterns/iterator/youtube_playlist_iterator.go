package iterator

type YoutubePlayListIterator struct {
	videos   []*Video
	position int // stores last accessed index
}

func NewYoutubePlayListIterator(videos []*Video) *YoutubePlayListIterator {
	return &YoutubePlayListIterator{
		videos:   videos,
		position: 0,
	}
}

func (pi *YoutubePlayListIterator) Next() *Video {
	if !pi.HasNext() {
		return nil
	}

	video := pi.videos[pi.position]
	pi.position++
	return video
}

func (pi *YoutubePlayListIterator) HasNext() bool {
	return pi.position < len(pi.videos)
}
