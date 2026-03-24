package iterator

type PlayList interface {
	AddVideoToPlaylist(video Video)
	CreateIterator() Iterator
}
